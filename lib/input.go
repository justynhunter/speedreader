package lib

import (
	"bufio"
	"os"
)

func ReadInput() (*WordProcessor, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input = input + scanner.Text()
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}

	return MakeTextProcessor(input), nil
}
