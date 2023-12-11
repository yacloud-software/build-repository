package archive

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/buildrepo"
	barchive "golang.yacloud.eu/apis/buildrepoarchive"
)

func DoesFileExist(ctx context.Context, req *buildrepo.GetFileRequest) (*barchive.FileExistResponse, error) {
	f := req.File
	key := fmt.Sprintf("%s/%s/%d", f.Repository, f.Branch, f.BuildID)
	path := fmt.Sprintf("artefacts/%s/%s/%d/%s", f.Repository, f.Branch, f.BuildID, f.Filename)
	fer := &barchive.FileExistRequest{
		DomainID: GetDomainID(),
		Key:      key,
		Path:     path,
	}
	fr, err := barchive.GetBuildRepoArchiveClient().DoesFileExist(ctx, fer)
	if *debug {
		if err != nil {
			fmt.Printf("Buildrepoarchive response to exists(%v): error=%s\n", fer, err)
		} else {
			fmt.Printf("Buildrepoarchive response to exists(%v): %v\n", fer, fr.Exists)
		}
	}
	return fr, err
}

























