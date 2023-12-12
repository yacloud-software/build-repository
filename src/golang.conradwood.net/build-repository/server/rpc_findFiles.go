package main

import (
	"fmt"
	pb "golang.conradwood.net/apis/buildrepo"
	"golang.conradwood.net/build-repository/helper"
	"golang.org/x/net/context"
	"io/ioutil"
)

// ListFiles : list all files for a given build
func (b *BuildRepoServer) FindFiles(ctx context.Context, req *pb.FilePattern) (*pb.FileList, error) {
	repos, err := b.ListRepos(ctx, &pb.ListReposRequest{})
	if err != nil {
		return nil, err
	}
	filenames := []string{req.Pattern}
	// now find the file:
	res := pb.FileList{}
	for _, r := range repos.Entries {
		branch := "master"
		latest, err := helper.GetLatestRepoVersion(r.Name, branch)
		if err != nil {
			fmt.Printf("findfiles: no latest version for %s/%s\n", r.Name, branch)
			// no latest version? bullshit, just continue
			continue
		}
		build := fmt.Sprintf("%d", latest)
		repo := r.Name
		if !helper.IsValidName(repo) {
			return nil, fmt.Errorf("Invalid repo name \"%s\"", repo)
		}
		if !helper.IsValidName(branch) {
			return nil, fmt.Errorf("Invalid branch name \"%s\"", branch)
		}
		if !helper.IsValidName(build) {
			return nil, fmt.Errorf("Invalid build name \"%s\"", build)
		}
		if *debug {
			fmt.Printf("Listing versions for repo %s and branch %s and build %s\n", repo, branch, build)
		}
		repodir := fmt.Sprintf("%s/%s/%s/%s", helper.GetBase(), repo, branch, build)
		files, err := findFilesInDir(repodir, filenames)
		if err != nil {
			return nil, fmt.Errorf("findfiles: %s", err)
		}
		for _, f := range files {
			res.Files = append(res.Files, &pb.File{
				Filename:   f,
				BuildID:    latest,
				Repository: repo,
			})
		}

	}
	return &res, nil
}

// return list of matching filenames in directory
func findFilesInDir(dir string, files []string) ([]string, error) {
	var res []string
	fmt.Printf("Searching in dir %s\n", dir)
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, fi := range fis {
		fmt.Printf("File: \"%s\"\n", fi.Name())
		if fi.IsDir() {
			f, err := findFilesInDir(dir+"/"+fi.Name(), files)
			if err != nil {
				return nil, err
			}
			res = append(res, f...)
		} else {
			for _, fm := range files {
				if fi.Name() == fm {
					res = append(res, dir+"/"+fi.Name())
					break
				}
			}
		}

	}
	return res, nil
}















































































