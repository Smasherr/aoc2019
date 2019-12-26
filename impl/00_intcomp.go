package impl

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Intcomp represents an int computer initially introduced on day 2 https://adventofcode.com/2019/day/2
type Intcomp struct {
	relativeBase int
	memory       []int
	in           io.Reader
	out          io.Writer
	instruction  []int
}

// NewIntcomp constructs an Intcomp
func NewIntcomp(memory []int, in io.Reader, out io.Writer) Intcomp {
	return Intcomp{0, memory, in, out, nil}
}

// ProcessInstructions sequentially processes instructions and their parameters in memory
func (ic Intcomp) ProcessInstructions() []int {
	step := 0
	for i := 0; i < len(ic.memory); i += step {
		ic.instruction = InstructionIntToIntArr(ic.memory[i])
		switch ic.instruction[0] {
		case 1:
			step = 4
			ic.processParameterModes(step-1, i)
			ic.memory[ic.instruction[3]] = ic.memory[ic.instruction[1]] + ic.memory[ic.instruction[2]]
		case 2:
			step = 4
			ic.processParameterModes(step-1, i)
			ic.memory[ic.instruction[3]] = ic.memory[ic.instruction[1]] * ic.memory[ic.instruction[2]]
		case 3:
			step = 2
			ic.processParameterModes(step-1, i)
			scanner := bufio.NewScanner(ic.in)
			scanner.Scan()
			text := scanner.Text()
			if len(text) > 0 {
				i, err := strconv.Atoi(text)
				if err == nil {
					ic.memory[ic.instruction[1]] = i
				} else {
					ic.memory[ic.instruction[1]] = int(text[0])
				}
			} else {
				ic.memory[ic.instruction[1]] = 10
			}
		case 4:
			step = 2
			ic.processParameterModes(step-1, i)
			fmt.Fprintln(ic.out, ic.memory[ic.instruction[1]])
		case 5:
			step = 3
			ic.processParameterModes(step-1, i)
			if ic.memory[ic.instruction[1]] != 0 {
				step = ic.memory[ic.instruction[2]] - i
			}
		case 6:
			step = 3
			ic.processParameterModes(step-1, i)
			if ic.memory[ic.instruction[1]] == 0 {
				step = ic.memory[ic.instruction[2]] - i
			}
		case 7:
			step = 4
			ic.processParameterModes(step-1, i)
			if ic.memory[ic.instruction[1]] < ic.memory[ic.instruction[2]] {
				ic.memory[ic.instruction[3]] = 1
			} else {
				ic.memory[ic.instruction[3]] = 0
			}
		case 8:
			step = 4
			ic.processParameterModes(step-1, i)
			if ic.memory[ic.instruction[1]] == ic.memory[ic.instruction[2]] {
				ic.memory[ic.instruction[3]] = 1
			} else {
				ic.memory[ic.instruction[3]] = 0
			}
		case 9:
			step = 2
			ic.processParameterModes(step-1, i)
			ic.relativeBase += ic.memory[ic.instruction[1]]
		case 99:
			step = 1
			return ic.memory
		}
	}
	return ic.memory
}

func (ic *Intcomp) processParameterModes(paramNumber int, index int) {
	for i := 1; i <= paramNumber; i++ {
		switch ic.instruction[i] {
		case 0:
			ic.instruction[i] = ic.memory[index+i]
		case 1:
			ic.instruction[i] = index + i
		case 2:
			ic.instruction[i] = ic.relativeBase + ic.memory[index+i]
		}
		if ic.instruction[i] >= len(ic.memory) {
			newMemory := make([]int, ic.instruction[i]+1)
			copy(newMemory, ic.memory)
			ic.memory = newMemory
		}
	}
}
