package cli

import (
	"io"
	os_exec "os/exec"

	"github.com/crispybaccoon/hayashi/exec"
	"github.com/crispybaccoon/hayashi/pkg"
	"github.com/crispybaccoon/hayashi/util"
)

func pull_pkg(p pkg.Pkg) error {

	printf("pulling " + COLOR_MAGENTA + p.Name + COLOR_RESET + " from " + COLOR_YELLOW + p.Url + COLOR_RESET + " ...")

	cmd := os_exec.Command("git", "pull")
	cmd.Dir = util.PathRepo(p.Name)
	stdout, err := cmd.StdoutPipe()
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

func update(p pkg.Pkg, force bool, deep_clone bool) error {

	printf("updating " + COLOR_MAGENTA + p.Name + COLOR_RESET + " ...")

	err := pull_pkg(p)
	if err != nil {
		return err
	}

	if len(p.Update) > 0 {
		err = exec.RunInRepo(p.Name, p.Update)
		if err != nil {
			return err
		}
	} else {
		err = Uninstall(p.Name)
		if err != nil {
			return err
		}
		err = install(p, force, deep_clone)
		if err != nil {
			return err
		}
	}

	err = p.CreatePack()
	if err != nil {
		return err
	}

	return nil
}

func Update(name string) error {
	p, err := pkg.GetPkg(name)
	if err != nil {
		return err
	}

	return update(p, cfg.force, cfg.config.DeepClone)
}
