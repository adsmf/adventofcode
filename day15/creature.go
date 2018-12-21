package main

import (
	"fmt"
)

type creature struct {
	hp       int
	race     creatureType
	location *gridSquare
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
