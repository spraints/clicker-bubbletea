package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type game struct {
	points int
	rate   int
	width  int
	start  time.Time
}

func (g game) Init() tea.Cmd {
	return func() tea.Msg { return time.Now() }
}

type target struct {
	key                string
	points, cost, rate int
}

var targets = []target{
	{
		key:    "a",
		points: 1,
	},
	{
		key:  "s",
		cost: 50,
		rate: 1,
	},
	{
		key:  "d",
		cost: 150,
		rate: 10,
	},
	{
		key:  "f",
		cost: 10000,
		rate: 1000,
	},
}

var targetIndex = func() map[string]target {
	res := make(map[string]target, len(targets))
	for _, t := range targets {
		res[t.key] = t
	}
	return res
}()

func (g game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		g.width = msg.Width
	case time.Time:
		g.start = msg
	case tickMsg:
		g.points += g.rate
		// todo: case tea.MouseMsg, inspect {X, Y, Type}
	case tea.KeyMsg:
		if clicked, ok := targetIndex[msg.String()]; ok {
			if clicked.cost <= g.points {
				g.points += clicked.points
				g.points -= clicked.cost
				g.rate += clicked.rate
			}
		}
	}
	if g.points > winningScore {
		return switchModel(win{time.Since(g.start)})
	}
	return g, nil
}

func (g game) View() string {
	var res strings.Builder
	seconds := time.Since(g.start) / time.Second
	fmt.Fprintf(&res, "points: %8d / rate: %8d / elapsed: %02d:%02d\n", g.points, g.rate, seconds/60, seconds%60)
	for _, target := range targets {
		fmt.Fprintf(&res, "\n+-+\n|")
		if g.points < target.cost {
			res.WriteString(" ")
		} else {
			res.WriteString(target.key)
		}
		fmt.Fprintf(&res, "|\n+-+\n")
	}
	return res.String()
}
