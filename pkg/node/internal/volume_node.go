package internal

import (
	"gonovel/configs"
	global "gonovel/internal"
	"gonovel/internal/inter"
	"gonovel/internal/utils"
	"regexp"
	"strings"
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
	reg := regexp.MustCompile("(?m)^(.+)")

	lineList := reg.FindAllStringIndex(s, -1)
	var index int = -1

	for i := 0; i < len(lineList); i++ {
		info := s[lineList[i][0]:lineList[i][1]]
		info = strings.Replace(info, " ", "", -1)

		index = this.parseVolumes(info)

		if -1 == index {
			continue
		} else if this.index == index {
			continue
		} else if this.index < index {
			this.endPos = lineList[i][0] - 1

			this.index = index
			this.nodeType = global.Volume
			this.startPos = lineList[i][1]
		}

	}

	return nil
}

func (this *VolumeNode) init() {
	this.childs = make([]inter.Node, 0)
	this.index = -1
}

func (this *VolumeNode) parseVolumes(s string) int {
	var index int = -1

	for _, v := range configs.VolumeRegexp {
		reg := regexp.MustCompile(v)

		volIds := reg.FindAllStringIndex(s, 1)

		if nil == volIds {
			continue
		}

		info := s[volIds[0][0]:volIds[0][1]]
		idx, _ := utils.GenNumberFromString(info)

		index = int(idx)

		break
	}

	return index
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
