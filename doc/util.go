package doc

import "strings"

func space(n int) string {
	return (strings.Repeat(" ", n))
}

var (
	spacing  int = 2
	bspacing int = 4
	shiftw   int = 20
	linelen  int = 45
)

func generateLineDoc(info [2]string) string {

	var b strings.Builder

	b.WriteString(space(bspacing))
	b.WriteString(info[0])
	if len(info[0]) > shiftw {
		b.WriteString("\n" + space(shiftw))
	} else {
		b.WriteString(space(shiftw - b.Len()))
	}

	usage := strings.ReplaceAll(info[1], "\n", "\n"+space(spacing))
	b.WriteString(strings.Repeat(" ", spacing))
	for i, s := range usage {
		if i+1%linelen == 0 {
			b.WriteString("\n" + space(shiftw+spacing))
		}
		b.WriteString(string(s))
	}

	return b.String()
}

func generateFlagDoc(fl Shortdoc) doc {
	return &Doc{
		Name: fl[0],
		Man: ManDoc{
			Short: fl[0],
			Long:  fl[1],
			Usage: fl[2],
		},
	}
}
