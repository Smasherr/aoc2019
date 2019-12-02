package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRocketEquation1(t *testing.T) {
	example1 := RocketEquation1(12)
	assert.Equal(t, 2, example1)

	example2 := RocketEquation1(14)
	assert.Equal(t, 2, example2)

	example3 := RocketEquation1(1969)
	assert.Equal(t, 654, example3)

	example4 := RocketEquation1(100756)
	assert.Equal(t, 33583, example4)
}

func TestRocketEquation2(t *testing.T) {
	example1 := RocketEquation2(14)
	assert.Equal(t, 2, example1)

	example2 := RocketEquation2(1969)
	assert.Equal(t, 966, example2)

	example3 := RocketEquation2(100756)
	assert.Equal(t, 50346, example3)
}

func RocketEquation1(input float64) int {
	toRet := int(math.Floor(input/3) - 2)
	return toRet
}

func RocketEquation2(input float64) int {
	fuel := RocketEquation1(input)
	if fuel <= 0 {
		return 0
	}
	return fuel + RocketEquation2(float64(fuel))
}

func TestMain(*testing.T) {
	input1, _ := readLines("01_rocket_equation_input1.txt")
	result1, result2 := 0, 0
	for _, mass := range input1 {
		f, _ := strconv.ParseFloat(mass, 64)
		result1 += RocketEquation1(f)
		result2 += RocketEquation2(f)
	}
	fmt.Println(result1)
	fmt.Println(result2)
}

func readLines(path string) ([]string, error) {
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
