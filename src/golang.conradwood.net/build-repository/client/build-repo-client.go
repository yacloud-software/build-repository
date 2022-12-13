package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/apis/common"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/tokens"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// static variables for flag parser
var (
	download    = flag.String("download", "", "if non empty, and a valid buildrepo url, then this will download a file")
	do_unfail   = flag.Bool("unfail", false, "if true, unfail diskscanner")
	tooldir     = flag.String("tooldir", "", "The location of directory to download tools to (overwrites existing files!).")
	artefact    = flag.String("artefact", "", "Fetch a specific artefact (use with tooldir)")
	gitlabUser  = flag.String("user", "", "The GitLab user.")
	reponame    = flag.String("repository", "", "name of repository")
	repoid      = flag.Uint64("repository_id", 0, "unique id of repository")
	branchname  = flag.String("branch", "", "branch of commit")
	commitid    = flag.String("commitid", "", "commit")
	commitmsg   = flag.String("commitmsg", "", "commit message")
	buildnumber = flag.Int("build", 0, "build number")
	distDir     = flag.String("distdir", "dist", "Default directory to upload")
	dryrun      = flag.Bool("n", false, "dry-run")
	versionfile = flag.String("versionfile", "", "filename of a versionfile to update with buildid")
	versiondir  = flag.String("versiondir", "", "directory to scan for buildversion.go files (update files with buildid)")
	info        = flag.Bool("info", false, "Get information about the repo")
	offset      = flag.Int("offset", -1, "if >0 read a block from the file in the repo beginning at the specified offset")
	blocksize   = flag.Int("blocksize", 512, "Default block size when reading block or file from repo")
	filename    = flag.String("filename", "", "Filename from which to retrieve a block")
	findname    = flag.Bool("find", false, "if true tries to find filename in a repos")
	reserve     = flag.Bool("next-build-number", false, "reserve a build number")
	grpcClient  buildrepo.BuildRepoManagerClient
)

const (
	MAX_UPLOAD_SECS = 180
)

func main() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting build-repository-client:", filepath.Dir(ex))

	flag.Parse()
	if *download != "" {
		DownloadURL()
		os.Exit(0)
	}
	grpcClient = buildrepo.GetBuildRepoManagerClient()

	if *do_unfail {
		_, err := grpcClient.UnfailDiskScanner(createContext(), &common.Void{})
		utils.Bail("failed to unfail", err)
		fmt.Printf("Unfailed\n")
		os.Exit(0)
	}
	if *reserve {
		ReserveBuildNumber()
		os.Exit(0)
	}
	if *findname {
		findfile()
		os.Exit(0)
	}
	if *tooldir != "" {
		if *artefact != "" {
			if *reponame == "" {
				fmt.Println("-reponame is required")
				os.Exit(1)
			}
			branch := "master"
			if *branchname != "" {
				branch = *branchname
			}
			DownloadOneTool(*tooldir, *reponame, branch, *artefact)
		} else {
			DownloadTools(*tooldir)
		}
		os.Exit(0)
	}

	if *versionfile != "" {
		updateVersionFile(*versionfile)
		os.Exit(0)
	}
	if *versiondir != "" {
		updateVersionDir(*versiondir)
		os.Exit(0)
	}
	if *info {
		getInfo()
		os.Exit(0)
	}
	if *offset >= 0 {
		getBlock()
		os.Exit(0)
	}

	files := flag.Args()
	if len(files) == 0 {
		//		fmt.Printf("No files specified on commandline, using \"%s\" as default\n", *distDir)
		df, err := ioutil.ReadDir(*distDir)
		if err != nil {
			fmt.Printf("Failed to read directory \"%s\": %s\n", *distDir, err)
			os.Exit(5)
		}
		for _, file := range df {
			//	fmt.Println(file.Name())
			files = append(files, fmt.Sprintf("%s/%s", *distDir, file.Name()))
		}
		df, err = ioutil.ReadDir("configs")
		if err == nil {
			for _, file := range df {
				//				fmt.Println(file.Name())
				files = append(files, fmt.Sprintf("configs/%s", file.Name()))
			}
		}
	}
	AddDirIfExists("deployment", &files)

	if *dryrun {
		for _, file := range files {
			fmt.Printf("Uploading file: %s\n", file)
		}
		return
	}

	ctx := createContext()

	fmt.Printf("New build %d in repo %s\n", *buildnumber, *reponame)
	req := buildrepo.CreateBuildRequest{
		UserEmail:    *gitlabUser,
		Repository:   *reponame,
		RepositoryID: *repoid,
		CommitID:     *commitid,
		Branch:       *branchname,
		BuildID:      uint64(*buildnumber),
		CommitMSG:    *commitmsg,
	}

	fmt.Printf("Updating buildrepo client...\n")
	fmt.Printf("Creating build...\n")
	resp, err := grpcClient.CreateBuild(ctx, &req)
	if err != nil {
		fmt.Printf("failed to create build: %v\n", err)
		os.Exit(10)
	}
	fmt.Printf("Response to createbuild was: %v\n", resp)

	storeid := resp.BuildStoreid
	if len(files) > 0 {
		fmt.Printf("Beginning uploads for %d files\n", len(files))
		err = uploadFiles(*gitlabUser, storeid, files)
		if err != nil {
			fmt.Printf("Failed to upload files, error reported from build-reposerver is: \"%v\"\n", err)
			os.Exit(5)
		}
	}
	//	fmt.Println("Calling grpc.CompleteUploads...")

	for {
		r, err := grpcClient.UploadsComplete(createContext(), &buildrepo.UploadDoneRequest{BuildStoreid: storeid})
		if err != nil {
			fmt.Printf("Failed to complete upload, error reported from build-reposerver is: \"%v\"\n", err)
			os.Exit(5)
		}
		if r.Closed {
			break
		}
		fmt.Printf("Upload not completed yet: %d uploading\n", r.Uploading)
		time.Sleep(3 * time.Second)
	}
}

