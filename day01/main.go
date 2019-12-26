package main

import (
	"fmt"
	"strconv"

	"github.com/adsmf/adventofcode2019/utils"
)

func main() {
	inputLines := utils.ReadInputLines("input.txt")
	fuel := calculateTotalFuel(inputLines, false)
	fmt.Printf("Part 1: %d\n", fuel)
	fuelR := calculateTotalFuel(inputLines, true)
	fmt.Printf("Part 2: %d\n", fuelR)
}

func calculateTotalFuel(modules []string, recursive bool) int {
	totalFuel := 0
	for _, module := range modules {
		moduleMass, err := strconv.Atoi(module)
		if err != nil {
			panic(err)
		}
		var fuel int
		if recursive {
			fuel = calculateFuelRecursive(moduleMass)
		} else {
			fuel = calculateFuel(moduleMass)
		}
		totalFuel = totalFuel + fuel
	}
	return totalFuel
}

func calculateFuel(mass int) int {
	return int(mass/3) - 2
}

func calculateFuelRecursive(mass int) int {
	totalFuel := 0
	for {
		fuel := calculateFuel(mass)

		if fuel > 0 {

			totalFuel = totalFuel + fuel
			mass = fuel
		} else {
			return totalFuel
		}
	}
}
