package book

import (
	"fmt"
	global "gonovel/internal"
	"gonovel/internal/inter"
	"gonovel/pkg/node"
)

type Book struct {
	BookInfomation *BookInfo
	Chapters       []inter.Node
}

func (this *Book) SetBookInfo(info *BookInfo) {
	this.BookInfomation = info
}

func (this *Book) Load(txt string) error {
	err := this.parseBookInfo(txt)

	if nil != err {
		return err
	}

	strList := this.parseSingleLine(txt)

	if 0 == len(strList) {
		return nil
	}

	//	this.Chapters, err = this.parseVolumes(txt)

	if nil != err {
		return err
	}

	return nil
}

func (this *Book) findNodeByIndex(index uint32) inter.Node {
	for _, v := range this.Chapters {
		if v.Index() == int(index) {
			return v
		}
	}

	return nil
}

func (this *Book) parseSingleLine(s string) []string {
	text := s
	var err error = nil

	this.Chapters = make([]inter.Node, 0)

	for 0 < len(text) {
		node := node.CreateNode(global.Volume)
		text, err = node.Parse(text)

		if nil != err {
			return nil
		}

		repeatNode := this.findNodeByIndex(uint32(node.Index()))

		if nil != repeatNode {
			repeatNode.Merge(node)
			global.Error(fmt.Sprintf("重复的卷,合并: %d", node.Index()))
		} else {
			this.Chapters = append(this.Chapters, node)
			global.Error(fmt.Sprintf("卷%d", node.Index()))
		}
	}

	return nil
}

func (this *Book) parseBookInfo(s string) error {
	// TODO

	return nil
}
