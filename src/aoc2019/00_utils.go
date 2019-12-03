package aoc2019

import (
	"bufio"
	"os"
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
