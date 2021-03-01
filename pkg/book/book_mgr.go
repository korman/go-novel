package book

import (
	"gonovel/internal/utils"
	"path/filepath"
)

func CreateBook(path string) (*Book, error) {
	txt, err := utils.LoadTxtFile(path)

	if nil != err {
		return nil, err
	}

	filename := filepath.Base(path)

	book := new(Book)
	info := new(BookInfo)

	info.BookName = filename
	book.SetBookInfo(info)

	err = book.Load(txt)

	if nil != err {
		return nil, err
	}

	return book, nil
}
