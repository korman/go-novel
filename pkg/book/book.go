package book

import (
	"gonovel/configs"
	"gonovel/internal/utils"
	"regexp"
	"strings"
)

type Book struct {
	BookInfomation *BookInfo
	Chapters       []*BookNode
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
	strlist := make([]string, 0)

	reg := regexp.MustCompile("(?m)^(.+)")

	lineList := reg.FindAllStringIndex(s, -1)

	var currentIndex int = -1

	chapters := make([]*BookNode, 0)
	var chapter *BookNode = nil

	for i := 0; i < len(lineList); i++ {
		info := s[lineList[i][0]:lineList[i][1]]
		info = strings.Replace(info, " ", "", -1)

		index := this.parseVolumes(info)

		if -1 == index {
			continue
		} else if currentIndex == index {
			continue
		} else if currentIndex < index {
			if nil != chapter {
				chapter.EndPos = lineList[i][0] - 1
				chapters = append(chapters, chapter)
			}

			chapter = new(BookNode)
			chapter.Index = index
			chapter.NodeType = "卷"
			chapter.StartPos = lineList[i][1]
		}
	}

	if nil != chapter {
		chapter.EndPos = len(s) - 1
		chapters = append(chapters, chapter)
		chapter = nil
	}

	return strlist
}

func (this *Book) parseVolumes(s string) int {
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

func (this *Book) parseSubInfo(s string) error {
	return nil
}

func (this *Book) parseBookInfo(s string) error {
	// TODO

	return nil
}
