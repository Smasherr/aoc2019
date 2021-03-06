package impl

import (
	"os"
)

// Main5 solves Day5
func Main5() {
	memory := ReadProgram("../res/05_test_diagnostics.txt")
	sr := NewReaderWriter([]int{1})
	ic := NewIntcomp(memory, &sr, os.Stdout)
	ic.ProcessInstructions()
}
