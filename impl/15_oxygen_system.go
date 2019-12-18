package impl

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tb "github.com/nsf/termbox-go"
)

var path []Point2D
var render bool

// Main15 solves Day15
func Main15() {
	tb.Init()
	tb.PollEvent()
	defer tb.Close()
	program := ReadProgram("../res/15_oxygen_system.txt")
	repairDroid := newRepairDroid()
	ic := NewIntcomp(program, repairDroid, repairDroid)
	ic.ProcessInstructions()
	tb.PollEvent()
}

type repairDroid struct {
	position  *Point2D
	theMap    map[Point2D]int
	command   int
	output    int
	commandCh chan int
	outputCh  chan int
	done      bool
}

func newRepairDroid() *repairDroid {
	theMap := make(map[Point2D]int)
	theMap[*new(Point2D)] = 1
	commandCh := make(chan int)
	outputCh := make(chan int)
	repairDroid := repairDroid{new(Point2D), theMap, 0, 0, commandCh, outputCh, false}
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
			render = false
			if d.output == 1 && possiblyNeedToRevert {
				previousPosition := neighbours[directionToIndex(revert(d.command))]
				possibleWays[previousPosition] = append(possibleWays[previousPosition], d.command)
				skipReadFromChannel = false
				possiblyNeedToRevert = false
				d.commandCh <- revert(d.command)
			} else if explorePointer < 4 {
				node := neighbours[explorePointer]
				_, knownPoint := d.theMap[node]
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
			render = true
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
	render = true
	neighbours := []Point2D{NewPoint2D(d.position.X, d.position.Y-1), NewPoint2D(d.position.X-1, d.position.Y),
		NewPoint2D(d.position.X, d.position.Y+1), NewPoint2D(d.position.X+1, d.position.Y)}
	back := path[len(path)-1]
	for i := 0; i < len(neighbours); i++ {
		if neighbours[i] == back {
			d.commandCh <- indexToDirection(i)
			break
		}
	}
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
	toRet := 0
	switch i {
	case 1:
		toRet = 2
	case 2:
		toRet = 1
	case 3:
		toRet = 4
	case 4:
		toRet = 3
	}
	return toRet
}

func indexToDirection(i int) int {
	toRet := 0
	switch i {
	case 0:
		toRet = 1
	case 1:
		toRet = 3
	case 2:
		toRet = 2
	case 3:
		toRet = 4
	}
	return toRet
}

func directionToIndex(i int) int {
	toRet := 0
	switch i {
	case 1:
		toRet = 0
	case 2:
		toRet = 2
	case 3:
		toRet = 1
	case 4:
		toRet = 3
	}
	return toRet
}

func (d *repairDroid) Write(data []byte) (int, error) {
	d.output, _ = strconv.Atoi(strings.TrimSpace(string(data)))
	switch d.output {
	case 0:
		// wall
		switch d.command {
		case 1:
			// north
			d.theMap[NewPoint2D(d.position.X, d.position.Y-1)] = d.output
		case 2:
			// south
			d.theMap[NewPoint2D(d.position.X, d.position.Y+1)] = d.output
		case 3:
			// west
			d.theMap[NewPoint2D(d.position.X-1, d.position.Y)] = d.output
		case 4:
			// east
			d.theMap[NewPoint2D(d.position.X+1, d.position.Y)] = d.output
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
		_, exists := d.theMap[*d.position]
		if !exists {
			d.theMap[*d.position] = d.output
		}
	}
	if !d.done {
		d.outputCh <- d.output
	} else {
		close(d.outputCh)
		close(d.commandCh)
	}
	if render {
		d.render()
	}
	return len(data), nil
}

func (d *repairDroid) Read(data []byte) (int, error) {
	d.command = <-d.commandCh
	b := []byte(strconv.Itoa(d.command))
	b = append(b, '\n')
	return copy(data, b), nil
}

func (d *repairDroid) render() {
	for y := 0; y <= 40; y++ {
		for x := 0; x <= 40; x++ {
			p := NewPoint2D(x-21, y-21)
			v, ok := d.theMap[p]
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
				}
			}
		}
	}
	tb.SetCell(d.position.X+21, d.position.Y+21, 'D', tb.ColorDefault, tb.ColorDefault)
	print(0, 41, fmt.Sprintf("X: %2d, Y: %2d, Path: %3d", d.position.X+21, d.position.Y+21, len(path)))
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
