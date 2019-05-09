package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ld86/go-aoc/util"
)

func FirstStar(input []byte) {
	sum := 0
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		d := strings.Split(line, "x")
		l, _ := strconv.Atoi(d[0])
		w, _ := strconv.Atoi(d[1])
		h, _ := strconv.Atoi(d[2])

		di := []int{l, w, h}
		sort.Ints(di)

		sum += l*w*2 + w*h*2 + l*h*2 + di[0]*di[1]
	}
	fmt.Println(sum)
}

func SecondStar(input []byte) {
	sum := 0
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		d := strings.Split(line, "x")
		l, _ := strconv.Atoi(d[0])
		w, _ := strconv.Atoi(d[1])
		h, _ := strconv.Atoi(d[2])

		di := []int{l, w, h}
		sort.Ints(di)

		sum += di[0]*di[1]*di[2] + di[0]*2 + di[1]*2
	}
	fmt.Println(sum)
}

func main() {
	input := util.ReadInput()
	FirstStar(input)
	SecondStar(input)
}
