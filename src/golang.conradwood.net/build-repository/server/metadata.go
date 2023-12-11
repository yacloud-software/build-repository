package main

import (
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/helper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

func saveCommitData(smd *StoreMetaData) error {
	if smd.BuildDate == 0 {
		smd.BuildDate = uint32(time.Now().Unix())
	}
	dir := fmt.Sprintf("%s/%s/%s/%d", helper.GetMetadir(), smd.Repository, smd.Branch, smd.BuildID)
	filename := fmt.Sprintf("%s/build.yaml", dir)
	st, err := os.Stat(dir)
	if (err == nil) && (st != nil) {
		return fmt.Errorf("Dir %s already exists. Trying to update an existing build??", dir)
	}
	err = os.MkdirAll(dir, 0700)
	if err != nil {
		fmt.Println("Failed to create directory ", dir, err)
		return err
	}
	fmt.Printf("Saving metadata in %s\n", filename)
	b, err := yaml.Marshal(smd)
	err = ioutil.WriteFile(filename, b, 0600)
	if err != nil {
		return err
	}
	return nil
}

func getMetadata(repo string, branch string, buildid uint64) (*pb.BuildMeta, error) {
	lm, err := loadMetadata(repo, branch, buildid)
	if err != nil {
		return nil, err
	}
	res := &pb.BuildMeta{
		RepositoryID: lm.RepositoryID,
		CommitID:     lm.CommitID,
		Branch:       lm.Branch,
	}
	return res, nil
}

// obsolete - use getmetadata (which returns a proto instead)
func loadMetadata(repo string, branch string, buildid uint64) (*StoreMetaData, error) {
	if !helper.IsValidName(repo) {
		return nil, fmt.Errorf("Invalid repo name \"%s\"", repo)
	}
	if !helper.IsValidName(branch) {
		return nil, fmt.Errorf("Invalid branch name \"%s\"", branch)
	}
	dir := fmt.Sprintf("%s/%s/%s/%d", helper.GetMetadir(), repo, branch, buildid)
	filename := fmt.Sprintf("%s/build.yaml", dir)
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	smd := &StoreMetaData{}
	err = yaml.Unmarshal(b, smd)
	if err != nil {
		return nil, err
	}
	if smd.BuildDate == 0 {
		st, err := os.Stat(filename)
		if err != nil {
			return nil, err
		}
		smd.BuildDate = uint32(st.ModTime().Unix())
	}

	return smd, nil
}


















