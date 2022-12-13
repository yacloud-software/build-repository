.PHONY:	client server
all:	client server
server:
	cd server && go install ${LDFLAGS} build-repository-server.go `ls -1 *.go | grep -v build-repository-server.go`
client:
	cd client && go install ${LDFLAGS} build-repo-client.go `ls -1 *.go|grep -v build-repo-client.go`
