package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	tea.NewProgram(mainMenu{}).Start()
}
