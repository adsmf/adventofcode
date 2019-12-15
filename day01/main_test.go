package main

import (
	"testing"

	"github.com/adsmf/adventofcode2019/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	tests := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}

	for mass, fuel := range tests {
		assert.Equal(t, fuel, calculateFuel(mass))
	}
}

func TestPart1Result(t *testing.T) {
	inputLines := utils.ReadInputLines("input.txt")
	fuel := calculateTotalFuel(inputLines, false)
	assert.Equal(t, 3367126, fuel)
}

func TestPart2Examples(t *testing.T) {
	tests := map[int]int{
		14:     2,
		1969:   966,
		100756: 50346,
	}

	for mass, fuel := range tests {
		assert.Equal(t, fuel, calculateFuelRecursive(mass))
	}
}
