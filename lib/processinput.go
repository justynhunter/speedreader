package lib

import (
	"fmt"
	"strings"
	"time"
)

func ProcessInput(input string, delayInMs int) error {
	words := strings.Fields(input)

	for _, word := range words {
		fmt.Println(word)
		fmt.Println("-")
		time.Sleep(time.Duration(delayInMs) * time.Millisecond)
	}

	return nil
}
