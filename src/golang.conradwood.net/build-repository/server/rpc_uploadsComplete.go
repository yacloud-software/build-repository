package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	dm "golang.conradwood.net/apis/deployminator"
	dmo "golang.conradwood.net/apis/deploymonkey"
	sb "golang.conradwood.net/apis/slackgateway"
	"golang.conradwood.net/build-repository/globals"
	"golang.conradwood.net/build-repository/helper"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/client"
	cl "golang.conradwood.net/go-easyops/cmdline"
	"golang.conradwood.net/go-easyops/utils"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	slack                   sb.SlackGatewayClient
	use_deploymonkey_binary = flag.Bool("use_deploymonkey_binary", false, "if true, use an external deploymonkey binary to call deploymonkey. usually just do an rpc call")
	deploymonkey            = flag.String("deploymonkey_binary", "/srv/build-repository/artefacts/autodeployer/master/latest/dist/linux/amd64/deploymonkey-client", "location of the binary deploymonkey-client")
)

// UploadsComplete : client claims it's all done, now call any hooks we might find
func (brs *BuildRepoServer) UploadsComplete(ctx context.Context, udr *pb.UploadDoneRequest) (*pb.UploadDoneResponse, error) {
	if *debug {
		fmt.Printf("Upload complete called\n")
	}
	diskScanner.Unclean()
	resp := &pb.UploadDoneResponse{Closed: false, Uploading: 0}

	if udr.BuildStoreid == "" {
		fmt.Println("BUILD STORE ID IS ZERO LENGTH")
		return resp, errors.New("missing build store id")
	}

	// deploy with deploymonkey-client
	store := brs.cache.GetStored(udr.BuildStoreid)
	if store == nil {
		return resp, fmt.Errorf("UNABLE TO RESOLVE BUILD ID: " + udr.BuildStoreid)
	}
	resp.Uploading = store.uploading
	if store.uploading > 0 {
		return resp, nil
	}
	fmt.Printf("Upload of %s completed (active: %d).\n", store.Repository, store.uploading)
	err := saveCommitData(store)
	if err != nil {
		fmt.Printf("Unable to save metadata: %s\n", err)
		return nil, err
	}
	if store.Branch != "master" {
		if *debug {
			fmt.Printf("branch %s is not eligible for automatic deployment. - no further action needed.\n", store.Branch)
		}
		resp.Closed = true
		return resp, nil
	}
	globals.UploadCompleted()
	diskScanner.Trigger()
	_, err = runHook(store)
	if err != nil {
		fmt.Printf("Failed to run hook: %s\n", err)
		return nil, err
	}
	tellAllDeployMonkeys(store)

	fmt.Printf("Deployment completed\n")
	resp.Closed = true
	return resp, nil
}

func getBuildRepoToken() (string, error) {
	b, err := ioutil.ReadFile("/etc/goeasyops/buildrepo.token")
	if err != nil {
		return "", err
	}
	s := string(b)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	return s, nil
}

// call all deploymonkeys for which there is a deployment-yaml file
// e.g. deploy-foo.yaml will invoke deploymonkey-client -registry=foo -configfile=deploy-foo.yaml
func tellAllDeployMonkeys(store *StoreMetaData) {
	fmt.Printf("checking Deploymonkey deployment\n")
	dir := store.StorePath + "/deployment"
	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("Failed to readdir %s: %s\n", dir, err)
		return
	}
	fr := make(map[string]string) // filename, registry
	for _, f := range dirs {
		n := f.Name()
		if n == "deploy.yaml" {
			fr[n] = cl.GetRegistryAddress()
			continue
		}
		if !strings.HasPrefix(n, "deploy-") {
			continue
		}
		if !strings.HasSuffix(n, ".yaml") {
			continue
		}
		f := strings.TrimSuffix(strings.TrimPrefix(n, "deploy-"), ".yaml")
		fr[n] = f
	}
	for k, v := range fr {
		go tellDeployMonkey(store, k, v)
		go tellDeployminator(store, k, v)
	}
}

