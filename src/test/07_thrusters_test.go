package test

import (
	"testing"

	"../aoc2019"

	"github.com/stretchr/testify/assert"
)

func TestCalculateThrusterSignal(t *testing.T) {
	i1 := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	p1 := []int{4, 3, 2, 1, 0}
	assert.EqualValues(t, 43210, aoc2019.CalculateThrusterSignal(i1, p1))
}

func TestCalculateMaxThrusterSignal(t *testing.T) {
	i1 := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	o11, o12 := aoc2019.CalculateMaxThrusterSignal(i1, 1)
	assert.EqualValues(t, 43210, o11)
	aoc2019.AssertEqual(t, []int{4, 3, 2, 1, 0}, o12)

	i2 := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	o21, o22 := aoc2019.CalculateMaxThrusterSignal(i2, 1)
	assert.EqualValues(t, 54321, o21)
	aoc2019.AssertEqual(t, []int{0, 1, 2, 3, 4}, o22)
}

func TestCalculateMaxThrusterSignalFeedback(t *testing.T) {
	i1 := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
	o11, o12 := aoc2019.CalculateMaxThrusterSignal(i1, 2)
	assert.EqualValues(t, 139629729, o11)
	aoc2019.AssertEqual(t, []int{9, 8, 7, 6, 5}, o12)

	i2 := []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
		-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53,
		1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}
	o21, o22 := aoc2019.CalculateMaxThrusterSignal(i2, 2)
	assert.EqualValues(t, 18216, o21)
	aoc2019.AssertEqual(t, []int{9, 7, 8, 5, 6}, o22)
}
