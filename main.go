package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/justynhunter/speedreader/lib"
	"github.com/justynhunter/speedreader/ui"
)

func main() {
	wordProcessor, err := lib.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(ui.UiModel{DelayInMs: 300, WordProcessor: *wordProcessor}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
