package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	g := loadGrid()
	counts := map[point]int{}
	for pos, tile := range g.tiles {
		if tile != tilePaper {
			continue
		}
		for xOff := -1; xOff <= 1; xOff++ {
			for yOff := -1; yOff <= 1; yOff++ {
				adj := point{pos.x + xOff, pos.y + yOff}
				if adj != pos && g.tiles[adj] == tilePaper {
					counts[adj]++
				}
			}
		}
	}
	accessible := 0
	for pos, tile := range g.tiles {
		if tile == tilePaper && counts[pos] < 4 {
			accessible++
		}

	}
	return accessible
}

func part2() int {
	g := loadGrid()
	removed := 0
	done := false
	for !done {
		counts := map[point]int{}
		done = true
		for pos, tile := range g.tiles {
			if tile != tilePaper {
				continue
			}
			for xOff := -1; xOff <= 1; xOff++ {
				for yOff := -1; yOff <= 1; yOff++ {
					adj := point{pos.x + xOff, pos.y + yOff}
					if adj != pos && g.tiles[adj] == tilePaper {
						counts[adj]++
					}
				}
			}
		}
		for pos, tile := range g.tiles {
			if tile == tilePaper && counts[pos] < 4 {
				removed++
				done = false
				g.tiles[pos] = tileEmpty
			}
		}
	}
	return removed
}

type point struct{ x, y int }

type gridTiles map[point]tileType

func loadGrid() grid {
	g := grid{
		tiles: make(gridTiles),
	}
	x, y := 0, 0
	for pos := 0; pos < len(input); pos++ {
		switch input[pos] {
		case '@':
			g.tiles[point{x, y}] = tilePaper
			x++
		case '.':
			g.tiles[point{x, y}] = tileEmpty
			x++
		case '\n':
			g.w = x
			x = 0
			y++
		}
	}
	g.h = y
	return g
}

type grid struct {
	h, w  int
	tiles gridTiles
}

func (g grid) inBound(p point) bool {
	return p.x >= 0 && p.x < g.w && p.y >= 0 && p.y < g.h
}

type tileType int

const (
	tileOOB tileType = iota
	tileEmpty
	tilePaper
)

var benchmark = false
