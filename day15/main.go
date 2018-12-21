package main

import (
	"fmt"
	"sort"

	"github.com/adsmf/adventofcode2018/utils"
)

var logger func(string, ...interface{}) (int, error)
var debugLogger func(string, ...interface{})

func debug(format string, args ...interface{}) {
	if debugLogger != nil {
		debugLogger(format, args...)
	}
}

func main() {
	cavern := loadFile("input.txt")
	logger = fmt.Printf
	runBattle(cavern, 1000)
}

func loadFile(filename string) *grid {
	lines := utils.ReadInputLines(filename)
	return load(lines)
}

func runBattle(cavern *grid, maxRounds int) int {
	completedRound := 0
	for !!!runRound(cavern) {
		completedRound++
		if completedRound > maxRounds {
			return 0
		}
	}

	remainingCreatures := cavern.creatures()
	remainingHP := 0
	for _, curCreature := range remainingCreatures {
		if curCreature.hp > 0 {
			remainingHP += curCreature.hp
		}
	}
	result := completedRound * remainingHP
	logger("Battle complete - %s win after %d rounds: %d (rounds) * %d (HP) = %d\n", remainingCreatures[0].toString(), completedRound, completedRound, remainingHP, result)
	return result
}

func runRound(cavern *grid) bool {
	creatures := cavern.creatures()
	sort.Slice(creatures, func(i, j int) bool {
		return compareCreatures(creatures[i], creatures[j])
	})
	for _, curCreature := range creatures {
		if curCreature.hp <= 0 {
			continue
		}
		enemies := curCreature.enemies(cavern)
		if len(enemies) == 0 {
			return true
		}

		// See if we're close to an enemy
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

		curCreature.attack(cavern)
	}
	return false
}

func load(lines []string) *grid {
	cavernMap := make(grid, len(lines))
	for y, line := range lines {
		cavernMap[y] = make(gridRow, len(line))
		for x, char := range line {
			cavernMap[y][x] = gridSquareFromChar(char)
			cavernMap[y][x].x = x
			cavernMap[y][x].y = y
		}
	}
	return &cavernMap
}

type readingOrder int

const (
	readingOrderNorth readingOrder = iota
	readingOrderWest
	readingOrderEast
	readingOrderSouth
	readingOrderEND
)
