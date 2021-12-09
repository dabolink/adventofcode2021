package main

import (
	"io/ioutil"
	"strings"
)

func main() {

}

func ParseFile(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	_ = strings.Split(string(bytes), "\r\n")
	return nil
}
