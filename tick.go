package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg struct{}

func tick(program *tea.Program) {
	t := time.NewTicker(time.Second)
	for range t.C {
		program.Send(tickMsg{})
	}
}
