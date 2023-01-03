package pkg

type Pkg struct {
	Name    string   `yaml:"pkg"`
	Url     string   `yaml:"url"`
	Desc    string   `yaml:"desc"`
	Install []string `yaml:"install,flow"`
	Remove  []string `yaml:"remove,flow"`
	Update  []string `yaml:"update,flow"`

	Bin string `yaml:"bin"`
}

func (pkg *Pkg) InferInfo() error {
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
