package impl

import (
	"fmt"
	"strings"
)

// Main6 solves Day6
func Main6() {
	lines, _ := ReadLines("../res/06_uom.txt")
	relationships := make([][]string, len(lines))
	for i, line := range lines {
		relationship := strings.Split(line, ")")
		relationships[i] = relationship
	}
	fmt.Println(Calculate(relationships))
}

// Calculate solutions for part 1 and 2
func Calculate(relationships [][]string) (int, int) {
	numberOfOrbits := 0
	orbits := CalculateOrbits(relationships)
	for _, v := range orbits {
		numberOfOrbits += len(v)
	}
	transfers := CalculateTransfers(&orbits)
	return numberOfOrbits, transfers
}

// CalculateTransfers calculates minimum number of orbital transfers required
// to move from the object YOU are orbiting to the object SAN is orbiting
func CalculateTransfers(orbits *map[string][]string) int {
	you := (*orbits)["YOU"]
	san := (*orbits)["SAN"]
	for i := 0; i < len(you); i++ {
		for k := 0; k < len(san); k++ {
			if you[i] == san[k] {
				return i + k
			}
		}
	}
	return 0
}

// CalculateOrbits calculates the total number of direct and indirect orbits
func CalculateOrbits(relationships [][]string) map[string][]string {
	orbits := make(map[string][]string)
	for _, orbit := range relationships {
		orbits[orbit[1]] = append(orbits[orbit[1]], orbit[0])
	}
	for l, k := range orbits {
		appendIndirect(l, k[0], &orbits)
	}
	return orbits
}

func appendIndirect(l string, k string, orbits *map[string][]string) {
	if k == "COM" {
		return
	}
	root := (*orbits)[k][0]
	(*orbits)[l] = append((*orbits)[l], root)
	appendIndirect(l, root, orbits)
}
