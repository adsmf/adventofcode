package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode2019/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	reactions := loadInput("input.txt")
	ore, _ := calculateRequired(chemicalQuantity{"FUEL", 1}, reactions, surplusMap{})
	return ore
}

func part2() int {
	reactions := loadInput("input.txt")
	return calcMaxHold(1000000000000, reactions)
}

func calcMaxHold(capacity int, reactions reactionMap) int {
	fuelStored := 0
	guessFuel, _ := calculateRequired(chemicalQuantity{"FUEL", 1}, reactions, surplusMap{})
	targetProduction := capacity / guessFuel
	surplus := surplusMap{}
	for capacity > 0 && targetProduction > 0 {
		trySurplus := surplusMap{}
		for chemical, quantity := range surplus {
			trySurplus[chemical] = quantity
		}
		var tryOre int
		tryOre, trySurplus = calculateRequired(chemicalQuantity{"FUEL", targetProduction}, reactions, trySurplus)
		if tryOre > capacity {
			targetProduction /= 2
		} else {
			surplus = trySurplus
			fuelStored += targetProduction
			capacity -= tryOre
		}
	}
	return fuelStored
}

func calculateRequired(target chemicalQuantity, reactions reactionMap, surplus surplusMap) (int, surplusMap) {
	if target.chemical == "ORE" {
		return target.quantity, surplus
	}
	if target.quantity <= surplus[target.chemical] {
		surplus[target.chemical] -= target.quantity
		return 0, surplus
	}

	ore := 0
	target.quantity -= surplus[target.chemical]
	surplus[target.chemical] = 0
	reaction := reactions[target.chemical]
	numReactions := int(math.Ceil(float64(target.quantity) / float64(reaction.result.quantity)))
	for _, chem := range reaction.reagents {
		required := chemicalQuantity{
			chemical: chem.chemical,
			quantity: chem.quantity * numReactions,
		}
		var newOre int
		newOre, surplus = calculateRequired(required, reactions, surplus)
		ore += newOre
	}
	surplus[target.chemical] += reaction.result.quantity*numReactions - target.quantity

	return ore, surplus
}

type surplusMap map[string]int

type chemicalQuantity struct {
	chemical string
	quantity int
}

type reactionDef struct {
	reagents []chemicalQuantity
	result   chemicalQuantity
}

type reactionMap map[string]reactionDef

func loadInput(filename string) reactionMap {
	reactions := reactionMap{}
	lines := utils.ReadInputLines(filename)
	for _, line := range lines {
		sides := strings.Split(line, " => ")
		reagentStrings := strings.Split(sides[0], ", ")

		prodChem, prodQty := splitChemQty(sides[1])
		reagents := []chemicalQuantity{}
		for _, reagent := range reagentStrings {
			reqChem, reqQty := splitChemQty(reagent)
			reagents = append(reagents, chemicalQuantity{reqChem, reqQty})
		}
		reaction := reactionDef{
			result:   chemicalQuantity{prodChem, prodQty},
			reagents: reagents,
		}
		reactions[prodChem] = reaction
	}
	return reactions
}

func splitChemQty(item string) (string, int) {
	parts := strings.Split(item, " ")
	qty, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(fmt.Errorf("Unable to parse %s as quanity", parts[0]))
	}
	return parts[1], qty
}
