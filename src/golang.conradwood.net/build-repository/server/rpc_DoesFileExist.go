package main

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/go-easyops/utils"
)

func (b *BuildRepoServer) DoesFileExist(ctx context.Context, req *pb.GetFileRequest) (*pb.FileExistsInfo, error) {

	filename, err := toLinuxFilename(req.File)

	if err != nil {
		return nil, fmt.Errorf("could not get filename: %v", err)
	}
	res := &pb.FileExistsInfo{
		Exists: utils.FileExists(filename),
	}
	return res, nil

}