func DownloadOneTool(tooldir string, reponame string, branch string, artefact string) {
	if _, err := os.Stat(tooldir); os.IsNotExist(err) {
		os.Mkdir(tooldir, 0755)
	}
	glvr := buildrepo.GetLatestVersionRequest{Repository: reponame, Branch: branch}
	glv, err := grpcClient.GetLatestVersion(createContext(), &glvr)
	utils.Bail("Tool not found", err)
	f := &buildrepo.File{
		Repository: reponame,
		Branch:     branch,
		BuildID:    glv.GetBuildID(),
		Filename:   artefact,
	}
	FetchOneFile(f, err)
}

// DownloadTools : download the toolkit to specified directory
func DownloadTools(tooldir string) {

	if _, err := os.Stat(tooldir); os.IsNotExist(err) {
		os.Mkdir(tooldir, 0755)
	}

	ltr, err := grpcClient.ListTools(createContext(), &common.Void{})

	utils.Bail("Failed to list tools", err)

	fmt.Println()
	fmt.Println("*** download started  ***")
	download_errors := 0
	for _, tool := range ltr.Tools {

		glvr := buildrepo.GetLatestVersionRequest{Repository: tool.GetRepository(), Branch: "master"}
		glv, err := grpcClient.GetLatestVersion(createContext(), &glvr)

		if err != nil {
			fmt.Printf("Failed to get latest version of tool %s %s\n", tool.Filename, err)
			download_errors++
			continue
		}

		f := &buildrepo.File{
			Repository: tool.GetRepository(),
			Branch:     "master",
			BuildID:    glv.GetBuildID(),
			Filename:   tool.GetFilename(),
		}
		FetchOneFile(f, err)
	}
	if download_errors > 0 {
		fmt.Printf("Got %d download errors\n", download_errors)
		os.Exit(1)

	}
	fmt.Println("*** download complete ***")
}

func FetchOneFile(f *buildrepo.File, err error) {

	gfr := &buildrepo.GetFileRequest{
		File:      f,
		Blocksize: uint32(*blocksize),
	}
	fmt.Printf("%v\n", gfr)
	stream, err := grpcClient.GetFileAsStream(createContext(), gfr)
	utils.Bail("Failed to get stream", err)
	fmt.Printf("Downloading %s from repo %s\n", f.GetFilename(), f.GetRepository())
	final_path := *tooldir + "/" + f.GetFilename()
	tmp_path := final_path + ".tmp"
	file, err := os.OpenFile(tmp_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	utils.Bail("opening "+tmp_path, err)
	var i = 0
	for {
		block, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("stream.Recv failed with %v  %v\n", stream, err)
			break
		}

		_, err = file.Write(block.Data[0:block.GetSize()])

		utils.Bail("writing to "+tmp_path, err)

		switch i {
		case 0:
			fmt.Print("\r[|]")
		case 500:
			fmt.Print("\r[/]")
		case 1000:
			fmt.Print("\r[-]")
		case 1500:
			fmt.Print("\r[\\]")
		}
		i++
		if i >= 2000 {
			i = 0
		}

	}
	fmt.Println("\r   ")
	file.Close()
	err = os.Rename(tmp_path, final_path)
	utils.Bail("Renaming "+tmp_path+" to "+final_path, err)
}

