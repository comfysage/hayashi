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
	if len(f.args) < n - 1 || len(f.args) < 1 {
		return ""
	}
	return f.args[n]
}

func (f FlagSet) Args() []string {
	return f.args
}

func (f FlagSet) UnquoteUsage(fl *Flag) (string, string) {
	flag := f.vars[fl.Name]
	return flag.Name, flag.Usage
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
func (f *FlagSet) parseOne() (bool, error) {
	if len(f.pargs) == 0 {
		return false, nil
	}
	for i, s := range f.pargs {
		f.pargs = f.pargs[i:]
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
				return false, fmt.Errorf("printing	help	message.")
			}
			return false, fmt.Errorf("flag	provided	but	not	defined:	-%s", name)
		}

		if *flag.Value {
			return true, nil // flag found in previous loop
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
		return true, nil
	}

	return false, nil
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
		if seen {
			f.pargs = f.pargs[1:]
			continue
		}
		break
	}
	return nil
}
