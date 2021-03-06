package test

import (
	"testing"

	. "github.com/Smasherr/aoc2019/impl"
	"github.com/stretchr/testify/assert"
)

func TestProcessInstructions(t *testing.T) {
	example1 := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	ic1 := NewIntcomp(example1, nil, nil)
	ic1.ProcessInstructions()
	AssertEqual(t, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, example1)

	example2 := []int{1, 0, 0, 0, 99}
	ic2 := NewIntcomp(example2, nil, nil)
	ic2.ProcessInstructions()
	AssertEqual(t, []int{2, 0, 0, 0, 99}, example2)

	example3 := []int{2, 3, 0, 3, 99}
	ic3 := NewIntcomp(example3, nil, nil)
	ic3.ProcessInstructions()
	AssertEqual(t, []int{2, 3, 0, 6, 99}, example3)

	example4 := []int{2, 4, 4, 5, 99, 0}
	ic4 := NewIntcomp(example4, nil, nil)
	ic4.ProcessInstructions()
	AssertEqual(t, []int{2, 4, 4, 5, 99, 9801}, example4)

	example5 := []int{3, 0, 4, 0, 99}
	rw := NewReaderWriter([]int{999})
	ic5 := NewIntcomp(example5, &rw, &rw)
	ic5.ProcessInstructions()
	o := <-rw.Ch
	assert.EqualValues(t, 999, o)
}
