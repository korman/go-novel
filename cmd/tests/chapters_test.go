package tests

import (
	global "gonovel/internal"
	"gonovel/pkg/book"
	"testing"

	"github.com/pkumza/numcn"
)

func TestNumberCN(t *testing.T) {
	s := "一万两千八百"
	n, err := numcn.DecodeToInt64(s)

	if nil != err {
		t.Error(err)
	}

	t.Log(n)
}

func TestLogSystem(t *testing.T) {
	global.Error("xxxxx")
	global.Fatal("xxxaaaa")
}

func TestExtractChapters(t *testing.T) {
	// s, err := utils.LoadTxtFile("txt_files/test_01.txt")
	// //[第卷]+[\s]*[^\x00-\xff一二三四五六七八九十零〇百千两]+[\s]*[章回部节集卷话]+
	// if nil != err {
	// 	t.Error(err)
	// }

	// reg := regexp.MustCompile(`(?m)^\s*[第]+[0-9一二三四五六七八九十零〇百千两]+[\s]*[卷]+.*$`)
	// r := reg.FindAllIndex([]byte(s), -1)
	// o := reg.FindAllStringIndex(s, -1)

	// t.Log(r)
	// t.Log(o)

	// for _, v := range r {
	// 	ms := string(s[v[0]:v[1]])
	// 	t.Log(ms)
	// }

	book, err := book.CreateBook("txt_files/test_01.txt")

	if nil == err {
		t.Error(err)
	}

	if nil == book {

	}
}
