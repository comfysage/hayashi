package pkg

import (
	"io"

	"gopkg.in/yaml.v3"
)

type StoreFile struct {
	Installed []string `yaml:"installed"`
}

func DefaultStoreFile() StoreFile {
	return StoreFile{
		Installed: []string{},
	}
}

func (c StoreFile) String() (string, error) {
	d, err := yaml.Marshal(c)
	return string(d), err
}

func (c *StoreFile) FromString(r io.Reader) error {
	d := yaml.NewDecoder(r)
	return d.Decode(c)
}

func (c *StoreFile) AddInstalled(pkg Pkg) error {
	pkgString, err := pkg.DisplayName()
	if err != nil {
		return err
	}

	for _, s := range c.Installed {
		if s == pkgString {
			return nil
		}
	}
	c.Installed = append(c.Installed, pkgString)
	SaveStoreFile(*c)

	return nil
}

func (c *StoreFile) RemoveInstalled(pkg Pkg) error {
	pkgString, err := pkg.DisplayName()
	if err != nil {
		return err
	}

	// NOTE: If you do not care about ordering, you have the much faster
	// possibility to replace the element to delete with the one at the end of
	// the slice and then return the n-1 first elements.
	//   ```go
	//	 func remove(s []int, i int) []int {
	//		 s[i] = s[len(s)-1]
	//		 return s[:len(s)-1]
	//	 }
	//   ```
	// If you want to keep your array ordered, you have to shift all of the
	// elements at the right of the deleting index by one to the left.
	for i, s := range c.Installed {
		if s == pkgString {
			c.Installed = append(c.Installed[:i], c.Installed[i+1:]...)
		}
	}
	SaveStoreFile(*c)

	return nil
}
