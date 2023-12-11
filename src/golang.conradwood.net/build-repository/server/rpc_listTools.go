package main

import (
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/apis/common"
	"golang.org/x/net/context"
)

// ListTools : list client tools that are regarded as part of the toolbox
func (brs *BuildRepoServer) ListTools(ctx context.Context, req *common.Void) (*pb.ListToolsResponse, error) {

	tools := []*pb.Tool{
		{Repository: "logservice", Filename: "logservice-client"},
		{Repository: "logservice", Filename: "glog"},
		{Repository: "autodeployer", Filename: "deploymonkey-client"},
		{Repository: "autodeployer", Filename: "autodeployer-client"},
		{Repository: "registry", Filename: "registry-client"},
		{Repository: "dc-tools", Filename: "vmmanager-client"},
		{Repository: "dc-tools", Filename: "gmodinfo"},
		{Repository: "auth-service", Filename: "auth-service-client"},
		{Repository: "auth-service", Filename: "login"},
		{Repository: "build-repository", Filename: "build-repo-client"},
		{Repository: "protoc-tools", Filename: "protoc-gen-skel"},
		{Repository: "ota-module", Filename: "otamodule-client"},
	}

	res := pb.ListToolsResponse{Tools: tools}

	return &res, nil
}












































































