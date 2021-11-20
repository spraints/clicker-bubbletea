package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var program tea.Model = mainMenu{}
	program = debug{model: program}
	tea.NewProgram(program).Start()
}
