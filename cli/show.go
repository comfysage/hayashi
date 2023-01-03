package cli

import (
	"github.com/crispybaccoon/hayashi/pkg"
)

func show(pkg pkg.Pkg) {
	printf(COLOR_MAGENTA + pkg.Name)

	if pkg.Desc != "" {
		printf(pkg.Desc)
	}
	if pkg.Url != "" {
		printf(COLOR_CYAN + "url  " + COLOR_RESET + pkg.Url)
	}
	if len(pkg.Install) > 0 {
		printf(COLOR_CYAN + "bash")
		for _, s := range pkg.Install {
			printf("  " + COLOR_GREEN + s)
		}
	}
}

func Show(name string) error {
	pkg, err := pkg.GetPkg(name)
	if err != nil {
		return err
	}

	show(pkg)

	return nil
}
