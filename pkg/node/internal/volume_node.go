package internal

import (
	global "gonovel/internal"
	"gonovel/internal/inter"
)

type VolumeNode struct {
	startPos int
	endPos   int
	index    int
	brief    int
	nodeType global.NodeType
	childs   []inter.Node
}

func (this *VolumeNode) Parse(s string) error {
	return nil
}

func (this *VolumeNode) StartPos() int {
	return 0
}

func (this *VolumeNode) SetStartPos(pos int) {
}

func (this *VolumeNode) EndPos() int {
	return 0
}

func (this *VolumeNode) SetEndPos(pos int) {

}

func (this *VolumeNode) Index() int {
	return 0
}

func (this *VolumeNode) SetIndex(index int) {

}

func (this *VolumeNode) Brief() string {
	return ""
}

func (this *VolumeNode) SetBrief(s string) {

}

func (this *VolumeNode) NodeType() global.NodeType {
	return global.Volume
}

func (this *VolumeNode) SetNodeType(t global.NodeType) {
	this.nodeType = t
}

func (this *VolumeNode) Childs() []inter.Node {
	return this.childs
}
