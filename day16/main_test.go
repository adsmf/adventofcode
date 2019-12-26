package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBehaveExamples(t *testing.T) {
	type countExample struct {
		example     exampleInput
		behavesLike []operation
	}
	examples := []countExample{
		countExample{
			example: exampleInput{
				input:       registers{3, 2, 1, 1},
				output:      registers{3, 2, 2, 1},
				instruction: instruction{9, 2, 1, 2},
			},
			behavesLike: []operation{mulr, addi, seti},
		},
	}
	for idx, example := range examples {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			debugLogger = t.Logf
			actual := findBehavesLike(example.example)
			expected := example.behavesLike
			sort.Slice(actual, func(i int, j int) bool {
				return actual[i] < actual[j]
			})
			sort.Slice(expected, func(i int, j int) bool {
				return expected[i] < expected[j]
			})
			assert.Equal(t, opListToString(expected), opListToString(actual))
		})
	}
}

func TestPart1(t *testing.T) {
	examples, _ := loadInput("input.txt")
	answer := part1(examples)
	assert.Equal(t, 509, answer)
}
