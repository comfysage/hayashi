package cli

import (
	"github.com/crispybaccoon/hayashi/util"
	"fmt"
	"os"
)

func Remove(name string) error {
	if util.PkgExists("custom", name) {
		printf("removing pkg " + COLOR_MAGENTA + name + COLOR_RESET + " ...")
		return	os.RemoveAll(util.PathPkg("custom", name))
	}

	return fmt.Errorf("pkg does not exist")
}
