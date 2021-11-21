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

	mouseDown       bool
	mouseDownTarget layout.Target
}

func (g game) Init() tea.Cmd {
	start := time.Now()
	layout := layout.New(g.width, g.height-2, targets)
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
		g.layout = layout.New(msg.Width, msg.Height-2, targets)
	case time.Time:
		g.start = msg
	case layout.Layout:
		g.layout = msg
	case tickMsg:
		g.points += g.rate
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseLeft:
			if !g.mouseDown {
				g.mouseDown = true
				t, ok := g.layout.GetTarget(c(msg))
				if ok {
					g.mouseDownTarget = t
				} else {
					g.mouseDownTarget = layout.Target{}
				}
			}
		case tea.MouseRelease:
			if g.mouseDown {
				t, ok := g.layout.GetTarget(c(msg))
				if ok && t.Key == g.mouseDownTarget.Key {
					g = g.click(t)
				}
			}
			g.mouseDown = false
		}
	case tea.KeyMsg:
		if clicked, ok := targetIndex[msg.String()]; ok {
			g = g.click(clicked)
		}
	}
	if g.points > winningScore {
		return switchModel(win{time.Since(g.start)})
	}
	return g, nil
}

func c(msg tea.MouseMsg) layout.Coords {
	return layout.Coords{
		Row: msg.Y - 2,
		Col: msg.X,
	}
}

func (g game) click(t layout.Target) game {
	if t.Cost <= g.points {
		g.points += t.Points
		g.points -= t.Cost
		g.rate += t.Rate
	}
	return g
}

func (g game) View() string {
	var res strings.Builder
	res.Grow(g.height * g.width)
	seconds := time.Since(g.start) / time.Second
	fmt.Fprintf(&res, "points: %8d / rate: %8d / elapsed: %02d:%02d\n\n", g.points, g.rate, seconds/60, seconds%60)
	g.layout.Render(&res, g.points)
	return res.String()
}
