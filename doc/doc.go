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
	Cmddoc  docs
	Flagdoc []Shortdoc
}

// document for (sub)commands or flags
type Shortdoc [3]string

// 0:	Name  string
// 1:	short string
// 2:	usage string

type doc *Doc
type docs []doc

type Docs struct {
	documents docs
}
