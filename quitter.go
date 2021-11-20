package main

import tea "github.com/charmbracelet/bubbletea"

type quitter struct {
	model tea.Model
}

func (q quitter) Init() tea.Cmd {
	return q.model.Init()
}

func (q quitter) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if isKey(msg, "q", "ctrl+c") {
		return q, tea.Quit
	}

	model, cmd := q.model.Update(msg)
	q.model = model
	return q, cmd
}

func (q quitter) View() string {
	return q.model.View()
}
