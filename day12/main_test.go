package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadExampleData(t *testing.T) {
	initial, mapping := loadData("example.txt")
	expectedMapping := make([]bool, 32)

	// ...## => #
	expectedMapping[3] = true
	// ..#.. => #
	expectedMapping[4] = true
	// .#... => #
	expectedMapping[8] = true
	// .#.#. => #
	expectedMapping[10] = true
	// .#.## => #
	expectedMapping[11] = true
	// .##.. => #
	expectedMapping[12] = true
	// .#### => #
	expectedMapping[15] = true
	// #.#.# => #
	expectedMapping[21] = true
	// #.### => #
	expectedMapping[23] = true
	// ##.#. => #
	expectedMapping[26] = true
	// ##.## => #
	expectedMapping[27] = true
	// ###.. => #
	expectedMapping[28] = true
	// ###.# => #
	expectedMapping[29] = true
	// ####. => #
	expectedMapping[30] = true

	assert.Equal(t, "1001010011000000111000111", initial)
	assert.Equal(t, expectedMapping, mapping)
}

func TestLoadInputData(t *testing.T) {
	initial, mapping := loadData("input.txt")
	expectedMapping := []bool{
		false,
		false,
		false,
		true,
		false,
		false,
		false,
		false,
		true,
		true,
		true,
		true,
		false,
		false,
		false,
		true,
		false,
		true,
		true,
		false,
		false,
		true,
		true,
		false,
		true,
		false,
		false,
		true,
		false,
		true,
		false,
		false,
	}
	assert.Equal(t, "11011110011110001011110011010011001111101101001000101110111000011110111000110010001101010001101100", initial)
	assert.Equal(t, expectedMapping, mapping)
}

func TestSimCount(t *testing.T) {
	type exampleCount struct {
		initial string
		ticks   int
		count   int
	}
	examples := []exampleCount{
		exampleCount{"1001010011000000111000111", 0, 145},
		exampleCount{"1001010011000000111000111", 1, 91},
		exampleCount{"1001010011000000111000111", 2, 132},
		exampleCount{"1001010011000000111000111", 4, 154},
		exampleCount{"1001010011000000111000111", 6, 174},
		exampleCount{"1001010011000000111000111", 8, 213},
		exampleCount{"1001010011000000111000111", 10, 213},
		exampleCount{"1001010011000000111000111", 12, 218},
		exampleCount{"1001010011000000111000111", 14, 235},
		exampleCount{"1001010011000000111000111", 16, 226},
		exampleCount{"1001010011000000111000111", 18, 280},
		exampleCount{"1001010011000000111000111", 20, 325},
		exampleCount{"1001010011000000111000111", 500, 9374},
		exampleCount{"1001010011000000111000111", 555, 10474},
		exampleCount{"1001010011000000111000111", 666, 12694},
		exampleCount{"1001010011000000111000111", 5000, 99374},
		exampleCount{"1001010011000000111000111", 6000, 119374},
	}
	initial, mapping := loadData("example.txt")

	for exID, example := range examples {
		t.Run(
			fmt.Sprintf("Example %d", exID),
			func(t *testing.T) {
				count := runSim(initial, mapping, example.ticks)
				assert.Equal(t, example.count, count)
			})
	}
}
