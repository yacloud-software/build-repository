package main

import (
	"fmt"
	dm "golang.conradwood.net/apis/deployminator"
	reg "golang.conradwood.net/apis/registry"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/client"
	"google.golang.org/grpc"
	"strings"
	"sync"
	"time"
)

var (
	regclients = make(map[string]reg.RegistryClient)
	dmclients  = make(map[string]*DMC)
	dmlock     sync.Mutex
	reglock    sync.Mutex
)

type DMC struct {
	dmc        dm.DeployminatorClient
	con        *grpc.ClientConn
	adr        string
	lastlookup time.Time
}

func GetDeployminatorClient(registry string) (*DMC, error) {
	regadr := registry
	if strings.HasSuffix(regadr, ":5000") {
		regadr = strings.TrimSuffix(regadr, ":5000")
	}

	regadr = regadr + ":5001"

	dmlock.Lock()
	defer dmlock.Unlock()
	dmc := dmclients[regadr]

	rc, err := getRegistryAtAddress(regadr)
	if err != nil {
		return nil, err
	}
	if dmc != nil {
		if time.Since(dmc.lastlookup) < time.Duration(15)*time.Second {
			return dmc, nil
		}
	}

	ips, err := ipsForClient(rc, "deployminator.Deployminator")
	if err != nil {
		return nil, err
	}

	if dmc != nil {
		// check if ip is still valid, if so re-use dmc
		for _, i := range ips {
			if dmc.adr == i {
				dmc.lastlookup = time.Now()
				return dmc, nil
			}
		}
		// it is no longer valid, close connection
		if dmc.con != nil {
			dmc.con.Close()
			dmc.con = nil
			dmc.dmc = nil
		}
	}

	dmc = &DMC{}
	dmc.adr = ips[0]
	con, err := client.ConnectWithIP(dmc.adr)
	if err != nil {
		return nil, err
	}
	dmc.con = con
	dmc.lastlookup = time.Now()
	dmc.dmc = dm.NewDeployminatorClient(dmc.con)
	dmclients[regadr] = dmc
	return dmc, nil
}
func getRegistryAtAddress(regadr string) (reg.RegistryClient, error) {
	rc := regclients[regadr]
	if rc != nil {
		return rc, nil
	}
	reglock.Lock()
	defer reglock.Unlock()
	rc = regclients[regadr]
	if rc != nil {
		return rc, nil
	}

	fmt.Printf("deployminator - Connecting to registry @%s...\n", regadr)
	con, err := client.ConnectWithIP(regadr)
	if err != nil {
		return nil, err
	}
	rc = reg.NewRegistryClient(con)
	regclients[regadr] = rc
	return rc, nil

}

func ipsForClient(rc reg.RegistryClient, service string) ([]string, error) {
	nbr := &reg.V2GetTargetRequest{ApiType: reg.Apitype_grpc, ServiceName: []string{service}}
	ctx := authremote.Context()
	//	fmt.Printf("Querying registry at %s for deployminator..\n", regadr)
	nb, err := rc.V2GetTarget(ctx, nbr)
	if err != nil {
		return nil, err
	}
	if len(nb.Targets) == 0 {
		return nil, fmt.Errorf("no service \"%s\" found", service)
	}
	var res []string
	for _, t := range nb.Targets {
		res = append(res, fmt.Sprintf("%s:%d", t.IP, t.Port))
	}
	return res, nil

}






































