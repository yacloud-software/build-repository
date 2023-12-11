package main

// don't use it in an untrusted environment!
// it expects clients to be authenticated
// (e.g. h2gproxy)
import (
	"context"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/diskscanner"
	"golang.conradwood.net/build-repository/helper"
	"golang.conradwood.net/go-easyops/cmdline"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"google.golang.org/grpc"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	CONST_RAND_ID_STRING_LEN = 64
)

// static variables
var (
	url         = flag.String("download", "", "if non empty, and a valid buildrepo url, then this will download a file")
	maxsize     = flag.Int("max_gb", 1000, "max gigabyte the repository may grow")
	port        = flag.Int("port", 5004, "The server port")
	httpport    = flag.Int("http_port", 5005, "The http server port")
	hooksdir    = flag.String("hooks", "/srv/build-repository/hooks", "Directory to search for hooks")
	reservedir  = flag.String("reservedir", "/srv/build-repository/reservedbuild", "Directory to save reserved builds in")
	src         = rand.NewSource(time.Now().UnixNano())
	debug       = flag.Bool("debug", false, "enable debug output")
	diskScanner *diskscanner.DiskScanner
)

// BuildRepoServer :
type BuildRepoServer struct {
	cache *Cache
}

// entry point
func main() {
	flag.Parse() // parse stuff. see "var" section above
	listenAddr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Starting build-repository Manager %s\n", listenAddr)
	dmc, err := GetDeployminatorClient(cmdline.GetClientRegistryAddress())
	if err != nil {
		fmt.Printf("Warning - unable to get deployminator client: %s\n", utils.ErrorString(err))
	} else {
		fmt.Printf("Found deployminator at address %s\n", dmc.adr)
	}
	diskScanner = diskscanner.NewDiskScanner()
	diskScanner.Dir = helper.GetBase()
	diskScanner.MaxSize = int64(*maxsize * 1024)
	diskScanner.Start()

	lis, err := net.Listen("tcp4", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		os.Exit(10)
	}
	/*
		fs, err := findFilesInDir("/srv/build-repository/artefacts/atomiccounter", []string{"atomiccounter.tar"})
		utils.Bail("Cannot find files", err)
		fmt.Printf("Found %d matching files\n", len(fs))
		for _, s := range fs {
			fmt.Printf("Found: %s\n", s)
		}
		os.Exit(0)
	*/
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	e := new(BuildRepoServer)
	e.cache = NewCache()
	pb.RegisterBuildRepoManagerServer(grpcServer, e) // created by proto

	go grpcServer.Serve(lis)
	sd := server.NewServerDef()
	sd.SetPort((*port + 1))
	sd.SetRegister(server.Register(
		func(server *grpc.Server) error {
			pb.RegisterBuildRepoManagerServer(server, e)
			return nil
		},
	))
	server.ServerStartup(sd)

}

// UpdateSymLink : remove symlink of old name and point to new
// (maintains a symlink 'latest' to point to next build)
func (brs *BuildRepoServer) UpdateSymLink(dir string, latestBuild int) error {
	linkName := fmt.Sprintf("%s/latest", dir)
	if *debug {
		fmt.Printf("linking \"latest\" in dir %s to %d\n", dir, latestBuild)
	}
	err := os.Chdir(dir)
	if err != nil {
		fmt.Printf("Failed to chdir to %s: %v\n", dir, err)
		return err
	}

	err = os.Symlink(fmt.Sprintf("%d", latestBuild), "latest")
	if err == nil {
		return nil
	}
	if os.IsExist(err) {
		os.Remove(linkName)
		err = os.Symlink(fmt.Sprintf("%d", latestBuild), "latest")
		if err != nil {
			fmt.Printf("Tried to remove symlink but still failed to create it: %s: %s\n", linkName, err)
			return err
		}
	} else {
		fmt.Printf("Failed to create symlink in %s: %v\n", dir, err)
		return err
	}
	return nil
}

// ReadEntries : return list of entries in dir - obsolete use ReadEntriesNew instead!
func ReadEntries(dir string) ([]*pb.RepoEntry, error) {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var res []*pb.RepoEntry
	for _, fi := range fis {
		re := &pb.RepoEntry{}
		re.Name = fi.Name()
		re.Type = 1
		if fi.IsDir() {
			re.Type = 2
		}
		re.Domain = helper.GetDomainForRepo(re)
		res = append(res, re)
	}
	return res, nil
}

