package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	img := parseImage("123456789012", 3, 2)
	expected := image{
		layer{1, 2, 3, 4, 5, 6},
		layer{7, 8, 9, 0, 1, 2},
	}
	assert.Equal(t, expected, img)
}

func TestPart2Examples(t *testing.T) {

}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 1572, part1())
	assert.Equal(t, 60, part2())
}

func TestMainRuns(t *testing.T) {
	assert.NotPanics(t, func() { main() })
}
