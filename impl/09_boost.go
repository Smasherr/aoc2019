package impl

import "os"

// Main9 solves Day9
func Main9() {
	memory := ReadProgram("../res/09_boost.txt")
	sr := NewReaderWriter([]int{2})
	ic := NewIntcomp(memory, &sr, os.Stdout)
	ic.ProcessInstructions()
}
