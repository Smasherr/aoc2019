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
	o11, o12 := aoc2019.CalculateMaxThrusterSignal(i1)
	assert.EqualValues(t, 43210, o11)
	aoc2019.AssertEqual(t, []int{4, 3, 2, 1, 0}, o12)

	i2 := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	o21, o22 := aoc2019.CalculateMaxThrusterSignal(i2)
	assert.EqualValues(t, 54321, o21)
	aoc2019.AssertEqual(t, []int{0, 1, 2, 3, 4}, o22)
}
