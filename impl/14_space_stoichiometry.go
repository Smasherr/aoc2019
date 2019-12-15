package impl

import (
	"fmt"
	"math"
	"strings"
)

// Main14 solves Day14
func Main14() {
	lines, _ := ReadLines("../res/14_space_stoichiometry.txt")
	fmt.Println(CalculateAmountOfOreForFuel(lines))
	fmt.Println(CalculateAmountOfFuel(lines, 720484))
}

type chemical struct {
	amount   int
	material string
}

var reactions map[string][]chemical
var basicChemicals map[string]int
var leftovers map[string]int

// CalculateAmountOfOreForFuel calculates the minimum amount of ORE required to produce exactly 1 FUEL
func CalculateAmountOfOreForFuel(input []string) int {
	prepareData(input)
	return calculateAmountOfOre("FUEL", 1)
}

// CalculateAmountOfFuel calculates the maximum amount of FUEL produced with 1000000000000 units of ORE
func CalculateAmountOfFuel(input []string, orePerFuel int) int {
	prepareData(input)
	ore := 1000000000000
	fuel := ore / orePerFuel
	ore -= calculateAmountOfOre("FUEL", fuel)
	for ; ore > 0; fuel++ {
		ore -= calculateAmountOfOre("FUEL", 1)
	}
	return fuel - 1
}

func prepareData(input []string) {
	reactions = make(map[string][]chemical)
	basicChemicals = make(map[string]int)
	leftovers = make(map[string]int)
	var materialList, inputMaterial, productMaterial string
	var inputAmount, productAmount int
	for _, s := range input {
		matches := strings.Split(s, " => ")
		materialList = matches[0]
		fmt.Sscanf(matches[1], "%d %s", &productAmount, &productMaterial)
		product := chemical{productAmount, productMaterial}
		materials := strings.Split(materialList, ", ")
		chemicals := make([]chemical, len(materials)+1)
		chemicals[0] = product
		for i, entry := range materials {
			fmt.Sscanf(entry, "%d %s", &inputAmount, &inputMaterial)
			chemicals[i+1] = chemical{inputAmount, inputMaterial}
			if inputMaterial == "ORE" {
				basicChemicals[productMaterial] = 0
			}
		}
		reactions[productMaterial] = chemicals
	}
}

func calculateAmountOfOre(target string, amount int) int {
	amountOfOre := 0
	materialList := reactions[target]
	for i := 1; i < len(materialList); i++ {
		if materialList[i].material == "ORE" {
			amountOfOre += materialList[i].amount * ceil(amount, materialList[0].amount)
		} else {
			neededAmount := ceil(amount, materialList[0].amount) * materialList[i].amount
			remnant := leftovers[materialList[i].material]
			if remnant >= neededAmount {
				leftovers[materialList[i].material] -= neededAmount
			} else {
				amountOfOre += calculateAmountOfOre(materialList[i].material, neededAmount-remnant)
				leftovers[materialList[i].material] = ceil(neededAmount-remnant, reactions[materialList[i].material][0].amount)*reactions[materialList[i].material][0].amount - (neededAmount - remnant)
			}
		}
	}
	return amountOfOre
}

func ceil(a int, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}
