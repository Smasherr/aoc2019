package test

import (
	"fmt"
	"testing"

	"../aoc2019"
	"github.com/stretchr/testify/assert"
)

func TestFindBestLocation(t *testing.T) {
	solutions1 := [5]aoc2019.Asteroid{aoc2019.Asteroid{X: 3, Y: 4}, aoc2019.Asteroid{X: 5, Y: 8},
		aoc2019.Asteroid{X: 1, Y: 2}, aoc2019.Asteroid{X: 6, Y: 3}, aoc2019.Asteroid{X: 11, Y: 13}}
	solutions2 := [5]int{8, 33, 35, 41, 210}
	for i := 0; i < 5; i++ {
		filename := fmt.Sprintf("../../res/10_monitoring_station_test%d.txt", i+1)
		lines, _ := aoc2019.ReadLines(filename)
		test1, test2, _ := aoc2019.FindBestLocation(lines)
		assert.True(t, solutions1[i] == test1)
		assert.EqualValues(t, solutions2[i], test2)
	}
}
