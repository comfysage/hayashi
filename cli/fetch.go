package cli

import (
	"github.com/crispybaccoon/hayashi/exec"
	"github.com/crispybaccoon/hayashi/pkg"
)

func fetch_pkg(p pkg.Pkg) error {
	printf("fetching " + COLOR_MAGENTA + p.Name + COLOR_RESET + " from " + COLOR_YELLOW + p.Url + COLOR_RESET + " ...")

	if err := exec.Fetch(p.Name); err != nil {
		return err
	}

	return nil
}

func Fetch(name string) error {
	p, err := pkg.GetPkg(name)
	if err != nil {
		return err
	}

	if err := fetch_pkg(p); err != nil {
		return err
	}

	return nil
}
