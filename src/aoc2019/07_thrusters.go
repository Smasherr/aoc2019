package aoc2019

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func Main7() {
	input := ReadProgram("../../res/07_thrusters.txt")
	fmt.Println(CalculateMaxThrusterSignal(input))
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CalculateMaxThrusterSignal(input []int) (int, []int) {
	max := 0
	var mp []int
	for i := 0; i <= 44444; i++ {
		s := fmt.Sprintf("%05d", i)
		split := strings.Split(s, "")
		phases := make([]int, 5)
		invalid := false
		for k := 0; k < 5; k++ {
			phase, _ := strconv.Atoi(split[k])
			if phase > 4 || contains(phases[:k], phase) {
				invalid = true
				break
			}
			phases[k] = phase
		}
		if invalid {
			continue
		}
		o := CalculateThrusterSignal(input, phases)
		if o > max {
			max = o
			mp = phases
		}
	}
	return max, mp
}

func CalculateThrusterSignal(i []int, p []int) int {
	toRet := 0
	for k := 0; k < 5; k++ {
		sw := NewStaticReader([]int{p[k], toRet})
		var buf bytes.Buffer
		ProcessInstructions(i, &sw, &buf)
		toRet, _ = strconv.Atoi(strings.TrimSpace(buf.String()))
	}
	return toRet
}
