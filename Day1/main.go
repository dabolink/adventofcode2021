package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	values, err := ReadFromInput("in.txt")
	if err != nil {
		panic(err)
	}
	increments := ThreeMeasurementValues(values)
	fmt.Println("Number of increments: ", increments)
}

func ReadFromInput(filename string) ([]int, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\r\n")
	values := make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		num, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		values = append(values, num)
	}
	return values, nil
}
