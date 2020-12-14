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

var areaBuffer1, areaBuffer2 areaGrid
var seatAdjacency = [][]int{}
var seatLinesOfSight = [][]int{}

func (a area) next(part2rules bool) (area, bool) {
	next := area{
		grid:   areaBuffer1,
		width:  a.width,
		height: a.height,
	}
	changes := false
	for loc := 0; loc < a.height*a.width; loc++ {
		tile := a.grid[loc]
		if tile == tileEmptyFloor {
			next.grid[loc] = tileEmptyFloor
			continue
		}
		occupied := 0
		var tolerance int
		if part2rules {
			occupied = a.countSeats(seatLinesOfSight, loc)
			tolerance = 5
		} else {
			occupied = a.countSeats(seatAdjacency, loc)
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
	areaBuffer1, areaBuffer2 = areaBuffer2, areaBuffer1
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

func (a area) countSeats(seatDirLokup [][]int, pos int) int {
	count := 0
	for _, index := range seatDirLokup[pos] {
		if a.grid[index] == tileOccupiedSeat {
			count++
		}
	}
	return count
}

func (a area) genSeatLookups() {
	seatLinesOfSight = make([][]int, a.height*a.width)
	seatAdjacency = make([][]int, a.height*a.width)
	for y := 0; y < a.height; y++ {
		for x := 0; x < a.width; x++ {
			pos := x + y*a.width
			seatLinesOfSight[pos] = make([]int, 0, 8)
			seatAdjacency[pos] = make([]int, 0, 8)
			for _, dir := range directions {
				for checkX, checkY, adjacent := x+dir.x, y+dir.y, true; ; checkX, checkY, adjacent = checkX+dir.x, checkY+dir.y, false {
					if checkX < 0 || checkX >= a.width {
						break
					}
					index := checkX + checkY*a.width
					if index < 0 || index >= len(a.grid) {
						break
					}
					tile := a.grid[index]
					if tile != tileEmptyFloor {
						seatLinesOfSight[pos] = append(seatLinesOfSight[pos], index)
						if adjacent {
							seatAdjacency[pos] = append(seatAdjacency[pos], index)
						}
						break
					}
				}
			}
		}
	}
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
	floorplan.genSeatLookups()
	areaBuffer1 = make(areaGrid, len(floorplan.grid))
	areaBuffer2 = make(areaGrid, len(floorplan.grid))
	return floorplan
}

var benchmark = false
