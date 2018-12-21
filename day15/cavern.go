package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode2018/utils"
)

type grid []gridRow
type gridRow []gridSquare
type gridSquare struct {
	isCavern   bool
	occupiedBy *creature
	isTarget   bool
	cost       int
}

func (g *grid) toString(withHealth bool) string {
	retString := ""
	for _, row := range *g {
		retString = fmt.Sprint(retString, row.toString(withHealth), "\n")
	}
	return retString
}

func (g *gridRow) toString(withHealth bool) string {
	retString := ""
	occupants := []*creature{}
	for _, col := range *g {
		retString = fmt.Sprint(retString, col.toString())
		if withHealth && col.occupiedBy != nil {
			occupants = append(occupants, col.occupiedBy)
		}
	}
	if len(occupants) == 0 {
		return retString
	}
	healths := []string{}
	for _, occupant := range occupants {
		healths = append(healths, occupant.fullString())
	}
	retString = fmt.Sprintf("%s   %s", retString, strings.Join(healths, ", "))
	return retString
}

func (g *gridSquare) toString() string {
	if g.isCavern {
		if g.isTarget {
			return "?"
		}
		if g.occupiedBy == nil {
			return "."
		}
		return g.occupiedBy.toString()
	}
	return "#"
}

func (g *grid) costString() string {
	retString := ""
	for _, row := range *g {
		retString = fmt.Sprint(retString, row.costString(), "\n")
	}
	return retString
}

func (g *gridRow) costString() string {
	retString := ""
	for _, col := range *g {
		retString = fmt.Sprint(retString, col.costString())
	}
	return retString
}

func (g *gridSquare) costString() string {
	if g.isCavern {
		switch {
		case g.occupiedBy != nil:
			return g.toString()
		case g.cost == utils.MaxInt:
			return "♾"
		case g.cost < 10:
			return strconv.Itoa(g.cost)
		default:
			return "?"
		}
	}
	return "#"
}

func gridSquareFromChar(char rune) gridSquare {
	square := gridSquare{}
	switch char {
	case '#':
	case '.':
		square.isCavern = true
	case 'G':
		square.isCavern = true
		square.occupiedBy = &creature{
			race: creatureTypeGoblin,
			hp:   200,
		}
	case 'E':
		square.isCavern = true
		square.occupiedBy = &creature{
			race: creatureTypeElf,
			hp:   200,
		}
	default:
		panic(fmt.Errorf("Unsupported char: %c", char))
	}
	return square
}
