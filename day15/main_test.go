package main

import (
	"fmt"
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

	cavern := loadFile(testFile)
	assert.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(cavern.toString(true)))
}

func TestLoadInput(t *testing.T) {
	testFile := "input.txt"
	expected, _ := ioutil.ReadFile(testFile)
	grid := loadFile(testFile)
	assert.Equal(t, string(expected), grid.toString(false))
}

func TestBattles(t *testing.T) {
	type exampleBattle struct {
		input    string
		expected string
		outcome  int
	}
	battles := []exampleBattle{
		exampleBattle{
			`#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`,
			`#######
#...#E#   E(200)
#E#...#   E(197)
#.E##.#   E(185)
#E..#E#   E(200), E(200)
#.....#
#######
`,
			36334,
		},
	}
	for idx, battle := range battles {
		t.Run(fmt.Sprint("Battle", idx+1), func(t *testing.T) {
			debugLogger = t.Logf
			logger = func(format string, args ...interface{}) (int, error) {
				t.Logf(format, args...)
				return 0, nil
			}
			cavern := load(strings.Split(battle.input, "\n"))
			outcome := runBattle(cavern, 100)
			assert.Equal(t, battle.expected, cavern.toString(true))
			assert.Equal(t, battle.outcome, outcome)
		})
	}
}

// func TestInput(t *testing.T) {
// 	debugLogger = t.Logf
// 	cavern := loadFile("input.txt")
// 	runBattle(cavern, 1)
// }
