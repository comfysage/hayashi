package pkg

import (
	"fmt"
	"os"

	"github.com/crispybaccoon/hayashi/util"
	"github.com/crispybaccoon/hayashi/exec"
)

type Pack struct {
	Path string `yaml:"path"`
	Type string `yaml:"type"`
}

func (pkg Pkg) CreatePack() error {
	_, err := util.Mkdir(util.PACK_ROOT)
	if err != nil {
		return err
	}

	for _, p := range pkg.Pack {
		err := exec.Pack(pkg.Name, p.Path, p.Type)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pkg Pkg) RemovePack() error {
	for _, p := range pkg.Pack {
		if len(p.Path) == 0 || len(p.Type) == 0 {
			continue
		}
		dst := util.PathPackFile(p.Type, p.Path)

		dstFileStat, err := os.Stat(dst)
		if err != nil {
			return err
		}

		if !dstFileStat.Mode().IsRegular() {
			return fmt.Errorf("%s is not a regular file", dst)
		}

		err = os.Remove(dst)
		if err != nil {
			return err
		}
	}
	return nil
}
