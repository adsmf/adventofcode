package main

import (
	"fmt"
	"sort"

	"github.com/adsmf/adventofcode2018/utils"
)

var mainLogger func(string, ...interface{}) (int, error)
var debugLogger func(string, ...interface{})

func logger(format string, args ...interface{}) {
	if mainLogger != nil {
		mainLogger(format, args...)
	} else if debugLogger != nil {
		debugLogger(format, args...)
	}
}

func debug(format string, args ...interface{}) {
	if debugLogger != nil {
		debugLogger(format, args...)
	}
}

func main() {
	mainLogger = fmt.Printf
	// cavern := loadFile("input.txt")
	// runBattle(cavern, 1000, 2, false)

	curPower := 3
	for {
		cavern := loadFile("input.txt", curPower)
		outcome := runBattle(cavern, 1000, true)
		if outcome >= 0 {
			break
		}
		curPower++
	}
}

func loadFile(filename string, elfPower int) *grid {
	lines := utils.ReadInputLines(filename)
	return load(lines, elfPower)
}

func runBattle(cavern *grid, maxRounds int, requireNoLoss bool) int {
	completedRound := 0
	for {
		battleComplete, elvesDied := runRound(cavern, requireNoLoss)
		if battleComplete {
			break
		}
		if requireNoLoss && elvesDied {
			logger("Lost an elf; abandoning trial with power after %d rounds\n", completedRound)
			return -1
		}
		completedRound++
		if completedRound > maxRounds {
			return -1
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

func runRound(cavern *grid, requireNoLoss bool) (bool, bool) {
	creatures := cavern.creatures()
	elves := []*creature{}
	for _, curCreature := range creatures {
		if curCreature.race == creatureTypeElf {
			elves = append(elves, curCreature)
		}
	}
	sort.Slice(creatures, func(i, j int) bool {
		return compareCreatures(creatures[i], creatures[j])
	})
	for _, curCreature := range creatures {
		if curCreature.hp <= 0 {
			continue
		}
		enemies := curCreature.enemies(cavern)
		if len(enemies) == 0 {
			return true, false
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
		if requireNoLoss {
			for _, elf := range elves {
				if elf.hp <= 0 {
					return false, true
				}
			}
		}
	}
	return false, false
}

func load(lines []string, elfPower int) *grid {
	cavernMap := make(grid, len(lines))
	for y, line := range lines {
		cavernMap[y] = make(gridRow, len(line))
		for x, char := range line {
			cavernMap[y][x] = gridSquareFromChar(char, elfPower)
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
