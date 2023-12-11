package diskscanner

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

type BuildDir struct {
	root  string
	repos []*Repo
}

func (b *BuildDir) Size() uint64 {
	res := uint64(0)
	for _, r := range b.repos {
		res = res + r.Size()
	}
	return res
}

// get versions whose branch got more than 3
func (b *BuildDir) Archivable() []*Version {
	var res []*Version
	for _, r := range b.repos {
		min_versions := 3
		if strings.Contains(r.name, "go-easyops") {
			// go-easyops is special, we need to keep those around for longer (modules)
			min_versions = 30
		}
		for _, b := range r.branches {
			if len(b.versions) < min_versions {
				continue
			}
			sort.Slice(b.versions, func(i, j int) bool {
				return b.versions[i].version < b.versions[j].version
			})
			arc := b.versions[:len(b.versions)-min_versions]
			for _, v := range arc {
				res = append(res, v)
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Created().Before(res[j].Created())
	})
	return res
}

/*****************************************************************
* Repo
*****************************************************************/
type Repo struct {
	builddir *BuildDir
	name     string
	branches []*Branch
}

func (r *Repo) Size() uint64 {
	res := uint64(0)

	for _, b := range r.branches {
		res = res + b.Size()
	}
	return res
}

func (r *Repo) Versions() []*Version {
	var res []*Version
	for _, b := range r.branches {
		for _, v := range b.versions {
			res = append(res, v)
		}
	}
	return res
}

/*****************************************************************
* Branch
*****************************************************************/
type Branch struct {
	repo     *Repo
	name     string
	versions []*Version
}

func (b *Branch) Size() uint64 {
	res := uint64(0)
	for _, v := range b.versions {
		if v.deleted {
			continue
		}
		res = res + v.Size()
	}
	return res
}

/*****************************************************************
* Version
*****************************************************************/
// a version is a single version, of a branch in a repository
type Version struct {
	branch  *Branch
	version int
	created time.Time
	size    uint64
	deleted bool
}

func (v *Version) Created() time.Time {
	path := fmt.Sprintf("%s/%s/%s/%d", v.branch.repo.builddir.root, v.branch.repo.name, v.branch.name, v.version)
	fileStat, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Unable to get time: %s\n", err)
		return time.Unix(0, 0)
	}
	return fileStat.ModTime()
}
func (v *Version) Size() uint64 {
	if v.size != 0 {
		return v.size
	}
	path := fmt.Sprintf("%s/%s/%s/%d", v.branch.repo.builddir.root, v.branch.repo.name, v.branch.name, v.version)

	res := disk_size(path)
	v.size = res
	return res
}

func (v *Version) Path() string {
	path := fmt.Sprintf("%s/%s/%s/%d", v.branch.repo.builddir.root, v.branch.repo.name, v.branch.name, v.version)
	return path
}
func disk_size(path string) uint64 {
	res := uint64(0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Failed to read dir: %s\n", err)
		return 0
	}
	//	fmt.Printf("%d files in %s\n", len(files), path)
	for _, f := range files {
		if f.IsDir() {
			res = res + disk_size(path+"/"+f.Name())
		} else if f.Mode().IsRegular() {
			//	fmt.Printf("File: %s\n", f.Name())
			res = res + uint64(f.Size())
		} else {
			fmt.Printf("Weird: %s (%d)\n", f.Name(), f.Mode())
		}
	}
	return res
}












































































