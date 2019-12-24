package impl

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tb "github.com/nsf/termbox-go"
)

var path []Point2D
var minutes int
var r bool

// Main15 solves Day15
func Main15() {
	tb.Init()
	defer tb.Close()
	program := ReadProgram("../res/15_oxygen_system.txt")
	repairDroid := newRepairDroid()
	ic := NewIntcomp(program, repairDroid, repairDroid)
	ic.ProcessInstructions()
	fillWithOxygen()
	tb.PollEvent()
}

type repairDroid struct {
	position  *Point2D
	command   int
	output    int
	commandCh chan int
	outputCh  chan int
	done      bool
}

func newRepairDroid() *repairDroid {
	theMap = make(map[Point2D]int)
	theMap[*new(Point2D)] = 1
	commandCh := make(chan int)
	outputCh := make(chan int)
	repairDroid := repairDroid{new(Point2D), 0, 0, commandCh, outputCh, false}
	go findSolution(&repairDroid)
	return &repairDroid
}

func findSolution(d *repairDroid) {
	skipReadFromChannel := true
	possibleWays := make(map[Point2D][]int)

	path = append(path, *new(Point2D))

	visited := make(map[Point2D]struct{})
	visited[*new(Point2D)] = struct{}{}

	explorePointer := 0
	possiblyNeedToRevert := false
	for skipReadFromChannel || <-d.outputCh != 2 && len(path) != 0 {
		neighbours := []Point2D{NewPoint2D(d.position.X, d.position.Y-1), NewPoint2D(d.position.X-1, d.position.Y),
			NewPoint2D(d.position.X, d.position.Y+1), NewPoint2D(d.position.X+1, d.position.Y)}
		if explorePointer < 5 {
			r = false
			if d.output == 1 && possiblyNeedToRevert {
				previousPosition := neighbours[directionToIndex(revert(d.command))]
				possibleWays[previousPosition] = append(possibleWays[previousPosition], d.command)
				skipReadFromChannel = false
				possiblyNeedToRevert = false
				d.commandCh <- revert(d.command)
			} else if explorePointer < 4 {
				node := neighbours[explorePointer]
				_, knownPoint := theMap[node]
				if knownPoint {
					explorePointer++
					skipReadFromChannel = true
					continue
				} else {
					skipReadFromChannel = false
					possiblyNeedToRevert = true
					d.commandCh <- indexToDirection(explorePointer)
					explorePointer++
				}
			} else {
				explorePointer++
				skipReadFromChannel = true
			}
		} else {
			r = true
			foundOutlet := false
			pw := possibleWays[*d.position]
			i := 0
			var point Point2D
			for ; !foundOutlet && i < len(pw); i++ {
				point = neighbours[directionToIndex(pw[i])]
				_, visitedPoint := visited[point]
				if !visitedPoint {
					foundOutlet = true
				}
			}
			if foundOutlet {
				path = append(path, point)
				visited[point] = struct{}{}
				d.commandCh <- pointToDirection(neighbours, &point)
			} else {
				back := path[len(path)-2]
				for i := 0; i < len(neighbours); i++ {
					if neighbours[i] == back {
						path = path[:len(path)-1]
						d.commandCh <- indexToDirection(i)
						break
					}
				}
			}
			explorePointer = 0
			possiblyNeedToRevert = false
			skipReadFromChannel = false
		}
	}
	d.done = true
	r = true
}

func pointToDirection(neighbours []Point2D, p *Point2D) int {
	for i := 0; i < len(neighbours); i++ {
		if neighbours[i] == *p {
			return indexToDirection(i)
		}
	}
	panic("Error")
}

func revert(i int) int {
	return [4]int{2, 1, 4, 3}[i-1]
}

func indexToDirection(i int) int {
	return [4]int{1, 3, 2, 4}[i]
}

func directionToIndex(i int) int {
	return [4]int{0, 2, 1, 3}[i-1]
}

