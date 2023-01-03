package cli

import (
	"fmt"
	"os"
	"strings"
)

func Err(err error) {
	if err != nil {
		var str strings.Builder
		fmt.Fprintf(&str, COLOR_RED+"there's been an error: %v", err)
		printf(str.String())
		os.Exit(1)
	}
}