// get an arbitrary block from a file from repo
// this is useful for OTA
func getBlock() {

	f := &buildrepo.File{
		Repository: *reponame,
		Branch:     *branchname,
		BuildID:    uint64(*buildnumber),
		Filename:   *filename,
	}

	glv := &buildrepo.GetBlockRequest{
		File:   f,
		Offset: uint64(*offset),
		Size:   uint32(*blocksize),
	}

	glr, err := grpcClient.GetBlock(createContext(), glv)
	utils.Bail("Failed to read block", err)

	fmt.Printf("Response:\n")
	fmt.Printf("Size=%d, Offset=%d\n", glr.Size, glr.Offset)
	fmt.Printf("Data: [%s]\n", glr.Data)
}
func ReserveBuildNumber() {
	rp := &buildrepo.RepoDef{
		Repository: *reponame,
		Branch:     *branchname,
	}
	bn, err := grpcClient.ReserveNextBuildNumber(createContext(), rp)
	if err != nil {
		fmt.Printf("Failed: %s\n", err)
		os.Exit(10)
	}
	fmt.Printf("%d\n", bn.BuildID)
}

// connect to server, get latest version information of a given repo
func getInfo() {
	b := *branchname
	if b == "" {
		b = "master"
	}
	glv := &buildrepo.GetLatestVersionRequest{
		Repository: *reponame,
		Branch:     b,
	}
	ctx := createContext()
	glr, err := grpcClient.GetLatestVersion(ctx, glv)
	utils.Bail("failed to get latest version", err)
	fmt.Printf("Latest Version: %d\n", glr.BuildID)
	files, err := grpcClient.ListFiles(ctx, &buildrepo.ListFilesRequest{
		Repository: glv.Repository,
		Branch:     glv.Branch,
		BuildID:    glr.BuildID,
		Dir:        "",
		Recursive:  true,
	})
	utils.Bail("failed to list files", err)
	for _, f := range files.Entries {
		if f.Type == 2 {
			continue
		}
		fmt.Printf("File: %s/%s\n", f.Dir, f.Name)
	}
	for _, f := range files.Entries {
		if f.Type != 2 {
			continue
		}
		fmt.Printf("Dir: %s/\n", f.Name)
	}
}

// end info stuff

// handles uploaded files
func uploadFiles(userEmail string, storeid string, filenames []string) error {
	if len(filenames) == 0 {
		return nil
	}
	//	fmt.Printf("Got %d files to upload...\n", len(filenames))
	for _, filename := range filenames {
		st, err := os.Stat(filename)
		if err != nil {
			fmt.Printf("Cannot stat %s: %s, skipping...\n", filename, err)
			return err
		}
		if st.Mode().IsDir() {
			var nfiles []string
			df, err := ioutil.ReadDir(filename)
			if err != nil {
				fmt.Printf("Failed to read directory \"%s\": %s\n,", filename, err)
				return errors.New("Failed to read directory")
			}
			for _, file := range df {
				nf := fmt.Sprintf("%s/%s", filename, file.Name())
				nfiles = append(nfiles, nf)
			}
			uploadFiles(userEmail, storeid, nfiles)
			continue

		}
		if !st.Mode().IsRegular() {
			fmt.Printf("Skipping %s - it's not a file\n", filename)
			continue
		}

		//	fmt.Printf("Preparing %s for Upload...\n", filename)

		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Unable to open \"%s\": %s\n", filename, err)
			return err
		}
		defer file.Close()

		ureq := &buildrepo.UploadTokenRequest{
			BuildStoreid: storeid,
			Filename:     filename,
			UserEmail:    userEmail,
		}
		resp, err := grpcClient.GetUploadToken(createContext(), ureq)
		if err != nil {
			fmt.Printf("Failed to upload %s: %v\n", filename, err)
			return err
		}

		ctx := authremote.ContextWithTimeout(time.Duration(300) * time.Second)
		s, err := grpcClient.PutFileAsStream(ctx)
		utils.Bail("Failed to get upload as stream", err)
		data := make([]byte, *blocksize)
		fmt.Printf("Uploading file %s...\n", filename)
		started := time.Now()
		var bytes uint64
		bytes = 0
		for {
			size, err := file.Read(data)
			if err != nil {
				if err == io.EOF {
					break
				}

				return fmt.Errorf("could not read file: %v", err)
			}
			bytes = bytes + uint64(size)
			// size check important, otherwise we'll get an EOF from server
			if size != 0 {
				err = s.Send(&buildrepo.FileBlock{UploadToken: resp.Token, Size: uint64(size), Data: data})
				if err != nil {
					diff := time.Now().Sub(started)
					printRate(diff, bytes)
					// we seem to be getting an eof on some timeout
					utils.Bail(fmt.Sprintf("failed to send %d bytes of file %s to buildrepo after %v (timeout set to %d seconds. too low?)", size, filename, diff, MAX_UPLOAD_SECS), err)
				}
			}
		}
		s.CloseAndRecv()
		/*
			diff := time.Now().Sub(started)
			printRate(diff, bytes)
				fmt.Printf("Completed upload of %s\n", filename)
		*/
	}
	//	fmt.Printf("All files uploaded successfully\n")
	return nil
}

