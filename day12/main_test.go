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

func TestSimMatches(t *testing.T) {
	pots, mapping := loadData("example.txt")

	pots = fmt.Sprintf("000%s00000000000", pots)

	expectedResults := []string{
		"000100101001100000011100011100000000000",
		"000100010000100000100100100100000000000",
		"000110011000110000100100100110000000000",
		"001010001001010000100100100010000000000",
		"000101001000101000100100110011000000000",
		"000010001100010100100100010001000000000",
		"000011010100001000100110011001100000000",
		"000100111010001100100010001000100000000",
		"000100001101010100110011001100110000000",
		"000110010011111000010001000100010000000",
		"001010010001011000011001100110011000000",
		"000100011000101000101000100010001000000",
		"000110101000010100010100110011001100000",
		"001001110100001010001000010001000100000",
		"001000011010000101001100011001100110000",
		"001100100101000010000100101000100010000",
		"010100100010100011000100010100110011000",
		"001000110001010101000110001000010001000",
		"001101010000111110101010001100011001100",
		"010011101001010111111101010100101000100",
		"010000110000111110001111111000010100110",
	}

	for ticks, expected := range expectedResults {
		t.Run(
			fmt.Sprintf("Tick %d", ticks),
			func(t *testing.T) {
				trimmedResult, _ := runSim(pots, mapping, ticks)
				assert.Equal(t, expected, trimmedResult)
			},
		)
	}
}

func TestSimCount(t *testing.T) {
	type exampleCount struct {
		initial string
		ticks   int
		count   int
	}
	examples := []exampleCount{
		exampleCount{"1001010011000000111000111", 20, 325},
	}
	initial, mapping := loadData("example.txt")

	for exID, example := range examples {
		t.Run(
			fmt.Sprintf("Example %d", exID),
			func(t *testing.T) {
				_, count := runSim(initial, mapping, example.ticks)
				assert.Equal(t, example.count, count)
			})
	}
}
