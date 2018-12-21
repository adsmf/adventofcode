package main

import (
	"fmt"
	"sort"

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

func runRound(cavern *grid) bool {
	creatures := cavern.creatures()
	sort.Slice(creatures, func(i, j int) bool {
		return compareCreatures(creatures[i], creatures[j])
	})
	for _, curCreature := range creatures {
		enemies := curCreature.enemies(cavern)
		if len(enemies) == 0 {
			return true
		}
		minDistToEnemy := utils.MaxInt
		for _, enemy := range enemies {
			dist := enemy.location.manhattenDistance(curCreature.location.point)
			if dist < minDistToEnemy {
				minDistToEnemy = dist
			}
			if minDistToEnemy == 1 {
				break
			}
		}
		if minDistToEnemy > 1 {
			setTargets(cavern, enemies)
			moveTowardBestTarget(cavern, curCreature.location)
		}
	}
	return false
}

func load(lines []string) grid {
	cavernMap := make(grid, len(lines))
	for y, line := range lines {
		cavernMap[y] = make(gridRow, len(line))
		for x, char := range line {
			cavernMap[y][x] = gridSquareFromChar(char)
			cavernMap[y][x].x = x
			cavernMap[y][x].y = y
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
