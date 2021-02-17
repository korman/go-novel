package book

type Book interface {
	BookInfo() BookInfo
	Chapters() []Node
}

type BookInfo interface {
	BookName() string
	Auther() string
	Version() string
}
