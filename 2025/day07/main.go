package main

import (
	_ "embed"
	"fmt"
	"slices"

	"github.com/adsmf/adventofcode/utils"
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
	g := grid{}
	g.load()
	beams, next := []int{g.startCol}, []int{}
	splits := 0
	for row := 0; row < g.h; row++ {
		for _, beam := range beams {
			beamPos := point{beam, row}
			if g.tiles[beamPos] == tileSplitter {
				next = append(next, beam-1, beam+1)
				splits++
			} else {
				next = append(next, beam)
			}
		}
		next = slices.Compact(next)
		next, beams = beams[0:0], next
	}
	return splits
}

func part2() int {
	g := grid{}
	g.load()
	beams, next := make([]int, g.w), make([]int, g.w)
	beams[g.startCol] = 1
	splits := 0
	for row := 0; row < g.h; row++ {
		for beam, count := range beams {
			beamPos := point{beam, row}
			if g.tiles[beamPos] == tileSplitter {
				next[beam-1] += count
				next[beam+1] += count
				splits++
			} else {
				next[beam] += count
			}
		}
		clear(beams)
		next, beams = beams, next
	}
	count := 0
	for _, beam := range beams {
		count += beam
	}
	return count
}

type grid struct {
	tiles    map[point]tileType
	h, w     int
	startCol int
}

func (g *grid) load() {
	g.tiles = map[point]tileType{}
	utils.EachLine(input, func(y int, line string) (done bool) {
		g.w = len(line)
		g.h = y
		for x, ch := range line {
			pos := point{x, y}
			switch ch {
			case '^':
				g.tiles[pos] = tileSplitter
			case 'S':
				g.startCol = x
			}
		}
		return
	})
}

type tileType byte

const (
	tileEmpty tileType = iota
	tileSplitter
	tileBeam
)

type point struct{ x, y int }

var benchmark = false
