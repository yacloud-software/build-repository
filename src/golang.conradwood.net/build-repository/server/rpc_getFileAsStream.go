package main

import (
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/archive"
	"golang.conradwood.net/go-easyops/utils"
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
		// is it perhaps on archive?
		aerr := archive.Retrieve(s.Context(), req.File, "/srv/build-repository/")
		if aerr != nil {
			fmt.Printf("Failed to retrieve from archive: %s\n", utils.ErrorString(err))
			return fmt.Errorf("could not open file: %v", err)
		}
		f, aerr = os.Open(filename)
		if aerr != nil {
			return fmt.Errorf("could not open file: %v", err)
		}
	}
	defer f.Close()
	if *debug {
		fmt.Printf("Sending file as stream: \"%s\"\n", filename)
	}
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











































































