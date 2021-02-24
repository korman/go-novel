package book

import (
	global "gonovel/internal"
	"gonovel/internal/inter"
	"gonovel/pkg/node"
)

type Book struct {
	BookInfomation *BookInfo
	Chapters       []*inter.Node
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

func (this *Book) parseSingleLine(s string) []string {
	node := node.CreateNode(global.Volume)

	node.Parse(s)

	return nil
}

func (this *Book) parseSubInfo(s string) error {
	return nil
}

func (this *Book) parseBookInfo(s string) error {
	// TODO

	return nil
}
