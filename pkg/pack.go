package pkg

import (
	"fmt"
	"io"
	"os"

	"github.com/crispybaccoon/hayashi/util"
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
		if len(p.Path) == 0 || len(p.Type) == 0 {
			continue
		}

		src := util.PathRepoFile(pkg.Name, p.Path)
		dst := util.PathPackFile(p.Type, p.Path)

		sourceFileStat, err := os.Stat(src)
		if err != nil {
			return err
		}

		if !sourceFileStat.Mode().IsRegular() {
			return fmt.Errorf("%s is not a regular file", src)
		}

		source, err := os.Open(src)
		if err != nil {
			return err
		}
		defer source.Close()

		_, err = util.Mkdir(util.PathPackDir(p.Type))
		if err != nil {
			return err
		}

		destination, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer destination.Close()
		_, err = io.Copy(destination, source)
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
