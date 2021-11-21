package layout

import (
	"log"
	"strings"
)

const (
	// +--------+
	// |  FREE  |
	// |   +1   |
	// +--------+
	//
	// +--------+
	// | -10000 |
	// | 1000/s |
	// +--------+
	// .. and one extra for a space around the buttons.
	buttonWidth  = 10
	buttonHeight = 5
)

func New(width, height int, targets []Target) Layout {
	log.Printf("layout.New(width=%d, height=%d, targets=%#v)", width, height, targets)

	gridWidth := width / buttonWidth
	gridHeight := height / buttonHeight
	if gridWidth*gridHeight < len(targets) {
		log.Printf("... too small to make real buttons")
		return Layout{targets: targets}
	}

	extents := make([]extent, 0, len(targets))
	for i := range targets {
		col := i % gridWidth
		row := i / gridWidth
		extents = append(extents, extent{
			top:    row * buttonHeight,
			left:   col * buttonWidth,
			right:  (col+1)*buttonWidth - 1,
			bottom: (row+1)*buttonHeight - 1,
		})
	}

	field := make([][]rune, height)
	for i := 0; i < height; i++ {
		field[i] = make([]rune, width)
		fill(field[i], ' ')
	}

	return Layout{
		width:   width,
		height:  height,
		targets: targets,
		field:   field,
		extents: extents,
	}
}

func fill(s []rune, r rune) {
	for i := range s {
		s[i] = r
	}
}

type Layout struct {
	width   int
	height  int
	targets []Target
	extents []extent
	field   [][]rune
}

type Target struct {
	Key                string
	Points, Cost, Rate int
}

type extent struct {
	top, left, right, bottom int
}

func (l Layout) Render(b *strings.Builder, points int) {
	if len(l.field) == 0 {
		for _, target := range l.targets {
			if target.Cost <= points {
				b.WriteString(target.Key)
			}
		}
		return
	}

	for i, target := range l.targets {
		if target.Cost <= points {
			l.show(l.extents[i], target)
		} else {
			l.hide(l.extents[i], target)
		}
	}

	for _, line := range l.field {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}
}

func (l Layout) show(e extent, t Target) {
	l.field[e.top][e.left] = []rune(t.Key)[0]
}

func (l Layout) hide(e extent, t Target) {
	l.field[e.top][e.left] = ' '
}
