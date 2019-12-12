package impl

import (
	"fmt"
	"math"
)

const format = "<x=%d, y=%d, z=%d>"
const steps = 1000

// Main12 solves Day12
func Main12() {
	lines, _ := ReadLines("../res/12_jupiter_moons.txt")
	moons := make([]*Moon, len(lines))
	for i := 0; i < len(lines); i++ {
		moons[i] = new(Moon)
		fmt.Sscanf(lines[i], format, &moons[i].Pos.X, &moons[i].Pos.Y, &moons[i].Pos.Z)
	}
	for i := 0; i < steps; i++ {
		TimeStep(moons)
		fmt.Println(TotalEnergy(moons))
	}
}

// Moon represents a moon with a position and velocity
type Moon struct {
	Pos Point3D
	Vel Point3D
}

// ApplyGravity applies gravity between two moons
func ApplyGravity(m1 *Moon, m2 *Moon) {
	if m2.Pos.X > m1.Pos.X {
		m1.Vel.X++
		m2.Vel.X--
	} else {
		m1.Vel.X--
		m2.Vel.X++
	}
	if m2.Pos.Y > m1.Pos.Y {
		m1.Vel.Y++
		m2.Vel.Y--
	} else {
		m1.Vel.Y--
		m2.Vel.Y++
	}
	if m2.Pos.Z > m1.Pos.Z {
		m1.Vel.Z++
		m2.Vel.Z--
	} else {
		m1.Vel.Z--
		m2.Vel.Z++
	}
}

// ApplyVelocity applies velocity to a moon
func ApplyVelocity(m *Moon) {
	m.Pos.X += m.Vel.X
	m.Pos.Y += m.Vel.Y
	m.Pos.Z += m.Vel.Z
}

// TimeStep applies gravity and velocity
func TimeStep(moons []*Moon) {
	for i := 0; i < len(moons); i++ {
		for k := i + 1; k < len(moons); k++ {
			ApplyGravity(moons[i], moons[k])
		}
	}
	for _, m := range moons {
		ApplyVelocity(m)
	}
}

// TotalEnergy calculates the total energy as the sum of
// potential energy multiplied by kinetic energy of every moon
func TotalEnergy(moons []*Moon) int {
	totalEnergy := 0
	for _, m := range moons {
		pot := int(math.Abs(float64(m.Pos.X)) + math.Abs(float64(m.Pos.Y)) + math.Abs(float64(m.Pos.Z)))
		kin := int(math.Abs(float64(m.Vel.X)) + math.Abs(float64(m.Vel.Y)) + math.Abs(float64(m.Vel.Z)))
		totalEnergy += pot * kin
	}
	return totalEnergy
}
