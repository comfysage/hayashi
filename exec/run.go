package exec

import (
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
)

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func run(script []string, pwd string) error {
	for _, c := range script {
		el := strings.Split(c, " ")
		if el[0] == "cd" {
			el[0] = pwd
			pwd = path.Join(el...)
			continue
		}

		/* var stdout []byte
		var errStdout error */
		cmd := exec.Command(el[0], el[1:]...)
		cmd.Dir = pwd
		stdoutpipe, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		err = cmd.Start()
		if err != nil {
			return err
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			_, err = copyAndCapture(os.Stdout, stdoutpipe)
			wg.Done()
		}()

		wg.Wait()

		if err != nil {
			return err
		}

		err = cmd.Wait()
		if err != nil {
			return err
		}
	}

	return nil
}

func runOne(cmd []string, pwd string) error {
	return run([]string{strings.Join(cmd, " ")}, pwd)
}
