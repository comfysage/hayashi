package cli

import (
	"github.com/crispybaccoon/hayashi/pkg"
)

func show(p pkg.Pkg) {
	printf(COLOR_MAGENTA + p.Name)

	if p.Desc != "" {
		printf(p.Desc)
	}
	if p.Url != "" {
		printf(COLOR_CYAN + "url  " + COLOR_RESET + p.Url)
	}
	if len(p.Install) > 0 {
		printf(COLOR_CYAN + "bash")
		for _, s := range p.Install {
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
