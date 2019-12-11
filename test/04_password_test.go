package test

import (
	"testing"

	. "github.com/Smasherr/aoc2019/impl"
	"github.com/stretchr/testify/assert"
)

func TestIsValidPassword(t *testing.T) {
	assert.True(t, IsValidPassword1(PasswordIntToIntArr(122345)))
	assert.True(t, IsValidPassword1(PasswordIntToIntArr(111123)))
	assert.True(t, IsValidPassword1(PasswordIntToIntArr(135699)))
	assert.True(t, IsValidPassword1(PasswordIntToIntArr(111111)))
	assert.False(t, IsValidPassword1(PasswordIntToIntArr(223450)))
	assert.False(t, IsValidPassword1(PasswordIntToIntArr(123789)))

	assert.True(t, IsValidPassword2(PasswordIntToIntArr(111122)))
}

func TestCountValidPasswords1(t *testing.T) {
	count1, count2 := CountValidPasswords(284639, 748759)
	assert.EqualValues(t, 895, count1)
	assert.EqualValues(t, 591, count2)
}
