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

func Add(name string, force bool) error {
	printf("adding " + COLOR_MAGENTA + name + COLOR_RESET + " ...")

	return add(pkg.Pkg{Name: name}, force)
}

func AddWithUrl(name string, url string, force bool) error {
	printf("adding " + COLOR_MAGENTA + name + COLOR_RESET + " ...")

	return add(pkg.Pkg{Name: name, Url: url}, force)
}
