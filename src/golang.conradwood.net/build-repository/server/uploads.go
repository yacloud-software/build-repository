package main

import (
	pb "golang.conradwood.net/apis/buildrepo"
	"time"
)

type StoreMetaData struct {
	StoreID      string
	StorePath    string
	BuildID      int
	CommitID     string
	Commitmsg    string
	Branch       string
	Repository   string
	UserEmail    string
	BuildDate    uint32
	uploading    uint32
	RepositoryID uint64
	ArtefactID   uint64
}

// UploadMetaData :
type UploadMetaData struct {
	*pb.UploadTokenRequest
	Token string
	// the path under which we store the files
	Storepath string
	Created   time.Time
}

// NewUploadMetaData :
func (brs *BuildRepoServer) NewUploadMetaData(token string, uploadTokenRequest *pb.UploadTokenRequest) {

	brs.cache.SetUpload(
		token,
		&UploadMetaData{
			uploadTokenRequest,
			token,
			"",
			time.Now(),
		},
	)
}




































































