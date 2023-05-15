package cli

import (
	"fmt"
	"os"

	"github.com/crispybaccoon/hayashi/pkg"
	"github.com/crispybaccoon/hayashi/util"
)

func Start(config pkg.Config) error {
	cfg = cliConfig{}
	cfg.config = &config

	cfg.force = false
	Flags.BoolVar(&cfg.force, "force", "Force install or override pkg config")
	cfg.local = false
	Flags.BoolVar(&cfg.local, "local", "Local install or add local pkg config")
	help := false
	Flags.BoolVar(&help, "help", "Show help message")
	print_prefix := false
	Flags.BoolVar(&print_prefix, "prefix", "Print prefix")
	err := Flags.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	if help {
		Help(0)
	}

	if print_prefix {
		fmt.Printf(util.PACK_ROOT)
		os.Exit(0)
	}

	Setup()

	if err := RunArgs(Flags.Args()); err != nil {
		return err
	}

	return nil
}
