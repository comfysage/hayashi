package exec

import "github.com/crispybaccoon/hayashi/util"

func Clone(url string, name string) error {
	return runOne(cloneCmd(url, util.PathRepo(name)), util.REPO_ROOT, nil)
}

func Changelog(name string) error {
	return runOne(logCmd(), util.PathRepo(name), nil)
}

func Pack(name string, path string, prefix string) error {
	cmd, err := packCmd(name, path, prefix)
	if err != nil {
		return err
	}
	return runOne(cmd, util.PACK_ROOT, nil)
}

func RunInRepo(name string, cmd []string) error {
	return run(cmd, util.PathRepo(name), nil)
}

func RunInPack(prefix string, cmd []string) error {
	return run(cmd, util.PathPackDir(prefix), nil)
}
