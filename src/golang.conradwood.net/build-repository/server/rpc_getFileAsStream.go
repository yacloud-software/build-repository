package main

import (
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"io"
	"os"
)

// GetFileAsStream : stream the requested file to the client
func (brs *BuildRepoServer) GetFileAsStream(req *pb.GetFileRequest, s pb.BuildRepoManager_GetFileAsStreamServer) error {
	if req.Blocksize == 0 {
		return fmt.Errorf("Invalid blocksize 0")
	}
	var filename string
	var err error

	filename, err = toLinuxFilename(req.File)

	if err != nil {
		return fmt.Errorf("could not get filename: %v", err)
	}

	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer f.Close()
	data := make([]byte, req.Blocksize)

	for {
		size, err := f.Read(data)
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return fmt.Errorf("could not read file: %v", err)
		}

		err = s.Send(&pb.FileBlock{Size: uint64(size), Data: data[:size]})
		if err != nil {
			return err
		}
	}
}
