package main

import "fmt"

func main() {
	level, x, y := bestInGrid(3613)
	fmt.Printf("Best pos (%d,%d) with level %d\n", x, y, level)
}

func calcLevel(serial, x, y int) int {
	rackID := x + 10
	level := rackID * y
	level += serial
	level = level * rackID
	level = int((level % 1000) / 100)
	level -= 5

	return level
}

func bestInGrid(serial int) (int, int, int) {
	bestLevel := 0
	bestX := 0
	bestY := 0

	levels := make([][]int, 300)
	for row := 0; row < 300; row++ {
		levels[row] = make([]int, 300)
		for col := 0; col < 300; col++ {
			levels[row][col] = calcLevel(serial, col, row)
		}
	}

	for y := 0; y < 298; y++ {
		for x := 0; x < 298; x++ {
			level := levels[y][x]
			level += levels[y][x+1]
			level += levels[y][x+2]

			level += levels[y+1][x]
			level += levels[y+1][x+1]
			level += levels[y+1][x+2]

			level += levels[y+2][x]
			level += levels[y+2][x+1]
			level += levels[y+2][x+2]

			if level > bestLevel {
				bestLevel = level
				bestX = x
				bestY = y
			}
		}
	}

	return bestLevel, bestX, bestY
}
