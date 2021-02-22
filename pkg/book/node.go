package book

type BookNode struct {
	StartPos  int
	EndPos    int
	Index     int
	Brief     string
	IndexText string
	NodeType  string
	Childs    []*BookNode
}
