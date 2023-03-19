package cli

import (
	"fmt"
	"os"

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

	cmd := util.CloneCommand(p.Url, util.PathRepo(p.Name))
	err := runOne(cmd, util.REPO_ROOT)
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

	pwd := util.PathRepo(p.Name)
	err := run(p.Install, pwd)
	if err != nil {
		return err
	}

	return nil
}

func startInstall(p pkg.Pkg, force bool, deep_clone bool) error {
	err := clone_pkg(p, force, deep_clone)
	if err != nil {
		return err
	}

	err = install(p, force, deep_clone)
	if err != nil {
		return err
	}

	err = AddInstalled(p)
	if err != nil {
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
