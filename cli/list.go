package cli

import (
	"fmt"
	"strings"

	"github.com/crispybaccoon/hayashi/pkg"
	"github.com/crispybaccoon/hayashi/util"
)

func list(c pkg.Config) error {

	printf("installed packages. " + COLOR_CYAN + util.PathConfig())

	for _, s := range c.Installed {
		sp := strings.Split(s, "/")
		if len(sp) < 2 {
			return fmt.Errorf("invalid pkg in config")
		}
		printf(" - " + COLOR_CYAN + sp[0] + "/" + COLOR_MAGENTA + sp[1])
	}

	return nil
}

func List() error {

	config, err := pkg.GetConfig()
	if err != nil {
		return err
	}

	err = list(config)
	if err != nil {
		return err
	}

	return nil
}
