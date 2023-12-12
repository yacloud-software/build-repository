package archive

import (
	"flag"
)

var (
	domain_id = flag.String("domain_id", "default_buildrepo_service", "domain id for buildrepo")
	debug     = flag.Bool("debug_archive", false, "debug buildrepoarchive")
)

func GetDomainID() string {
	return *domain_id
}
















































































