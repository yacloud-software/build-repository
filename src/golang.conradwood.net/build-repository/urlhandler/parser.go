package urlhandler

import (
	"fmt"
	"golang.conradwood.net/apis/buildrepo"
	"strconv"
	"strings"
)

const (
	URL_PROTOCOL = "buildrepo://"
)

type buildrepourl struct {
	url     string
	Host    string
	Repo    string
	Branch  string
	Version uint64
	Path    string
}

func parse(url string) (*buildrepourl, error) {
	if !strings.HasPrefix(url, URL_PROTOCOL) {
		return nil, fmt.Errorf("invalid url (%s), has no buildrepo:// prefix", url)
	}
	res := &buildrepourl{url: url}
	url = strings.TrimPrefix(url, URL_PROTOCOL)
	parts := strings.SplitN(url, "/", 5)
	/*
		for i, p := range parts {
			fmt.Printf("part %d: \"%s\"\n", i, p)
		}
	*/
	if len(parts) < 5 {
		return nil, fmt.Errorf("invalid url (%s), missing components (%d)", url, len(parts))
	}
	res.Host = parts[0]
	res.Repo = parts[1]
	res.Branch = parts[2]
	v, err := strconv.ParseUint(parts[3], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid url (%s), invalid version (%s) %s", url, parts[3], err)
	}
	res.Version = uint64(v)
	res.Path = parts[4]
	return res, nil
}
func (b *buildrepourl) KeyString() string {
	return fmt.Sprintf("host=%s, repo=%s, branch=%s, version=%d, path=%s", b.Host, b.Repo, b.Branch, b.Version, b.Path)
}
func (b *buildrepourl) ToGetFileRequest() *buildrepo.GetFileRequest {
	res := &buildrepo.GetFileRequest{
		Blocksize: 4096,
		File: &buildrepo.File{
			Repository: b.Repo,
			Branch:     b.Branch,
			BuildID:    b.Version,
			Filename:   b.Path,
		},
	}
	return res
}













































































