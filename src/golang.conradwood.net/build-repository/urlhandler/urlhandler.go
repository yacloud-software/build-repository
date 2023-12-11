package urlhandler

// urls like so buildrepo://server/reponame/branch/version/path

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/go-easyops/utils"
	//	"io"
)

type DownloadReader interface {
	Read([]byte) (int, error)
}
type downloader struct {
	parsedurl *buildrepourl
	stream    buildrepo.BuildRepoManager_GetFileAsStreamClient
	read_chan chan *readdata
}
type readdata struct {
	block *buildrepo.FileBlock
	err   error
}

func NewDownloadStreamForURL(ctx context.Context, url string) (DownloadReader, error) {
	res := &downloader{}
	br, err := parse(url)
	if err != nil {
		return nil, err
	}
	res.parsedurl = br
	fmt.Printf("new stream for \"%s\"\n", url)
	fmt.Printf("url: %s\n", res.parsedurl.KeyString())
	bs, err := get_build_repo_by_hostname(res.parsedurl.Host)
	if err != nil {
		return nil, err
	}
	gfr := res.parsedurl.ToGetFileRequest()
	// does it exist??
	er, err := bs.DoesFileExist(ctx, gfr)
	if err != nil {
		return nil, err
	}
	if !er.Exists {
		return nil, fmt.Errorf("file %s not found", url)
	}
	// yes->create stream
	res.stream, err = bs.GetFileAsStream(ctx, gfr)
	if err != nil {
		return nil, err
	}
	res.read_chan = make(chan *readdata, 3)
	go res.download_loop()
	return res, nil
}
func (d *downloader) Read(buf []byte) (int, error) {
	rd := <-d.read_chan
	n := 0
	if rd.block != nil {
		n = int(rd.block.Size)
		for i := 0; i < n; i++ {
			buf[i] = rd.block.Data[i]
		}
	}
	return n, rd.err
}
func (d *downloader) download_loop() {
	p := utils.ProgressReporter{Prefix: "Receiver"}
	for {
		p.Print()
		dt, err := d.stream.Recv()
		p.Add(1)
		rd := &readdata{block: dt, err: err}
		d.read_chan <- rd
		if err != nil {
			break
		}
	}
	fmt.Printf("download_loop completed\n")
}













