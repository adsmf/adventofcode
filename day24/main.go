package main

import (
	"fmt"
	"math"
	"sort"

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
	p := loadMap("input.txt")
	return p.iterRecursiveN(200)
}

type planes map[int]plane
type plane map[point]tile

type planet struct {
	grids      planes
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

func (p *planet) iterRecursiveN(steps int) int {
	midX := p.maxX / 2
	midY := p.maxY / 2

	delete(p.grids[0], point{midX, midY, 0})

	lastBugs := 0
	for i := 0; i < steps; i++ {
		lastBugs = p.iterRecursive()
	}

	return lastBugs
}

func (p *planet) iter() {
	newGrid := plane{}

	for pos, t := range p.grids[0] {
		infested := t.tileType == tileTypeInfested
		surrounding := 0
		for _, neighbour := range pos.neighbours() {
			if p.grids[0][neighbour].tileType == tileTypeInfested {
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
	p.grids[0] = newGrid
}

func (p *planet) iterRecursive() int {

	width := p.maxX - p.minX + 1
	height := p.maxY - p.minY + 1

	levels := []int{}
	for level := range p.grids {
		levels = append(levels, level)
	}
	sort.Ints(levels)
	lowestLevel := levels[0]
	highestLevel := levels[len(levels)-1]
	p.set(point{0, 0, lowestLevel - 1}, tile{tileType: tileTypeEmpty})
	p.set(point{0, 0, highestLevel + 1}, tile{tileType: tileTypeEmpty})

	newGrids := planes{}
	bugs := 0
	for level, grid := range p.grids {
		newGrid := plane{}

		for pos, t := range grid {
			infested := t.tileType == tileTypeInfested
			surrounding := 0
			for _, neighbour := range pos.recursiveNeighbours(width, height) {
				if p.infested(neighbour) {
					surrounding++
				}
			}
			if surrounding == 1 ||
				(!infested && surrounding == 2) {
				newGrid[pos] = tile{tileType: tileTypeInfested}
				bugs++
			} else {
				newGrid[pos] = tile{tileType: tileTypeEmpty}
			}
		}
		newGrids[level] = newGrid

	}
	p.grids = newGrids
	return bugs
}

func (p *planet) set(pos point, val tile) {
	val.pos = pos
	if p.grids[pos.level] == nil {
		p.grids[pos.level] = plane{}
		for pos0 := range p.grids[0] {
			posNew := pos0
			posNew.level = pos.level
			p.grids[posNew.level][posNew] = tile{tileType: tileTypeEmpty, pos: posNew}
		}
	}
	p.grids[pos.level][pos] = val
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
func (p *planet) infested(pos point) bool {
	if p.grids[pos.level] == nil {
		return false
	}
	return p.grids[pos.level][pos].tileType == tileTypeInfested
}

func (p planet) String() string {
	newString := ""
	levels := []int{}
	for level := range p.grids {
		levels = append(levels, level)
	}
	sort.Ints(levels)
	for _, level := range levels {
		newString += fmt.Sprintf("Level %d\n", level)
		grid := p.grids[level]
		for y := p.minY; y <= p.maxY; y++ {
			for x := p.minX; x <= p.maxX; x++ {
				newString += fmt.Sprintf("%v", grid[point{x, y, level}])
			}
			newString += fmt.Sprintln()
		}
		newString += fmt.Sprintln()
	}
	return newString
}

func (p planet) biodiversity() int {
	bio := 0
	for pos, t := range p.grids[0] {
		if t.tileType == tileTypeInfested {
			id := pos.x + pos.y*(p.maxX+1)
			bio += int(math.Pow(2, float64(id)))
		}
	}
	return bio
}

type point struct {
	x, y  int
	level int
}

func (p point) neighbours() []point {
	return []point{
		point{p.x - 1, p.y, p.level},
		point{p.x + 1, p.y, p.level},
		point{p.x, p.y - 1, p.level},
		point{p.x, p.y + 1, p.level},
	}
}

func (p point) String() string {
	return fmt.Sprintf("{%d, %d, %d}", p.x, p.y, p.level)
}

func (p point) recursiveNeighbours(w, h int) []point {
	midX := (w - 1) / 2
	midY := (h - 1) / 2

	basePoints := []point{
		point{p.x - 1, p.y, p.level},
		point{p.x + 1, p.y, p.level},
		point{p.x, p.y - 1, p.level},
		point{p.x, p.y + 1, p.level},
	}
	realPoints := []point{}

	for _, pos := range basePoints {
		normal := true
		// Check adjacency to center
		if pos.x == midX && pos.y == midY {
			if p.y == midY+1 {
				normal = false
				for x2 := 0; x2 < w; x2++ {
					realPoints = append(realPoints, point{x2, h - 1, pos.level - 1})
				}
			} else if p.y == midY-1 {
				normal = false
				for x2 := 0; x2 < w; x2++ {
					realPoints = append(realPoints, point{x2, 0, pos.level - 1})
				}
			} else if p.x == midX+1 {
				normal = false
				for y2 := 0; y2 < h; y2++ {
					realPoints = append(realPoints, point{w - 1, y2, pos.level - 1})
				}
			} else if p.x == midX-1 {
				normal = false
				for y2 := 0; y2 < h; y2++ {
					realPoints = append(realPoints, point{0, y2, pos.level - 1})
				}
			}
		}

		// Check adjacency to edge
		if pos.x == -1 {
			normal = false
			realPoints = append(realPoints, point{midX - 1, midY, pos.level + 1})
		} else if pos.x == w {
			normal = false
			realPoints = append(realPoints, point{midX + 1, midY, pos.level + 1})
		}

		if pos.y == -1 {
			normal = false
			realPoints = append(realPoints, point{midX, midY - 1, pos.level + 1})
		} else if pos.y == h {
			normal = false
			realPoints = append(realPoints, point{midX, midY + 1, pos.level + 1})
		}

		if normal {
			realPoints = append(realPoints, pos)
		}
	}
	return realPoints
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
		grids: planes{},
	}

	lines := utils.ReadInputLines(filename)
	for y, line := range lines {
		for x, char := range line {
			pos := point{x, y, 0}
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
