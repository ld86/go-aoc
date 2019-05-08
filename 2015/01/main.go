package main

import (
	"fmt"

	"github.com/ld86/go-aoc/util"
)

func FirstStar(input []byte) {
	level := 0
	for _, c := range input {
		switch c {
		case '(':
			level++
		case ')':
			level--
		}
	}
	fmt.Println(level)
}

func SecondStar(input []byte) {
	level := 0
	for i, c := range input {
		switch c {
		case '(':
			level++
		case ')':
			level--
		}
		if level == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}

func main() {
	input := util.ReadInput()
	FirstStar(input)
	SecondStar(input)
}
