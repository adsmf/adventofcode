package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	floorplan := load("input.txt")
	p1 := seatLife(floorplan, false)
	p2 := seatLife(floorplan, true)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func seatLife(floorplan area, part2rules bool) int {
	iteration := 0
	states := map[string]int{}
	states[floorplan.save()]++
	for {
		floorplan = floorplan.next(part2rules)
		state := floorplan.save()
		iteration++
		if _, found := states[state]; found {
			return floorplan.countOccupied()
		}
		states[state]++
	}
}

type area struct {
	grid   map[pos]floorTileType
	height int
	width  int
}

func (a area) next(part2rules bool) area {
	next := area{
		grid:   map[pos]floorTileType{},
		width:  a.width,
		height: a.height,
	}
	for loc, tile := range a.grid {
		if tile == tileEmptyFloor {
			continue
		}
		occupiedAdjacent := 0
		var tolerance int
		var surroundings []pos
		if part2rules {
			surroundings = loc.sight(a)
			tolerance = 5
		} else {
			surroundings = loc.adjacent()
			tolerance = 4
		}
		for _, surLoc := range surroundings {
			if a.grid[surLoc] == tileOccupiedSeat {
				occupiedAdjacent++
			}
		}
		switch {
		case tile == tileEmptySeat && occupiedAdjacent == 0:
			next.grid[loc] = tileOccupiedSeat
		case tile == tileOccupiedSeat && occupiedAdjacent >= tolerance:
			next.grid[loc] = tileEmptySeat
		default:
			next.grid[loc] = tile
		}
	}
	return next
}

func (a area) save() string {
	state := ""
	for y := 0; y <= a.height; y++ {
		for x := 0; x <= a.width; x++ {
			switch a.grid[pos{x, y}] {
			case tileEmptyFloor:
				state += "."
			case tileOccupiedSeat:
				state += "#"
			case tileEmptySeat:
				state += "L"
			}
		}
		state += "\n"
	}
	return state
}

func (a area) countOccupied() int {
	count := 0
	for y := 0; y <= a.height; y++ {
		for x := 0; x <= a.width; x++ {
			if a.grid[pos{x, y}] == tileOccupiedSeat {
				count++
			}
		}
	}
	return count
}

func (a *area) fixSize() {
	for loc := range a.grid {
		if loc.x > a.width {
			a.width = loc.x
		}
		if loc.y > a.height {
			a.height = loc.y
		}
	}
}

type floorTileType int

const (
	tileEmptyFloor floorTileType = iota
	tileEmptySeat
	tileOccupiedSeat
)

type pos struct {
	x, y int
}

func (p pos) adjacent() []pos {
	surroundings := []pos{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			surroundings = append(surroundings, pos{p.x + x, p.y + y})
		}
	}
	return surroundings
}

func (p pos) sight(floorplan area) []pos {
	surroundings := []pos{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			for i := 1; i < 99; i++ {
				tryPos := pos{p.x + x*i, p.y + y*i}
				if tryPos.x < 0 || tryPos.y < 0 || tryPos.x > floorplan.width || tryPos.y > floorplan.height {
					break
				}
				if floorplan.grid[tryPos] != tileEmptyFloor {
					surroundings = append(surroundings, tryPos)
					break
				}
			}
		}
	}
	return surroundings
}

func load(filename string) area {
	inputBytes, _ := ioutil.ReadFile(filename)
	x, y := 0, 0
	floorplan := area{
		grid: map[pos]floorTileType{},
	}
	for _, char := range inputBytes {
		switch char {
		case '.':
			floorplan.grid[pos{x, y}] = tileEmptyFloor
			x++
		case 'L':
			floorplan.grid[pos{x, y}] = tileEmptySeat
			x++
		case '#':
			floorplan.grid[pos{x, y}] = tileOccupiedSeat
			x++
		case '\n':
			x = 0
			y++
		}
	}
	floorplan.fixSize()
	return floorplan
}

var benchmark = false
