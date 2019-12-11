package test

import (
	"strconv"
	"testing"

	. "github.com/Smasherr/aoc2019/impl"
	"github.com/stretchr/testify/assert"
)

func TestProcessInstructionsDay9(t *testing.T) {
	example1 := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	rw1 := NewReaderWriter([]int{})
	ic1 := NewIntcomp(example1, nil, &rw1)
	ic1.ProcessInstructions()
	o1 := make([]int, len(example1))
	for i := 0; i < len(o1); i++ {
		o1[i] = <-rw1.Ch
	}
	AssertEqual(t, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, o1)

	example2 := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	rw2 := NewReaderWriter([]int{})
	ic2 := NewIntcomp(example2, nil, &rw2)
	ic2.ProcessInstructions()
	o2 := <-rw2.Ch
	assert.EqualValues(t, 16, len(strconv.Itoa(o2)))

	example3 := []int{104, 1125899906842624, 99}
	rw3 := NewReaderWriter([]int{})
	ic3 := NewIntcomp(example3, nil, &rw3)
	ic3.ProcessInstructions()
	o3 := <-rw3.Ch
	assert.EqualValues(t, 1125899906842624, o3)
}
