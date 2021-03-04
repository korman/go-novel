package node

import (
	global "gonovel/internal"
	"gonovel/internal/inter"
	"gonovel/pkg/node/internal"
)

func CreateNode(nodeType global.NodeType) inter.Node {
	var node inter.Node = nil

	switch nodeType {
	case global.Volume:
		node = new(internal.VolumeNode)
		node.Init()
	case global.Chapter:
		node = new(internal.ChapterNode)
		node.Init()
	}

	return node
}
