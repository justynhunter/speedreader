package main

import (
	"github.com/justynhunter/speedreader/lib"
)

func main() {
	input, err := lib.ReadInput()
	if err != nil {
		panic("error reading input from stdin")
	}

	lib.ProcessInput(input, 100)
}
