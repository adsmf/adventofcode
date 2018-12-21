package main

import (
	"fmt"

	"github.com/adsmf/adventofcode2018/utils"
)

func main() {
	cavern := loadFile("testData/examplecombat1.txt")
	fmt.Print(cavern.toString(true))
}

func loadFile(filename string) grid {
	lines := utils.ReadInputLines(filename)
	return load(lines)
}

func load(lines []string) grid {
	cavernMap := make(grid, len(lines))
	for y, line := range lines {
		cavernMap[y] = make(gridRow, len(line))
		for x, char := range line {
			cavernMap[y][x] = gridSquareFromChar(char)
		}
	}
	return cavernMap
}

type readingOrder int

const (
	readingOrderNorth readingOrder = iota
	readingOrderWest
	readingOrderEast
	readingOrderSouth
	readingOrderEND
)
