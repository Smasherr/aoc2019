package aoc2019

import "fmt"

func ProcessOpcode(input []int) []int {
	step := 0
	for i := 0; i < len(input); i += step {
		if input[i] == 1 {
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
			step = 4
			continue
		}
		if input[i] == 2 {
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
			step = 4
			continue
		}
		if input[i] == 99 {
			step = 1
			return input
		}
	}
	return input
}

func main2() {
	input := []int{1, 0, 0, 3,
		1, 1, 2, 3,
		1, 3, 4, 3,
		1, 5, 0, 3,
		2, 6, 1, 19,
		1, 19, 10, 23,
		2, 13, 23, 27,
		1, 5, 27, 31,
		2, 6, 31, 35,
		1, 6, 35, 39,
		2, 39, 9, 43,
		1, 5, 43, 47,
		1, 13, 47, 51,
		1, 10, 51, 55,
		2, 55, 10, 59,
		2, 10, 59, 63,
		1, 9, 63, 67,
		2, 67, 13, 71,
		1, 71, 6, 75,
		2, 6, 75, 79,
		1, 5, 79, 83,
		2, 83, 9, 87,
		1, 6, 87, 91,
		2, 91, 6, 95,
		1, 95, 6, 99,
		2, 99, 13, 103,
		1, 6, 103, 107,
		1, 2, 107, 111,
		1, 111, 9, 0,
		99,
		2, 14, 0, 0}
	result1 := make([]int, len(input))
	copy(result1, input)
	result1[1] = 12
	result1[2] = 2
	fmt.Println(ProcessOpcode(result1)[0])

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			result2 := make([]int, len(input))
			copy(result2, input)
			result2[1] = noun
			result2[2] = verb
			if ProcessOpcode(result2)[0] == 19690720 {
				fmt.Println(100*noun + verb)
				break
			}
		}
	}
}
