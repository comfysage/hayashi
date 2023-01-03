package main

import (
	"github.com/crispybaccoon/hayashi/cli"
	"os"
)

func main() {

	force := false
	cli.Flags.BoolVar(&force, "force", "Force install or override pkg config")
	help := false
	cli.Flags.BoolVar(&help, "help", "Show help message")
	err := cli.Flags.Parse(os.Args[1:])
	if err != nil {
		cli.Err(err)
	}

	if help {
		cli.Help()
	}

	cmd, argv := cli.Flags.Arg(0), cli.Flags.Arg(1)
	args := cli.Flags.Args()

	if len(args) == 0 {
		cli.Help()
	}

	switch cmd {

	// .. pkg <> <>
	case "pkg":
		args = args[2:]
		if len(args) < 1 {
			os.Exit(1)
		}

		switch argv {

		// .. pkg add <>
		case "add":
			var err error
			if len(args[1]) > 0 {
				err = cli.AddWithUrl(args[0], args[1], force)
			} else {
				err = cli.Add(args[0], force)
			}
			cli.Err(err)
			break

		// .. pkg remove <>
		case "remove":
			err := cli.Remove(args[0])
			cli.Err(err)
			break
		}
		break

		// .. config <>
	case "config":
		switch argv {
		// .. config init
		case "init":
			err := cli.Init()
			cli.Err(err)
			break
		// .. config create
		case "create":
			err := cli.Create()
			cli.Err(err)
			break
		}
		break

	// .. show <...>
	case "show":
		err := cli.Show(argv)
		cli.Err(err)
		if len(args) > 2 {
			for _, s := range args[2:] {
				err = cli.Show(s)
				cli.Err(err)
			}
		}
		break

	// .. add <...>
	case "add":
		err := cli.Install(argv, force, false)
		cli.Err(err)
		if len(args) > 2 {
			for _, s := range args[2:] {
				err := cli.Install(s, force, false)
				cli.Err(err)
			}
		}
		break
	// .. update <...>
	case "update":
		err := cli.Update(argv, force, false)
		cli.Err(err)
		if len(args) > 2 {
			for _, s := range args[2:] {
				err := cli.Update(s, force, false)
				cli.Err(err)
			}
		}
		break
	// .. remove <...>
	case "remove":
		err := cli.Uninstall(argv)
		cli.Err(err)
		if len(args) > 2 {
			for _, s := range args[2:] {
				err := cli.Uninstall(s)
				cli.Err(err)
			}
		}
		break
	default:
		cli.Help()
	}

	os.Exit(0)
}
