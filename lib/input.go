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

func ReadFile(file string) (*WordProcessor, error) {
	contentBytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return MakeTextProcessor(string(contentBytes)), nil
}
