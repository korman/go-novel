package utils

import (
	"io/ioutil"
	"os"
)

func LoadTxtFile(path string) (text string, err error) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModeAppend)

	if nil != err {
		return "", err
	}

	context, err := ioutil.ReadAll(f)

	if nil != err {
		return "", err
	}

	s := string(context)

	return s, nil
}
