package main

import (
	pb "golang.conradwood.net/apis/buildrepo"
)

func getDomainForRepo(req *pb.RepoEntry) string {
	return *default_domain
}
