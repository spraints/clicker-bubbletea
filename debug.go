package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const debugHistorySize = 3

type debug struct {
	show bool
	i    int
	msgs []tea.Msg

	model tea.Model
}

func (d debug) Init() tea.Cmd {
	return d.model.Init()
}

func (d debug) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if len(d.msgs) < debugHistorySize {
		d.msgs = append(d.msgs, msg)
	} else {
		d.msgs[d.i] = msg
		d.i = (d.i + 1) % len(d.msgs)
	}

	if key, ok := msg.(tea.KeyMsg); ok {
		if key.String() == "`" {
			d.show = !d.show
		}
	}

	model, cmd := d.model.Update(msg)
	d.model = model
	return d, cmd
}

func (d debug) View() string {
	if !d.show {
		return d.model.View()
	}

	var res strings.Builder
	for i := 0; i < debugHistorySize; i++ {
		p := (d.i - i - 1 + debugHistorySize) % debugHistorySize
		if p < len(d.msgs) {
			msg := d.msgs[p]
			fmt.Fprintf(&res, "[%d] %T %v\n%#v\n", i, msg, msg, msg)
		} else {
			fmt.Fprintf(&res, "\n\n")
		}
	}
	return res.String()
}
