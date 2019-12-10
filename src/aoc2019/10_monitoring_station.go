package aoc2019

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func Main10() {
	lines, _ := ReadLines("../../res/10_monitoring_station.txt")
	fmt.Println(FindBestLocation(lines))
}

type Asteroid struct {
	X int
	Y int
}

type asteroidRelation struct {
	a Asteroid
	d float64
}

// FindBestLocation Asteroid map as string slice and calculates the best location for a new monitoring station
func FindBestLocation(amap []string) (Asteroid, int, Asteroid) {
	var asteroids []Asteroid
	for y := 0; y < len(amap); y++ {
		split := strings.Split(amap[y], "")
		for x := 0; x < len(split); x++ {
			if split[x] == "#" {
				asteroids = append(asteroids, Asteroid{x, y})
			}
		}
	}
	asteroidsAngles := make(map[Asteroid]map[float64][]asteroidRelation)
	for _, a1 := range asteroids {
		asteroidsAngles[a1] = make(map[float64][]asteroidRelation)
		for _, a2 := range asteroids {
			if a1 == a2 {
				continue
			}
			angle := math.Atan2(float64(a2.Y-a1.Y), float64(a2.X-a1.X))*180/math.Pi + 450
			angle = float64(int(math.Round(angle*100))%36000) / 100
			distance12 := math.Sqrt(math.Pow(float64(a2.X-a1.X), 2) + math.Pow(float64(a2.Y-a1.Y), 2))

			//fmt.Printf("Asteroid 1: %d,%d; Asteroid 2: %d,%d; Angle: %f; Distance: %f\n", a1.X, a1.Y, a2.X, a2.Y, angle, distance)
			asteroidsAngles[a1][angle] = append(asteroidsAngles[a1][angle], asteroidRelation{a2, distance12})
			sort.Slice(asteroidsAngles[a1][angle], func(i, j int) bool {
				return asteroidsAngles[a1][angle][i].d < asteroidsAngles[a1][angle][j].d
			})
		}
	}

	max := 0
	var best Asteroid
	var twoHundredth Asteroid
	for asteroid1, angles := range asteroidsAngles {
		if len(angles) > max {
			max = len(angles)
			best = asteroid1
		}
	}

	station := asteroidsAngles[best]
	keys := make([]float64, 0, len(station))
	for angle := range station {
		keys = append(keys, angle)
	}
	sort.Float64s(keys)
	for i, k := 0, 0; i < 200; i++ {
		angle := keys[k]
		ar := station[angle]
		if len(ar) > 0 {
			twoHundredth = ar[0].a
			station[angle] = ar[1:]
		}
		k++
		k %= len(keys)
	}
	return best, max, twoHundredth
}
