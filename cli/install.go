package cli

import (
	"crispybaccoon/hayashi/pkg"
	"crispybaccoon/hayashi/util"
	"fmt"
	"io"
	"os"
	"os/exec"
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

	cmd := exec.Command("git", "clone", "--depth", "1", p.Url, util.PathRepo(p.Name))
	stdout, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	s, err := io.ReadAll(stdout)
	if err != nil {
		return err
	}
	printf(string(s))

	err = cmd.Wait()
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

func Install(name string, force bool, deep_clone bool) error {

	p, err := pkg.GetPkg(name)
	if err != nil {
		return err
	}

	err = clone_pkg(p, force, deep_clone)
	if err != nil {
		return err
	}

	return install(p, force, deep_clone)
}
