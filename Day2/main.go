package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Direction string

const (
	Forward Direction = "forward"
	Down    Direction = "down"
	Up      Direction = "up"
)

type Command struct {
	Direction Direction
	Distance  int
}

func main() {
	commands, err := ReadInputFile("in.txt")
	if err != nil {
		panic(err)
	}
	output := CalculateCourseWithAim(commands)
	fmt.Println("Distance is:", output)
}

func ReadInputFile(filename string) ([]Command, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\r\n")
	values := make([]Command, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		split := strings.Split(l, " ")
		if len(split) != 2 {
			return nil, errors.New("split not of size 2")
		}
		distance, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		command := Command{
			Direction: Direction(split[0]),
			Distance:  distance,
		}

		values = append(values, command)
	}
	return values, nil
}
