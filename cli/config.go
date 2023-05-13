package cli

import (
	"fmt"
	"io"
	os_exec "os/exec"

	"github.com/crispybaccoon/hayashi/pkg"
	"github.com/crispybaccoon/hayashi/util"
	"github.com/crispybaccoon/hayashi/exec"
)

func AddInstalled(p pkg.Pkg) error {
	store, err := pkg.GetStoreFile()
	if err != nil {
		return err
	}

	err = store.AddInstalled(p)
	if err != nil {
		return err
	}
	return nil
}

func RemoveInstalled(p pkg.Pkg) error {
	store, err := pkg.GetStoreFile()
	if err != nil {
		return err
	}

	err = store.RemoveInstalled(p)
	if err != nil {
		return err
	}
	return nil
}

type cliConfig struct {
	config *pkg.Config
	force  bool
	local  bool
}

var cfg cliConfig

func Read() pkg.Config {

	config, _ := pkg.GetConfig()

	return config
}

func createPack() error {
	if _, err := util.Mkdir(util.PACK_ROOT); err != nil {
		return err
	}

	if _, err := util.Mkdir(util.PathPackDir("man")); err != nil {
		return err
	}

	if _, err := util.Mkdir(util.PathPackDir("share")); err != nil {
		return err
	}

	if !util.PathExists(util.PathPackDir("share/man")) {
		if err := exec.RunInPack("share", []string{"ln -s ../man man"}); err != nil {
			return err
		}
	}

	return nil
}

func Init() error {
	_, err := util.Mkdir(util.HAYASHI_ROOT)
	if err != nil {
		return err
	}
	_, err = util.Mkdir(util.PKG_ROOT)
	if err != nil {
		return err
	}
	_, err = util.Mkdir(util.REPO_ROOT)
	if err != nil {
		return err
	}

	err = createPack()
	if err != nil {
		return err
	}

	_, err = util.Mkdir(util.PathCl("custom"))
	if err != nil {
		return err
	}

	if !util.PkgExists("custom", "core") {
		printf("fetching " + COLOR_MAGENTA + "core" + COLOR_RESET + " from " + COLOR_YELLOW + "https://github.com/crispybaccoon/hayashi" + COLOR_RESET + " ...")
		p := util.PathPkg("custom", "core")

		cmd := os_exec.Command("curl", "-fsSL", "https://raw.githubusercontent.com/CrispyBaccoon/hayashi/mega/core.yaml", "-o", p)
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
		p, err := pkg.GetPkgFromPath(util.PathPkg("custom", "core"))
		if err != nil {
			return err
		}
		err = startInstall(p, true, cfg.config.DeepClone)
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
