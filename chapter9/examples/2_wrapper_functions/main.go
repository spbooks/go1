package main

import (
	"io/ioutil"
	"os"
)

func main() {
	str := "How do I test this message is being written?"
	loc := "test.txt"
	err := WriteStuff(str, loc)
	if err != nil {
		panic(err)
	}
}

func WriteStuff(str, loc string) error {
	err := writeFile(loc, []byte(str), os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}

var writeFile = ioutil.WriteFile
