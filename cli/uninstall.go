package cli

import (
	"fmt"
	"os"

	"github.com/crispybaccoon/hayashi/pkg"
	"github.com/crispybaccoon/hayashi/util"
)

func remove_clone(p pkg.Pkg) error {
	repoPath := util.PathRepo(p.Name)
	if !util.PathExists(repoPath) {
		return fmt.Errorf("repo `%s` not created.", repoPath)
	}

	printf("removing repo " + COLOR_BLUE + util.PathRepo(p.Name) + COLOR_RESET + " ...")
	err := os.RemoveAll(repoPath)
	if err != nil {
		return err
	}
	return nil
}

func uninstall(p pkg.Pkg) error {
	printf("removing " + COLOR_MAGENTA + p.Name + COLOR_RESET + " at " + COLOR_BLUE + util.PathRepo(p.Name) + COLOR_RESET + " ...")

	pwd := util.PathRepo(p.Name)
	err := run(p.Remove, pwd)
	if err != nil {
		return err
	}

	return nil
}

func startUninstall(p pkg.Pkg) error {
	err := uninstall(p)
	if err != nil {
		return err
	}

	err = remove_clone(p)
	if err != nil {
		return err
	}

	err = p.RemovePack()
	if err != nil {
		return err
	}

	err = RemoveInstalled(p)
	if err != nil {
		return err
	}

	return nil
}

func Uninstall(name string) error {
	var p pkg.Pkg
	var err error

	if cfg.local {
		p, err = pkg.GetPkgFromPath(name)
	} else {
		p, err = pkg.GetPkg(name)
	}
	if err != nil {
		return err
	}

	err = startUninstall(p)
	if err != nil {
		return err
	}

	return nil
}
