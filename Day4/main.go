package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	sequence, boards, err := ParseFile("in.txt")
	if err != nil {
		panic(err)
	}

	for _, board := range boards {
		if len(*board) != MAX_LENGTH {
			panic("max length problem")
		}
		for _, row := range *board {
			if len(row) != MAX_LENGTH {
				panic("max length problem")
			}
		}
	}

	hasWinner := false
	for _, num := range sequence {
		for _, board := range boards {
			board.Select(num)
			if board.HasBingo() {
				hasWinner = true
				fmt.Println(*board)
				fmt.Println(board.CalcPoints() * num)
				break
			}
		}
		if hasWinner {
			break
		}
	}

}

func ParseFile(filename string) ([]int, []*Board, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	lines := strings.Split(string(bytes), "\r\n")

	if len(lines) < 7 {
		return nil, nil, errors.New("not enough lines in input")
	}

	ints := GetIntArray(lines[0])

	boardInput := lines[2:]
	boards := GenerateBoards(boardInput)

	return ints, boards, nil
}

func GetIntArray(input string) []int {
	vals := strings.Split(input, ",")
	output := make([]int, 0, len(vals))
	for _, val := range vals {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic("failed to parse num")
		}
		output = append(output, num)
	}
	return output
}

func GenerateBoards(input []string) []*Board {
	boards := []*Board{}
	boardStrings := make([]string, 5)
	i := 0
	for _, str := range input {
		if str == "" {
			i = 0
			boards = append(boards, NewBoard(boardStrings))
			boardStrings = make([]string, 5)
			continue
		}
		boardStrings[i] = str
		i++
	}
	boards = append(boards, NewBoard(boardStrings))
	return boards
}
