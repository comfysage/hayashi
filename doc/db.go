package doc

import "fmt"

func (db Docs) FindQuery(query []string, flags []string) (*Doc, error) {

	q := query[0]
	for _, d := range db.documents {
		if d.Name != q {
			continue
		}
		for _, flag := range flags {
			for _, fl := range d.Man.Flagdoc {
				if flag == fl[0] {
					return generateFlagDoc(fl), nil
				}
			}

		}
		return d, nil
	}

	return nil, fmt.Errorf("document not found.")
}
