package main

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/helper"
)

// ListVersions : given a repo, list all versions we have (all build numbers)
func (brs *BuildRepoServer) ListVersions(ctx context.Context, req *pb.ListVersionsRequest) (*pb.ListVersionsResponse, error) {
	repo := req.Repository
	if !helper.IsValidName(repo) {
		return nil, fmt.Errorf("Invalid repo name \"%s\"", repo)
	}
	branch := req.Branch
	if !helper.IsValidName(branch) {
		return nil, fmt.Errorf("Invalid branch name \"%s\"", branch)
	}
	if *debug {
		fmt.Printf("Listing versions for repo %s and branch %s\n", repo, branch)
	}
	repodir := fmt.Sprintf("%s/%s/%s", helper.GetBase(), repo, branch)
	res := pb.ListVersionsResponse{}
	e, err := helper.GetVersionsFromDir(repodir, false)
	res.Entries = e
	if err != nil {
		return nil, err
	}
	return &res, nil
}
















































