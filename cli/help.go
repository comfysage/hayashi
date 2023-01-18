package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/crispybaccoon/hayashi/doc"
)

func Help(exitcode int) {
	Flags.usage()
	os.Exit(exitcode)
}

func (f *FlagSet) usage() {

	space := func(n int) string {
		return (strings.Repeat(" ", n))
	}

	fmt.Println("hayashi. a tiny distro-independent package manager written in Go.")

	spacing := 2
	bspacing := spacing
	shiftw := 20
	linelen := 45

	fmt.Printf("\n" + space(bspacing) +"usage: \n\n")

	Flags.VisitAll(func(f *Flag) {
		var b strings.Builder
		if len(f.Name) > 1 {
			fmt.Fprintf(&b, space(bspacing) + "--%s", f.Name)
		} else {
			fmt.Fprintf(&b, space(bspacing) + "-%s", f.Name)
		}
		name, usage := Flags.UnquoteUsage(f)
		if len(name) > 0 {
			b.WriteString(" ")
			b.WriteString(name)
		}

		if b.Len() > shiftw {
			b.WriteString("\n" + space(shiftw))
		} else {
			b.WriteString(space(shiftw - b.Len()))
		}
		b.WriteString(space(spacing))
		usage = strings.ReplaceAll(usage, "\n", "\n"+space(spacing))
		for i, s := range usage {
			if i+1%linelen == 0 {
				b.WriteString("\n" + space(shiftw+spacing))
			}
			b.WriteString(string(s))
		}

		fmt.Fprintf(&b, " (default %v)", f.DefValue)
		fmt.Fprint(os.Stderr, b.String(), "\n")
	})

}

func GetHelp(query []string, allflags []string) error {

	doc, err := doc.DOCS.FindQuery(query, allflags)
	if err != nil {
		return err
	}

	printf(doc.String())

	return nil
}
