package cli

import (
	"fmt"
	"os"

	"github.com/crispybaccoon/hayashi/pkg"
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
	err := Flags.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	if help {
		Help(0)
	}

	cmd := Flags.Arg(0)
	args := Flags.Args()

	if len(args) == 0 {
		Help(1)
	}

	switch cmd {

	// .. help
	case "help":
		if len(args) > 1 {
			err := GetHelp(args[1:], Flags.AllFlags())
			return err
		}
		Help(0)

		return nil

	case "list":
		err := List()
		return err

	// .. pkg <> <>
	case "pkg":
		argv := Flags.Arg(1)
		switch argv {

		// .. pkg add <>
		case "add":
			if len(args) < 3 {
				return fmt.Errorf("not enough arguments to call")
			}
			args = args[2:]
			var err error
			// .. pkg add <pkg file>
			if cfg.local {
				err = AddLocal(args[0])
				return err
			}
			// .. pkg add <...pkg info>
			if len(args) > 1 {
				err = AddWithUrl(args[0], args[1])
			} else {
				err = Add(args[0])
			}
			return err

		// .. pkg remove <>
		case "remove":
			if len(args) < 3 {
				return fmt.Errorf("not enough arguments to call")
			}
			args = args[2:]
			err := Remove(args[0])
			return err

		// .. pkg show <>
		case "show":
			if len(args) < 3 {
				return fmt.Errorf("not enough arguments to call")
			}
			args = args[2:]
			err := ShowPkg(args[0])
			return err
		}
		return fmt.Errorf("no matching subcommand for pkg command")

		// .. config <>
	case "config":
		argv := Flags.Arg(1)
		switch argv {
		// .. config init
		case "init":
			err := Init()
			return err
		// .. config create
		case "create":
			err := Create()
			return err
		}
		return fmt.Errorf("no matching subcommand for config command")

	// .. show <...>
	case "show":
		args = args[1:]
		if len(args) < 1 {
			Err(fmt.Errorf("not enough arguments"))
		}
		for _, s := range args {
			err = Show(s)
			return err
		}
		return nil

	// .. add <...>
	case "add":
		args = args[1:]
		if len(args) < 1 {
			Err(fmt.Errorf("not enough arguments"))
		}
		var err error
		for _, s := range args {
			err = Install(s)
			return err
		}
		return nil
	// .. update <...>
	case "update":
		args = args[1:]
		if len(args) < 1 {
			Err(fmt.Errorf("not enough arguments"))
		}
		for _, s := range args {
			err := Update(s)
			return err
		}
		return nil
	// .. remove <...>
	case "remove":
		args = args[1:]
		if len(args) < 1 {
			Err(fmt.Errorf("not enough arguments"))
		}
		for _, s := range args {
			err := Uninstall(s)
			return err
		}
		return nil
	default:
		Help(1)
	}

	return nil
}
