package cli

import (
	"io"
	"os/exec"
	"path"
	"strings"
)

func run(script []string, pwd string) error {
	for _, c := range script {
		el := strings.Split(c, " ")
		if el[0] == "cd" {
			el[0] = pwd
			pwd = path.Join(el...)
			continue
		}
		cmd := exec.Command(el[0], el[1:]...)
		cmd.Dir = pwd
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		err = cmd.Start()
		if err != nil {
			return err
		}

		s, err := io.ReadAll(stdout)
		if err != nil {
			return err
		}
		printf(string(s))

		err = cmd.Wait()
		if err != nil {
			return err
		}
	}

	return nil
}
