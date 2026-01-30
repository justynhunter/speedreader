package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/justynhunter/speedreader/lib"
)

type UiModel struct {
	DelayInMs     int
	WordProcessor lib.WordProcessor
}

type tickMsg = time.Time

func (m UiModel) Init() tea.Cmd {
	return tick(m.DelayInMs)
}

func (m UiModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg:
		eof := m.WordProcessor.Next()
		if eof {
			return m, tea.Quit
		}
		return m, tick(m.DelayInMs)
	}

	return m, nil
}

func (m UiModel) View() string {
	return m.WordProcessor.CurrentWord
}

func tick(delayInMs int) tea.Cmd {
	return tea.Tick(time.Millisecond*time.Duration(delayInMs), func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
