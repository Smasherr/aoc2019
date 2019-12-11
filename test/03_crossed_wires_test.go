package test

import (
	"testing"

	. "github.com/Smasherr/aoc2019/impl"
	"github.com/stretchr/testify/assert"
)

func TestCalculateDistanceToClosestIntersection(t *testing.T) {
	wire11 := []string{"R8", "U5", "L5", "D3"}
	wire12 := []string{"U7", "R6", "D4", "L4"}
	r11, r12 := CalculateDistancesToClosestIntersection(wire11, wire12)
	assert.EqualValues(t, 6, r11)
	assert.EqualValues(t, 30, r12)

	wire21 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	wire22 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}
	r21, r22 := CalculateDistancesToClosestIntersection(wire21, wire22)
	assert.EqualValues(t, 159, r21)
	assert.EqualValues(t, 610, r22)

	wire31 := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	wire32 := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}
	r31, r32 := CalculateDistancesToClosestIntersection(wire31, wire32)
	assert.EqualValues(t, 135, r31)
	assert.EqualValues(t, 410, r32)
}
