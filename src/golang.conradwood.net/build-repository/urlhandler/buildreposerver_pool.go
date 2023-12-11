package urlhandler

import (
	"fmt"
	"golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/go-easyops/client"
	"strings"
	"sync"
)

var (
	block sync.Mutex
	cons  = make(map[string]*con)
)

type con struct {
	server buildrepo.BuildRepoManagerClient
}

func get_build_repo_by_hostname(host string) (buildrepo.BuildRepoManagerClient, error) {
	h := strings.ToLower(host)
	block.Lock()
	defer block.Unlock()
	ccon, found := cons[h]
	if found {
		return ccon.server, nil
	}
	ccon = &con{}
	adr := fmt.Sprintf("%s:5005", h)
	c, err := client.ConnectWithIP(adr)
	if err != nil {
		return nil, err
	}
	res := buildrepo.NewBuildRepoManagerClient(c)
	ccon.server = res
	cons[h] = ccon
	return res, nil

}





















