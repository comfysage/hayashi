package cli

import (
	"fmt"
	"os"

	"github.com/crispybaccoon/hayashi/doc"
)

func Help(exitcode int) {
	printf(doc.HELP_DOC.String())
	Flags.usage()
	os.Exit(exitcode)
}

func (f *FlagSet) usage() {
	flagdocs := []doc.Shortdoc{}

	Flags.VisitAll(func(f *Flag) {
		flagdoc := doc.Shortdoc{}
		if len(f.Name) > 1 {
			flagdoc[0] = fmt.Sprintf("--%s", f.Name)
		} else {
			flagdoc[0] = fmt.Sprintf("-%s", f.Name)
		}
		_, usage := Flags.UnquoteUsage(f)
		flagdoc[1] = usage
		flagdoc[2] = ""

		flagdocs = append(flagdocs, flagdoc)
	})

	helpman := doc.ManDoc{
		Flagdoc: flagdocs,
	}.String()

	printf(helpman)
}

func GetHelp(query []string, allflags []string) error {

	doc, err := doc.DOCS.FindQuery(query, allflags)
	if err != nil {
		return err
	}

	printf(doc.String())

	return nil
}
