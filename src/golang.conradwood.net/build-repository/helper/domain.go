package helper

import (
	"flag"
	pb "golang.conradwood.net/apis/buildrepo"
)

var (
	default_domain = flag.String("default_domain", "", "default domain for this build repository")
)

func GetDefaultDomain() string {
	return *default_domain
}
func GetDomainForRepo(req *pb.RepoEntry) string {
	return *default_domain
}

















































