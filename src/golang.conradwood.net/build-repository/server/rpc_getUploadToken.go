package main

import (
	"errors"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/helper"
	"golang.org/x/net/context"
	"os"
	"path/filepath"
	"strings"
)

// GetUploadToken :
// generate a random ID for a given file to be uploaded
// the client basically says: "I got a file for build X,
// please give me temporary upload URL" and then uploads
// to that URL.
// (we don't expose a directory structure to the client,
// because we might store the files elsewhere in future)
func (brs *BuildRepoServer) GetUploadToken(ctx context.Context, pr *pb.UploadTokenRequest) (*pb.UploadTokenResponse, error) {

	token, _ := helper.RandString(CONST_RAND_ID_STRING_LEN)
	res := &pb.UploadTokenResponse{
		Token: token,
	}

	fname := pr.Filename
	fname = filepath.Clean(fname)
	if filepath.IsAbs(fname) {
		return res, errors.New("file must be relative")
	}

	sp := brs.cache.GetStored(pr.BuildStoreid).StorePath
	if !strings.HasPrefix(sp, helper.GetBase()) {
		if *debug {
			fmt.Printf("Base=\"%s\", but token sent was: \"%s\"\n", helper.GetBase(), sp)
		}
		return res, errors.New("storeid is invalid")
	}
	fbase := filepath.Dir(fname)

	absDir := fmt.Sprintf("%s/%s", sp, fbase)
	//fmt.Printf("Filebase: \"%s\" (%s)\n", fbase, absDir)
	err := os.MkdirAll(absDir, 0777)
	if err != nil {
		fmt.Println("Failed to create directory ", absDir, err)
		return res, err
	}

	brs.NewUploadMetaData(res.Token, pr)

	return res, nil
}






































