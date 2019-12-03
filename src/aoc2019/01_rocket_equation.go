package aoc2019

import (
	"fmt"
	"math"
	"strconv"
)

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

func main1() {
	input1, _ := ReadLines("01_rocket_equation_input1.txt")
	result1, result2 := 0, 0
	for _, mass := range input1 {
		f, _ := strconv.ParseFloat(mass, 64)
		result1 += RocketEquation1(f)
		result2 += RocketEquation2(f)
	}
	fmt.Println(result1)
	fmt.Println(result2)
}
