package impl

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func Main7() {
	input := ReadProgram("../res/07_thrusters.txt")
	fmt.Println(CalculateMaxThrusterSignal(input, 1))
	fmt.Println(CalculateMaxThrusterSignal(input, 2))
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CalculateMaxThrusterSignal(input []int, mode int) (int, []int) {
	max := 0
	var mp []int
	for i := 0; i <= 99999; i++ {
		s := fmt.Sprintf("%05d", i)
		split := strings.Split(s, "")
		phases := make([]int, 5)
		invalid := false
		for k := 0; k < 5; k++ {
			phase, _ := strconv.Atoi(split[k])
			if (mode == 1 && phase > 4) || (mode == 2 && phase < 5) || contains(phases[:k], phase) {
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

func CalculateThrusterSignal(memory []int, p []int) int {
	rw := make([]ReaderWriter, 5)
	for k := 0; k < 5; k++ {
		rw[k] = NewReaderWriter([]int{p[k]})
		rw[k].Name = fmt.Sprintf("Amplifier %d", k)
	}
	rw[0].Ch <- 0
	var wg sync.WaitGroup
	wg.Add(5)
	for k := 0; k < 5; k++ {
		go func(l int) {
			memcpy := make([]int, len(memory))
			copy(memcpy, memory)
			intcomp := NewIntcomp(memcpy, &rw[l], &rw[(l+1)%5])
			intcomp.ProcessInstructions()
			wg.Done()
		}(k)
	}
	wg.Wait()
	return <-rw[0].Ch
}