func printRate(t time.Duration, bytes uint64) {
	mbytes := bytes / 1024 / 1024
	var rate float64
	secs := uint64(t / time.Second)
	if secs == 0 {
		fmt.Printf("%d MiB transferred almost immediately\n", mbytes)
		return
	}
	rate = float64(bytes) / float64(secs)
	// in mbit/s
	rate = rate / 1024 / 1024 * 8
	fmt.Printf("%d MiB in %d seconds (%f mbits/sec)\n", mbytes, secs, rate)
}

// AddDirIfExists :
func AddDirIfExists(dirname string, files *[]string) error {
	if !exists(dirname) {
		fmt.Printf("%s does not exist. skipping\n", dirname)
		return nil
	}
	df, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Printf("Failed to read directory \"%s\": %s\n", dirname, err)
		return err
	}
	for _, file := range df {
		fmt.Println(file.Name())
		*files = append(*files, fmt.Sprintf("%s/%s", dirname, file.Name()))
	}
	return nil
}

// recursively go through directory and process all files called buildversion.go
func updateVersionDir(dname string) {
	fos, err := ioutil.ReadDir(dname)
	utils.Bail("Unable to read dir", err)
	for _, file := range fos {
		if file.IsDir() {
			updateVersionDir(fmt.Sprintf("%s/%s", dname, file.Name()))
			continue
		}
		if file.Name() != "buildversion.go" {
			continue
		}
		fullname := fmt.Sprintf("%s/%s", dname, file.Name())
		fmt.Printf("File: %s\n", fullname)
		updateVersionFile(fullname)

	}
}

func updateVersionFile(fname string) {
	bs, err := ioutil.ReadFile(fname)
	utils.Bail("Failed to readfile", err)
	lines := string(bs)
	var buffer bytes.Buffer
	changed := false
	for _, line := range strings.Split(lines, "\n") {
		if !strings.Contains(line, "// AUTOMATIC VERSION UPDATE: OK") {
			buffer.WriteString(line)
			buffer.WriteString("\n")
			continue
		}
		if strings.Contains(line, "Buildnumber") {
			changed = true
			line = strings.Replace(line, "0", fmt.Sprintf("%d", *buildnumber), 1)
		} else if strings.Contains(line, "Build_date_string") {
			changed = true
			line = strings.Replace(line, "today", time.Now().UTC().Format("2006-01-02T15:04:05-0700"), 1)
		} else if strings.Contains(line, "Build_date") {
			changed = true
			line = strings.Replace(line, "0", fmt.Sprintf("%d", time.Now().Unix()), 1)
		}
		buffer.WriteString(line)
		buffer.WriteString("\n")

	}
	if !changed {
		fmt.Printf("File %s was not changed\n", fname)
		return
	}
	s := buffer.String()
	if *buildnumber != 0 {
		err := ioutil.WriteFile(fname, []byte(s), 0777)
		utils.Bail("Failed to write versionfile", err)
		fmt.Printf("File %s updated\n", fname)
	} else {
		fmt.Printf("File %s would have been updated to:\n%s\n", fname, s)
	}

}

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func findfile() {
	fp := &buildrepo.FilePattern{Pattern: *filename}
	r, err := grpcClient.FindFiles(createContext(), fp)
	utils.Bail("Failed to find files", err)
	fmt.Printf("Found %d files\n", len(r.Files))
	for _, rx := range r.Files {
		fmt.Printf("repo: %s, build %d: %s\n", rx.Repository, rx.BuildID, rx.Filename)
	}
}
func createContext() context.Context {
	if 1 == 0 {
		return tokens.ContextWithToken()
	}
	return authremote.Context()
}
