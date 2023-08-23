package pkg

import (
	"fmt"
	fpath "path"
	"path/filepath"
	"strings"

	"github.com/crispybaccoon/hayashi/util"
)

type Pkg struct {
	Name       string   `yaml:"pkg"`
	Url        string   `yaml:"url,omitempty"`
	Desc       string   `yaml:"desc,omitempty"`
	Clone      bool     `yaml:"clone,omitempty"`
	Collection string   `yaml:"collection,omitempty"`
	Install    []string `yaml:"install,omitempty"`
	Remove     []string `yaml:"remove,omitempty"`
	Update     []string `yaml:"update,omitempty"`
	Pack       []Pack   `yaml:"pack,omitempty"`
}

func NewPkg(name string, url string) Pkg {
	return Pkg {
		Name: name,
		Url: url,
		Clone: true,
	}
}

func (pkg Pkg) DisplayName() (string, error) {
	if pkg.Collection == "" {
		return "", fmt.Errorf("no collection defined for pkg")
	}
	if pkg.Name == "" {
		return "", fmt.Errorf("no name defined for pkg")
	}
	return pkg.Collection + "/" + pkg.Name, nil
}

func (pkg *Pkg) inferCollection(path string) error {
	if len(pkg.Collection) > 0 {
		return nil
	}
	cwd, err := util.GetCwd()
	if err != nil {
		return err
	}
	abspath, err := filepath.Abs(fpath.Join(cwd, path))
	if err != nil {
		return err
	}

	sp := strings.Split(abspath, "/")
	if len(sp) < 2 {
		return fmt.Errorf("path to pkg not long enough")
	}
	pkg.Collection = sp[len(sp)-2]
	return nil
}

func (pkg *Pkg) InferInfo(path string) error {
	// infer collection
	err := pkg.inferCollection(path)
	if err != nil {
		return err
	}

	// infer url
	// founddot := false
	foundslash := false
	founddb := false
	for _, s := range pkg.Url {
		if string(s) == "." {
			// founddot = true

			// found address without protocol
			if foundslash == false {
				pkg.Url = "https://" + pkg.Url
				break
			}

			continue
		}

		if string(s) == "/" {
			foundslash = true

			// found protocol
			if founddb {
				break
			}

			// potential address already catched
			pkg.Url = "https://github.com/" + pkg.Url
			break
		}

		if string(s) == ":" {
			founddb = true
		}
	}

	return nil
}

// error checking for pkg to avoid unwanted behaviour
func (pkg Pkg) SafeGuard() error {

	// require pkg name
	if len(pkg.Name) < 1 {
		return fmt.Errorf("no pkg name specified.")
	}

	// require url
	if len(pkg.Url) < 1 {
		return fmt.Errorf("no url specified.")
	}

	return nil
}
