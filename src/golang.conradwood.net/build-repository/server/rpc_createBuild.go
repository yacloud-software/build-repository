package main

import (
	"context"
	"errors"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/helper"
	"google.golang.org/grpc/peer"
	"os"
)

// CreateBuild :
func (brs *BuildRepoServer) CreateBuild(ctx context.Context, cr *pb.CreateBuildRequest) (*pb.CreateBuildResponse, error) {

	peer, ok := peer.FromContext(ctx)
	if !ok {
		fmt.Println("Error getting peer ")
	}
	if cr.Repository == "" {
		return nil, errors.New("Missing repository name")
	}
	if cr.CommitID == "" {
		return nil, errors.New("Missing commit id")
	}
	if cr.CommitMSG == "" {
		return nil, errors.New("Missing commit message")
	}
	if cr.Branch == "" {
		return nil, errors.New("Missing branch name")
	}
	if cr.BuildID == 0 {
		return nil, errors.New("Missing build id")
	}
	if cr.RepositoryID != 0 {
		err := setRepositoryMetaRepositoryID(ctx, cr.Repository, cr.RepositoryID)
		if err != nil {
			return nil, err
		}
	}
	resp := pb.CreateBuildResponse{}

	dir := fmt.Sprintf("%s/%s/%s/%d", helper.GetBase(), cr.Repository, cr.Branch, cr.BuildID)
	st, err := os.Stat(dir)
	if (err == nil) && (st != nil) {
		return nil, fmt.Errorf("Dir %s already exists. Trying to update an existing build??", dir)
	}
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println("Failed to create directory ", dir, err)
		return &resp, err
	}
	fmt.Println("Created directory:", dir)
	fmt.Println(peer.Addr, "called createbuild")

	id, _ := helper.RandString(CONST_RAND_ID_STRING_LEN)

	brs.cache.SetStored(
		id,
		&StoreMetaData{
			ArtefactID:   cr.ArtefactID,
			RepositoryID: cr.RepositoryID,
			UserEmail:    cr.UserEmail,
			BuildID:      int(cr.BuildID),
			CommitID:     cr.CommitID,
			Commitmsg:    cr.CommitMSG,
			Branch:       cr.Branch,
			Repository:   cr.Repository,
			StorePath:    dir,
			StoreID:      id,
		},
	)

	resp.BuildStoreid = id

	linkdir := fmt.Sprintf("%s/%s/%s", helper.GetBase(), cr.Repository, cr.Branch)
	err = brs.UpdateSymLink(linkdir, int(cr.BuildID))
	if err != nil {
		fmt.Printf("Failed to create symlink in %s: %v\n", linkdir, err)
		return nil, err
	}
	return &resp, nil
}





































