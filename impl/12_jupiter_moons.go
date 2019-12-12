package impl

import (
	"fmt"
	"math"
)

const format = "<x=%d, y=%d, z=%d>"

// Main12 solves Day12
func Main12() {
	lines, _ := ReadLines("../res/12_jupiter_moons.txt")
	moons := make([]*Moon, len(lines))
	initState := make([]*Moon, len(lines))
	for i := 0; i < len(lines); i++ {
		moons[i] = new(Moon)
		initState[i] = new(Moon)
		fmt.Sscanf(lines[i], format, &moons[i].Pos[0], &moons[i].Pos[1], &moons[i].Pos[2])
		fmt.Sscanf(lines[i], format, &initState[i].Pos[0], &initState[i].Pos[1], &initState[i].Pos[2])
	}
	for i := 0; i < 1000; i++ {
		for a := 0; a < 3; a++ {
			TimeStep(moons, a)
		}
	}
	fmt.Println(TotalEnergy(moons))
	steps := make(chan int, 3)
	for a := 0; a < 3; a++ {
		go func(axis int) {
			for i := 0; true; i++ {
				TimeStep(moons, axis)
				if StatesEqual(moons, initState, axis) {
					steps <- i + 1001
					return
				}
			}
		}(a)
	}
	fmt.Println(LCM(<-steps, <-steps, <-steps))
}

// StatesEqual checks the equality of two states of the universes for an axis
func StatesEqual(state1 []*Moon, state2 []*Moon, axis int) bool {
	toRet := true
	for i := 0; i < len(state1); i++ {
		if state1[i].Pos[axis] != state2[i].Pos[axis] || state1[i].Vel[axis] != state2[i].Vel[axis] {
			toRet = false
		}
	}
	return toRet
}

// Moon represents a moon with a position and velocity
type Moon struct {
	Pos [3]int
	Vel [3]int
}

// ApplyGravity applies gravity between two moons
func ApplyGravity(m1 *Moon, m2 *Moon, axis int) {
	if m2.Pos[axis] > m1.Pos[axis] {
		m1.Vel[axis]++
		m2.Vel[axis]--
	} else if m2.Pos[axis] < m1.Pos[axis] {
		m1.Vel[axis]--
		m2.Vel[axis]++
	}
}

// ApplyVelocity applies velocity to a moon
func ApplyVelocity(m *Moon, axis int) {
	m.Pos[axis] += m.Vel[axis]
}

// TimeStep applies gravity and velocity
func TimeStep(moons []*Moon, axis int) {
	for i := 0; i < len(moons); i++ {
		for k := i + 1; k < len(moons); k++ {
			ApplyGravity(moons[i], moons[k], axis)
		}
	}
	for _, m := range moons {
		ApplyVelocity(m, axis)
	}
}

// TotalEnergy calculates the total energy as the sum of
// potential energy multiplied by kinetic energy of every moon
func TotalEnergy(moons []*Moon) int {
	totalEnergy := 0
	for _, m := range moons {
		pot := int(math.Abs(float64(m.Pos[0])) + math.Abs(float64(m.Pos[1])) + math.Abs(float64(m.Pos[2])))
		kin := int(math.Abs(float64(m.Vel[0])) + math.Abs(float64(m.Vel[1])) + math.Abs(float64(m.Vel[2])))
		totalEnergy += pot * kin
	}
	return totalEnergy
}
