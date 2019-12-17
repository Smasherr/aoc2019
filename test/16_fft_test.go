package test

import (
	"testing"

	"github.com/Smasherr/aoc2019/impl"
	"github.com/stretchr/testify/assert"
)

func TestFFT1(t *testing.T) {
	assert.EqualValues(t, "01029498", impl.FFT("12345678", 4))
}

func TestFFT2(t *testing.T) {
	assert.EqualValues(t, "24176176", impl.FFT("80871224585914546619083218645595", 100)[:8])
}

func TestFFT3(t *testing.T) {
	assert.EqualValues(t, "73745418", impl.FFT("19617804207202209144916044189917", 100)[:8])
}

func TestFFT4(t *testing.T) {
	assert.EqualValues(t, "52432133", impl.FFT("69317163492948606335995924319873", 100)[:8])
}
