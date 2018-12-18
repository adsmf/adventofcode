package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
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
		logger := t.Logf
		t.Run(fmt.Sprintf("Tick %d", curTick), func(t *testing.T) {
			for id, testCart := range carts {
				t.Run(fmt.Sprintf("Cart %d", id), func(t *testing.T) {
					assert.Equal(t, testCart.facing, simCarts[id].facing)
					assert.Equal(t, testCart.x, simCarts[id].x)
					assert.Equal(t, testCart.y, simCarts[id].y)
					assert.Equal(t, testCart.crashed, simCarts[id].crashed)
				})
			}
			tick(logger, grid, simCarts)
		})
	}
}
