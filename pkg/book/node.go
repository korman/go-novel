package book

type Node interface {
	Text() string
	ChildNodes() []Node
}
