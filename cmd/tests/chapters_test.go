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
	book, err := book.CreateBook("txt_files/test_01.txt")

	if nil != err {
		t.Error(err)
	}

	if nil == book {

	}

	err = book.ConvertToMd(".")

	if nil != err {
		t.Error(err)
	}
}

func TestExportMarkdown(t *testing.T) {
	book, err := book.CreateBook("txt_files/test_01.txt")

	if nil != err {
		t.Error(err)
	}

	if nil == book {

	}

	err = book.ConvertToSingleMd(".")

	if nil != err {
		t.Error(err)
	}
}
