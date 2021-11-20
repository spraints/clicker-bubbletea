package main

import (
	"io"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if logFile := os.Getenv("CLICKER_LOG"); logFile != "" {
		f, err := os.Create(logFile)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(f)
	} else {
		log.SetOutput(io.Discard)
	}

	var model tea.Model = quitter{model: window{model: mainMenu{}}}
	model = debug{model: model}
	program := tea.NewProgram(model, tea.WithMouseCellMotion())
	go tick(program)
	program.Start()
}

func switchModel(model tea.Model) (tea.Model, tea.Cmd) {
	return model, model.Init()
}
