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
	fmt.Println(FFT(lines[0], 100, false)[:8])
	part2 := ""
	for i := 0; i < 10000; i++ {
		part2 += lines[0]
	}
	part2 = FFT(part2, 100, true)
	offset, _ := strconv.Atoi(part2[:7])
	fmt.Println(part2[offset : offset+8])
}

var basePattern [4]int = [4]int{0, 1, 0, -1}

// FFT calculates FFT
func FFT(input string, phases int, output bool) string {
	signal := stringToIntArr(input)
	for phase := 0; phase < phases; phase++ {
		if output {
			fmt.Printf("Phase %d\n", phase)
		}
		phaseOutput := make([]int, len(signal))
		for y := range signal {
			for x, i := range signal {
				phaseOutput[y] += i * basePattern[((x+1)/(y+1))%4]
			}
			phaseOutput[y] = int(math.Copysign(float64(phaseOutput[y]), 1))
			phaseOutput[y] %= 10
		}
		signal = phaseOutput
	}
	return strings.Trim(strings.Replace(fmt.Sprint(signal), " ", "", -1), "[]")
}

func negMod(divisor int, modulo int) int {
	if divisor < 0 {
		return modulo + divisor/modulo
	}
	return divisor % modulo
}

func stringToIntArr(s string) []int {
	toRet := make([]int, len(s))
	for i, c := range s {
		toRet[i], _ = strconv.Atoi(string(c))
	}
	return toRet
}
