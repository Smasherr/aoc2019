package impl

import (
	"strconv"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

// Main13 solves Day13
func Main13() {
	program := ReadProgram("../res/13_care_package.txt")
	game := newGame()
	ic := NewIntcomp(program, &game, &game)
	program[0] = 2
	ic.ProcessInstructions()
	game.render()
}

type game struct {
	paintedPoints map[Point2D]int
	writePointer  int
	currX         int
	currY         int
	blocks        int
	joystick      int
	paddle        Point2D
	score         int
}

func newGame() game {
	return game{make(map[Point2D]int), 0, 0, 0, 0, 0, *new(Point2D), 0}
}

func (g *game) Write(data []byte) (int, error) {
	val, _ := strconv.Atoi(strings.TrimSpace(string(data)))
	switch g.writePointer {
	case 0:
		g.currX = val
	case 1:
		g.currY = val
	case 2:
		if g.currX == -1 {
			g.score = val
		} else {
			if val == 0 && g.paintedPoints[NewPoint2D(g.currX, g.currY)] == 2 {
				g.blocks++
			}
			if val == 3 {
				g.paddle = NewPoint2D(g.currX, g.currY)
			}
			p := NewPoint2D(g.currX, g.currY)
			g.paintedPoints[p] = val
			if val == 4 {
				if p.X > g.paddle.X {
					g.joystick = 1
				} else if p.X < g.paddle.X {
					g.joystick = -1
				} else {
					g.joystick = 0
				}
			}
		}
	}
	g.writePointer++
	g.writePointer %= 3
	return len(data), nil
}

func (g *game) Read(data []byte) (int, error) {
	b := []byte(strconv.Itoa(g.joystick))
	b = append(b, '\n')
	g.render()
	return copy(data, b), nil
}

func (g *game) render() {
	tm.Clear()
	tm.Printf("Score: %d\tBlocks: %d\n", g.score, g.blocks)
	for y := 0; y <= 19; y++ {
		for x := 0; x <= 43; x++ {
			switch g.paintedPoints[NewPoint2D(x, y)] {
			case 0:
				tm.Print(" ")
			case 1:
				tm.Print("|")
			case 2:
				tm.Print("#")
			case 3:
				tm.Print("_")
			case 4:
				tm.Print("o")
			}
		}
		tm.Println()
	}
	time.Sleep(25 * time.Millisecond)
	tm.Flush()
}
