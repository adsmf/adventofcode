package main

import (
	"fmt"
	"sort"
)

type creature struct {
	hp       int
	power    int
	race     creatureType
	location *gridSquare
}

func (c *creature) attack(cavern *grid) {
	adjacentEnemies := []*creature{}
	for testDir := readingOrderNorth; testDir < readingOrderEND; testDir++ {
		var occupant *creature
		switch testDir {
		case readingOrderNorth:
			occupant = (*cavern)[c.location.y-1][c.location.x].occupiedBy
		case readingOrderEast:
			occupant = (*cavern)[c.location.y][c.location.x+1].occupiedBy
		case readingOrderSouth:
			occupant = (*cavern)[c.location.y+1][c.location.x].occupiedBy
		case readingOrderWest:
			occupant = (*cavern)[c.location.y][c.location.x-1].occupiedBy
		default:
			panic("Direction out of range!")
		}
		if occupant != nil {
			if occupant.hp > 0 && occupant.race != c.race {
				adjacentEnemies = append(adjacentEnemies, occupant)
			}
		}
	}
	if len(adjacentEnemies) == 0 {
		return
	}
	sort.Slice(adjacentEnemies, func(i, j int) bool {
		return compareCreaturesForAttack(adjacentEnemies[i], adjacentEnemies[j])
	})
	target := adjacentEnemies[0]
	target.hp -= c.power
	if target.hp <= 0 {
		debug("%s killed by %s\n", target.toString(), c.toString())
		target.location.occupiedBy = nil
		target.location = nil
	}
}

func (c *creature) toString() string {
	switch c.race {
	case creatureTypeElf:
		return "E"
	case creatureTypeGoblin:
		return "G"
	default:
		panic("Unknown creature type")
	}
}

func (c *creature) fullString() string {
	return fmt.Sprintf("%s(%d)", c.toString(), c.hp)
}

func compareCreatures(i, j *creature) bool {
	lowerReadingOrder := i.location.y < j.location.y || (i.location.y == j.location.y && i.location.x < j.location.x)
	return lowerReadingOrder
}

func compareCreaturesForAttack(i, j *creature) bool {
	lowerHP := i.hp < j.hp
	equalHP := i.hp == j.hp

	lowerReadingOrder := i.location.y < j.location.y || (i.location.y == j.location.y && i.location.x < j.location.x)
	return lowerHP || (equalHP && lowerReadingOrder)
}

func (c *creature) enemies(cavern *grid) []*creature {
	enemies := []*creature{}
	for y, row := range *cavern {
		for x := range row {
			candidate := (*cavern)[y][x].occupiedBy
			if candidate != nil && candidate.race != c.race {
				enemies = append(enemies, candidate)
			}
		}
	}
	return enemies
}

type creatureType int

const (
	creatureTypeElf creatureType = iota
	creatureTypeGoblin
	creatureTypeEND
)
