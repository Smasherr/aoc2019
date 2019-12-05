package test

import (
	"testing"

	"../aoc2019"
)

func TestOpcode(t *testing.T) {
	example1 := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	aoc2019.AssertEqual(t, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, aoc2019.ProcessOpcode(example1))

	example2 := []int{1, 0, 0, 0, 99}
	aoc2019.AssertEqual(t, []int{2, 0, 0, 0, 99}, aoc2019.ProcessOpcode(example2))

	example3 := []int{2, 3, 0, 3, 99}
	aoc2019.AssertEqual(t, []int{2, 3, 0, 6, 99}, aoc2019.ProcessOpcode(example3))

	example4 := []int{2, 4, 4, 5, 99, 0}
	aoc2019.AssertEqual(t, []int{2, 4, 4, 5, 99, 9801}, aoc2019.ProcessOpcode(example4))

	example5 := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	aoc2019.AssertEqual(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, aoc2019.ProcessOpcode(example5))
}
