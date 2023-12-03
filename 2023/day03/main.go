package main

import (
	_ "embed"
	"fmt"
	"sort"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	sch := loadSchematic()
	p1, p2 := sch.analyse()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func loadSchematic() schematic {
	sch := schematic{}
	utils.EachLine(input, func(row int, line string) (done bool) {
		for col, ch := range line {
			sch[point{col, row}] = byte(ch)
		}
		return false
	})
	return sch
}

type schematic map[point]byte

func (s schematic) analyse() (int, int) {
	totalPartNum := 0
	totalGearRatio := 0
	gears := map[point][]partNum{}
	parts := s.getPartNums()
	for _, part := range parts {
		partGears := part.gears()
		if part.valid() {
			totalPartNum += part.number
			for gear := range partGears {
				gears[gear] = append(gears[gear], part)
			}
		}
	}
	for _, gearParts := range gears {
		if len(gearParts) != 2 {
			continue
		}
		totalGearRatio += gearParts[0].number * gearParts[1].number
	}
	return totalPartNum, totalGearRatio
}

func (s schematic) getPartNums() []partNum {
	toSearch := map[point]bool{}
	for pos, ch := range s {
		if ch < '0' || ch > '9' {
			continue
		}
		toSearch[pos] = true
	}
	parts := []partNum{}
	for pos := range toSearch {
		part := partNum{
			schematic: s,
		}
		part.points = append(part.points, pos)
		for curPos := pos.left(); s.isNum(curPos); curPos = curPos.left() {
			part.points = append(part.points, curPos)
			delete(toSearch, curPos)
		}
		for curPos := pos.right(); s.isNum(curPos); curPos = curPos.right() {
			part.points = append(part.points, curPos)
			delete(toSearch, curPos)
		}
		part.calcNeighbours()
		if !part.valid() {
			continue
		}
		part.calcNumber()
		parts = append(parts, part)
	}
	return parts
}

func (s schematic) isNum(pos point) bool {
	ch := s[pos]
	return (ch >= '0' && ch <= '9')
}

type partNum struct {
	schematic  schematic
	points     []point
	neighbours []point
	number     int
}

func (p partNum) gears() pointSet {
	gears := pointSet{}
	for _, neighbour := range p.neighbours {
		if p.schematic[neighbour] == '*' {
			gears[neighbour] = true
		}
	}
	return gears
}

func (p partNum) valid() bool {
	return len(p.neighbours) > 0
}

func (p *partNum) calcNumber() {
	sort.Slice(p.points, func(i, j int) bool { return p.points[i].x < p.points[j].x })
	num := 0
	for _, pos := range p.points {
		num *= 10
		num += int(p.schematic[pos] - '0')
	}
	p.number = num
}

func (p *partNum) calcNeighbours() {
	for _, pos := range p.points {
		for _, neighbour := range pos.neighbours() {
			nCh := p.schematic[neighbour]
			if nCh == 0 || nCh == '.' || (nCh >= '0' && nCh <= '9') {
				continue
			}
			p.neighbours = append(p.neighbours, neighbour)
		}
	}
}

type pointSet map[point]bool

type point struct {
	x, y int
}

func (p point) add(a point) point {
	return point{
		x: p.x + a.x,
		y: p.y + a.y,
	}
}

func (p point) left() point {
	return p.add(point{-1, 0})
}

func (p point) right() point {
	return p.add(point{1, 0})
}

func (p point) neighbours() []point {
	return []point{
		{p.x - 1, p.y - 1},
		{p.x + 0, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y + 0},
		{p.x + 1, p.y + 0},
		{p.x - 1, p.y + 1},
		{p.x + 0, p.y + 1},
		{p.x + 1, p.y + 1},
	}
}

var benchmark = false
