package pkg

import (
	"fmt"
	"strings"
)

type Pkg struct {
	Name       string   `yaml:"pkg"`
	Url        string   `yaml:"url"`
	Desc       string   `yaml:"desc"`
	Collection string   `yaml:"collection"`
	Install    []string `yaml:"install,flow"`
	Remove     []string `yaml:"remove,flow"`
	Update     []string `yaml:"update,flow"`

	Bin string `yaml:"bin"`
}

func (pkg *Pkg) inferCollection(path string) error {
	if len(pkg.Collection) > 0 {
		return nil
	}
	sp := strings.Split(path, "/")
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
