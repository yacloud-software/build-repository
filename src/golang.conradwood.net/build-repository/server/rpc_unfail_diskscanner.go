package main

import (
	"context"
	"golang.conradwood.net/apis/common"
	"golang.conradwood.net/go-easyops/errors"
)

func (brs *BuildRepoServer) UnfailDiskScanner(ctx context.Context, req *common.Void) (*common.Void, error) {
	if diskScanner == nil {
		return nil, errors.NotFound(ctx, "diskscanner not available")
	}
	diskScanner.Unfail()

	return &common.Void{}, nil
}
































