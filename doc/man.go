package doc

import (
	"fmt"
	"strings"
)

func (d Shortdoc) ShortString() string {
	return generateLineDoc([2]string{d[0], d[1]})
}

func (d Doc) ShortString() string {
	return generateLineDoc([2]string{d.Name, d.Man.Short})
}

func (d ManDoc) String() string {
	var s strings.Builder

	if len(d.Short) > 0 {
		fmt.Fprintln(&s, d.Short)
		fmt.Fprintln(&s, "")
	}
	if len(d.Long) > 0 {
		fmt.Fprintln(&s, d.Long)
		fmt.Fprintln(&s, "")
	}
	if len(d.Usage) > 0 {
		fmt.Fprintln(&s, "Usage:")
		fmt.Fprintln(&s, space(bspacing), d.Usage)
		fmt.Fprintln(&s, "")
	}
	if len(d.Cmddoc) > 0 {
		fmt.Fprintln(&s, "Commands:")
		for _, c := range d.Cmddoc {
			fmt.Fprintln(&s, (*c).ShortString())
		}
	}
	if len(d.Flagdoc) > 0 {
		fmt.Fprintln(&s, "Flags:")
		for _, c := range d.Flagdoc {
			fmt.Fprintln(&s, c.ShortString())
		}
	}

	return s.String()
}

func (d Doc) String() string {
	return d.Man.String()
}
