package main

import (
	"fmt"
	"testing"
)

func TestOpcode(t *testing.T) {
	example1 := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	AssertEqual(t, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, ProcessOpcode(example1))

	example2 := []int{1, 0, 0, 0, 99}
	AssertEqual(t, []int{2, 0, 0, 0, 99}, ProcessOpcode(example2))

	example3 := []int{2, 3, 0, 3, 99}
	AssertEqual(t, []int{2, 3, 0, 6, 99}, ProcessOpcode(example3))

	example4 := []int{2, 4, 4, 5, 99, 0}
	AssertEqual(t, []int{2, 4, 4, 5, 99, 9801}, ProcessOpcode(example4))

	example5 := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	AssertEqual(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, ProcessOpcode(example5))
}

func AssertEqual(t *testing.T, a, b []int) {
	if len(a) != len(b) {
		t.Error("lengths don't match")
		return
	}
	for i, v := range a {
		if v != b[i] {
			t.Errorf("%d != %d at index %d", v, b[i], i)
			return
		}
	}
}

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

func TestMain2(*testing.T) {
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
