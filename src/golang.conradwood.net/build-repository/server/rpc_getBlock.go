package main

import (
	"io"
	"os"

	pb "golang.conradwood.net/apis/buildrepo"
	"golang.org/x/net/context"
)

// GetBlock :
// this one is a bit funny - given a specific file in a repo
// and build, take an offset and size and return the chunk of that
// file.
// (some binaries are capable of being streamed over the air to
// remote IoT Devices).
func (app *BuildRepoServer) GetBlock(ctx context.Context, req *pb.GetBlockRequest) (*pb.GetBlockResponse, error) {
	filename, err := toFilename(req.File)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := make([]byte, req.Size)
	size, err := f.ReadAt(buf, int64(req.Offset))
	if err != io.EOF && err != nil {
		return nil, err
	}

	resp := &pb.GetBlockResponse{
		File:   req.File,
		Offset: req.Offset,
		Size:   uint32(size),
		Data:   buf[:size],
	}
	return resp, nil
}
































