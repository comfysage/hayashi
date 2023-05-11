package util

import "io/ioutil"

func GetPkgs() [2][]string {
	names := [2][]string{}

	collections, err := ioutil.ReadDir(PKG_ROOT)
	if err != nil {
		return names
	}
	for _, d := range collections {
		clName := d.Name()

		fd, err := ioutil.ReadDir(PathCl(clName))
		if err != nil {
			continue
		}
		for _, p := range fd {
			if p.IsDir() {
				continue
			}
			pName := p.Name()
			names[0] = append(names[0], clName)
			names[1] = append(names[1], RemoveExtension(pName))
		}
	}

	return names
}

func SearchPkgs(str string) [][2]string {
	pkgs := GetPkgs()
	results := SearchStrings(pkgs[1], str)

	p_results := [][2]string{}
	for _, i := range results {
		p_results = append(p_results, [2]string{pkgs[0][i], pkgs[1][i]})
	}
	return p_results
}
