package main

import (
	"context"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/helper"
)

// ListRepos : list names of all repositories on this build server
func (brs *BuildRepoServer) ListRepos(ctx context.Context, req *pb.ListReposRequest) (*pb.ListReposResponse, error) {
	res := pb.ListReposResponse{}
	e, err := ReadEntries(helper.GetBase())
	res.Entries = e
	if err != nil {
		return nil, err
	}
	// get repoid(s)
	for _, entry := range res.Entries {
		lvr, err := brs.GetLatestVersion(ctx, &pb.GetLatestVersionRequest{Repository: entry.Name, Branch: "master"})
		if err != nil {
			return nil, err
		}

		entry.LatestBuild = lvr

	}
	return &res, nil
}





















































































