package main

import (
	"context"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/helper"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"os"
)

func (b *BuildRepoServer) GetRepositoryMeta(ctx context.Context, req *pb.GetRepoMetaRequest) (*pb.RepoMetaInfo, error) {
	if !helper.IsValidName(req.Path) {
		return nil, errors.InvalidArgs(ctx, "not a valid path (%s)", req.Path)
	}
	return loadRepoMeta(ctx, req.Path)
}
func loadRepoMeta(ctx context.Context, repo string) (*pb.RepoMetaInfo, error) {
	filename := fmt.Sprintf("%s/%s/repometa.proto", helper.GetMetadir(), repo)
	if !utils.FileExists(filename) {
		return nil, errors.NotFound(ctx, "\"%s\" not found", filename)
	}
	bp, err := utils.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	res := &pb.RepoMetaInfo{}
	err = utils.UnmarshalBytes(bp, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func saveRepoMeta(ctx context.Context, repo string, meta *pb.RepoMetaInfo) error {
	dir := fmt.Sprintf("%s/%s", helper.GetMetadir(), repo)
	filename := fmt.Sprintf("%s/repometa.proto", dir)
	bs, err := utils.MarshalBytes(meta)
	if err != nil {
		return err
	}
	os.MkdirAll(dir, 0777)
	err = utils.WriteFile(filename, bs)
	if err != nil {
		return err
	}
	return nil
}

func setRepositoryMetaRepositoryID(ctx context.Context, repo string, repoid uint64) error {
	cur, err := loadRepoMeta(ctx, repo)
	if err != nil || cur == nil {
		cur = &pb.RepoMetaInfo{}
	}
	cur.RepositoryID = repoid
	return saveRepoMeta(ctx, repo, cur)

}
































































