package main

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type window struct {
	width, height int
	model         tea.Model
}

func (w window) Init() tea.Cmd {
	return nil
}

func (w window) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if sz, ok := msg.(tea.WindowSizeMsg); ok {
		w.width = sz.Width
		w.height = sz.Height
	}
	model, cmd := w.model.Update(msg)
	w.model = model
	return w, cmd
}

func (w window) View() string {
	if w.width < 3 || w.height < 3 {
		log.Printf("can't draw in %dx%d", w.width, w.height)
		return ""
	}

	log.Printf("draw in %dx%d", w.width, w.height)

	const (
		// https://en.wikipedia.org/wiki/Box-drawing_character
		topLeft     = 0x250c
		topRight    = 0x2510
		bottomLeft  = 0x2514
		bottomRight = 0x2518
		horizontal  = 0x2500
		vertical    = 0x2502
		space       = ' '
		nl          = '\n'
	)

	lines := strings.Split(w.model.View(), "\n")

	cw := w.width - 2
	ch := w.height - 2
	log.Printf("content: %dx%d", cw, ch)

	var res strings.Builder
	res.Grow(w.width*w.height + 4)

	res.WriteRune(topLeft)
	writeN(&res, horizontal, cw)
	res.WriteRune(topRight)
	res.WriteRune(nl)

	for i := 0; i < ch; i++ {
		res.WriteRune(vertical)
		if off := i; off >= len(lines) {
			writeN(&res, space, cw)
		} else if line := []rune(lines[off]); len(line) >= cw {
			res.WriteString(string(line[:cw]))
		} else {
			res.WriteString(string(line))
			writeN(&res, space, cw-len(line))
		}
		res.WriteRune(vertical)
		res.WriteRune(nl)
	}

	res.WriteRune(bottomLeft)
	if len(lines) > ch {
		writeN(&res, 'x', cw)
	} else {
		writeN(&res, horizontal, cw)
	}
	res.WriteRune(bottomRight)

	return res.String()
}

func writeN(buf *strings.Builder, r rune, count int) {
	for i := 0; i < count; i++ {
		buf.WriteRune(r)
	}
}
