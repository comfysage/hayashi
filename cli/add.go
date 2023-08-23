package cli

import (
	"github.com/crispybaccoon/hayashi/pkg"
	"github.com/crispybaccoon/hayashi/util"
	"fmt"
)

func add(p pkg.Pkg, force bool) error {
	if util.PkgExists("custom", p.Name) {
		if force {
			err := Remove(p.Name)
			if err != nil {
				return err
			}
			printf(COLOR_YELLOW + "overwritting existing pkg " +
				COLOR_MAGENTA + p.Name + COLOR_YELLOW + "...")
		} else {
		return fmt.Errorf("pkg already exists. try again with --force.")
		}
	}

	return pkg.SavePkg(p)
}

func AddLocal(path string) error {
	printf("adding " + COLOR_MAGENTA + path + COLOR_RESET + " ...")

	p, err := pkg.GetPkgFromPath(path)
	if err != nil {
		return err
	}

	err = add(p, cfg.force)
	if err != nil {
		return err
	}

	return nil
}

func Add(name string) error {
	printf("adding " + COLOR_MAGENTA + name + COLOR_RESET + " ...")

	return add(pkg.NewPkg(name, ""), cfg.force)
}

func AddWithUrl(name string, url string) error {
	printf("adding " + COLOR_MAGENTA + name + COLOR_RESET + " ...")

	return add(pkg.NewPkg(name, url), cfg.force)
}
