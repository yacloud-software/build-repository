package main

import (
	"crypto/rand"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"io/ioutil"
	"strconv"
	"strings"
)

// Generate a random string of length 'n'.
func randString(n int) (string, error) {
	b := make([]byte, int((n+1)/2))
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", b)[:n], nil
}

// check if it's a valid name for a repo or branch,
// basically no / or .. or . or so allowed
func isValidName(path string) bool {
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

func getLatestRepoVersion(repo string, branch string) (uint64, error) {
	if !isValidName(repo) {
		return 0, fmt.Errorf("Invalid repo name \"%s\"", repo)
	}
	if !isValidName(branch) {
		return 0, fmt.Errorf("Invalid branch name \"%s\"", branch)
	}
	if *debug {
		fmt.Printf("Listing versions for repo %s and branch %s\n", repo, branch)
	}
	repodir := fmt.Sprintf("%s/%s/%s", base, repo, branch)
	e, err := ReadEntries(repodir)
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

// get the versions in a dir
// dir, for example: /srv/build-repository/artefacts/build-repository/master/
func GetVersionsFromDir(dir string, include_incomplete bool) ([]*pb.RepoEntry, error) {
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
		re.Domain = getDomainForRepo(re)
		res = append(res, re)
	}
	return res, nil
}
