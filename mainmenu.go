package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type mainMenu struct {
	width, height int
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
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m mainMenu) View() string {
	return fmt.Sprintf("todo: render inside of %d x %d\n", m.width, m.height)
}
