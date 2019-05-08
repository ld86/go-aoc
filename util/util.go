package util

import (
	"io/ioutil"
	"os"
)

// ReadInput is getting task's input from file
func ReadInput() []byte {
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	return input
}
