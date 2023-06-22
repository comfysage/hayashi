package exec

import (
	"fmt"
	"os"
	"strconv"

	"github.com/crispybaccoon/hayashi/util"
)

func cloneCmd(url string, output string) []string {
	return []string{"git", "clone", "--filter=blob:none", url, output}
}

func fetchCmds() []string {
	return []string{"git fetch -q", "git checkout -q"}
}

func logCmd() []string {
	return []string{"git", "-c", "pager.show=false",
		"show", "--format=' - %C(yellow)%h%C(reset) %<(80,trunc)%s'", "-q", "@@{1}..@@{0}"}
}

func installCmd(src string, dst string, perm uint64) []string {
	mode := strconv.FormatUint(perm, 8)
	if dst == util.PathPackFile("bin", "hayashi") {
		tmpdst := util.PathPackFile("bin", "hayashi.new")
		return []string{"install", "-m", mode, src, tmpdst, ";", "mv", tmpdst, dst}
	}
	return []string{"install", "-m", mode, src, dst}
}

func packCmd(name string, path string, prefix string) ([]string, error) {
	var cmd []string
	var err error
	if len(path) == 0 || len(prefix) == 0 {
		return nil, fmt.Errorf("pack struct has missing properties")
	}
	src := util.PathRepoFile(name, path)
	dst := util.PathPackFile(prefix, path)

	if _, err = os.Stat(src); err != nil {
		return nil, err
	}

	/* if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	} */

	/* source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close() */

	if _, err = util.Mkdir(util.PathPackDir(prefix)); err != nil {
		return nil, err
	}

	var destination *os.File
	if destination, err = os.Create(dst); err != nil {
		return nil, err
	}
	defer destination.Close()

	if prefix == "bin" {
		cmd = installCmd(src, dst, 0755)
	} else {
		cmd = installCmd(src, dst, 0644)
	}
	return cmd, nil
}
