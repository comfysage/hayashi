package pkg

import (
	"bufio"
	"crispybaccoon/hayashi/ini"
	"strings"
)

func (pkg Pkg) iniString() string {
	o := ini.NewObject()
	o.Add("pkg", "name", pkg.Name)
	o.Add("pkg", "url", pkg.Url)
	o.Add("pkg", "description", pkg.Desc)

	o.Add("script", "install", strings.Join(pkg.Install, " ; "))
	o.Add("script", "remove", strings.Join(pkg.Remove, " ; "))
	o.Add("script", "update", strings.Join(pkg.Update, " ; "))

	return o.ToString()
}

func (pkg *Pkg) IniFromString(str bufio.Scanner) {
	i := ini.NewIni()
	i.Parse(str)
	pkg.Name = i.Object.Get("pkg", "name")
	pkg.Url = i.Object.Get("pkg", "url")
	pkg.Desc = i.Object.Get("pkg", "description")

	install := i.Object.Get("script", "install")
	pkg.Install = strings.Split(install, " ; ")
	remove := i.Object.Get("script", "remove")
	pkg.Remove = strings.Split(remove, " ; ")
	update := i.Object.Get("script", "update")
	pkg.Update = strings.Split(update, " ; ")
}
