package impl

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Main16 solves Day16
func Main16() {
	lines, _ := ReadLines("../res/16_fft.txt")
	fmt.Print(FFT(lines[0], 100)[:8])
}

var pattern [][]int

var basePattern [4]int = [4]int{0, 1, 0, -1}

// FFT calculates FFT
func FFT(input string, phases int) string {
	pattern = make([][]int, len(input))
	for i := 0; i < len(input); i++ {
		pattern[i] = make([]int, len(input))
		for k := 0; k < ceil(len(input), (i+1)*4)*(i+1)*4; k += (i + 1) * 4 {
			for b := 1; b <= 4*(i+1) && k+(i+1)*((b-1)/(i+1))+((b-1)%(i+1)) < len(input); b++ {
				b1 := b % (i + 1)
				b2 := b / (i + 1)
				pattern[i][k+(i+1)*b2+(b1-1)] = basePattern[b2%4]
			}
		}
	}
	signal := stringToIntArr(input)
	for phase := 0; phase < phases; phase++ {
		phaseOutput := make([]int, len(signal))
		for y := range signal {
			for x, i := range signal {
				phaseOutput[y] += i * pattern[y][x]
			}
			phaseOutput[y] = int(math.Copysign(float64(phaseOutput[y]), 1))
			phaseOutput[y] %= 10
		}
		signal = phaseOutput
	}
	return strings.Trim(strings.Replace(fmt.Sprint(signal), " ", "", -1), "[]")
}

func stringToIntArr(s string) []int {
	toRet := make([]int, len(s))
	for i, c := range s {
		toRet[i], _ = strconv.Atoi(string(c))
	}
	return toRet
}
