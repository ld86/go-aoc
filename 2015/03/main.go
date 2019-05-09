package main

import (
	"fmt"

	"github.com/fatih/set"
	"github.com/ld86/go-aoc/util"
)

type Point struct {
	X int
	Y int
}

type Santa struct {
	x int
	y int
	s set.Interface
}

func NewSanta() *Santa {
	s := set.New(set.ThreadSafe)
	s.Add(Point{0, 0})
	return &Santa{0, 0, s}
}

func (santa *Santa) Move(direction byte) {
	x, y := santa.x, santa.y
	switch direction {
	case '>':
		x++
	case '<':
		x--
	case '^':
		y++
	case 'v':
		y--
	}
	santa.x, santa.y = x, y
	santa.s.Add(Point{x, y})
}

func (santa *Santa) Houses() set.Interface {
	return santa.s
}

func FirstStar(input []byte) {
	santa := NewSanta()
	for _, c := range input {
		santa.Move(c)
	}
	fmt.Println(santa.Houses().Size())
}

func SecondStar(input []byte) {
	santa := NewSanta()
	roboSanta := NewSanta()
	for i, c := range input {
		if i%2 == 0 {
			santa.Move(c)
		} else {
			roboSanta.Move(c)
		}
	}
	houses := santa.Houses().Copy()
	houses.Merge(roboSanta.Houses())
	fmt.Println(houses.Size())
}

func main() {
	input := util.ReadInput()
	FirstStar(input)
	SecondStar(input)
}
