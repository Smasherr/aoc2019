package aoc2019

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const width = 25
const height = 6
const resolution = width * height

func Main8() {
	file, _ := ReadLines("../../res/08_rover_bios.txt")
	imagedata := strings.Split(file[0], "")
	layers := make([][height][width]int, 100)
	zeroi, zerok, zerol := 0, math.MaxInt32, 0
	onei, twoi := 0, 0
	onesandtwos := make([][2]int, 100)
	var image [height][width]int
	for i := 0; i < len(imagedata); i++ {
		l := i / resolution
		h := i / width % height
		w := i % width
		id, _ := strconv.Atoi(imagedata[i])
		layers[l][h][w] = id

		switch id {
		case 0:
			zeroi++
		case 1:
			onei++
		case 2:
			twoi++
		}

		if (i+1)%resolution == 0 {
			if zeroi < zerok {
				zerol = l
				zerok = zeroi
			}
			onesandtwos[l] = [2]int{onei, twoi}
			zeroi, onei, twoi = 0, 0, 0
		}

		if l == 0 || image[h][w] == 2 {
			image[h][w] = id
		}
	}
	fmt.Println(onesandtwos[zerol][0] * onesandtwos[zerol][1])
	for i := 0; i < height; i++ {
		for k := 0; k < width; k++ {
			p := image[i][k]
			switch p {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("■")
			}
		}
		fmt.Println()
	}
}
