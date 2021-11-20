package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type mainMenu struct{}

func (m mainMenu) Init() tea.Cmd {
	return nil
}

func (m mainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if isKey(msg, "enter") {
		log.Print("starting game!")
		return switchModel(game{})
	}
	return m, nil
}

func (m mainMenu) View() string {
	return (`==== CLICKER ====

Press <enter> to start, <q> to quit.

Type the letters you see to get some points!
`)
}
