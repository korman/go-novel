package inter

import global "gonovel/internal"

type Node interface {
	Init()
	StartPos() int
	SetStartPos(int)
	EndPos() int
	SetEndPos(int)
	Index() int
	SetIndex(int)
	Brief() string
	SetBrief(string)
	Merge(Node) error
	NodeType() global.NodeType
	SetNodeType(global.NodeType)
	Childs() []Node
	GenMarkdownFormat() (string, error)
	Text() string
	Parse(string) (string, error)
}
