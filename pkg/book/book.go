package book

import (
	"fmt"
	global "gonovel/internal"
	"gonovel/internal/inter"
	"gonovel/internal/utils"
	"gonovel/pkg/node"
	"os"
	"path/filepath"
	"strconv"
)

type Book struct {
	BookInfomation *BookInfo
	Chapters       []inter.Node
}

func (this *Book) SetBookInfo(info *BookInfo) {
	this.BookInfomation = info
}

func (this *Book) ConvertToMd(outpath string) error {
	bookPath := filepath.Join(outpath, this.BookInfomation.BookName)

	println("创建书籍目录：" + bookPath)

	for _, v := range this.Chapters {
		volumeIndex := v.Index()

		volumePath := filepath.Join(bookPath, strconv.Itoa(volumeIndex))

		err := os.MkdirAll(volumePath, os.ModePerm)

		if nil != err {
			return err
		}

		for _, c := range v.Childs() {
			// chapterPath := filepath.Join(volumePath, strconv.Itoa(c.Index()))

			// err := os.MkdirAll(chapterPath, os.ModePerm)

			// if nil != err {
			// 	return err
			// }

			fpath := filepath.Join(volumePath, fmt.Sprintf("%d.md", c.Index()))

			text, err := c.GenMarkdownFormat()

			if nil != err {
				return err
			}

			err = utils.WriteFile(fpath, text)

			if nil != err {
				return err
			}
		}
	}

	return nil
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
