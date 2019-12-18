package test

import (
	"strconv"
	"testing"

	"github.com/Smasherr/aoc2019/impl"
	"github.com/stretchr/testify/assert"
)

func TestFFT1(t *testing.T) {
	assert.EqualValues(t, "01029498", impl.FFT("12345678", 4, 0))
}

func TestFFT2(t *testing.T) {
	assert.EqualValues(t, "24176176", impl.FFT("80871224585914546619083218645595", 100, 0))
}

func TestFFT3(t *testing.T) {
	assert.EqualValues(t, "73745418", impl.FFT("19617804207202209144916044189917", 100, 0))
}

func TestFFT4(t *testing.T) {
	assert.EqualValues(t, "52432133", impl.FFT("69317163492948606335995924319873", 100, 0))
}

func TestFFT5(t *testing.T) {
	lines, _ := impl.ReadLines("../res/16_fft.txt")
	assert.EqualValues(t, "84487724", impl.FFT(lines[0], 100, 0))
}

func TestFFT6(t *testing.T) {
	lines, _ := impl.ReadLines("../res/16_fft.txt")
	part2 := ""
	for i := 0; i < 10000; i++ {
		part2 += lines[0]
	}
	offset, _ := strconv.Atoi(part2[:7])
	assert.EqualValues(t, "84692524", impl.FFT(part2, 100, offset))
}
