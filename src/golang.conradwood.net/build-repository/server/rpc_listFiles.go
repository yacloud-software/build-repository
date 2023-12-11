package main

import (
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/helper"
	"golang.org/x/net/context"
	"io/ioutil"
	"strings"
)

// ListFiles : list all files for a given build
func (brs *BuildRepoServer) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	repo := req.Repository
	if !helper.IsValidName(repo) {
		return nil, fmt.Errorf("Invalid repo name \"%s\"", repo)
	}
	branch := req.Branch
	if !helper.IsValidName(branch) {
		return nil, fmt.Errorf("Invalid branch name \"%s\"", branch)
	}
	build := req.BuildID
	if *debug {
		fmt.Printf("Listing versions for repo %s and branch %s and build %d\n", repo, branch, build)
	}
	repodir := fmt.Sprintf("%s/%s/%s/%d", helper.GetBase(), repo, branch, build)
	res := pb.ListFilesResponse{}
	x, err := ReadEntriesNew(repodir, req.Dir, true)
	if err != nil {
		return nil, err
	}
	res.Entries = x
	return &res, nil
}

// ReadEntries : return list of entries in a repo and a directory within it
// example repo: "/srv/build-repository/artefacts/skel-go/master/11/"
// example dir: "dist"
func ReadEntriesNew(repo string, dir string, recurse bool) ([]*pb.RepoEntry, error) {
	dir = strings.TrimPrefix(dir, "/")
	fis, err := ioutil.ReadDir(repo + "/" + dir)
	if err != nil {
		return nil, err
	}
	var res []*pb.RepoEntry
	for _, fi := range fis {
		re := &pb.RepoEntry{Dir: dir}
		re.Name = fi.Name()
		re.Type = 1
		if fi.IsDir() {
			re.Type = 2
			if recurse {
				ad, err := ReadEntriesNew(repo, dir+"/"+fi.Name(), true)
				if err != nil {
					return nil, err
				}
				res = append(res, ad...)
			}
		}
		re.Domain = helper.GetDomainForRepo(re)
		res = append(res, re)
	}
	return res, nil
}




























































