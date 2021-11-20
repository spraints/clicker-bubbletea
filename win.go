package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const winningScore = 1_000_000

type win struct {
	elapsed time.Duration
}

func (w win) Init() tea.Cmd {
	return nil
}

func (w win) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return w, nil
}

func (w win) View() string {
	return fmt.Sprintf("🎉 you won in %v! 🎉", w.elapsed)
}
