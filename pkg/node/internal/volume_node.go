package internal

import (
	"errors"
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
	brief    string
	nodeType global.NodeType
	childs   []inter.Node
	text     string
}

func (this *VolumeNode) Parse(s string) (string, error) {
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
		}

		if -1 < this.index && -1 < this.startPos {
			this.endPos = lineList[i][0]
			break
		}

		this.index = index
		this.startPos = lineList[i][1]
	}

	if -1 == this.index {
		return "", errors.New("没有找到卷")
	}

	this.text = s[this.startPos:this.endPos]

	return s[this.endPos:], nil
}

func (this *VolumeNode) Text() string {
	return this.text
}

func (this *VolumeNode) Init() {
	this.childs = make([]inter.Node, 0)
	this.index = -1
	this.startPos = -1
	this.endPos = -1
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
	return this.startPos
}

func (this *VolumeNode) SetStartPos(pos int) {
	this.startPos = pos
}

func (this *VolumeNode) EndPos() int {
	return this.endPos
}

func (this *VolumeNode) SetEndPos(pos int) {
	this.endPos = pos
}

func (this *VolumeNode) Index() int {
	return this.index
}

func (this *VolumeNode) SetIndex(index int) {
	this.index = index
}

func (this *VolumeNode) Brief() string {
	return this.brief
}

func (this *VolumeNode) SetBrief(s string) {
	this.brief = s
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
