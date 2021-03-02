package main

import (
	"flag"
	"fmt"
	"gonovel/pkg/book"
	"time"
)

func main() {
	path := flag.String("path", "tests/txt_files/test_01.txt", "输入路径")
	outpath := flag.String("outpath", ".", "输出书的路径")

	flag.Parse()

	println(fmt.Sprintf("准备加载%s...", *path))

	t := time.Now()
	book, err := book.CreateBook(*path)

	if nil != err {
		println(err)
	}

	if nil == book {

	}

	elapsed := time.Since(t)

	println(fmt.Sprintf("耗时%s秒", elapsed))

	err = book.ConvertToMd(*outpath)

	if nil != err {
		println(err)
	}
}
