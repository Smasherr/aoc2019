package test

import (
	"testing"

	. "github.com/Smasherr/aoc2019/impl"
	"github.com/stretchr/testify/assert"
)

func TestInstructionIntToIntArr(t *testing.T) {
	output1 := InstructionIntToIntArr(1002)
	assert.EqualValues(t, output1[0], 2)
	assert.EqualValues(t, output1[1], 0)
	assert.EqualValues(t, output1[2], 1)
	assert.EqualValues(t, output1[3], 0)

	output2 := InstructionIntToIntArr(11002)
	assert.EqualValues(t, output2[0], 2)
	assert.EqualValues(t, output2[1], 0)
	assert.EqualValues(t, output2[2], 1)
	assert.EqualValues(t, output2[3], 1)

	output3 := InstructionIntToIntArr(3)
	assert.EqualValues(t, output3[0], 3)
	assert.EqualValues(t, output3[1], 0)
	assert.EqualValues(t, output3[2], 0)
	assert.EqualValues(t, output3[3], 0)
}
