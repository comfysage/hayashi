package cli

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/crispybaccoon/hayashi/pkg"
	"github.com/crispybaccoon/hayashi/util"
)

func AddInstalled(p pkg.Pkg) error {
	config, err := pkg.GetConfig()
	if err != nil {
		return err
	}

	err = config.AddInstalled(p)
	if err != nil {
		return err
	}
	return nil
}

func Read() pkg.Config {

	config, err := pkg.GetConfig()
	if err != nil {
		Err(fmt.Errorf(COLOR_RED + "could not read config file " + COLOR_CYAN +
			util.PathConfig() + COLOR_RED + ". try running hayashi config create."))
	}

	return config
}

// mkdir `path` if directory does not already exists.
// if path is created return true else return false
func mkdir(path string) (bool, error) {
	exists := util.PathExists(path)
	if exists {
		return false, nil
	}
	err := os.MkdirAll(path, os.ModeAppend.Perm())
	if err != nil {
		return false, err
	}
	return true, nil
}

func Init() error {
	_, err := mkdir(util.HAYASHI_ROOT)
	if err != nil {
		return err
	}
	_, err = mkdir(util.PKG_ROOT)
	if err != nil {
		return err
	}
	_, err = mkdir(util.REPO_ROOT)
	if err != nil {
		return err
	}
	_, err = mkdir(util.PathCl("custom"))
	if err != nil {
		return err
	}

	if !util.PkgExists("custom", "core") {
		printf("fetching " + COLOR_MAGENTA + "core" + COLOR_RESET + " from " + COLOR_YELLOW + "https://github.com/crispybaccoon/hayashi" + COLOR_RESET + " ...")
		p := util.PathPkg("custom", "core")

		cmd := exec.Command("curl", "-L", "https://raw.githubusercontent.com/CrispyBaccoon/hayashi/mega/core.yaml", "-o", p)
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
	}

	if !util.PathExists(util.PathCl("core")) {
		c := Read()
		err = InstallLocal(util.PathPkg("custom", "core"), true, c.DeepClone)
		if err != nil {
			return err
		}
	}

	return nil
}

func Create() error {

	if util.PathExists(util.PathConfig()) {
		return fmt.Errorf(COLOR_RED + "config file " + COLOR_CYAN +
			util.PathConfig() + COLOR_RED + " already exists.")
	}
	printf("creating config file at " + COLOR_CYAN +
		util.PathConfig() + COLOR_RESET + " ...")
	err := pkg.SaveConfig(pkg.DefaultConfig())
	if err != nil {
		return err
	}

	return nil
}
