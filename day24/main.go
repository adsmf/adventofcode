package main

import (
	"fmt"
	"math"

	"github.com/adsmf/adventofcode2019/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	p := loadMap("input.txt")
	return p.findRepeat()
}

func part2() int {
	return 0
}

type planet struct {
	grid       map[point]tile
	minX, minY int
	maxX, maxY int
}

func (p *planet) findRepeat() int {
	seen := map[int]bool{}

	for {
		curBio := p.biodiversity()
		if seen[curBio] {
			return curBio
		}
		seen[curBio] = true
		p.iter()
	}
}

func (p *planet) iter() {
	newGrid := map[point]tile{}

	for pos, t := range p.grid {
		infested := t.tileType == tileTypeInfested
		surrounding := 0
		for _, neighbour := range pos.neighbours() {
			if p.grid[neighbour].tileType == tileTypeInfested {
				surrounding++
			}
		}
		if surrounding == 1 ||
			(!infested && surrounding == 2) {
			newGrid[pos] = tile{tileType: tileTypeInfested}
		} else {
			newGrid[pos] = tile{tileType: tileTypeEmpty}
		}
	}
	p.grid = newGrid
}

func (p *planet) set(pos point, val tile) {
	val.pos = pos
	p.grid[pos] = val
	if p.minX > pos.x {
		p.minX = pos.x
	}
	if p.maxX < pos.x {
		p.maxX = pos.x
	}
	if p.minY > pos.y {
		p.minY = pos.y
	}
	if p.maxY < pos.y {
		p.maxY = pos.y
	}
}

func (p planet) String() string {
	newString := ""
	for y := p.minY; y <= p.maxY; y++ {
		for x := p.minX; x <= p.maxX; x++ {
			newString += fmt.Sprintf("%v", p.grid[point{x, y}])
		}
		newString += fmt.Sprintln()
	}
	return newString
}

func (p planet) biodiversity() int {
	bio := 0
	for pos, t := range p.grid {
		if t.tileType == tileTypeInfested {
			id := pos.x + pos.y*(p.maxX+1)
			bio += int(math.Pow(2, float64(id)))
		}
	}
	return bio
}

type point struct {
	x, y int
}

func (p point) neighbours() []point {
	return []point{
		point{p.x - 1, p.y},
		point{p.x + 1, p.y},
		point{p.x, p.y - 1},
		point{p.x, p.y + 1},
	}
}

type tile struct {
	tileType tileType
	pos      point
}

func (t tile) String() string {
	switch t.tileType {
	case tileTypeEmpty:
		return "."
	case tileTypeInfested:
		return "#"
	case tileTypeUnkown:
		return "?"
	default:
		return "!"
	}
}

type tileType int

const (
	tileTypeUnkown tileType = iota
	tileTypeEmpty
	tileTypeInfested
)

func loadMap(filename string) planet {
	newPlanet := planet{
		grid: map[point]tile{},
	}

	lines := utils.ReadInputLines(filename)
	for y, line := range lines {
		for x, char := range line {
			pos := point{x, y}
			switch {
			case char == '.':
				newPlanet.set(pos, tile{tileType: tileTypeEmpty})
			case char == '#':
				newPlanet.set(pos, tile{tileType: tileTypeInfested})
			}
		}
	}
	return newPlanet
}
