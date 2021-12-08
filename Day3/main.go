package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Input []BitArray

func main() {
	inputs, err := ParseFile("in.txt")
	if err != nil {
		panic(err)
	}
	gamma := GetMostCommon(inputs)
	epsilon := GetLeastCommon(inputs)
	fmt.Println(gamma.ToDecimal() * epsilon.ToDecimal())

	if len(inputs) == 0 {
		panic(errors.New("no inputs?"))
	}
	var newInputs Input
	newInputs = inputs

	for i := range inputs[0] {
		mostCommon := GetMostCommon(newInputs)
		newInputs = newInputs.Filter(i, func(index int, ba BitArray) bool {
			return mostCommon[i] == ba[i]
		})
		if len(newInputs) == 1 {
			break
		}
	}

	oxygen := newInputs[0].ToDecimal()

	newInputs = inputs
	for i := range inputs[0] {
		leastCommon := GetLeastCommon(newInputs)
		newInputs = newInputs.Filter(i, func(index int, ba BitArray) bool {
			return leastCommon[i] == ba[i]
		})
		if len(newInputs) == 1 {
			break
		}
	}
	c02 := newInputs[0].ToDecimal()
	fmt.Println(oxygen * c02)
}

func ParseFile(filename string) (Input, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\r\n")
	byteArrays := make([]BitArray, 0, len(lines))
	for _, line := range lines {
		byteArray := NewBitArray(line)
		byteArrays = append(byteArrays, byteArray)
	}
	return byteArrays, nil
}
