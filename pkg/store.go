package pkg

import (
	"fmt"
	"io"
	"strings"

	"github.com/crispybaccoon/hayashi/util"
	"gopkg.in/yaml.v3"
)

type StoreFile struct {
	Installed []string `yaml:"installed"`
	installed []*Pkg
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

func pkgFromStoreFormat(format string) (*Pkg, error) {

	split_format := strings.Split(format, "/")
	if len(split_format) < 2 {
		return nil, fmt.Errorf("incorrect format of store item")
	}
	collection, name := split_format[0], split_format[1]

	pkg, err := GetPkgFromPath(util.PathPkg(collection, name))
	if err != nil {
		return nil, err
	}

	return &pkg, nil
}

func (c *StoreFile) fixLength() {
	for {
		if len(c.installed) < len(c.Installed) {
			c.installed = append(c.installed, nil)
			continue
		}
		break
	}
}

func (c *StoreFile) at(index uint) (*Pkg, error) {
	if len(c.Installed) < 1 {
		return nil, fmt.Errorf("list of installed packages is empty")
	}
	if len(c.Installed)-1 < int(index) {
		return nil, fmt.Errorf("index of pkg outside memory")
	}

	pkg, err := pkgFromStoreFormat(c.Installed[index])
	if err != nil {
		return nil, err
	}

	c.fixLength()
	c.installed[index] = pkg

	return pkg, nil
}

func (c StoreFile) ForEach(fn func(p Pkg) error) error {

	for index := range c.Installed {
		pkg, err := pkgFromStoreFormat(c.Installed[index])
		if err != nil {
			fmt.Printf("error while retrieving pkg\n  %v", err)
		}

		c.fixLength()
		c.installed[index] = pkg

		if err := fn(*pkg); err != nil {
			fmt.Printf("error while processing hooks (%v)\n  %v", pkg.Name, err)
		}
	}

	return nil
}
