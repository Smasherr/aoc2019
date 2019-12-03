package aoc2019

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Main3() {
	lines, _ := ReadLines("../../res/03_crossed_wires.txt")
	wire1 := strings.Split(lines[0], ",")
	wire2 := strings.Split(lines[1], ",")
	fmt.Println(CalculateDistancesToClosestIntersection(wire1, wire2))
}

func CalculateDistancesToClosestIntersection(wire1 []string, wire2 []string) (int, int) {
	positions1 := calculatePositions(wire1)
	positions2 := calculatePositions(wire2)
	intersections := calculateIntersections(positions1, positions2)
	closestIntersectionManhattan, closestIntersectionDelay := int(^uint(0)>>1), int(^uint(0)>>1)
	for _, i := range intersections {
		closestIntersectionManhattan = int(math.Min(float64(closestIntersectionManhattan), float64(calculateManhattan(i))))
		closestIntersectionDelay = int(math.Min(float64(closestIntersectionDelay), float64(i[2]+i[3])))
	}
	return closestIntersectionManhattan, closestIntersectionDelay
}

/*
 return value contains:
  x
  y
  steps for first wire
  steps for second wire
*/
func calculateIntersections(positions1 [][]int, positions2 [][]int) [][]int {
	var intersections [][]int
	for i := 0; i < len(positions1); i++ {
		for k := 0; k < len(positions2); k++ {
			if positions1[i][0] == positions2[k][0] &&
				positions1[i][1] == positions2[k][1] {
				intersections = append(intersections, []int{positions1[i][0], positions1[i][1], i + 1, k + 1})
			}
		}
	}
	return intersections
}

func calculateManhattan(position []int) int {
	return int(math.Abs(float64(position[0])) + math.Abs(float64(position[1])))
}

func calculatePositions(wire []string) [][]int {
	var toRet [][]int
	currentPosition := []int{0, 0}
	for i := 0; i < len(wire); i++ {
		runes := []rune(wire[i])
		direction := string(runes[0:1])
		length, _ := strconv.Atoi(string(runes[1:]))
		for k := 0; k < length; k++ {
			slice := make([]int, 2)

			switch direction {
			case "R":
				currentPosition[0]++
			case "U":
				currentPosition[1]++
			case "L":
				currentPosition[0]--
			case "D":
				currentPosition[1]--
			}

			slice[0] = currentPosition[0]
			slice[1] = currentPosition[1]
			toRet = append(toRet, slice)
		}
	}
	return toRet
}
