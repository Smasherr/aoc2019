package impl

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Main11 solves Day11
func Main11() {
	program := ReadProgram("../res/11_painting_robot.txt")
	r := newRobot()
	ic := NewIntcomp(program, &r, &r)
	ic.ProcessInstructions()
	fmt.Println(len(r.paintedPoints))

	program = ReadProgram("../res/11_painting_robot.txt")
	r = newRobot()
	r.paintedPoints[*new(Point2D)] = 1
	ic = NewIntcomp(program, &r, &r)
	ic.ProcessInstructions()
	maxX, maxY := 0, 0
	for p := range r.paintedPoints {
		maxX = int(math.Max(float64(maxX), float64(p.X)))
		maxY = int(math.Max(float64(maxY), float64(p.Y)))
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if r.paintedPoints[NewPoint2D(x, y)] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("■")
			}
		}
		fmt.Println()
	}
}

type robot struct {
	paintedPoints       map[Point2D]int
	currentPosition     Point2D
	writePointer, angle int
}

func newRobot() robot {
	return robot{make(map[Point2D]int), *new(Point2D), 0, 0}
}

func (r *robot) Read(data []byte) (int, error) {
	b := []byte(strconv.Itoa(r.paintedPoints[r.currentPosition]))
	b = append(b, '\n')
	return copy(data, b), nil
}

func (r *robot) Write(data []byte) (int, error) {
	val, _ := strconv.Atoi(strings.TrimSpace(string(data)))
	switch r.writePointer {
	case 0:
		r.paintedPoints[r.currentPosition] = val
	case 1:
		switch val {
		case 0:
			switch r.angle {
			case 0:
				r.currentPosition.X--
			case 90:
				r.currentPosition.Y--
			case 180:
				r.currentPosition.X++
			case 270:
				r.currentPosition.Y++
			}
			r.angle += 270
		case 1:
			switch r.angle {
			case 0:
				r.currentPosition.X++
			case 90:
				r.currentPosition.Y++
			case 180:
				r.currentPosition.X--
			case 270:
				r.currentPosition.Y--
			}
			r.angle += 90
		}
		r.angle %= 360
	}
	r.writePointer++
	r.writePointer %= 2
	return len(data), nil
}
