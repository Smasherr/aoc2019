package impl

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

// Main16 solves Day16
func Main16() {
	lines, _ := ReadLines("../res/16_fft.txt")
	fmt.Println(FFT(lines[0], 100, 0))
	part2 := ""
	for i := 0; i < 10000; i++ {
		part2 += lines[0]
	}
	offset, _ := strconv.Atoi(part2[:7])
	part2 = FFT(part2, 100, offset)
	fmt.Println(part2)
}

var basePattern [4]int = [4]int{0, 1, 0, -1}

// FFT calculates FFT
func FFT(input string, phases int, offset int) string {
	signal := stringToIntArr(input[offset:])
	if offset < len(input)/2 {
		for phase := 0; phase < phases; phase++ {
			phaseOutput := make([]int, len(signal))
			for y := 0; y < len(signal); y++ {
				for x := y; x < len(signal); x++ {
					a := signal[x]
					b := basePattern[((x+offset+1)/(y+offset+1))%4]
					phaseOutput[y] += a * b
				}
				phaseOutput[y] = int(math.Abs(float64(phaseOutput[y])))
				phaseOutput[y] %= 10
			}
			signal = phaseOutput
		}
	} else {
		for phase := 0; phase < phases; phase++ {
			phaseOutput := make([]int, len(signal))
			phaseOutput[len(signal)-1] = signal[len(signal)-1]
			for y := len(signal) - 2; y >= 0; y-- {
				phaseOutput[y] += signal[y] + phaseOutput[y+1]
				phaseOutput[y] = int(math.Abs(float64(phaseOutput[y])))
				phaseOutput[y] %= 10
			}
			signal = phaseOutput
		}
	}
	return strings.Trim(strings.Replace(fmt.Sprint(signal), " ", "", -1), "[]")[:8]
}

/* func stringToIntArr(s string) []int {
	toRet := make([]int, len(s))
	for i, c := range s {
		toRet[i], _ = strconv.Atoi(string(c))
	}
	return toRet
} */

func stringToIntArr(s string) []int {
	toRet := make([]int, len(s))
	var wg sync.WaitGroup
	wg.Add(4)
	for t := 0; t < 4; t++ {
		start := t * ceil(len(s), 4)
		var stop int
		if t < 3 {
			stop = start + ceil(len(s), 4)
		} else {
			stop = len(s)
		}
		go func() {
			for i := start; i < stop; i++ {
				toRet[i], _ = strconv.Atoi(string(s[i]))
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return toRet
}
