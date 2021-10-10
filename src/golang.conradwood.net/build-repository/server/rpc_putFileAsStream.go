package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"os"
	"time"
)

var (
	auto_create_md5    = flag.Bool("auto_create_md5", true, "automatically create md5 sums for uploaded files")
	upload_throttle_ms = flag.Int("upload_throttle_ms", 0, "each block will be stalled by `milliseconds`")
)

// PutFileAsStream : Receives a file, identifies the location by token,
// which it previously generated, and responded to with a response in an RPC call)
func (brs *BuildRepoServer) PutFileAsStream(s pb.BuildRepoManager_PutFileAsStreamServer) error {
	token := ""
	first := true
	filename := ""
	var file *os.File
	for {
		if *upload_throttle_ms != 0 {
			time.Sleep(time.Duration(*upload_throttle_ms) * time.Millisecond)
		}
		block, err := s.Recv()
		if err == io.EOF {
			if *debug {
				fmt.Printf("EOF received (%s)\n", filename)
			}
			break
		}
		/********************* first block only *****************/
		if first {
			token = block.UploadToken
			if token == "" {
				return fmt.Errorf("token missing")
			}
			meta := brs.cache.GetUpload(token)
			if meta == nil {
				return fmt.Errorf("Invalid token")
			}
			c := brs.cache.GetStored(meta.BuildStoreid)
			meta.Storepath = c.StorePath
			c.uploading++
			defer func() {
				fmt.Printf("Upload complete (%s)\n", meta.Filename)
				c.uploading--
			}()
			filename = fmt.Sprintf("%s/%s", meta.Storepath, meta.Filename)
			if *debug {
				fmt.Printf("Receiving file %s => %s token=%s\n", meta.Filename, filename, token)
			}

			file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
			if err != nil {
				return fmt.Errorf("could not open file %s: %v", filename, err)
			}
			defer file.Close()
			first = false
		}
		/********************* all blocks *****************/
		if err != nil {
			if *debug {
				fmt.Printf("stream.Recv (%s): %s\n", filename, err)
			}
			return fmt.Errorf("stream.Recv failed with %v  %v", s, err)
		}
		_, err = file.Write(block.Data[0:block.GetSize()])

		if err != nil {
			if *debug {
				fmt.Printf("Write failed (%s): %s\n", filename, err)
			}
			return fmt.Errorf("could not write to file %s %v", filename, err)
		}
	}

	if *auto_create_md5 {
		// now do the md5sum for "filename":
		fname := filename + ".md5"
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		m := md5.New()
		if _, err := io.Copy(m, f); err != nil {
			return err
		}
		fmd5 := fmt.Sprintf("%x", m.Sum(nil))
		err = utils.WriteFile(fname, []byte(fmd5))
		if err != nil {
			return err
		}
	}
	return nil
}
