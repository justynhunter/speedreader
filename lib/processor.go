package lib

import (
	"strings"
)

type WordProcessor struct {
	CurrentWord  string
	words        []string
	currentIndex int
}

func (t *WordProcessor) Next() bool {
	t.currentIndex = t.currentIndex + 1
	if t.currentIndex >= len(t.words) {
		t.CurrentWord = ""
		return true
	}

	t.CurrentWord = t.words[t.currentIndex]
	return false
}

func MakeTextProcessor(text string) *WordProcessor {
	return &WordProcessor{
		words:        strings.Fields(text),
		currentIndex: -1,
	}
}
