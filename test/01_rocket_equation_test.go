package test

import (
	"testing"

	. "github.com/Smasherr/aoc2019/impl"

	"github.com/stretchr/testify/assert"
)

func TestRocketEquation1(t *testing.T) {
	example1 := RocketEquation1(12)
	assert.Equal(t, 2, example1)

	example2 := RocketEquation1(14)
	assert.Equal(t, 2, example2)

	example3 := RocketEquation1(1969)
	assert.Equal(t, 654, example3)

	example4 := RocketEquation1(100756)
	assert.Equal(t, 33583, example4)
}

func TestRocketEquation2(t *testing.T) {
	example1 := RocketEquation2(14)
	assert.Equal(t, 2, example1)

	example2 := RocketEquation2(1969)
	assert.Equal(t, 966, example2)

	example3 := RocketEquation2(100756)
	assert.Equal(t, 50346, example3)
}
