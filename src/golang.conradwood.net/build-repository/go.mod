module golang.conradwood.net/build-repository

go 1.21.1

toolchain go1.22.2

require (
	golang.conradwood.net/apis/buildrepo v1.1.1466
	golang.conradwood.net/apis/common v1.1.2976
	golang.conradwood.net/apis/deployminator v1.1.2960
	golang.conradwood.net/apis/deploymonkey v1.1.2963
	golang.conradwood.net/apis/registry v1.1.2963
	golang.conradwood.net/apis/slackgateway v1.1.2960
	golang.conradwood.net/go-easyops v0.1.28926
	golang.org/x/net v0.27.0
	golang.org/x/sys v0.22.0
	golang.yacloud.eu/apis/buildrepoarchive v1.1.2960
	google.golang.org/grpc v1.65.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/grafana/pyroscope-go v1.1.1 // indirect
	github.com/grafana/pyroscope-go/godeltaprof v0.1.7 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/client_golang v1.19.1 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	golang.conradwood.net/apis/auth v1.1.2976 // indirect
	golang.conradwood.net/apis/autodeployer v1.1.2963 // indirect
	golang.conradwood.net/apis/echoservice v1.1.2963 // indirect
	golang.conradwood.net/apis/errorlogger v1.1.2963 // indirect
	golang.conradwood.net/apis/framework v1.1.2963 // indirect
	golang.conradwood.net/apis/goeasyops v1.1.2976 // indirect
	golang.conradwood.net/apis/grafanadata v1.1.2963 // indirect
	golang.conradwood.net/apis/h2gproxy v1.1.2960 // indirect
	golang.conradwood.net/apis/objectstore v1.1.2963 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.yacloud.eu/apis/autodeployer2 v1.1.2963 // indirect
	golang.yacloud.eu/apis/fscache v1.1.2963 // indirect
	golang.yacloud.eu/apis/session v1.1.2976 // indirect
	golang.yacloud.eu/apis/unixipc v1.1.2963 // indirect
	golang.yacloud.eu/apis/urlcacher v1.1.2963 // indirect
	golang.yacloud.eu/unixipc v0.1.26852 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace golang.conradwood.net/apis/buildrepo => ../../golang.conradwood.net/apis/buildrepo
