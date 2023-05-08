package cli

import (
	"github.com/crispybaccoon/hayashi/exec"
	"github.com/crispybaccoon/hayashi/pkg"
)

func changelog(p pkg.Pkg) error {
	return exec.Changelog(p.Name)
}

func Changelog(name string) error {
	p, err := pkg.GetPkg(name)
	if err != nil {
		return err
	}

	return changelog(p)
}
