package inter

import global "gonovel/internal"

type Node interface {
	StartPos() int
	SetStartPos(int)
	EndPos() int
	SetEndPos(int)
	Index() int
	SetIndex(int)
	Brief() string
	SetBrief(string)
	NodeType() global.NodeType
	SetNodeType(global.NodeType)
	Childs() []Node
	Parse(string) (string, error)
}
