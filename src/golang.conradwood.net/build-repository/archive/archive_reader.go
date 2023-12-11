package archive

import (
	"context"
	"fmt"
	"golang.conradwood.net/apis/buildrepo"
	barchive "golang.yacloud.eu/apis/buildrepoarchive"
	"io"
	"strings"
)

type Retriever struct {
	DomainID string
	Key      string
}

/*
retrieve from objectarchive. this _must_ be externally synchronised
*/
func Retrieve(ctx context.Context, file *buildrepo.File, targetdir string) error {
	domain := GetDomainID()
	key := fmt.Sprintf("%s/%s/%d", file.Repository, file.Branch, file.BuildID)
	retriever := &Retriever{DomainID: domain, Key: key}
	err := retriever.Retrieve(ctx, targetdir)
	if err != nil {
		retriever.Printf("retriever failed: %s\n", err)
	}
	fmt.Printf("Set marker to %s\n", targetdir)
	return err
}

func (r *Retriever) Retrieve(ctx context.Context, targetdir string) error {
	fmt.Printf("Retrieving key \"%s\" from domain \"%s\" to %s from buildrepoarchive\n", r.Key, r.DomainID, targetdir)
	dr := &barchive.DownloadRequest{
		DomainID: r.DomainID,
		Key:      r.Key,
	}
	use_lock_key := dr.DomainID + "_" + dr.Key
	lock_key(use_lock_key)
	defer unlock_key(use_lock_key)
	srv, err := barchive.GetBuildRepoArchiveClient().Download(ctx, dr)
	if err != nil {
		return err
	}
	targetdir = strings.TrimSuffix(targetdir, "/")
	w := &writer{targetdir: targetdir}
	for {
		dr, err := srv.Recv()
		if dr != nil {
			if dr.Filename != "" {
				err = w.NewFile(dr.Filename)
				if err != nil {
					r.Printf("newfile failed: %s\n", err)
					return err
				}
			}
			err := w.Write(dr.Data)
			if err != nil {
				r.Printf("write failed: %s\n", err)
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	if w != nil {
		err = w.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
func (r *Retriever) Printf(format string, args ...interface{}) {
	if !*debug {
		return
	}
	s := fmt.Sprintf(format, args...)
	fmt.Print(s)
}

// lock a single key (TODO: limit to a single key, it ignores keys and just locks atm)
func lock_key(key string) {
	file_map_lock.Lock()
}

// unlock a single key (TODO: limit to a single key, it ignores keys and just locks atm)
func unlock_key(key string) {
	file_map_lock.Unlock()
}











































