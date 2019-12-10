package test

import (
	"fmt"
	"testing"

	"../aoc2019"
	"github.com/stretchr/testify/assert"
)

func TestFindBestLocation(t *testing.T) {
	solutions1 := [5][]int{{3, 4}, {5, 8}, {1, 2}, {6, 3}, {11, 13}}
	solutions2 := [5]int{8, 33, 35, 41, 210}
	for i := 0; i < 5; i++ {
		filename := fmt.Sprintf("../../res/10_monitoring_station_test%d", i+1)
		lines, _ := aoc2019.ReadLines(filename)
		test1, test2 := aoc2019.FindBestLocation(lines)
		aoc2019.AssertEqual(t, solutions1[i], test1)
		assert.EqualValues(t, solutions2[i], test2)
	}
}
