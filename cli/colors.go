package cli

import "fmt"

var (
	COLOR_RED     string = "\x1b[31m"
	COLOR_GREEN   string = "\x1b[32m"
	COLOR_YELLOW  string = "\x1b[33m"
	COLOR_BLUE    string = "\x1b[34m"
	COLOR_MAGENTA string = "\x1b[35m"
	COLOR_CYAN    string = "\x1b[36m"
	COLOR_RESET   string = "\x1b[0m"
)

func printf(str string) {
	fmt.Printf("%v" + COLOR_RESET + "\n", str)
}
