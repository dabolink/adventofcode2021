package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Square struct {
	Selected bool
	Value    int
}

func (s Square) String() string {
	mark := "x"
	if s.Selected {
		mark = "."
	}
	return fmt.Sprintf("%d%s", s.Value, mark)
}

type Board [][]*Square

func (b Board) String() string {
	s := ""
	for _, row := range b {
		for _, sqr := range row {
			s += fmt.Sprintf("%v ", sqr)
		}
		s += "\n"
	}
	return s
}

const (
	MAX_LENGTH = 5
)

func NewBoard(input []string) *Board {
	board := &Board{}
	for i := 0; i < MAX_LENGTH; i++ {
		*board = append(*board, make([]*Square, 5))
	}
	i, j := 0, 0
	for _, str := range input {
		if i >= MAX_LENGTH {
			panic("col length not 5")
		}
		if str == "" {
			continue
		}
		vals := strings.Split(str, " ")
		for _, val := range vals {
			if val == "" {
				continue
			}
			if j >= MAX_LENGTH {
				panic("col length not 5")
			}
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(errors.New("failed to parse value"))
			}
			(*board)[i][j] = &Square{
				Selected: false,
				Value:    num,
			}
			j++
		}
		j = 0
		i++
	}
	return board
}

func (b Board) HasBingo() bool {
	for _, row := range b {
		currentRowBingo := true
		for _, square := range row {
			if !square.Selected {
				currentRowBingo = false
			}
		}
		if currentRowBingo {
			return true
		}
	}
	for i := 0; i < MAX_LENGTH; i++ {
		currentColBingo := true
		for j := 0; j < MAX_LENGTH; j++ {
			square := b[i][j]
			if !square.Selected {
				currentColBingo = false
			}
		}
		if currentColBingo {
			return true
		}
	}
	return false
}

func (b Board) CalcPoints() int {
	sum := 0
	for _, row := range b {
		for _, square := range row {
			if !square.Selected {
				sum += square.Value
			}
		}
	}
	return sum
}

func (b *Board) Select(value int) {
	for _, row := range *b {
		for _, square := range row {
			if square.Value == value {
				square.Selected = true
			}
		}
	}
}
