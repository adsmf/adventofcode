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

type areaGrid []floorTileType

type area struct {
	grid   areaGrid
	height int
	width  int
}

func (a area) next(part2rules bool) (area, bool) {
	next := area{
		grid:   make(areaGrid, len(a.grid)),
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
			occupied = a.countSight(loc)
			tolerance = 5
		} else {
			occupied = a.countAdjacent(loc)
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

func (a area) countOccupied() int {
	count := 0
	for _, tile := range a.grid {
		if tile == tileOccupiedSeat {
			count++
		}
	}
	return count
}

func (a area) countAdjacent(index int) int {
	count := 0
	indX := index % a.width
	indY := (index - indX) / a.width
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			checkY := indY + y
			checkX := indX + x
			if checkX < 0 || checkX >= a.width || checkY < 0 || checkY >= a.height {
				continue
			}
			checkIndex := index + x + y*a.width
			if a.grid[checkIndex] == tileOccupiedSeat {
				count++
			}
		}
	}
	return count
}

func (a area) countSight(index int) int {
	count := 0
	indX := index % a.width
	indY := (index - indX) / a.width
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			for checkX, checkY := indX+x, indY+y; ; checkX, checkY = checkX+x, checkY+y {
				if checkX < 0 || checkX >= a.width {
					break
				}
				checkIndex := checkX + checkY*a.width
				if checkIndex < 0 || checkIndex >= len(a.grid) {
					break
				}
				if a.grid[checkIndex] != tileEmptyFloor {
					if a.grid[checkIndex] == tileOccupiedSeat {
						count++
					}
					break
				}
			}
		}
	}
	return count
}

type floorTileType byte

const (
	tileEmptyFloor   floorTileType = '.'
	tileEmptySeat    floorTileType = 'L'
	tileOccupiedSeat floorTileType = '#'
)

func load(filename string) area {
	inputBytes, _ := ioutil.ReadFile(filename)
	x, y := 0, 0
	floorplan := area{
		grid: make(areaGrid, 0, len(inputBytes)),
	}
	maxX := 0
	maxY := 0
	for _, char := range inputBytes {
		tile := floorTileType(char)
		switch tile {
		case tileEmptyFloor, tileEmptySeat, tileOccupiedSeat:
			floorplan.grid = append(floorplan.grid, tile)
			x++
			if x > maxX {
				maxX = x
			}
		case '\n':
			x = 0
			y++
			maxY = y
		}
	}
	floorplan.width = maxX
	floorplan.height = maxY
	return floorplan
}

var benchmark = false
