package cli

import (
	"fmt"
	"strings"

	"github.com/crispybaccoon/hayashi/doc"
)

type ArgListener struct {
	Argv    []string
	Fn      func([]string) error
	Doc     *doc.Doc
	IsAlias bool
}

var Listeners []ArgListener

func ListenerMan() string {
	listenerDocs := []*doc.Doc{}

	for _, listener := range Listeners {
		if len(listener.Argv) > 1 || listener.IsAlias {
			continue
		}
		listenerDocs = append(listenerDocs, listener.Doc)
	}

	return doc.ManDoc{
		Cmddoc: listenerDocs,
	}.String()
}

func NewListener(usage string, args string, long string, f func([]string) error, argv ...string) {
	tag := strings.Join(argv, "-")
	name := ""
	if len(argv) > 0 {
		name = argv[len(argv)-1]
	}
	usageArray := []string{"yki"}
	usageArray = append(usageArray, argv...)
	if len(args) > 0 {
		usageArray = append(usageArray, args)
	}
	usageString := strings.Join(usageArray, " ")
	listener_doc := doc.Doc{
		Name: tag,
		Man: doc.ManDoc{
			Name:  name,
			Short: usage,
			Long:  long,
			Usage: usageString,
		},
	}

	AddDoc(listener_doc)

	Listeners = append(Listeners,
		ArgListener{
			Argv:    argv,
			Fn:      f,
			Doc:     &listener_doc,
			IsAlias: false,
		})
}

func NewAlias(cmdTag string, argv ...string) {
	var cmd *ArgListener = nil
	for _, l := range Listeners {
		if l.Doc.Name == cmdTag {
			cmd = &l
		}
	}
	Listeners = append(Listeners,
		ArgListener{
			Argv:    argv,
			Fn:      cmd.Fn,
			Doc:     cmd.Doc,
			IsAlias: true,
		})
}

func Setup() {
	NewListener("Show help message", "",
		"", func(args []string) error {
			if len(args) > 0 {
				err := GetHelp(args, Flags.AllFlags())
				return err
			}
			Help(0)
			return nil
		}, "help")
	NewListener("Show environment script", "",
		"", func(s []string) error {
			if err := Env(); err != nil {
				return err
			}
			return nil
		}, "env")
	NewListener("List installed packages", "",
		"", func(args []string) error {
			err := List()
			return err
		}, "list")
	NewListener("Add a pkg configuration", "",
		"", func(args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("not enough arguments to call")
			}
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
		}, "pkg", "add")
	NewListener("Remove a pkg configuration", "",
		"", func(args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("not enough arguments to call")
			}
			err := Remove(args[0])
			return err
		}, "pkg", "remove")
	NewListener("Show a pkg configuration", "",
		"", func(args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("not enough arguments to call")
			}
			err := ShowPkg(args[0])
			return err
		}, "pkg", "show")
	NewListener("", "",
		"", func(s []string) error {
			err := Init()
			return err
		}, "config", "init")
	NewListener("", "",
		"", func(s []string) error {
			err := Create()
			return err
		}, "config", "create")
	NewListener("Show package details", "",
		"", func(args []string) error {
			if len(args) < 1 {
				Err(fmt.Errorf("not enough arguments"))
			}
			for _, s := range args {
				if err := Show(s); err != nil {
					return err
				}
			}
			return nil
		}, "show")
	NewListener("Search for a package", "",
		"", func(args []string) error {
			if len(args) < 1 {
				Err(fmt.Errorf("not enough arguments"))
			}
			if err := Search(args[0]); err != nil {
				return err
			}
			return nil
		}, "search")
	NewListener("Install a package", "",
		"", func(args []string) error {
			if len(args) < 1 {
				Err(fmt.Errorf("not enough arguments"))
			}
			for _, s := range args {
				if err := Install(s); err != nil {
					return err
				}
			}
			return nil
		}, "add")
	NewAlias("add", "install")
	NewListener("Update a package", "",
		"", func(args []string) error {
			if len(args) < 1 {
				Err(fmt.Errorf("not enough arguments"))
			}
			for _, s := range args {
				if err := Update(s); err != nil {
					return err
				}
			}
			return nil
		}, "update")
	NewListener("Remove a package", "",
		"", func(args []string) error {
			if len(args) < 1 {
				Err(fmt.Errorf("not enough arguments"))
			}
			for _, s := range args {
				if err := Uninstall(s); err != nil {
					return err
				}
			}
			return nil
		}, "remove")
	NewAlias("remove", "uninstall")
	NewListener("Show changelog for a package", "",
		"", func(args []string) error {
			if len(args) < 1 {
				return Changelog("hayashi")
			}
			for _, s := range args {
				if err := Changelog(s); err != nil {
					return err
				}
			}
			return nil
		}, "changelog")
	NewListener("Run clone task on a package", "",
		"", func(args []string) error {
			if len(args) < 1 {
				Err(fmt.Errorf("not enough arguments"))
			}
			return Clone(args[0])
		}, "task", "clone")
	NewListener("Run build task on a package", "",
		"", func(args []string) error {
			if len(args) < 1 {
				Err(fmt.Errorf("not enough arguments"))
			}
			return Build(args[0])
		}, "task", "build")
	NewListener("Run pack task on a package", "",
		"", func(args []string) error {
			if len(args) < 1 {
				Err(fmt.Errorf("not enough arguments"))
			}
			return Pack(args[0])
		}, "task", "pack")

	// if len(args) < 1 {
	// 	return GetHelp([]string{"task"}, Flags.AllFlags())
	// }
}

func RunArgs(argv []string) error {
	if len(argv) == 0 {
		Help(1)
	}

	results := Listeners
	result := ArgListener{}
	found := false
	i := 0
	for i = range argv {
		next := []ArgListener{}
		for _, l := range results {
			if len(l.Argv) > i {
				if l.Argv[i] == argv[i] {
					next = append(next, l)
				}
			} else {
				// previously collected a listener that has less arguments than provided
				result = l
				found = true
				break
			}
		}
		if found {
			break
		}
		results = next
	}

	if len(results) == 0 {
		Help(1)
		return nil
	}

	if !found {
		if len(results) > 1 {
			return GetHelp(argv, Flags.AllFlags())
		}
		result = results[0]
	}

	i = len(result.Argv)

	args := []string{}
	if len(argv) > i {
		args = argv[i:]
	}
	if err := result.Fn(args); err != nil {
		return err
	}

	return nil
}
