package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type mainMenu struct {
	size tea.WindowSizeMsg
}

func (m mainMenu) Init() tea.Cmd {
	return nil
}

func (m mainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if sz, ok := msg.(tea.WindowSizeMsg); ok {
		m.size = sz
	}
	if isKey(msg, "enter") {
		log.Printf("starting game! (window size = %v)", m.size)
		return switchModel(game{
			width:  m.size.Width,
			height: m.size.Height,
		})
	}
	return m, nil
}

func (m mainMenu) View() string {
	return (`==== CLICKER ====

Press <enter> to start, <q> to quit.

Type the letters you see to get some points!
`)
}
