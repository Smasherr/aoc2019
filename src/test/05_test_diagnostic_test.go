package test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"../aoc2019"
	"github.com/stretchr/testify/assert"
)

func TestProcessInstructions(t *testing.T) {
	example1 := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	aoc2019.AssertEqual(t, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, aoc2019.ProcessInstructions(example1, os.Stdin, os.Stdout))

	example2 := []int{1, 0, 0, 0, 99}
	aoc2019.AssertEqual(t, []int{2, 0, 0, 0, 99}, aoc2019.ProcessInstructions(example2, os.Stdin, os.Stdout))

	example3 := []int{2, 3, 0, 3, 99}
	aoc2019.AssertEqual(t, []int{2, 3, 0, 6, 99}, aoc2019.ProcessInstructions(example3, os.Stdin, os.Stdout))

	example4 := []int{2, 4, 4, 5, 99, 0}
	aoc2019.AssertEqual(t, []int{2, 4, 4, 5, 99, 9801}, aoc2019.ProcessInstructions(example4, os.Stdin, os.Stdout))

	example5 := []int{3, 0, 4, 0, 99}
	reader := Static999Reader{}
	var buf bytes.Buffer
	aoc2019.ProcessInstructions(example5, reader, &buf)
	assert.EqualValues(t, "999", strings.TrimSpace(buf.String()))
}

type Static999Reader struct{}

func (sw Static999Reader) Read(data []byte) (int, error) {
	data[0] = 0x39
	data[1] = 0x39
	data[2] = 0x39
	return 3, io.EOF
}
