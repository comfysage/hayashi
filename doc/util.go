package doc

func generateFlagDoc(fl Shortdoc) doc {
	return &Doc{
		Name: fl[0],
		Man: ManDoc{
			Short: fl[0],
			Long:  fl[1],
			Usage: fl[2],
		},
	}
}
