package main

import (
	"github.com/crispybaccoon/hayashi/cli"
	"os"
)

func main() {

	config := cli.Read()

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

	cmd := cli.Flags.Arg(0)
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

		argv := cli.Flags.Arg(1)
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
		argv := cli.Flags.Arg(1)
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
		argv := cli.Flags.Arg(1)
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
		argv := cli.Flags.Arg(1)
		err := cli.Install(argv, force, config.DeepClone)
		cli.Err(err)
		if len(args) > 2 {
			for _, s := range args[2:] {
				err := cli.Install(s, force, config.DeepClone)
				cli.Err(err)
			}
		}
		break
	// .. update <...>
	case "update":
		argv := cli.Flags.Arg(1)
		err := cli.Update(argv, force, config.DeepClone)
		cli.Err(err)
		if len(args) > 2 {
			for _, s := range args[2:] {
				err := cli.Update(s, force, config.DeepClone)
				cli.Err(err)
			}
		}
		break
	// .. remove <...>
	case "remove":
		argv := cli.Flags.Arg(1)
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
