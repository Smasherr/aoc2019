package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateNumberOfOrbits(t *testing.T) {
	testobjects := make([][]string, 10)

	testobjects[0] = make([]string, 2)
	testobjects[0][0] = "COM"
	testobjects[0][1] = "B"

	testobjects[1] = make([]string, 2)
	testobjects[1][0] = "B"
	testobjects[1][1] = "G"

	testobjects[2] = make([]string, 2)
	testobjects[2][0] = "G"
	testobjects[2][1] = "H"

	testobjects[3] = make([]string, 2)
	testobjects[3][0] = "B"
	testobjects[3][1] = "D"

	testobjects[4] = make([]string, 2)
	testobjects[4][0] = "D"
	testobjects[4][1] = "I"

	testobjects[5] = make([]string, 2)
	testobjects[5][0] = "D"
	testobjects[5][1] = "E"

	testobjects[6] = make([]string, 2)
	testobjects[6][0] = "E"
	testobjects[6][1] = "J"

	testobjects[7] = make([]string, 2)
	testobjects[7][0] = "J"
	testobjects[7][1] = "K"

	testobjects[8] = make([]string, 2)
	testobjects[8][0] = "K"
	testobjects[8][1] = "L"

	testobjects[9] = make([]string, 2)
	testobjects[9][0] = "E"
	testobjects[9][1] = "F"

	assert.EqualValues(t, 42, aoc2019.CalculateNumberOfOrbits(testobjects))
}
