package main

import (
	"fmt"
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
		point{3, 3},
		point{4, 2},
		point{5, 3},
	}
	cavern := load(strings.Split(strings.TrimSpace(startingGrid), "\n"))
	for _, target := range targets {
		cavern[target.y][target.x].isTarget = true
	}
	t.Logf("Targets:\n%+v", cavern.toString(false))

	startingPoint := &cavern[1][2]
	expectedCosts := []int{3, 3, 5}

	cavern = moveTowardBestTarget(cavern, startingPoint)
	t.Logf("Costs:\n%+v", cavern.costString())

	for idx, target := range targets {
		assert.Equal(t, expectedCosts[idx], cavern[target.y][target.x].cost)
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
		point{1, 3},
		point{2, 2},
		point{3, 3},
	}
	cavern := load(strings.Split(strings.TrimSpace(startingGrid), "\n"))
	for _, target := range targets {
		cavern[target.y][target.x].isTarget = true
	}
	t.Logf("Targets:\n%+v", cavern.toString(false))

	startingPoint := &cavern[1][2]
	expectedCosts := []int{3, utils.MaxInt, 7}

	cavern = moveTowardBestTarget(cavern, startingPoint)
	t.Logf("Costs:\n%+v", cavern.costString())

	for idx, target := range targets {
		t.Run(fmt.Sprint("Target", idx+1), func(t *testing.T) {
			assert.Equal(t, expectedCosts[idx], cavern[target.y][target.x].cost)
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
		cavern[target.y][target.x].isTarget = true
	}
	t.Logf("Targets:\n%+v", cavern.toString(false))

	startingPoint := &cavern[1][2]
	expectedCosts := []int{3, 1, 3}

	cavern = moveTowardBestTarget(cavern, startingPoint)
	t.Logf("Costs:\n%+v", cavern.costString())

	for idx, target := range targets {
		t.Run(fmt.Sprint("Target", idx+1), func(t *testing.T) {
			assert.Equal(t, expectedCosts[idx], cavern[target.y][target.x].cost)
		})
	}
	assert.Equal(t, utils.MaxInt, cavern[3][4].cost)
}
