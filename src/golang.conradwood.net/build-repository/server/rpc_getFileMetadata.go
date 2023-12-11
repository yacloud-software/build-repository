package main

import (
	"fmt"
	"os"

	pb "golang.conradwood.net/apis/buildrepo"
	"golang.org/x/net/context"
)

// GetFileMetaData : get information about a file, e.g. size
func (brs *BuildRepoServer) GetFileMetaData(ctx context.Context, req *pb.GetMetaRequest) (*pb.GetMetaResponse, error) {
	filename, err := toLinuxFilename(req.File)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if *debug {
		fmt.Printf("GetFileMetaData: The file is %d bytes long\n", fi.Size())
	}

	resp := &pb.GetMetaResponse{
		File: req.File,
		Size: uint64(fi.Size()),
	}
	return resp, nil
}




































































