package pkg

import (
	"io"

	yaml "gopkg.in/yaml.v3"
)

func (pkg Pkg) String() (string, error) {
	d, err := yaml.Marshal(&pkg)
	return string(d), err
}

func (pkg *Pkg) FromString(r io.Reader) error {
	d := yaml.NewDecoder(r)
	return d.Decode(pkg)
}
