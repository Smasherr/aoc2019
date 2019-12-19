package impl

import (
	"fmt"
	"math"
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
	scanWidth = 1000
	for beamDroid.count < scanWidth*scanWidth-1 {
		ic.ProcessInstructions()
		beamDroid.render()
	}
	scanWidth = 1000
	beamDroid = newBeamDroid()
	ic = NewIntcomp(program, beamDroid, beamDroid)
	for i := 0; i < scanWidth*scanWidth-1 || done; i++ {
		ic.ProcessInstructions()
		if beamDroid.position.Y > 500 && lineBreak {
			done, x, y = beamDroid.squareFits()
			lineBreak = false
		}
	}
	fmt.Println(10000*x + y)
	tb.PollEvent()
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
	d.affectedPoints += d.output
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
	print(0, 0, fmt.Sprintf("Affected points: %2d", d.affectedPoints))
	maxX, offsetX, offsetY := 0, 0, -1
	for p, v := range theMap {
		if v == 1 {
			maxX = int(math.Max(float64(maxX), float64(p.X)))
		}
	}
	if maxX > 200 {
		offsetX = maxX - 200
	}
	if d.position.Y > 60 {
		offsetY = d.position.Y - 60 - 1
	}
	for p, v := range theMap {
		if p.Y < d.position.Y-60 || p.X < maxX-200 {
			continue
		}
		tb.SetCell(p.X-offsetX, p.Y-offsetY, vtor(v), tb.ColorDefault, tb.ColorDefault)
	}
	tb.Flush()
}

func vtor(v int) rune {
	return [3]rune{'.', '#', 'O'}[v]
}

func (d *beamDroid) squareFits() (bool, int, int) {
	x, y := 0, d.position.Y-11
	fits := false
	if d.affectedOnX[y-11] >= 100 {
		x = firstXforY[y-11] + 1
		for i := 0; i <= d.affectedOnX[y-11]-100; i++ {
			x += i
			p := NewPoint2D(x, y)
			pVal := theMap[p]
			pVal = pVal
			pDown := theMap[NewPoint2D(p.X, p.Y+99)]
			pDownRight := theMap[NewPoint2D(p.X+99, p.Y+99)]
			if pDown == 1 && pDownRight == 1 {
				fits = true
				x = p.X
				y = p.Y
				break
			}
		}
	}
	return fits, x, y
}
