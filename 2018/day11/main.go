package main

import "fmt"

func main() {
	serial := 3613
	_, x, y := bestForSerial(serial, 3)
	p1 := fmt.Sprintf("%d,%d", x, y)

	_, x, y, size := bestGrid(serial)
	p2 := fmt.Sprintf("%d,%d,%d", x, y, size)
	if !benchmark {
		fmt.Println("Part 1:", p1)
		fmt.Println("Part 2:", p2)
	}
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

func bestForSerial(serial, size int) (int, int, int) {
	levels := calcLevels(serial)
	return bestInGrid(levels, size)
}

func calcLevels(serial int) [][]int {
	levels := make([][]int, 300)
	for row := 0; row < 300; row++ {
		levels[row] = make([]int, 300)
		for col := 0; col < 300; col++ {
			levels[row][col] = calcLevel(serial, col, row)
		}
	}
	return levels
}

func bestInGrid(levels [][]int, size int) (int, int, int) {
	bestLevel := 0
	bestX := 0
	bestY := 0

	for y := 0; y <= 300-size; y++ {
		for x := 0; x <= 300-size; x++ {
			level := 0
			for dy := 0; dy < size; dy++ {
				for dx := 0; dx < size; dx++ {
					level += levels[y+dy][x+dx]
				}
			}

			if level > bestLevel {
				bestLevel = level
				bestX = x
				bestY = y
			}
		}
	}

	return bestLevel, bestX, bestY
}

func bestGrid(serial int) (int, int, int, int) {
	bestLevel := 0
	bestX := 0
	bestY := 0
	bestSize := 0

	levels := calcLevels(serial)

	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			level := 0
			for size := 1; size <= 300; size++ {
				if (x+size >= 300) || (y+size >= 300) {
					break
				}
				for dx := 0; dx < size-1; dx++ {
					level += levels[y+size-1][x+dx]
				}
				for dy := 0; dy < size-1; dy++ {
					level += levels[y+dy][x+size-1]
				}
				level += levels[y+size-1][x+size-1]

				if level > bestLevel {
					bestLevel = level
					bestX = x
					bestY = y
					bestSize = size
				}
			}
		}
	}
	return bestLevel, bestX, bestY, bestSize
}

var benchmark = false
