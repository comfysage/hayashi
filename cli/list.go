package cli

import (
	"fmt"
	"strings"

	"github.com/crispybaccoon/hayashi/pkg"
	"github.com/crispybaccoon/hayashi/util"
)

func list(c pkg.StoreFile) error {

	printf("installed packages. " + COLOR_CYAN + util.PathStoreFile())

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
	store, err := pkg.GetStoreFile()
	if err != nil {
		return err
	}

	err = list(store)
	if err != nil {
		return err
	}

	return nil
}
