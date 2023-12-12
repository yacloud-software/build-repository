package main

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/helper"
	"strconv"
)

// GetLatestVersion : give the latest build number of a given repo/branch
func (brs *BuildRepoServer) GetLatestVersion(ctx context.Context, req *pb.GetLatestVersionRequest) (*pb.GetLatestVersionResponse, error) {
	repo := req.Repository
	fmt.Printf("Getting latest version of repository \"%s\"\n", repo)
	if !helper.IsValidName(repo) {
		return nil, fmt.Errorf("Invalid repo name \"%s\"", repo)
	}
	branch := req.Branch
	if !helper.IsValidName(branch) {
		return nil, fmt.Errorf("Invalid branch name \"%s\"", branch)
	}
	/*
		// really noisy - get's called regularly by ota module for updates to hubs
				if *debug {
					fmt.Printf("getting latest version for repo %s and branch %s\n", repo, branch)
				}
	*/
	repodir := fmt.Sprintf("%s/%s/%s", helper.GetBase(), repo, branch)
	e, err := ReadEntries(repodir)
	if err != nil {
		return nil, err
	}

	bid := -1
	for _, en := range e {

		if en.Type != 2 {
			continue
		}
		x, er := strconv.Atoi(en.Name)
		if er != nil {
			continue
		}
		if x > bid {
			bid = x
		}
	}

	if bid == -1 {
		return nil, fmt.Errorf("could not get latest version for %s", repodir)
	}
	bm, err := getMetadata(repo, branch, uint64(bid))
	if err != nil {
		return nil, err
	}
	res := &pb.GetLatestVersionResponse{
		BuildID:   uint64(bid),
		BuildMeta: bm,
	}
	return res, nil
}



















































































