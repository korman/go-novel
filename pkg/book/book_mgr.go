package book

import "gonovel/internal/utils"

func CreateBook(path string) (*Book, error) {
	txt, err := utils.LoadTxtFile(path)

	if nil != err {
		return nil, err
	}

	book := new(Book)

	err = book.Load(txt)

	if nil != err {
		return nil, err
	}

	return book, nil
}
