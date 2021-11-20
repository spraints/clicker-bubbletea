package main

import tea "github.com/charmbracelet/bubbletea"

func isKey(msg tea.Msg, keys ...string) bool {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return false
	}
	actual := keyMsg.String()
	for _, key := range keys {
		if key == actual {
			return true
		}
	}
	return false
}