func (d *repairDroid) Write(data []byte) (int, error) {
	d.output, _ = strconv.Atoi(strings.TrimSpace(string(data)))
	switch d.output {
	case 0:
		// wall
		switch d.command {
		case 1:
			// north
			theMap[NewPoint2D(d.position.X, d.position.Y-1)] = d.output
		case 2:
			// south
			theMap[NewPoint2D(d.position.X, d.position.Y+1)] = d.output
		case 3:
			// west
			theMap[NewPoint2D(d.position.X-1, d.position.Y)] = d.output
		case 4:
			// east
			theMap[NewPoint2D(d.position.X+1, d.position.Y)] = d.output
		}
	case 1, 2:
		// one step
		switch d.command {
		case 1:
			// north
			d.position.Y--
		case 2:
			// south
			d.position.Y++
		case 3:
			// west
			d.position.X--
		case 4:
			// east
			d.position.X++
		}
		_, exists := theMap[*d.position]
		if !exists {
			theMap[*d.position] = d.output
		}
	}
	if !d.done {
		d.outputCh <- d.output
	} else {
		close(d.outputCh)
		close(d.commandCh)
	}
	if r {
		render(d)
	}
	return len(data), nil
}

func (d *repairDroid) Read(data []byte) (int, error) {
	if d.output != 2 {
		d.command = <-d.commandCh
	}
	b := []byte(strconv.Itoa(d.command))
	b = append(b, '\n')
	return copy(data, b), nil
}

func render(d *repairDroid) {
	for y := 0; y <= 40; y++ {
		for x := 0; x <= 40; x++ {
			p := NewPoint2D(x-21, y-21)
			v, ok := theMap[p]
			if ok {
				switch v {
				case 0:
					tb.SetCell(x, y, '#', tb.ColorDefault, tb.ColorDefault)
				case 1:
					if pathContains(&p) {
						tb.SetCell(x, y, ' ', tb.ColorDefault, tb.ColorRed)
					} else {
						tb.SetCell(x, y, '.', tb.ColorDefault, tb.ColorDefault)
					}
				case 2:
					tb.SetCell(x, y, 'O', tb.ColorCyan, tb.ColorDefault)
				case 3:
					tb.SetCell(x, y, ' ', tb.ColorDefault, tb.ColorBlue)
				}
			}
		}
	}
	if d != nil {
		if theMap[*d.position] != 2 {
			tb.SetCell(d.position.X+21, d.position.Y+21, 'D', tb.ColorDefault, tb.ColorDefault)
		}
		print(0, 41, fmt.Sprintf("X: %2d, Y: %2d, Path: %3d", d.position.X+21, d.position.Y+21, len(path)))
	} else {
		print(23, 41, fmt.Sprintf(", Minutes: %2d", minutes))
	}
	time.Sleep(25 * time.Millisecond)
	tb.Flush()
}

func pathContains(p *Point2D) bool {
	contains := false
	for _, node := range path {
		if node == *p {
			contains = true
			break
		}
	}
	return contains
}
func print(x, y int, s string) {
	for _, r := range s {
		c := tb.ColorDefault
		tb.SetCell(x, y, r, tb.ColorDefault, c)
		x++
	}
}

func fillWithOxygen() {
	start := NewPoint2D(35-21, 35-21)
	waves := []Point2D{start}
	for minutes = 0; len(waves) > 0; minutes++ {
		var newWave []Point2D
		for len(waves) > 0 {
			w := waves[0]
			theMap[w] = 3
			waves = waves[1:]
			newWave = append(newWave, findOutlets([]Point2D{NewPoint2D(w.X, w.Y-1), NewPoint2D(w.X-1, w.Y),
				NewPoint2D(w.X, w.Y+1), NewPoint2D(w.X+1, w.Y)})...)
		}
		waves = newWave
		render(nil)
	}
}

func findOutlets(n []Point2D) []Point2D {
	var outlets []Point2D
	for _, p := range n {
		if theMap[p] == 1 {
			outlets = append(outlets, p)
		}
	}
	return outlets
}
