package main

import (
	"context"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/apis/common"
	//"golang.conradwood.net/build-repository/archive"
	//"golang.conradwood.net/go-easyops/utils"
	"golang.conradwood.net/build-repository/helper"
)

func (b *BuildRepoServer) GetManagerInfo(ctx context.Context, req *common.Void) (*pb.ManagerInfo, error) {
	res := &pb.ManagerInfo{
		Domain: helper.GetDefaultDomain(),
	}
	return res, nil
}










































