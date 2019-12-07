package aoc2019

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

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

func PasswordIntToIntArr(i int) []int {
	arr := make([]int, 6)
	split := strings.Split(strconv.Itoa(i), "")
	for k := 0; k < 6; k++ {
		arr[k], _ = strconv.Atoi(split[k])
	}
	return arr
}

func InstructionIntToIntArr(i int) []int {
	arr := make([]int, 4)
	arr[0] = (i % 100)
	for k := 2; k < 5; k++ {
		arr[k-1] = (i / int(math.Pow10(k))) % 10
	}
	return arr
}

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

type ReaderWriter struct {
	Name string
	Ch   chan int
}

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
