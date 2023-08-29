package diskscanner

import (
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
)

func checkIntegrity() error {
	err := utils.DirWalk("/srv/artefacts", func(root, rel string) error {
		if !utils.FileExists("/srv/metadata/" + rel) {
			return fmt.Errorf("no metadata found for \"%s\"", rel)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
