package impl

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/nsf/termbox-go"
	tb "github.com/nsf/termbox-go"
)

// Main25 solves Day25
func Main25() {
	tb.Init()
	defer tb.Close()
	program := ReadProgram("../res/25_cryostasis.txt")
	droid25 := newDroid25()
	ic := NewIntcomp(program, droid25, droid25)
	ic.ProcessInstructions()
	tb.PollEvent()
}

type command int

const (
	north command = iota
	west
	south
	east
	take
	drop
	inv
)

func (c command) String() string {
	names := [...]string{
		"north",
		"west",
		"south",
		"east",
		"take %s",
		"drop %s",
		"inv"}
	return names[c]
}

type droid25 struct {
	readPointer int
	sb          strings.Builder
	regexp      *regexp.Regexp
	position    *Point2D
	command     command
	output      string
}

func newDroid25() *droid25 {
	theMap = make(map[Point2D]int)
	theMap[*new(Point2D)] = 1
	d := new(droid25)
	d.regexp = regexp.MustCompile(`(?ms).*== (.*) ==\n(.*)\n+.*Doors here lead:\n(.*?)\n^$\n(?:Items here:\n(.*)^$)?`)
	d.position = new(Point2D)
	d.command = -1
	return d
}

func (d *droid25) Write(data []byte) (int, error) {
	asciiCode, _ := strconv.Atoi(strings.TrimSpace(string(data)))
	d.output = string(rune(asciiCode))
	d.sb.WriteRune(rune(asciiCode))
	return len(data), nil
}

func (d *droid25) Read(data []byte) (int, error) {
	if d.readPointer == 0 {
		neighbours := []Point2D{NewPoint2D(d.position.X, d.position.Y-1), NewPoint2D(d.position.X-1, d.position.Y),
			NewPoint2D(d.position.X, d.position.Y+1), NewPoint2D(d.position.X+1, d.position.Y)}
		d.parseOutput(&neighbours)
		render25(d)
		for d.command == -1 {
			switch ev := tb.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Ch {
				case 'w':
					d.command = north
					if (theMap[neighbours[d.command]]) != 0 {
						d.position.Y--
					}
				case 'a':
					d.command = west
					if (theMap[neighbours[d.command]]) != 0 {
						d.position.X--
					}
				case 's':
					d.command = south
					if (theMap[neighbours[d.command]]) != 0 {
						d.position.Y++
					}
				case 'd':
					d.command = east
					if (theMap[neighbours[d.command]]) != 0 {
						d.position.X++
					}
				case 'e':
					d.command = take
				case 't':
					d.command = drop
				case 'i':
					d.command = inv
				}
			}
		}
	}
	var b []byte
	if d.readPointer < len(d.command.String()) {
		b = []byte{d.command.String()[d.readPointer]}
		d.readPointer++
	} else {
		b = []byte{10}
		d.readPointer = 0
	}
	b = append(b, '\n')
	return copy(data, b), nil
}

func (d *droid25) parseOutput(neighbours *[]Point2D) {
	output := d.sb.String()
	if strings.Contains(output, "ejected back to the checkpoint") {
		d.position = &(*neighbours)[revert25(d.command)]
	} else {
		matches := d.regexp.FindStringSubmatch(output)
		if len(matches) > 0 {
			directions := matches[3]
			//_ := matches[4]
			for i, n := range *neighbours {
				if strings.Contains(directions, command(i).String()) {
					theMap[n] = 1
				} else {
					theMap[n] = 0
				}
			}
		}
	}
	d.command = -1
}

func revert25(c command) command {
	return [...]command{south, east, north, west}[c]
}

func render25(d *droid25) {
	tb.Clear(tb.ColorDefault, tb.ColorDefault)
	offsetX, offsetY := 150, 20
	for y := -20; y <= 20; y++ {
		for x := -20; x <= 20; x++ {
			p := NewPoint2D(x, y)
			v, ok := theMap[p]
			if ok {
				switch v {
				case 0:
					tb.SetCell(x+offsetX, y+offsetY, '#', tb.ColorDefault, tb.ColorDefault)
				case 1:
					tb.SetCell(x+offsetX, y+offsetY, '.', tb.ColorDefault, tb.ColorDefault)
				}
			}
		}
	}
	tb.SetCell(d.position.X+offsetX, d.position.Y+offsetY, 'D', tb.ColorDefault, tb.ColorDefault)
	print(0, 0, d.sb.String())
	tb.Flush()
	d.sb.Reset()
}
