package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Hash(s string) string {
	h := md5.New()
	fmt.Fprint(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Zeros(s string) int {
	zerosCount := 0
	for _, c := range s {
		if c != '0' {
			break
		}
		zerosCount++
	}
	return zerosCount
}

func FirstStar(input string) {
	for i := 0; ; i++ {
		var ss strings.Builder
		fmt.Fprint(&ss, input)
		fmt.Fprint(&ss, i)
		if Zeros(Hash(ss.String())) == 5 {
			fmt.Println(i)
			break
		}
	}
}

func SecondStar(input string) {
	for i := 0; ; i++ {
		var ss strings.Builder
		fmt.Fprint(&ss, input)
		fmt.Fprint(&ss, i)
		if Zeros(Hash(ss.String())) == 6 {
			fmt.Println(i)
			break
		}
	}
}

func main() {
	input := "ckczppom"
	FirstStar(input)
	SecondStar(input)
}
