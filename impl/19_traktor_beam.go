package impl

import (
	"fmt"
	"strconv"
	"strings"

	tb "github.com/nsf/termbox-go"
)

// Main19 solves Day19
func Main19() {
	tb.Init()
	theMap = make(map[Point2D]int)
	theMap[*new(Point2D)] = 1
	firstXforY = make(map[int]int)
	program := ReadProgram("../res/19_traktor_beam.txt")
	beamDroid := newBeamDroid()
	ic := NewIntcomp(program, beamDroid, beamDroid)
	done := false
	x, y := 0, 0
	scanWidth = 2000
	for !done && beamDroid.count < scanWidth*scanWidth-1 {
		ic.ProcessInstructions()
		if beamDroid.count%100 == 0 {
			beamDroid.render()
		}
		if beamDroid.position.Y > 500 && lineBreak {
			done, x, y = beamDroid.squareFits()
			lineBreak = false
		}
	}
	for i := y; i < y+100; i++ {
		for k := x; k < x+100; k++ {
			p := NewPoint2D(k, i)
			theMap[p] = 2
		}
	}
	beamDroid.render()
	print(29, 0, fmt.Sprintf(", coordinates for square: %d:%d", x, y))
	tb.Flush()
	tb.PollEvent()
	tb.Close()
}

var scanWidth int
var firstXforY map[int]int
var lineBreak bool

type beamDroid struct {
	count          int
	position       Point2D
	readerPointer  int
	affectedPoints int
	output         int
	affectedOnX    map[int]int
	affectedOnY    map[int]int
}

func newBeamDroid() *beamDroid {
	beamDroid := beamDroid{1, NewPoint2D(1, 0), 0, 1, 0, make(map[int]int), make(map[int]int)}
	return &beamDroid
}

func (d *beamDroid) Write(data []byte) (int, error) {
	d.output, _ = strconv.Atoi(strings.TrimSpace(string(data)))
	theMap[d.position] = d.output
	if d.position.Y < 50 {
		d.affectedPoints += d.output
	}
	d.affectedOnX[d.position.Y] += d.output
	d.affectedOnY[d.position.X] += d.output
	return len(data), nil
}

func (d *beamDroid) Read(data []byte) (int, error) {
	input := -1
	if d.output == 0 && theMap[NewPoint2D(d.position.X-1, d.position.Y)] == 1 {
		firstXforY[d.position.Y] = d.position.X - d.affectedOnX[d.position.Y]
		d.count += scanWidth - 1 - d.affectedOnX[d.position.Y]
		lineBreak = true
	}
	if d.count == scanWidth {
		d.count += scanWidth*3 + 3
	}
	switch d.readerPointer {
	case 0:
		input = d.count % scanWidth
		d.position.X = input
	case 1:
		input = d.count / scanWidth
		d.position.Y = input
		d.count++
	}
	d.readerPointer++
	d.readerPointer %= 2
	b := []byte(strconv.Itoa(input))
	b = append(b, '\n')
	return copy(data, b), nil
}

func (d *beamDroid) render() {
	tb.Clear(tb.ColorDefault, tb.ColorDefault)
	print(0, 0, fmt.Sprintf("Affected points in 50x50: %2d", d.affectedPoints))
	maxX, offsetX, offsetY := firstXforY[d.position.Y-1]+d.affectedOnX[d.position.Y-1], -4, -1
	if maxX > 200 {
		offsetX = maxX - 200 - 4
	}
	if d.position.Y > 100 {
		offsetY = d.position.Y - 100 - 1
	}
	for p, v := range theMap {
		if p.Y < d.position.Y-100 || p.X < maxX-200 {
			continue
		}
		tb.SetCell(p.X-offsetX, p.Y-offsetY, vtor(v), tb.ColorDefault, tb.ColorDefault)
	}
	for y := 1; y <= 101; y++ {
		print(0, y, fmt.Sprintf("%4d ", y+offsetY))
	}
	tb.Flush()
}

func vtor(v int) rune {
	return [3]rune{'.', '#', 'O'}[v]
}

func (d *beamDroid) squareFits() (bool, int, int) {
	x, y := 0, d.position.Y-1
	fits := false
	if d.affectedOnX[y] >= 100 {
		x = firstXforY[y]
		if d.affectedOnY[x+99] >= 100 {
			fits = true
		}
	}
	return fits, x, y - 99
}
