package test

import (
	"testing"

	. "github.com/Smasherr/aoc2019/impl"
)

func TestOpcode(t *testing.T) {
	example1 := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	AssertEqual(t, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, NewIntcomp(example1, nil, nil).ProcessInstructions())

	example2 := []int{1, 0, 0, 0, 99}
	AssertEqual(t, []int{2, 0, 0, 0, 99}, NewIntcomp(example2, nil, nil).ProcessInstructions())

	example3 := []int{2, 3, 0, 3, 99}
	AssertEqual(t, []int{2, 3, 0, 6, 99}, NewIntcomp(example3, nil, nil).ProcessInstructions())

	example4 := []int{2, 4, 4, 5, 99, 0}
	AssertEqual(t, []int{2, 4, 4, 5, 99, 9801}, NewIntcomp(example4, nil, nil).ProcessInstructions())

	example5 := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	AssertEqual(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, NewIntcomp(example5, nil, nil).ProcessInstructions())
}
