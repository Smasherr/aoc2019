package impl

import (
	"fmt"
	"strconv"
	"strings"
)

// Main13 solves Day13
func Main13() {
	program := ReadProgram("../res/13_care_package.txt")
	painter := newPainter()
	ic := NewIntcomp(program, &painter, &painter)
	ic.ProcessInstructions()
	fmt.Println(painter.blocks)
	program[0] = 2
	ic.ProcessInstructions()
	fmt.Println(painter.score)
}

type painter struct {
	paintedPoints map[Point2D]int
	writePointer  int
	currX         int
	currY         int
	blocks        int
	joystick      int
	ball          Point2D
	score         int
}

func newPainter() painter {
	return painter{make(map[Point2D]int), 0, 0, 0, 0, 0, *new(Point2D), 0}
}

func (r *painter) Write(data []byte) (int, error) {
	val, _ := strconv.Atoi(strings.TrimSpace(string(data)))
	switch r.writePointer {
	case 0:
		r.currX = val
	case 1:
		r.currY = val
	case 2:
		if r.currX == -1 {
			r.score = val
		} else {
			if val == 2 {
				r.blocks++
			}
			p := NewPoint2D(r.currX, r.currY)
			r.paintedPoints[p] = val
			if val == 4 {
				if p.X > r.ball.X {
					r.joystick = 1
				} else if p.X < r.ball.X {
					r.joystick = -1
				} else {
					r.joystick = 0
				}
				r.ball = p
			}
		}
	}
	r.writePointer++
	r.writePointer %= 3
	return len(data), nil
}

func (r *painter) Read(data []byte) (int, error) {
	b := []byte(strconv.Itoa(r.joystick))
	b = append(b, '\n')
	return copy(data, b), nil
}
