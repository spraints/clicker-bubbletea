package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spraints/clicker-bubbletea/layout"
)

type game struct {
	height int
	width  int
	layout layout.Layout

	points int
	rate   int

	start time.Time
}

func (g game) Init() tea.Cmd {
	start := time.Now()
	layout := layout.New(g.width, g.height, targets)
	return tea.Batch(
		func() tea.Msg { return start },
		func() tea.Msg { return layout },
	)
}

var targets = []layout.Target{
	{
		Key:    "a",
		Points: 1,
	},
	{
		Key:  "s",
		Cost: 50,
		Rate: 1,
	},
	{
		Key:  "d",
		Cost: 150,
		Rate: 10,
	},
	{
		Key:  "f",
		Cost: 10000,
		Rate: 1000,
	},
}

var targetIndex = func() map[string]layout.Target {
	res := make(map[string]layout.Target, len(targets))
	for _, t := range targets {
		res[t.Key] = t
	}
	return res
}()

func (g game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		g.width = msg.Width
		g.height = msg.Height
		g.layout = layout.New(msg.Width-2, msg.Height, targets)
	case time.Time:
		g.start = msg
	case layout.Layout:
		g.layout = msg
	case tickMsg:
		g.points += g.rate
		// todo: case tea.MouseMsg, inspect {X, Y, Type}
	case tea.KeyMsg:
		if clicked, ok := targetIndex[msg.String()]; ok {
			if clicked.Cost <= g.points {
				g.points += clicked.Points
				g.points -= clicked.Cost
				g.rate += clicked.Rate
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
	res.Grow(g.height * g.width)
	seconds := time.Since(g.start) / time.Second
	fmt.Fprintf(&res, "points: %8d / rate: %8d / elapsed: %02d:%02d\n\n", g.points, g.rate, seconds/60, seconds%60)
	g.layout.Render(&res, g.points)
	return res.String()
}
