package impl

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

// ReadLines from a file
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// AssertEqual checks if 2 slices are equal
func AssertEqual(t *testing.T, a, b []int) {
	if len(a) != len(b) {
		t.Error("lengths don't match")
		return
	}
	for i, v := range a {
		if v != b[i] {
			t.Errorf("%d != %d at index %d", v, b[i], i)
			return
		}
	}
}

// PasswordIntToIntArr converts a 6-digit password to an array with 6 ints
func PasswordIntToIntArr(i int) []int {
	arr := make([]int, 6)
	split := strings.Split(strconv.Itoa(i), "")
	for k := 0; k < 6; k++ {
		arr[k], _ = strconv.Atoi(split[k])
	}
	return arr
}

/*
InstructionIntToIntArr converts an instruction into an array. Positions are used as following:
  0 - opcode
  1 - mode of 1st parameter
  2 - mode of 2nd parameter
  3 - mode of 3rd parameter
*/
func InstructionIntToIntArr(i int) []int {
	arr := make([]int, 4)
	arr[0] = (i % 100)
	for k := 2; k < 5; k++ {
		arr[k-1] = (i / int(math.Pow10(k))) % 10
	}
	return arr
}

/*
ReadProgram reads a file with a program for intcode computer into an int array
*/
func ReadProgram(path string) []int {
	inputText, _ := ReadLines(path)
	inputText = strings.Split(inputText[0], ",")
	toRet := make([]int, len(inputText))
	for i := 0; i < len(inputText); i++ {
		value, _ := strconv.Atoi(inputText[i])
		toRet[i] = value
	}
	return toRet
}

// ReaderWriter implements io.Reader and io.Writer and uses a buffered channel for reading and writing int values
type ReaderWriter struct {
	Name string
	Ch   chan int
}

// NewReaderWriter consctructs a ReaderWriter using some init values inserted into the channel
func NewReaderWriter(init []int) ReaderWriter {
	rw := ReaderWriter{Ch: make(chan int, 100)}
	for _, v := range init {
		rw.Ch <- v
	}
	return rw
}

func (rw *ReaderWriter) Read(data []byte) (int, error) {
	val := <-rw.Ch
	b := []byte(strconv.Itoa(val))
	b = append(b, '\n')
	return copy(data, b), nil
}

func (rw *ReaderWriter) Write(data []byte) (int, error) {
	val, _ := strconv.Atoi(strings.TrimSpace(string(data)))
	rw.Ch <- val
	return len(data), nil
}

// Point is a simple struct with two integers coordinates X and Y
type Point struct {
	X int
	Y int
}

// NewPoint constructs Point
func NewPoint(x int, y int) Point {
	return Point{x, y}
}
