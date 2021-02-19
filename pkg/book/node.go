package book

type BookNode struct {
	Index     int
	Brief     string
	IndexText string
	Childs    []*BookNode
}
