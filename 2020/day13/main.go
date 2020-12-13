package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	lines := utils.ReadInputLines("input.txt")
	p1 := part1(lines)
	p2 := part2(lines[1])
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(lines []string) int {
	departure := utils.GetInts(lines[0])[0]
	busses := utils.GetInts(lines[1])
	bestBus := 0
	bestDelay := utils.MaxInt
	for _, bus := range busses {
		delay := bus - departure%bus
		if delay < bestDelay {
			bestDelay = delay
			bestBus = bus
		}
	}
	return bestBus * bestDelay
}

func part2(busDef string) int {
	busStrings := strings.Split(busDef, ",")
	busses := make([]*big.Int, 0, len(busStrings))
	delays := make([]*big.Int, 0, len(busStrings))
	runningMultiple := 1
	for index, busString := range busStrings {
		if busString == "x" {
			continue
		}
		bus, _ := strconv.Atoi(busString)
		runningMultiple *= bus
		delays = append(delays, big.NewInt(int64(bus-index)))
		busses = append(busses, big.NewInt(int64(bus)))
	}

	departureTime := new(big.Int)
	commonMultiple := new(big.Int).SetInt64(int64(runningMultiple))
	for index, bus := range busses {
		remainder := new(big.Int).Div(commonMultiple, bus)
		var coeff, divisor big.Int
		divisor.GCD(nil, &coeff, bus, remainder)
		departureTime.Add(departureTime, coeff.Mul(delays[index], coeff.Mul(&coeff, remainder)))
	}

	return int(departureTime.Mod(departureTime, commonMultiple).Int64())
}

var benchmark = false
