package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const debugHistorySize = 10

type debug struct {
	show bool
	next int
	msgs []tea.Msg

	model tea.Model
}

func (d debug) Init() tea.Cmd {
	return d.model.Init()
}

func (d debug) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if d.next < debugHistorySize {
		d.msgs = append(d.msgs, msg)
	} else {
		d.msgs[d.next%debugHistorySize] = msg
	}
	d.next++

	if isKey(msg, "`") {
		d.show = !d.show
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
	for i := d.next - debugHistorySize; i < d.next; i++ {
		if i < 0 {
			continue
		}
		msg := d.msgs[i%debugHistorySize]
		fmt.Fprintf(&res, "[%d] %T %v\n%#v\n", i, msg, msg, msg)
	}
	return res.String()
}