// sanity check for filenames
func toFilename(f *pb.File) (string, error) {
	filename := f.Filename
	if strings.Contains(filename, "~") {
		return "", fmt.Errorf("Filename must not contain '~' (%s)", filename)
	}
	filename = fmt.Sprintf("%s/%s/%s/%d/dist/%s", helper.GetBase(), f.Repository, f.Branch, f.BuildID, filename)

	/*
		if *debug {
			fmt.Printf("Filename: %s\n", filename)
		}
	*/
	return filename, nil
}

// sanity check for filenames & add os - specific path
func toLinuxFilename(f *pb.File) (string, error) {
	if f == nil {
		return "", fmt.Errorf("missing filename")
	}
	filename := f.Filename
	if strings.Contains(filename, "~") {
		return "", fmt.Errorf("Filename must not contain '~' (%s)", filename)
	}
	if strings.Contains(filename, "..") {
		return "", fmt.Errorf("Filename must not contain '..' (%s)", filename)
	}
	filename = fmt.Sprintf("%s/%s/%s/%d/%s", helper.GetBase(), f.Repository, f.Branch, f.BuildID, filename)
	if *debug {
		fmt.Printf("Filename: %s\n", filename)
	}
	return filename, nil
}

// sanity check for filenames & add os - specific path
func toDarwinFilename(f *pb.File) (string, error) {
	filename := f.Filename
	if strings.Contains(filename, "~") {
		return "", fmt.Errorf("Filename must not contain '~' (%s)", filename)
	}
	filename = fmt.Sprintf("%s/%s/%s/%d/dist/darwin/amd64/%s", helper.GetBase(), f.Repository, f.Branch, f.BuildID, filename)
	if *debug {
		fmt.Printf("Filename: %s\n", filename)
	}
	return filename, nil
}

// sanity check for filenames & add os - specific path
func toWindowsFilename(f *pb.File) (string, error) {
	filename := f.Filename
	if strings.Contains(filename, "~") {
		return "", fmt.Errorf("Filename must not contain '~' (%s)", filename)
	}
	filename = fmt.Sprintf("%s/%s/%s/%d/dist/windows/amd64/%s", helper.GetBase(), f.Repository, f.Branch, f.BuildID, filename)
	if *debug {
		fmt.Printf("Filename: %s\n", filename)
	}
	return filename, nil
}

func (b *BuildRepoServer) ReserveNextBuildNumber(ctx context.Context, req *pb.RepoDef) (*pb.BuildNumber, error) {
	i, err := helper.GetLatestRepoVersion(req.Repository, req.Branch)
	if err != nil {
		return nil, fmt.Errorf("Unable to get latest version for %s/%s: %s", req.Repository, req.Branch, err)
	}
	newv := i + 1
	dir := fmt.Sprintf("%s/%s/%s", *reservedir, req.Repository, req.Branch)
	filename := fmt.Sprintf("%s/reserved.txt", dir)
	body, err := ioutil.ReadFile(filename)
	if err == nil {
		latest_reserved, err := strconv.Atoi(string(body))
		if err == nil {
			if latest_reserved >= int(newv) {
				newv = uint64(latest_reserved + 1)
			}
		}
	}
	os.MkdirAll(dir, 0700)
	err = ioutil.WriteFile(filename, []byte(fmt.Sprintf("%d", newv)), 0700)
	if err != nil {
		return nil, fmt.Errorf("Failed to write %s: %s", filename, err)
	}
	return &pb.BuildNumber{BuildID: newv}, nil
}

func (b *BuildRepoServer) GetBuildInfo(ctx context.Context, req *pb.BuildDef) (*pb.BuildInfo, error) {
	res := &pb.BuildInfo{}
	s, err := loadMetadata(req.Repository, req.Branch, req.BuildID)
	if err != nil {
		return nil, err
	}
	res.CommitID = s.CommitID
	res.CommitMessage = s.Commitmsg
	res.UserEmail = s.UserEmail
	res.BuildDate = s.BuildDate
	return res, nil
}



