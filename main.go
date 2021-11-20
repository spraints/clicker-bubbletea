package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := os.Create("clicker.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	// - or -
	// log.SetOutput(io.Discard)

	var program tea.Model = mainMenu{}
	program = window{model: mainMenu{}}
	program = debug{model: program}
	tea.NewProgram(program).Start()
}
