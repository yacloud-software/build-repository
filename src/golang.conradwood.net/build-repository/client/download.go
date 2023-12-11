package main

import (
	"fmt"
	"golang.conradwood.net/build-repository/urlhandler"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"path/filepath"
)

func DownloadURL() {
	url := *download
	outfile := "/tmp/" + filepath.Base(url)
	fmt.Printf("Downloading...\"%s\" to \"%s\"\n", url, outfile)
	ctx := authremote.Context()
	stream, err := urlhandler.NewDownloadStreamForURL(ctx, url)
	utils.Bail("failed to open download stream", err)
	f, err := utils.OpenWriteFile(outfile)
	_, err = io.Copy(f, stream)
	utils.Bail("failed to copy", err)
	f.Close()
	fmt.Printf("File saved to %s\n", outfile)
}




































