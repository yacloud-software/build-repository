package main

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/archive"
	"golang.conradwood.net/go-easyops/utils"
)

func (b *BuildRepoServer) DoesFileExist(ctx context.Context, req *pb.GetFileRequest) (*pb.FileExistsInfo, error) {

	filename, err := toLinuxFilename(req.File)

	if err != nil {
		return nil, err
	}

	if utils.FileExists(filename) {
		return &pb.FileExistsInfo{Exists: true}, nil
	}

	// perhaps it is in archive?
	fr, err := archive.DoesFileExist(ctx, req)
	if err != nil {
		return nil, err
	}
	if fr.Exists {
		return &pb.FileExistsInfo{Exists: true}, nil
	}
	return nil, fmt.Errorf("could not get filename: %v", err)

}
