package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type mainMenu struct {
	width, height int

	debug bool
	msgs  []tea.Msg
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
		case "d":
			m.debug = !m.debug
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	m.msgs = append(m.msgs, msg)
	return m, nil
}

func (m mainMenu) View() string {
	if m.debug {
		const show = 3
		s := make([]string, 0, show)
		for i := 0; i < show; i++ {
			if i < len(m.msgs) {
				msg := m.msgs[len(m.msgs)-i-1]
				s = append(s, fmt.Sprintf("[%d] %T %v\n%#v\n", i, msg, msg, msg))
			} else {
				s = append(s, "\n\n")
			}
		}
		return strings.Join(s, "")
	}
	return fmt.Sprintf("todo: render inside of %d x %d\n", m.width, m.height)
}
