package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/adsmf/adventofcode2018/utils"

	"github.com/stretchr/testify/assert"
)

func TestSimpleDistance(t *testing.T) {
	startingGrid := `
#######
#.E...#
#.....#
#...G.#
#######
`
	targets := []point{
		point{4, 3},
	}
	cavern := load(strings.Split(strings.TrimSpace(startingGrid), "\n"))
	for _, target := range targets {
		(*cavern)[target.y][target.x].isTarget = true
	}
	t.Logf("Targets:\n%+v", cavern.toString(false))

	startingPoint := (*cavern)[1][2]
	expectedCosts := []int{4}

	moveTowardBestTarget(cavern, startingPoint)
	t.Logf("Costs:\n%+v", cavern.costString())

	for idx, target := range targets {
		assert.Equal(t, expectedCosts[idx], (*cavern)[target.y][target.x].cost)
	}
}

func TestWallDistance(t *testing.T) {
	startingGrid := `
#######
#.E...#
#.###.#
#.G...#
#######
`
	targets := []point{
		point{2, 3},
	}
	cavern := load(strings.Split(strings.TrimSpace(startingGrid), "\n"))
	for _, target := range targets {
		(*cavern)[target.y][target.x].isTarget = true
	}
	t.Logf("Targets:\n%+v", cavern.toString(false))

	startingPoint := (*cavern)[1][2]
	expectedCosts := []int{4}

	moveTowardBestTarget(cavern, startingPoint)
	t.Logf("Costs:\n%+v", cavern.costString())

	for idx, target := range targets {
		t.Run(fmt.Sprint("Target", idx+1), func(t *testing.T) {
			assert.Equal(t, expectedCosts[idx], (*cavern)[target.y][target.x].cost)
		})
	}
}

func TestLazyDistance(t *testing.T) {
	startingGrid := `
#######
#.E...#
#.....#
#.G...#
#######
`
	targets := []point{
		point{1, 3},
		point{2, 2},
		point{3, 3},
	}
	cavern := load(strings.Split(strings.TrimSpace(startingGrid), "\n"))
	for _, target := range targets {
		(*cavern)[target.y][target.x].isTarget = true
	}
	t.Logf("Targets:\n%+v", cavern.toString(false))

	startingPoint := (*cavern)[1][2]
	expectedCosts := []int{3, 1, 3}

	moveTowardBestTarget(cavern, startingPoint)
	t.Logf("Costs:\n%+v", (*cavern).costString())

	for idx, target := range targets {
		t.Run(fmt.Sprint("Target", idx+1), func(t *testing.T) {
			assert.Equal(t, expectedCosts[idx], (*cavern)[target.y][target.x].cost)
		})
	}
	assert.Equal(t, utils.MaxInt, (*cavern)[3][4].cost)
}

func TestExampleMovement2(t *testing.T) {
	cavern := loadFile("testData/examplemovement2.0.txt")
	var exampleRounds []*grid
	for i := 1; i <= 3; i++ {
		exampleRounds = append(exampleRounds, loadFile(fmt.Sprintf("testData/examplemovement2.%d.txt", i)))
	}
	for roundNum, exampleRound := range exampleRounds {
		t.Run(fmt.Sprint("Round", roundNum+1), func(t *testing.T) {
			runRound(cavern, 2)
			t.Logf("Cavern after round %d:\n%s", roundNum+1, cavern.toString(false))
			assert.Equal(t, exampleRound.toString(false), cavern.toString(false))
		})
	}
}

func TestExampleCombat1(t *testing.T) {
	cavern := loadFile("testData/examplecombat1.txt")
	var exampleRounds []string

	roundIDs := []int{1, 2, 23, 24, 25, 26, 27, 28, 47}
	for _, round := range roundIDs {
		roundText, err := ioutil.ReadFile(fmt.Sprintf("testData/examplecombat1.%d.txt", round))
		assert.NoError(t, err)
		exampleRounds = append(exampleRounds, string(roundText))

	}
	curRound := 0
	for idx, roundID := range roundIDs {
		exampleRound := exampleRounds[idx]
		t.Run(fmt.Sprintf("Test %d - round %d", idx+1, roundID), func(t *testing.T) {
			for curRound < roundID {
				curRound++
				runRound(cavern, 2)
			}
			t.Logf("Cavern after round %d:\n%s", curRound, cavern.toString(true))
			assert.Equal(t, exampleRound, cavern.toString(true))
		})
	}
}
