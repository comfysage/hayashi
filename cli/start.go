package cli

import (
	"fmt"
	"os"

	"github.com/crispybaccoon/hayashi/pkg"
)

func Start(config pkg.Config) error {

	force := false
	Flags.BoolVar(&force, "force", "Force install or override pkg config")
	help := false
	Flags.BoolVar(&help, "help", "Show help message")
	err := Flags.Parse(os.Args[1:])
	if err != nil {
		Err(err)
	}

	if help {
		Help()
	}

	cmd := Flags.Arg(0)
	args := Flags.Args()

	if len(args) == 0 {
		Help()
	}

	switch cmd {

	// .. help
	case "help":
		Help()
		break

	case "list":
		err := List()
		Err(err)
		break

	// .. pkg <> <>
	case "pkg":
		args = args[2:]
		if len(args) < 1 {
			os.Exit(1)
		}

		argv := Flags.Arg(1)
		switch argv {

		// .. pkg add <>
		case "add":
			var err error
			if len(args[1]) > 0 {
				err = AddWithUrl(args[0], args[1], force)
			} else {
				err = Add(args[0], force)
			}
			Err(err)
			break

		// .. pkg remove <>
		case "remove":
			err := Remove(args[0])
			Err(err)
			break
		}
		break

		// .. config <>
	case "config":
		argv := Flags.Arg(1)
		switch argv {
		// .. config init
		case "init":
			err := Init()
			Err(err)
			break
		// .. config create
		case "create":
			err := Create()
			Err(err)
			break
		}
		break

	// .. show <...>
	case "show":
		args = args[1:]
		if len(args) < 0 {
			Err(fmt.Errorf("not enough arguments"))
		}
		for _, s := range args {
			err = Show(s)
			Err(err)
		}
		break

	// .. add <...>
	case "add":
		args = args[1:]
		if len(args) < 0 {
			Err(fmt.Errorf("not enough arguments"))
		}
		for _, s := range args {
			err := Install(s, force, config.DeepClone)
			Err(err)
		}
		break
	// .. update <...>
	case "update":
		args = args[1:]
		if len(args) < 0 {
			Err(fmt.Errorf("not enough arguments"))
		}
		for _, s := range args {
			err := Update(s, force, config.DeepClone)
			Err(err)
		}
		break
	// .. remove <...>
	case "remove":
		args = args[1:]
		if len(args) < 0 {
			Err(fmt.Errorf("not enough arguments"))
		}
		for _, s := range args {
			err := Uninstall(s)
			Err(err)
		}
		break
	default:
		Help()
	}

	return nil
}
