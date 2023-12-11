// client create: BuildRepoManagerClient
/*
  Created by /home/cnw/devel/go/yatools/src/golang.yacloud.eu/yatools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.conradwood.net/apis/buildrepo/buildrepo.proto
   gopackage : golang.conradwood.net/apis/buildrepo
   importname: ai_0
   clientfunc: GetBuildRepoManager
   serverfunc: NewBuildRepoManager
   lookupfunc: BuildRepoManagerLookupID
   varname   : client_BuildRepoManagerClient_0
   clientname: BuildRepoManagerClient
   servername: BuildRepoManagerServer
   gsvcname  : buildrepo.BuildRepoManager
   lockname  : lock_BuildRepoManagerClient_0
   activename: active_BuildRepoManagerClient_0
*/

package buildrepo

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_BuildRepoManagerClient_0 sync.Mutex
  client_BuildRepoManagerClient_0 BuildRepoManagerClient
)

func GetBuildRepoManagerClient() BuildRepoManagerClient { 
    if client_BuildRepoManagerClient_0 != nil {
        return client_BuildRepoManagerClient_0
    }

    lock_BuildRepoManagerClient_0.Lock() 
    if client_BuildRepoManagerClient_0 != nil {
       lock_BuildRepoManagerClient_0.Unlock()
       return client_BuildRepoManagerClient_0
    }

    client_BuildRepoManagerClient_0 = NewBuildRepoManagerClient(client.Connect(BuildRepoManagerLookupID()))
    lock_BuildRepoManagerClient_0.Unlock()
    return client_BuildRepoManagerClient_0
}

func BuildRepoManagerLookupID() string { return "buildrepo.BuildRepoManager" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.

func init() {
   client.RegisterDependency("buildrepo.BuildRepoManager")
   AddService("buildrepo.BuildRepoManager")
}


































































