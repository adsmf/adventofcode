package main

import "fmt"

type creature struct {
	hp   int
	race creatureType
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

type creatureType int

const (
	creatureTypeElf creatureType = iota
	creatureTypeGoblin
	creatureTypeEND
)
