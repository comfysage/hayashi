package cli

import (
	"fmt"
	"os"

	"github.com/crispybaccoon/hayashi/exec"
	"github.com/crispybaccoon/hayashi/pkg"
	"github.com/crispybaccoon/hayashi/util"
)

func clone_pkg(p pkg.Pkg, force bool, deep_clone bool) error {
	if util.PathExists(util.PathRepo(p.Name)) {
		if force {
			err := os.RemoveAll(util.PathRepo(p.Name))
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("repo already cloned. try again with --force.")
		}
	}

	printf("cloning " + COLOR_MAGENTA + p.Name + COLOR_RESET + " from " + COLOR_YELLOW + p.Url + COLOR_RESET + " ...")

	err := exec.Clone(p.Url, p.Name)
	if err != nil {
		return err
	}

	return nil
}

func install(p pkg.Pkg, force bool, deep_clone bool) error {

	printf("building " + COLOR_MAGENTA + p.Name + COLOR_RESET + " at " + COLOR_BLUE + util.PathRepo(p.Name) + COLOR_RESET + " ...")
	/* for _, s := range p.Install {
			printf("  " + COLOR_GREEN + s)
	} */ // TODO: on verbose

	err := exec.RunInRepo(p.Name, p.Install)
	if err != nil {
		return err
	}

	return nil
}

func startInstall(p pkg.Pkg, force bool, deep_clone bool) error {
	if p.Clone {
		if err := clone_pkg(p, force, deep_clone); err != nil {
			return err
		}
	} else {
		if _, err := util.Mkdir(util.PathRepo(p.Name)); err != nil {
			return err
		}
	}

	if err := install(p, force, deep_clone); err != nil {
		return err
	}

	if err := p.CreatePack(); err != nil {
		return err
	}

	if err := AddInstalled(p); err != nil {
		return err
	}

	return nil
}

func Install(name string) error {
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

	err = startInstall(p, cfg.force, cfg.config.DeepClone)
	if err != nil {
		return err
	}

	return nil
}
