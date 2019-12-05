package test

import (
	"testing"

	"../aoc2019"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPassword(t *testing.T) {
	assert.True(t, aoc2019.IsValidPassword1(aoc2019.IntToIntArr(122345)))
	assert.True(t, aoc2019.IsValidPassword1(aoc2019.IntToIntArr(111123)))
	assert.True(t, aoc2019.IsValidPassword1(aoc2019.IntToIntArr(135699)))
	assert.True(t, aoc2019.IsValidPassword1(aoc2019.IntToIntArr(111111)))
	assert.False(t, aoc2019.IsValidPassword1(aoc2019.IntToIntArr(223450)))
	assert.False(t, aoc2019.IsValidPassword1(aoc2019.IntToIntArr(123789)))

	assert.True(t, aoc2019.IsValidPassword2(aoc2019.IntToIntArr(111122)))
}

func TestCountValidPasswords1(t *testing.T) {
	count1, count2 := aoc2019.CountValidPasswords(284639, 748759)
	assert.EqualValues(t, 895, count1)
	assert.EqualValues(t, 591, count2)
}
