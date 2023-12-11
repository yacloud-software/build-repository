package helper

import (
	"crypto/rand"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/go-easyops/utils"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	base = "/srv/build-repository/artefacts"
)

var (
	metadir = flag.String("metadir", "/srv/build-repository/metadata", "Directory to save metadata in")
)

func GetBase() string {
	return base
}
func GetMetadir() string {
	return *metadir
}

// Generate a random string of length 'n'.
func RandString(n int) (string, error) {
	b := make([]byte, int((n+1)/2))
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", b)[:n], nil
}

// check if it's a valid name for a repo or branch,
// basically no / or .. or . or so allowed
func IsValidName(path string) bool {
	if path == "" {
		return false
	}
	if strings.Contains(path, "/") {
		return false
	}
	if strings.Contains(path, ".") {
		return false
	}
	if strings.Contains(path, "~") {
		return false
	}
	return true
}

func GetLatestRepoVersion(repo string, branch string) (uint64, error) {
	if !IsValidName(repo) {
		return 0, fmt.Errorf("Invalid repo name \"%s\"", repo)
	}
	if !IsValidName(branch) {
		return 0, fmt.Errorf("Invalid branch name \"%s\"", branch)
	}
	repodir := fmt.Sprintf("%s/%s/%s", base, repo, branch)
	e, err := GetVersionsFromDir(repodir, false)
	if err != nil {
		return 0, err
	}
	v := 0
	for _, r := range e {
		vv, err := strconv.Atoi(r.Name)
		if err != nil {
			continue
		}
		if vv > v {
			v = vv
		}
	}
	return uint64(v), nil
}

type pathentry struct {
	root    string // e.g. "/srv/build-repository"
	data    string // e.g. artefacts|metadata
	repo    string // e.g. "autodeployer"
	branch  string // e.g. "master"
	version string // e.g. "21231" or ""
}

func (pe *pathentry) String() string {
	return fmt.Sprintf(`Root: "%s", Data: "%s", Repo: "%s", Branch: "%s", Version: "%s"`, pe.root, pe.data, pe.repo, pe.branch, pe.version)
}

/*
parses a fully qualified path into its components
for example: /srv/build-repository/artefacts/autodeployer/master/
version and branch is optional, it parses as much as is there
*/
func ParsePath(path string) *pathentry {
	res := &pathentry{
		data: "artefacts",
	}
	match := "/artefacts/"
	idx := strings.Index(path, match)

	if idx == -1 {
		match := "/metadata/"
		idx = strings.Index(path, match)
		if idx == -1 {
			return res
		}
		res.data = "metadata"
	}
	res.root = strings.TrimSuffix(path[:idx], "/")
	rem_s := strings.Split(path[idx+len(match):], "/")
	if len(rem_s) > 0 {
		res.repo = rem_s[0]
	}
	if len(rem_s) > 1 {
		res.branch = rem_s[1]
	}
	if len(rem_s) > 2 {
		res.version = rem_s[2]
	}
	return res
}

// get the versions in a dir
// dir, for example: /srv/build-repository/artefacts/autodeployer/master/
func GetVersionsFromDir(dir string, include_incomplete bool) ([]*pb.RepoEntry, error) {
	pe := ParsePath(dir)
	fmt.Printf("[versionsfromdir] %s\n", pe.String())
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
		if !include_incomplete {
			metadatadir := pe.root + "/" + "metadata/" + pe.repo + "/" + pe.branch + "/" + re.Name
			if !utils.FileExists(metadatadir) {
				continue
			}
		}
		re.Domain = GetDomainForRepo(re)
		res = append(res, re)
	}
	return res, nil
}
































