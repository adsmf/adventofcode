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
	var changed bool
	for {
		floorplan, changed = floorplan.next(part2rules)
		if !changed {
			return floorplan.countOccupied()
		}
		iteration++
	}
}

type area struct {
	grid   map[pos]floorTileType
	height int
	width  int
}

func (a area) next(part2rules bool) (area, bool) {
	next := area{
		grid:   map[pos]floorTileType{},
		width:  a.width,
		height: a.height,
	}
	changes := false
	for loc, tile := range a.grid {
		if tile == tileEmptyFloor {
			next.grid[loc] = tileEmptyFloor
			continue
		}
		occupied := 0
		var tolerance int
		if part2rules {
			occupied = loc.countSight(a)
			tolerance = 5
		} else {
			occupied = loc.countAdjacent(a)
			tolerance = 4
		}
		switch {
		case tile == tileEmptySeat && occupied == 0:
			next.grid[loc] = tileOccupiedSeat
			changes = true
		case tile == tileOccupiedSeat && occupied >= tolerance:
			next.grid[loc] = tileEmptySeat
			changes = true
		default:
			next.grid[loc] = tile
		}
	}
	return next, changes
}

func (a area) save() string {
	state := make([]byte, 0, a.height*(a.width+1))
	for y := 0; y <= a.height; y++ {
		for x := 0; x <= a.width; x++ {
			state = append(state, byte(a.grid[pos{x, y}]))
		}
		state = append(state, '\n')
	}
	return string(state)
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

type floorTileType byte

const (
	tileOutside      floorTileType = 0
	tileEmptyFloor   floorTileType = '.'
	tileEmptySeat    floorTileType = 'L'
	tileOccupiedSeat floorTileType = '#'
)

type pos struct {
	x, y int
}

func (p pos) countAdjacent(floorplan area) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			if floorplan.grid[pos{p.x + x, p.y + y}] == tileOccupiedSeat {
				count++
			}
		}
	}
	return count
}

func (p pos) countSight(floorplan area) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			for i := 1; i < 99; i++ {
				tryPos := pos{p.x + x*i, p.y + y*i}
				if floorplan.grid[tryPos] == tileOutside {
					break
				}
				if floorplan.grid[tryPos] != tileEmptyFloor {
					if floorplan.grid[tryPos] == tileOccupiedSeat {
						count++
					}
					break
				}
			}
		}
	}
	return count
}

func load(filename string) area {
	inputBytes, _ := ioutil.ReadFile(filename)
	x, y := 0, 0
	floorplan := area{
		grid: map[pos]floorTileType{},
	}
	for _, char := range inputBytes {
		tile := floorTileType(char)
		switch tile {
		case tileEmptyFloor, tileEmptySeat, tileOccupiedSeat:
			floorplan.grid[pos{x, y}] = tile
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
