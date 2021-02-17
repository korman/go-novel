package book

type Book struct {
	BookInfomation *BookInfo
	Chapters       []*BookNode
}

func (this *Book) init() {
	this.BookInfomation = new(BookInfo)
}

func (this *Book) Load(txt string) error {
	return nil
}
