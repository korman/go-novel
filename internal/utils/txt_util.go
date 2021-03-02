package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/pkumza/numcn"
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

func GenNumberFromString(s string) (int64, error) {
	var num int64 = -1

	reg := regexp.MustCompile(`[一二三四五六七八九十零百千两]+`)

	numInfo := reg.FindAllStringIndex(s, 1)

	if nil == numInfo {
		reg = regexp.MustCompile(`[0-9]+`)

		numInfo = reg.FindAllStringIndex(s, 1)

		if nil == numInfo {
			return -1, errors.New("NONE")
		}

		numString := s[numInfo[0][0]:numInfo[0][1]]

		nm, err := strconv.Atoi(numString)

		if nil != err {
			return -1, err
		}

		return int64(nm), nil
	}

	numString := s[numInfo[0][0]:numInfo[0][1]]

	num, err := numcn.DecodeToInt64(numString)

	if nil != err {
		return -1, err
	}

	return num, nil
}

func WriteFile(path string, text string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)

	if nil != err {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(text)

	if nil != err {
		return err
	}

	return nil
}
