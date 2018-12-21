package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	testFile := "testData/examplemovement1.txt"
	expected, _ := ioutil.ReadFile(testFile)
	grid := loadFile(testFile)
	assert.Equal(t, string(expected), grid.toString(false))
}

func TestHealthDisplay(t *testing.T) {
	testFile := "testData/examplemovement1.txt"
	expected := `
#######
#.G.E.#   G(200), E(200)
#E.G.E#   E(200), G(200), E(200)
#.G.E.#   G(200), E(200)
#######`

	grid := loadFile(testFile)
	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(grid.toString(true)))
}
