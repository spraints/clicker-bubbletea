package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type mainMenu struct {
}

func (m mainMenu) Init() tea.Cmd {
	return nil
}

func (m mainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m mainMenu) View() string {
	return fmt.Sprintf("todo: put a menu here")
}
