package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 74,87
	//Part 2: 29,74
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestP1Example(t *testing.T) {
	cartPositions := [][]cart{
		[]cart{
			cart{facing: facingEast, x: 2, y: 0},
			cart{facing: facingSouth, x: 9, y: 3},
		},
		[]cart{
			cart{facing: facingEast, x: 3, y: 0},
			cart{facing: facingEast, x: 9, y: 4},
		},
		[]cart{
			cart{facing: facingSouth, x: 4, y: 0},
			cart{facing: facingEast, x: 10, y: 4},
		},
		[]cart{
			cart{facing: facingSouth, x: 4, y: 1},
			cart{facing: facingEast, x: 11, y: 4},
		},
		[]cart{
			cart{facing: facingEast, x: 4, y: 2},
			cart{facing: facingNorth, x: 12, y: 4},
		},
		[]cart{
			cart{facing: facingEast, x: 5, y: 2},
			cart{facing: facingNorth, x: 12, y: 3},
		},
		[]cart{
			cart{facing: facingEast, x: 6, y: 2},
			cart{facing: facingNorth, x: 12, y: 2},
		},
		[]cart{
			cart{facing: facingEast, x: 7, y: 2},
			cart{facing: facingWest, x: 12, y: 1},
		},
		[]cart{
			cart{facing: facingEast, x: 8, y: 2},
			cart{facing: facingWest, x: 11, y: 1},
		},
		[]cart{
			cart{facing: facingSouth, x: 9, y: 2},
			cart{facing: facingWest, x: 10, y: 1},
		},
		[]cart{
			cart{facing: facingSouth, x: 9, y: 3},
			cart{facing: facingWest, x: 9, y: 1},
		},
		[]cart{
			cart{facing: facingWest, x: 9, y: 4},
			cart{facing: facingWest, x: 8, y: 1},
		},
		[]cart{
			cart{facing: facingWest, x: 8, y: 4},
			cart{facing: facingSouth, x: 7, y: 1},
		},
		[]cart{
			cart{facing: facingNorth, x: 7, y: 4},
			cart{facing: facingSouth, x: 7, y: 2},
		},
		[]cart{
			cart{facing: facingNorth, x: 7, y: 3, crashed: true},
			cart{facing: facingSouth, x: 7, y: 3, crashed: true},
		},
	}
	grid, simCarts := loadData("example.txt")
	for curTick, carts := range cartPositions {
		t.Run(fmt.Sprintf("Tick %d", curTick), func(t *testing.T) {
			for id, testCart := range carts {
				t.Run(fmt.Sprintf("Cart %d", id), func(t *testing.T) {
					testLoc := fmt.Sprintf("%d,%d  %d", testCart.x, testCart.y, testCart.facing)
					simLoc := fmt.Sprintf("%d,%d  %d", simCarts[id].x, simCarts[id].y, simCarts[id].facing)
					assert.Equal(t, testLoc, simLoc)
				})
			}
			tick(grid, simCarts)
		})
	}
}

func TestP1Answer(t *testing.T) {
	x, y, tick := part1("input.txt")
	assert.Equal(t, 74, x)
	assert.Equal(t, 87, y)
	assert.Equal(t, 114, tick)
}

func TestP2Example(t *testing.T) {
	cartPositions := [][]cart{
		[]cart{
			cart{facing: facingEast, x: 1, y: 0},
			cart{facing: facingWest, x: 3, y: 0},
			cart{facing: facingWest, x: 3, y: 2},
			cart{facing: facingSouth, x: 6, y: 3},
			cart{facing: facingEast, x: 1, y: 4},
			cart{facing: facingWest, x: 3, y: 4},
			cart{facing: facingNorth, x: 6, y: 5},
			cart{facing: facingWest, x: 3, y: 6},
			cart{facing: facingEast, x: 5, y: 6},
		},
		[]cart{
			cart{facing: facingSouth, x: 2, y: 2},
			cart{facing: facingNorth, x: 2, y: 6},
			cart{facing: facingNorth, x: 6, y: 6},
		},
		[]cart{
			cart{facing: facingSouth, x: 2, y: 3},
			cart{facing: facingNorth, x: 2, y: 5},
			cart{facing: facingNorth, x: 6, y: 5},
		},
		[]cart{
			cart{facing: facingNorth, x: 6, y: 4},
		},
	}
	grid, simCarts := loadData("example2.txt")
	for curTick, carts := range cartPositions {
		t.Run(fmt.Sprintf("Tick %d", curTick), func(t *testing.T) {
			for id, testCart := range carts {
				t.Run(fmt.Sprintf("Cart %d", id), func(t *testing.T) {
					testLoc := fmt.Sprintf("%d,%d  %d", testCart.x, testCart.y, testCart.facing)
					simLoc := fmt.Sprintf("%d,%d  %d", simCarts[id].x, simCarts[id].y, simCarts[id].facing)
					assert.Equal(t, testLoc, simLoc)
				})
			}
			_, simCarts = tick(grid, simCarts)
		})
	}
}

func TestP2ExampleAnswer(t *testing.T) {
	x, y, _ := part2("example2.txt")
	pos := fmt.Sprintf("%d,%d", x, y)
	assert.Equal(t, "6,4", pos)
}
