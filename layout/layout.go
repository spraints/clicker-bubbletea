package layout

import (
	"fmt"
	"log"
	"math/rand"
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
	buttonWidth  = 9
	buttonHeight = 4
)

func New(width, height int, targets []Target) Layout {
	log.Printf("layout.New(width=%d, height=%d, targets=%#v)", width, height, targets)

	gridWidth := width / (buttonWidth + 1)
	gridHeight := height / (buttonHeight + 1)
	if gridWidth*gridHeight < len(targets) {
		log.Printf("... too small to make real buttons")
		return Layout{targets: targets}
	}

	positions := make([]int, gridWidth*gridHeight)
	for i := 0; i < len(positions); i++ {
		positions[i] = i
	}
	rand.Shuffle(len(positions), func(i, j int) { positions[i], positions[j] = positions[j], positions[i] })

	extents := make([]extent, 0, len(targets))
	for i := range targets {
		row := positions[i] / gridWidth
		col := positions[i] % gridWidth
		extents = append(extents, extent{
			top:    row * (buttonHeight + 1),
			left:   col * (buttonWidth + 1),
			right:  (col+1)*(buttonWidth+1) - 2,
			bottom: (row+1)*(buttonHeight+1) - 2,
		})
	}

	field := make([][]rune, height)
	for i := 0; i < len(field); i++ {
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

type Coords struct {
	Row, Col int
}

func (t Target) Line1() []rune {
	if t.Cost == 0 {
		return []rune("FREE")
	}
	return []rune(fmt.Sprintf("-%d", t.Cost))
}

func (t Target) Line2() []rune {
	if t.Rate > 0 {
		return []rune(fmt.Sprintf("%d/s", t.Rate))
	}
	return []rune(fmt.Sprintf("+%d", t.Points))
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

	for i, line := range l.field {
		if i != 0 {
			b.WriteRune('\n')
		}
		b.WriteString(string(line))
	}
}

func (l Layout) show(e extent, t Target) {
	l.field[e.top][e.left] = '+'
	l.field[e.bottom][e.left] = '+'
	l.field[e.top][e.right] = '+'
	l.field[e.bottom][e.right] = '+'
	for i := 1; i < buttonWidth-1; i++ {
		l.field[e.top][e.left+i] = '-'
		l.field[e.bottom][e.left+i] = '-'
	}

	l.showLine(l.field[e.top+1], e.left, t.Line1())
	l.showLine(l.field[e.top+2], e.left, t.Line2())
}

func (l Layout) showLine(row []rune, col int, s []rune) {
	row[col] = '|'
	row[col+buttonWidth-1] = '|'
	pad := buttonWidth - len(s)
	prePad := pad - pad/2
	for i := 0; i < len(s); i++ {
		row[col+prePad+i] = s[i]
	}
}

func (l Layout) hide(e extent, t Target) {
	for i := e.top; i <= e.bottom; i++ {
		for j := e.left; j <= e.right; j++ {
			l.field[i][j] = ' '
		}
	}
}

func (l Layout) GetTarget(coords Coords) (Target, bool) {
	for i, e := range l.extents {
		if coords.Row < e.top {
			continue
		}
		if coords.Row > e.bottom {
			continue
		}
		if coords.Col < e.left {
			continue
		}
		if coords.Col > e.right {
			continue
		}
		return l.targets[i], true
	}
	return Target{}, false
}
