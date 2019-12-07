package aoc2019

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Main5() {
	input := ReadProgram("../../res/05_test_diagnostics.txt")
	sr := NewReaderWriter([]int{1})
	ProcessInstructions(input, &sr, os.Stdout)
}

func ProcessInstructions(input []int, in io.Reader, out io.Writer) []int {
	step := 0
	for i := 0; i < len(input); i += step {
		instruction := InstructionIntToIntArr(input[i])
		switch instruction[0] {
		case 1:
			step = 4
			ProcessParameterModes(&input, &instruction, step-1, i)
			input[instruction[3]] = input[instruction[1]] + input[instruction[2]]
		case 2:
			step = 4
			ProcessParameterModes(&input, &instruction, step-1, i)
			input[instruction[3]] = input[instruction[1]] * input[instruction[2]]
		case 3:
			step = 2
			ProcessParameterModes(&input, &instruction, step-1, i)
			scanner := bufio.NewScanner(in)
			scanner.Scan()
			input[instruction[1]], _ = strconv.Atoi(scanner.Text())
		case 4:
			step = 2
			ProcessParameterModes(&input, &instruction, step-1, i)
			fmt.Fprintln(out, input[instruction[1]])
		case 5:
			step = 3
			ProcessParameterModes(&input, &instruction, step-1, i)
			if input[instruction[1]] != 0 {
				step = input[instruction[2]] - i
			}
		case 6:
			step = 3
			ProcessParameterModes(&input, &instruction, step-1, i)
			if input[instruction[1]] == 0 {
				step = input[instruction[2]] - i
			}
		case 7:
			step = 4
			ProcessParameterModes(&input, &instruction, step-1, i)
			if input[instruction[1]] < input[instruction[2]] {
				input[instruction[3]] = 1
			} else {
				input[instruction[3]] = 0
			}
		case 8:
			step = 4
			ProcessParameterModes(&input, &instruction, step-1, i)
			if input[instruction[1]] == input[instruction[2]] {
				input[instruction[3]] = 1
			} else {
				input[instruction[3]] = 0
			}
		case 99:
			step = 1
			return input
		}
	}
	return input
}

func ProcessParameterModes(input *[]int, instruction *[]int, paramNumber int, index int) {
	for i := 1; i <= paramNumber; i++ {
		switch (*instruction)[i] {
		case 0:
			(*instruction)[i] = (*input)[index+i]
		case 1:
			(*instruction)[i] = index + i
		}
	}
}
