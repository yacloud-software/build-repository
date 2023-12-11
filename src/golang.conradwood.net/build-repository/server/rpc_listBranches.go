package main

import (
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/helper"
	"golang.org/x/net/context"
)

// ListBranches : given a repo, list all branches for which we have builds
func (brs *BuildRepoServer) ListBranches(ctx context.Context, req *pb.ListBranchesRequest) (*pb.ListBranchesResponse, error) {
	repo := req.Repository
	if *debug {
		fmt.Printf("Listing branches of repository %s\n", repo)
	}
	if !helper.IsValidName(repo) {
		return nil, fmt.Errorf("Invalid name \"%s\"", repo)
	}
	repodir := fmt.Sprintf("%s/%s", helper.GetBase(), repo)
	res := pb.ListBranchesResponse{}
	e, err := ReadEntries(repodir)
	res.Entries = e
	if err != nil {
		return nil, err
	}
	return &res, nil
}





















