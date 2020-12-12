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
	for y := 0; y < a.height; y++ {
		for x := 0; x < a.width; x++ {
			loc := x + y*a.width
			tile := a.grid[loc]
			if tile == tileEmptyFloor {
				next.grid[loc] = tileEmptyFloor
				continue
			}
			occupied := 0
			var tolerance int
			if part2rules {
				occupied = a.countSight(x, y)
				tolerance = 5
			} else {
				occupied = a.countAdjacent(x, y)
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

func (a area) countAdjacent(x, y int) int {
	count := 0
	for _, dir := range directions {
		checkX := x + dir.x
		if checkX < 0 || checkX >= a.width {
			continue
		}
		checkY := y + dir.y
		checkIndex := checkX + checkY*a.width
		if checkIndex < 0 || checkIndex >= len(a.grid) {
			continue
		}
		if a.grid[checkIndex] == tileOccupiedSeat {
			count++
		}
	}
	return count
}

func (a area) countSight(x, y int) int {
	count := 0
	for _, dir := range directions {
		for checkX, checkY := x+dir.x, y+dir.y; ; checkX, checkY = checkX+dir.x, checkY+dir.y {
			if checkX < 0 || checkX >= a.width {
				break
			}
			index := checkX + checkY*a.width
			if index < 0 || index >= len(a.grid) {
				break
			}
			tile := a.grid[index]
			if tile != tileEmptyFloor {
				if tile == tileOccupiedSeat {
					count++
				}
				break
			}
		}
	}
	return count
}

var directions []vector = []vector{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

type vector struct{ x, y int }

type floorTileType byte

const (
	tileEmptyFloor   floorTileType = '.'
	tileEmptySeat    floorTileType = 'L'
	tileOccupiedSeat floorTileType = '#'
)

func load(filename string) area {
	inputBytes, _ := ioutil.ReadFile(filename)
	floorplan := area{
		grid: make(areaGrid, 0, len(inputBytes)),
	}
	height := 0
	for _, char := range inputBytes {
		tile := floorTileType(char)
		switch tile {
		case tileEmptyFloor, tileEmptySeat, tileOccupiedSeat:
			floorplan.grid = append(floorplan.grid, tile)
		case '\n':
			height++
		}
	}
	floorplan.width = len(inputBytes)/height - 1
	floorplan.height = height
	return floorplan
}

var benchmark = false
