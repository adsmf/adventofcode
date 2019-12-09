package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {

}

func TestPart2Examples(t *testing.T) {

}

func TestAnswers(t *testing.T) {
	assert.Equal(t, 0, part1())
	assert.Equal(t, 0, part2())
}

func TestMainRuns(t *testing.T) {
	assert.NotPanics(t, func() { main() })
}
