package doc

type Doc struct {
	Name string
	Man  ManDoc
}

// document for bigger items ( "pkg", "config" )
type ManDoc struct {
	Name    string
	Short   string
	Long    string
	Usage   string
	Cmddoc  []*Doc
	Flagdoc []Shortdoc
}

// document for (sub)commands or flags
type Shortdoc [3]string

// 0:	Name  string
// 1:	short string
// 2:	usage string

type Docs struct {
	documents []*Doc
}
