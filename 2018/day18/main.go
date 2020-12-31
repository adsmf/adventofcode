package main

import (
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := lumberLife(10)
	p2 := lumberLife(1000000000)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func lumberLife(maxIter int) int {
	forestry := load("input.txt")
	states := map[string]int{}

	for i := maxIter; i > 0; i-- {
		forestry = forestry.iterate()
		state := forestry.String()
		if prevTurn, found := states[state]; found {
			cycleLength := prevTurn - i
			i %= cycleLength
			for ; i > 1; i-- {
				forestry = forestry.iterate()
			}
			return forestry.value()
		}
		states[state] = i
	}
	return forestry.value()
}

type forestryMap struct {
	grid          map[point]forestryTile
	height, width int
}

func (f forestryMap) iterate() forestryMap {
	nextMap := forestryMap{
		height: f.height,
		width:  f.width,
		grid:   make(map[point]forestryTile, len(f.grid)),
	}

	adjacentTrees := map[point]int{}
	adjacentYards := map[point]int{}
	for pos, tile := range f.grid {
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				if x == 0 && y == 0 {
					continue
				}
				adjPos := point{pos.x + x, pos.y + y}
				switch tile {
				case tileLumberyard:
					adjacentYards[adjPos]++
				case tileTrees:
					adjacentTrees[adjPos]++
				}
			}
		}
	}
	for pos, tile := range f.grid {
		switch tile {
		case tileOpenGround:
			if adjacentTrees[pos] >= 3 {
				nextMap.grid[pos] = tileTrees
			} else {
				nextMap.grid[pos] = tileOpenGround
			}
		case tileTrees:
			if adjacentYards[pos] >= 3 {
				nextMap.grid[pos] = tileLumberyard
			} else {
				nextMap.grid[pos] = tileTrees
			}
		case tileLumberyard:
			if adjacentTrees[pos] >= 1 && adjacentYards[pos] >= 1 {
				nextMap.grid[pos] = tileLumberyard
			} else {
				nextMap.grid[pos] = tileOpenGround
			}
		}
	}

	return nextMap
}

func (f forestryMap) value() int {
	trees, lumberyards := 0, 0

	for _, tile := range f.grid {
		switch tile {
		case tileLumberyard:
			lumberyards++
		case tileTrees:
			trees++
		}
	}

	return trees * lumberyards
}

func (f forestryMap) String() string {
	sb := &strings.Builder{}
	sb.Grow(f.height * (f.width + 1))
	for y := 0; y < f.height; y++ {
		for x := 0; x < f.width; x++ {
			sb.WriteByte(tileEncode[f.grid[point{x, y}]])
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

type forestryTile int

const (
	tileOpenGround forestryTile = iota
	tileTrees
	tileLumberyard
)

var tileDecode = map[byte]forestryTile{'.': tileOpenGround, '|': tileTrees, '#': tileLumberyard}
var tileEncode = map[forestryTile]byte{tileOpenGround: '.', tileTrees: '|', tileLumberyard: '#'}

type point struct{ x, y int }

func load(filename string) forestryMap {
	forestry := forestryMap{
		grid: map[point]forestryTile{},
	}

	lines := utils.ReadInputLines(filename)
	for y, line := range lines {
		for x, char := range line {
			pos := point{x, y}
			forestry.grid[pos] = tileDecode[byte(char)]
		}
	}
	forestry.height = len(lines)
	forestry.width = len(lines[0])

	return forestry
}

var benchmark = false
