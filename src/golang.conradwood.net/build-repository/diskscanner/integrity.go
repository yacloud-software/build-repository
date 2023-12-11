package diskscanner

import (
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"io/ioutil"
	"os"
)

func checkIntegrity() error {
	if *debug {
		fmt.Printf("Checking repository integrity...\n")
	}
	err := RepoWalk("/srv/build-repository/artefacts", func(root, rel string) error {
		if !utils.FileExists("/srv/build-repository/metadata/" + rel) {
			return fmt.Errorf("no metadata found for \"%s\"", rel)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func RepoWalk(dir string, f func(root, rel string) error) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	// repos
	for _, repo := range files {
		r_filename := repo.Name()
		if do_skip_repowalk(dir + "/" + r_filename) {
			continue
		}

		files2, err := ioutil.ReadDir(dir + "/" + r_filename)
		if err != nil {
			return err
		}
		for _, branch := range files2 {
			b_filename := r_filename + "/" + branch.Name()
			if do_skip_repowalk(dir + "/" + b_filename) {
				continue
			}

			files3, err := ioutil.ReadDir(dir + "/" + b_filename)
			if err != nil {
				return err
			}
			for _, version := range files3 {
				v_filename := b_filename + "/" + version.Name()
				if do_skip_repowalk(dir + "/" + v_filename) {
					continue
				}
				//fmt.Printf("Filename: %s\n", v_filename)
				err = f(dir, v_filename)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func do_skip_repowalk(filename string) bool {
	fileInfo, err := os.Lstat(filename)
	if err != nil {
		fmt.Printf("[diskscanner] stat() of file \"%s\" failed: %s\n", filename, err)
		return false
	}
	if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
		//symlink
		return true
	}
	if fileInfo.Mode()&os.ModeDir != os.ModeDir {
		//not a dir
		return true
	}
	return false
}











