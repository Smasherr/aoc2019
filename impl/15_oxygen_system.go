package impl

import (
	"math"
	"strconv"
	"strings"
	"time"

	tm "github.com/buger/goterm"
)

// Main15 solves Day15
func Main15() {
	program := ReadProgram("../res/15_oxygen_system.txt")
	repairDroid := newRepairDroid()
	ic := NewIntcomp(program, &repairDroid, &repairDroid)
	ic.ProcessInstructions()
}

type repairDroid struct {
	position  *Point2D
	theMap    map[Point2D][2]int
	command   int
	output    int
	commandCh chan int
	outputCh  chan int
}

func newRepairDroid() repairDroid {
	theMap := make(map[Point2D][2]int)
	theMap[*new(Point2D)] = [2]int{1, 0}
	commandCh := make(chan int)
	outputCh := make(chan int)
	repairDroid := repairDroid{new(Point2D), theMap, 1, 1, commandCh, outputCh}
	go findSolution(&repairDroid)
	return repairDroid
}

func findSolution(d *repairDroid) {
	d.commandCh <- 1

	var path []Point2D
	path = append(path, *new(Point2D))
	path = append(path, NewPoint2D(0, -1))

	visited := make(map[Point2D]struct{})
	visited[*new(Point2D)] = struct{}{}
	visited[NewPoint2D(0, -1)] = struct{}{}

	for <-d.outputCh != 2 && len(path) != 0 {
		direction := d.theMap[*d.position][1]
		var command int
		possibleWays := []Point2D{NewPoint2D(d.position.X, d.position.Y-1), NewPoint2D(d.position.X-1, d.position.Y),
			NewPoint2D(d.position.X, d.position.Y+1), NewPoint2D(d.position.X+1, d.position.Y)}
		if direction < 4 {

			path = append(path, possibleWays[direction])
			visited[*d.position] = struct{}{}
			d.theMap[*d.position] = [2]int{d.theMap[*d.position][0], direction + 1}
			command = remap(direction)
		} else {
			back := path[len(path)-1]
			i := 1
			for i < 5 && back != possibleWays[i-1] {
				i++
			}
			path = path[:len(path)-1]
			command = remap(i)
		}
		d.commandCh <- command
		/*foundOutlet := false
		possibleWays := []Point2D{NewPoint2D(d.position.X, d.position.Y-1), NewPoint2D(d.position.X-1, d.position.Y),
			NewPoint2D(d.position.X, d.position.Y+1), NewPoint2D(d.position.X+1, d.position.Y)}
		i := 0
		for i < 4 && !foundOutlet {
			i++
			_, visitedPoint := visited[possibleWays[i-1]]
			pointOnTheMap, knownPoint := d.theMap[possibleWays[i-1]]
			if (!knownPoint || pointOnTheMap != 0) && !visitedPoint {
				foundOutlet = true
			}
		}
		if foundOutlet {
			path = append(path, possibleWays[i-1])
						visitedPoint := true
						for _, p := range possibleWays {
							_, knownPoint := d.theMap[p]
							if !knownPoint {
								visitedPoint = false
								break
							}
						}
						if visitedPoint {
			visited[*d.position] = struct{}{}
			}
		} else {
			back := path[len(path)-1]
			i = 1
			for i < 5 && back != possibleWays[i-1] {
				i++
			}
			path = path[:len(path)-1]
		}
		d.commandCh <- i*/
	}
}

func remap(i int) int {
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

func (d *repairDroid) Write(data []byte) (int, error) {
	d.output, _ = strconv.Atoi(strings.TrimSpace(string(data)))
	switch d.output {
	case 0:
		// wall
		switch d.command {
		case 1:
			// north
			d.theMap[NewPoint2D(d.position.X, d.position.Y-1)] = [2]int{d.output, 0}
		case 2:
			// south
			d.theMap[NewPoint2D(d.position.X, d.position.Y+1)] = [2]int{d.output, 0}
		case 3:
			// west
			d.theMap[NewPoint2D(d.position.X-1, d.position.Y)] = [2]int{d.output, 0}
		case 4:
			// east
			d.theMap[NewPoint2D(d.position.X+1, d.position.Y)] = [2]int{d.output, 0}
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
			d.theMap[*d.position] = [2]int{d.output, 0}
		}
	}
	d.render()
	d.outputCh <- d.output
	return len(data), nil
}

func (d *repairDroid) Read(data []byte) (int, error) {
	d.command = <-d.commandCh
	b := []byte(strconv.Itoa(d.command))
	b = append(b, '\n')
	return copy(data, b), nil
}

func (d *repairDroid) render() {
	tm.Clear()
	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, math.MinInt64, math.MinInt64
	for p := range d.theMap {
		minX = int(math.Min(float64(minX), float64(p.X)))
		minY = int(math.Min(float64(minY), float64(p.Y)))
		maxX = int(math.Max(float64(maxX), float64(p.X)))
		maxY = int(math.Max(float64(maxY), float64(p.Y)))
	}
	for y := d.position.Y - 11; y <= d.position.Y+11; y++ {
		for x := d.position.X - 20; x <= d.position.X+20; x++ {
			p := NewPoint2D(x, y)
			if p == *d.position {
				tm.Print("D")
			} else {
				v, ok := d.theMap[p]
				if ok {
					switch v[0] {
					case 0:
						tm.Print("#")
					case 1:
						tm.Print(".")
					case 2:
						tm.Print("O")
					}
				} else {
					tm.Print(" ")
				}
			}
		}
		tm.Println()
	}
	time.Sleep(25 * time.Millisecond)
	tm.Flush()
}
