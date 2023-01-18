package cli

import (
	"fmt"
)

type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    *bool  // value as set
	DefValue string // default value (as text); for usage message
}

type FlagSet struct {
	name   string
	parsed bool
	vars   map[string]*Flag
	argv   []string
	args   []string // arguments without flags
	pargs  []string // copy of argv used for parsing
}

func (f FlagSet) Arg(n int) string {
	if len(f.args) < n + 1 || len(f.args) < 1 {
		return ""
	}
	return f.args[n]
}

func (f FlagSet) Args() []string {
	return f.args
}

func (f FlagSet) AllFlags() []string {
	allflags := []string{}

	f.VisitAll(func(f *Flag) {
		if !*f.Value {
			return
		}

		if len(f.Name) > 1 {
		allflags = append(allflags, fmt.Sprintf("--%v", f.Name))
			return
		}
		allflags = append(allflags, fmt.Sprintf("-%v", f.Name))
	})

	return allflags
}

func (f FlagSet) UnquoteUsage(fl *Flag) (string, string) {
	// Look for a back-quoted name, but avoid the strings package.
	name := ""
	usage := fl.Usage
	for i := 0; i < len(usage); i++ {
		if usage[i] == '`' {
			for j := i + 1; j < len(usage); j++ {
				if usage[j] == '`' {
					name = usage[i+1 : j]
					usage = usage[:i] + name + usage[j+1:]
					return name, usage
				}
			}
			break // Only one back quote; use type name.
		}
	}
	// No explicit name, so use type if we can find one.
	/* name := "value"
	switch flag.Value.(type) {
	case boolFlag:
		name = ""
	case *durationValue:
		name = "duration"
	case *float64Value:
		name = "float"
	case *intValue, *int64Value:
		name = "int"
	case *stringValue:
		name = "string"
	case *uintValue, *uint64Value:
		name = "uint"
	} */
	return name, usage
}

func (f FlagSet) VisitAll(fn func(f *Flag)) {
	for _, f := range f.vars {
		fn(f)
	}
}

var Flags FlagSet

func (f *FlagSet) BoolVar(pt *bool, name string, usage string) error {
	*pt = false
	flag := Flag{
		Name:     name,
		Usage:    usage,
		Value:    pt,
		DefValue: "false",
	}
	if Flags.vars == nil {
		Flags.vars = make(map[string]*Flag)
	}
	Flags.vars[name] = &flag
	return nil
}

// parseOne parses one flag. It reports whether a flag was seen.
func (f *FlagSet) parseOne() (int, error) {
	if len(f.pargs) == 0 {
		return -1, nil
	}
	for i, s := range f.pargs {
		if len(s) < 2 || s[0] != '-' {
			continue
		}
		numMinuses := 1
		if s[1] == '-' {
			numMinuses++
			if len(s) == 2 { //	"--"	terminates	the	flags
				continue
			}
		}
		name := s[numMinuses:]
		if len(name) == 0 || name[0] == '-' || name[0] == '=' {
			continue
		}

		flag, ok := f.vars[name]
		if !ok {
			if name == "help" || name == "h" { //	special	case	for	nice	help	message.
				f.usage()
				return -1, fmt.Errorf("printing	help	message.")
			}
			return -1, fmt.Errorf("flag	provided	but	not	defined:	-%s", name)
		}

		if *flag.Value {
			return i, nil // flag found in previous loop
		}

		// remove all flags equivalent to
		nargs := []string{} // new empty array for args
		for _, g := range f.args {
			if g[1:] == name || g[2:] == name {
				continue
			}
			nargs = append(nargs, g)
		}
		f.args = nargs

		*flag.Value = true
		return i, nil
	}

	return -1, nil
}

// Parse parses flag definitions from the argument list, which should not
// include the command name. Must be called after all flags in the FlagSet
// are defined and before flags are accessed by the program.
// The return value will be ErrHelp if -help or -h were set but not defined.
func (f *FlagSet) Parse(arguments []string) error {
	f.parsed = true
	f.argv = arguments
	f.args = f.argv
	f.pargs = f.argv
	for {
		seen, err := f.parseOne()
		if err != nil {
			return err
		}
		if seen > -1 {
			f.pargs = f.pargs[seen+1:]
			continue
		}
		break
	}
	return nil
}
