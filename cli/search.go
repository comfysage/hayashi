package cli

import (
	"fmt"

	"github.com/crispybaccoon/hayashi/pkg"
	"github.com/crispybaccoon/hayashi/util"
)

func Search(str string) error {
	pkgs := util.SearchPkgs(str)
	if len(pkgs) == 0 {
		return fmt.Errorf("no packages matching that name were found")
	}

	for _, p := range pkgs {
		path := util.PathPkg(p[0], p[1])
		pkg, err := pkg.GetPkgFromPath(path)
		if err != nil {
			return err
		}

		show(pkg)
	}

	return nil
}
