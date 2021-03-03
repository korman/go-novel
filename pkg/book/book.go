package book

import (
	"errors"
	"fmt"
	global "gonovel/internal"
	"gonovel/internal/inter"
	"gonovel/internal/utils"
	"gonovel/pkg/node"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Book struct {
	BookInfomation *BookInfo
	Chapters       []inter.Node
}

func (this *Book) SetBookInfo(info *BookInfo) {
	this.BookInfomation = info
}

func (this *Book) ConvertToSingleMd(outpath string) error {
	var err error = nil

	var info string = ""

	if "" == this.BookInfomation.BookName {
		return errors.New("没有书名")
	}

	info += fmt.Sprintf("# %s\n\n", this.BookInfomation.BookName)

	for _, v := range this.Chapters {
		header := fmt.Sprintf("## 第%d卷\n\n", v.Index())

		info += header

		for _, c := range v.Childs() {
			chapterHeader := fmt.Sprintf("### 第%d章\n\n", c.Index())

			info += chapterHeader
			text := strings.Replace(c.Text(), " ", "", -1)
			info += text

			info += "\n\n"
		}
	}

	fullout := filepath.Join(outpath, fmt.Sprintf("%s.md", this.BookInfomation.BookName))

	err = utils.WriteFile(fullout, info)

	return err
}

func (this *Book) ConvertToMd(outpath string) error {
	bookPath := filepath.Join(outpath, this.BookInfomation.BookName)

	println("创建书籍目录：" + bookPath)

	err := os.MkdirAll(bookPath, os.ModePerm)

	readmeString, err := this.GenReadmeMarkdownString()

	if nil != err {
		return err
	}

	err = utils.WriteFile(filepath.Join(bookPath, "README.md"), readmeString)

	if nil != err {
		return err
	}

	for _, v := range this.Chapters {
		volumeIndex := v.Index()

		volumePath := filepath.Join(bookPath, strconv.Itoa(volumeIndex))

		err := os.MkdirAll(volumePath, os.ModePerm)

		if nil != err {
			return err
		}

		volumeReadme, err := v.GenReadmeMarkdownString()

		if nil != err {
			return err
		}

		volumeReadmePath := filepath.Join(volumePath, "README.md")
		err = utils.WriteFile(volumeReadmePath, volumeReadme)

		if nil != err {
			return err
		}

		for _, c := range v.Childs() {
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

func (this *Book) GenReadmeMarkdownString() (string, error) {
	var readme string = ""

	bookinfo, err := this.genMarkdownBookInfo()

	if nil != err {
		return "", err
	}

	readme += bookinfo

	catalog, err := this.generateCatalogMarkdown()

	if nil != err {
		return "", err
	}

	readme += catalog

	return readme, nil
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

func (this *Book) generateCatalogMarkdown() (string, error) {
	var info string = "## 小说目录\n\n"

	for _, v := range this.Chapters {
		info += fmt.Sprintf("* [第%d卷](%d/README.md)\n", v.Index(), v.Index())
	}

	info += "\n"

	return info, nil
}

func (this *Book) genMarkdownBookInfo() (string, error) {
	var info string = ""

	if "" != this.BookInfomation.BookName {
		info += fmt.Sprintf("# %s\n\n", this.BookInfomation.BookName)
	}

	var authors string = "## 作者:\n\n"

	if 0 < len(this.BookInfomation.Author) {
		authors = "## 作者:\n\n"

		for _, v := range this.BookInfomation.Author {
			authors += fmt.Sprintf("* %s\n", v)
		}

		authors += "\n\n"
	} else {
		authors += "未知\n\n"
	}

	info += authors

	if "" != this.BookInfomation.Intro {
		info += fmt.Sprintf("## 简介\n\n %s\n", this.BookInfomation.Intro)
	} else {
		info += fmt.Sprintf("## 简介\n\n %s\n", "无简介")
	}

	return info, nil
}
