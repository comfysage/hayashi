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