func tellDeployminator(store *StoreMetaData, filename string, registry string) {
	fmt.Printf("Telling deployminator @ registry %s\n", registry)
	deployFile := fmt.Sprintf("%s/deployment/%s", store.StorePath, filename)
	df, err := utils.ReadFile(deployFile)
	if err != nil {
		fmt.Printf("deployminator: Failed to read file: %s\n", utils.ErrorString(err))
		return
	}
	nbr := &dm.NewBuildRequest{
		DeployFile: df,
		//RepositoryID:  store.RepositoryID,
		Branch:        store.Branch,
		BuildNumber:   uint64(store.BuildID),
		CommitID:      store.CommitID,
		ArtefactName:  store.Repository,
		BuildserverID: helper.GetDefaultDomain(),
		ArtefactID:    store.ArtefactID,
	}
	ctx := authremote.Context()
	dmc, err := GetDeployminatorClient(registry)
	if err != nil {
		fmt.Printf("cannot get deployminator at registry %s: %s\n", registry, err)
		return
	}
	fmt.Printf("Using deployminator at address %s\n", dmc.adr)
	_, err = dmc.dmc.NewBuildAvailable(ctx, nbr)
	if err != nil {
		fmt.Printf("deployminator error: %s\n", utils.ErrorString(err))
		return
	}
	fmt.Printf("Told deployminator (registry=%s) about new build %d in repo %d (artefact %s)\n", registry, nbr.BuildNumber, nbr.ArtefactID, nbr.ArtefactName)
}
func tellDeployMonkey(store *StoreMetaData, filename string, registry string) {
	if *use_deploymonkey_binary {
		// old style
		tellDeployMonkeyUsingExternalClient(store, filename, registry)
		return
	}
	// make a grpc call - the default
	deployFile := fmt.Sprintf("%s/deployment/%s", store.StorePath, filename)
	if !utils.FileExists(deployFile) {
		// no deploy.yaml -> nothing to do
		return
	}
	yaml, err := utils.ReadFile(deployFile)
	if err != nil {
		fmt.Printf("failed to read file: %s\n", err)
		return
	}
	if !strings.Contains(registry, ":") {
		registry = registry + ":5000"
	}
	c := client.ConnectAt(registry, "deploymonkey.DeployMonkey")
	depl := dmo.NewDeployMonkeyClient(c)
	defer c.Close()

	nbar := &dmo.NewBuildAvailableRequest{
		DeployYaml:   yaml,
		ArtefactID:   store.ArtefactID,
		BuildRepoID:  helper.GetDefaultDomain(),
		BuildID:      uint64(store.BuildID),
		CommitID:     store.CommitID,
		Branch:       store.Branch,
		RepositoryID: store.RepositoryID,
	}
	ctx := authremote.Context()
	_, err = depl.NewBuildAvailable(ctx, nbar)
	if err != nil {
		fmt.Printf("failed to tell deploymonkey about new build: %s\n", utils.ErrorString(err))
		return
	}
	fmt.Printf("Told deploymonkey at registry %s about new build %d\n", registry, nbar.BuildID)
}
func tellDeployMonkeyUsingExternalClient(store *StoreMetaData, filename string, registry string) {
	var err error
	fmt.Printf("Sending %s to deploymonkey (repo %s) (registry=%s)\n", filename, store.Repository, registry)
	tok, err := getBuildRepoToken()
	if err != nil {
		fmt.Printf("Could not load deploymonkey token: %s\n", err)
		return
	}
	args := []string{
		"-groupname=testing",
		fmt.Sprintf("-token=%s", tok),
		fmt.Sprintf("-buildid=%d", store.BuildID),
		fmt.Sprintf("-repository=%d", store.RepositoryID),
		fmt.Sprintf("-registry=%s", registry),
	}

	// get deploy yaml file
	deployFile := fmt.Sprintf("%s/deployment/%s", store.StorePath, filename)
	_, err = os.Stat(deployFile)
	if err == nil {
		fmt.Println("found deploy yaml: " + deployFile)
		args = append(
			args,
			fmt.Sprintf("-configfile=%s", deployFile),
		)

		command := strings.Join(args, " ")
		fmt.Printf("Executing: deploymonkey-client %s\n", command)
		out, err := exec.Command(*deploymonkey, args...).CombinedOutput()
		if err != nil {
			fmt.Printf("Command failed - %s: %s\n", command, err)
			x := fmt.Sprintf("Command output: %s", out)
			s := fmt.Sprintf("Buildrepo executed deploymonkey-client binary, and the binary failed: %s", err)
			fmt.Println(x)
			fmt.Println(s)

			go tellUsers(fmt.Sprintf("%s\n%s\n", x, s))
			return
		}
		fmt.Println("Deploymonkey-client executed sucessfully")
	} else {
		fmt.Printf("No deployment descriptor in repo %s: %s does not exist\n", store.Repository, deployFile)
	}
}

// true if it exists
func runHook(store *StoreMetaData) (bool, error) {
	hook := fmt.Sprintf("/srv/build-repository/hooks/%s/master/post-upload", store.Repository)
	_, err := os.Stat(hook)
	if os.IsNotExist(err) {
		return false, nil
	}

	fmt.Println("executing hook:", hook)
	cmd := exec.Command(hook)
	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("USER_EMAIL=%s", store.UserEmail),
		fmt.Sprintf("COMMIT_ID=%s", store.CommitID),
		fmt.Sprintf("COMMIT_MSG=%s", store.Commitmsg),
		fmt.Sprintf("GIT_BRANCH=%s", store.Branch),
		fmt.Sprintf("BUILD_NUMBER=%d", store.BuildID),
		fmt.Sprintf("PROJECT_NAME=%s", store.Repository),
		fmt.Sprintf("REPOSITORY=%s", store.Repository),
		fmt.Sprintf("DIST=%s", store.StorePath),
		fmt.Sprintf("BUILDDIR=%s", store.StorePath),
	)

	out, err := cmd.CombinedOutput()
	if *debug {
		fmt.Printf("hook output: %s\n", out)
	}
	if err != nil {
		return true, err
	}
	return true, nil
}

func tellUsers(msg string) {
	var err error
	if slack == nil {
		slack = sb.NewSlackGatewayClient(client.Connect("slackgateway.SlackGateway"))
	}
	ctx := authremote.Context()
	pm := &sb.PublishMessageRequest{OriginService: "originservicenotfilledinyet",
		Channel: "deployments",
		Text:    msg,
	}
	_, err = slack.PublishMessage(ctx, pm)
	if err != nil {
		fmt.Printf("Failed to post slack message: %s\n", err)
	} else {
		fmt.Printf("Posted slack message: %s\n", msg)
	}
}











































































