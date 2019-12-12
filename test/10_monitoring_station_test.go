package test

import (
	"fmt"
	"testing"

	. "github.com/Smasherr/aoc2019/impl"
	"github.com/stretchr/testify/assert"
)

func TestFindBestLocation(t *testing.T) {
	solutions1 := [5]Point2D{NewPoint2D(3, 4), NewPoint2D(5, 8), NewPoint2D(1, 2), NewPoint2D(6, 3), NewPoint2D(11, 13)}
	solutions2 := [5]int{8, 33, 35, 41, 210}
	for i := 0; i < 5; i++ {
		filename := fmt.Sprintf("../res/10_monitoring_station_test%d.txt", i+1)
		lines, _ := ReadLines(filename)
		test1, test2, _ := FindBestLocation(lines)
		assert.True(t, solutions1[i] == test1)
		assert.EqualValues(t, solutions2[i], test2)
	}
}
