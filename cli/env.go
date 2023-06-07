package cli

import (
	"fmt"
	"os"
	"regexp"

	"github.com/crispybaccoon/hayashi/util"
)

func Env() error {
	script := fmt.Sprintf(`
export PATH="$PATH:%s"
export XDG_DATA_DIRS="${XDG_DATA_DIRS:-/usr/share:/usr/local/share}:%s"
`, util.PathPackDir("bin"), util.PathPackDir("share"))
	var re = regexp.MustCompile(`\n\n`)
	script = re.ReplaceAllString(script, "\n")
	fmt.Fprintf(os.Stdout, script)

	return nil
}
