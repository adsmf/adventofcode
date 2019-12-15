package main

import (
	"fmt"
	"sync"
)

var interactive bool
var autopilot bool

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	inputString := loadInputString()
	return findOxygen(inputString)
}

func part2() int {
	return 0
}

func findOxygen(program string) int {
	region := area{
		grid: grid{
			point{0, 0}: tileEmpty,
		},
	}

	routes := [][]int64{
		[]int64{1},
		[]int64{2},
		[]int64{3},
		[]int64{4},
	}

	for i := 1; i < 400; i++ {
		nextRoutes := [][]int64{}
		for _, route := range routes {
			wall, oxygen := tryRoute(program, &region, route)
			if oxygen {
				return i
			}
			if !wall {
				lastDir := route[len(route)-1]
				if lastDir != 1 {
					copy := append(route[:0:0], route...)
					nextRoutes = append(nextRoutes, append(copy, 2))
				}
				if lastDir != 2 {
					copy := append(route[:0:0], route...)
					nextRoutes = append(nextRoutes, append(copy, 1))
				}
				if lastDir != 3 {
					copy := append(route[:0:0], route...)
					nextRoutes = append(nextRoutes, append(copy, 4))
				}
				if lastDir != 4 {
					copy := append(route[:0:0], route...)
					nextRoutes = append(nextRoutes, append(copy, 3))
				}
			}
		}
		if len(nextRoutes) == 0 {
			panic("No routes left!")
		}
		routes = nextRoutes
	}

	fmt.Printf("Region:\n%#v\n%v\n", region, region)
	return 0
}

func tryRoute(program string, region *area, inputs []int64) (bool, bool) {
	cpu := newMachine(program, nil, nil)
	mapper := robot{
		tiles:     region,
		cpu:       &cpu,
		inputList: inputs,
		position:  point{0, 0},
	}

	cpu.inputCallback = mapper.guided
	cpu.outputCallback = mapper.outputCallback

	cpu.run()

	return mapper.hitWall, mapper.foundOxygen
}

type tile int

const (
	tileUnknown tile = iota
	tileEmpty
	tileWall
	tileOxygen
)

type point struct {
	x, y int
}

type area struct {
	grid grid

	minX, minY int
	maxX, maxY int
}

func (a area) String() string {
	newOutput := ""
	for y := a.minY; y <= a.maxY; y++ {
		for x := a.minX; x <= a.maxX; x++ {
			switch a.grid[point{x, y}] {
			case tileEmpty:
				if x == 0 && y == 0 {
					newOutput += fmt.Sprint("*")
				} else {
					newOutput += fmt.Sprint(".")
				}
			case tileWall:
				newOutput += fmt.Sprint("#")
			case tileOxygen:
				newOutput += fmt.Sprint("o")
			default:
				newOutput += fmt.Sprint("?")
			}
		}
		newOutput += fmt.Sprintln()
	}
	return newOutput
}

func (a *area) set(pos point, tileType tile) {
	if pos.x > a.maxX {
		a.maxX = pos.x
	}
	if pos.x < a.minX {
		a.minX = pos.x
	}
	if pos.y > a.maxY {
		a.maxY = pos.y
	}
	if pos.y < a.minY {
		a.minY = pos.y
	}
	a.grid[pos] = tileType
}

type grid map[point]tile

type direction int

const (
	directionNorth direction = iota + 1
	directionSouth
	directionWest
	directionEast
)

type robot struct {
	tiles       *area
	cpu         *machine
	inputList   []int64
	heading     direction
	position    point
	hitWall     bool
	foundOxygen bool
	oxygenPos   point
}

func (r *robot) guided() (int64, bool) {
	if len(r.inputList) > 0 {
		nextInput := r.inputList[0]
		r.inputList = r.inputList[1:]
		r.heading = direction(nextInput)
		return nextInput, false
	}
	return 99, true
}

func (r *robot) outputCallback(out int64) {
	pos := r.position
	switch r.heading {
	case directionNorth:
		pos.y--
	case directionSouth:
		pos.y++
	case directionEast:
		pos.x++
	case directionWest:
		pos.x--
	}
	switch out {
	case 0:
		// 0: The repair droid hit a wall. Its position has not changed.
		r.tiles.set(pos, tileWall)
		r.hitWall = true
	case 1:
		// 1: The repair droid has moved one step in the requested direction.
		r.tiles.set(pos, tileEmpty)
		r.position = pos
	case 2:
		// 2: The repair droid has moved one step in the requested direction; its new position is the location of the oxygen system.
		r.tiles.set(pos, tileOxygen)
		r.position = pos
		r.foundOxygen = true
		r.oxygenPos = pos
		// fmt.Printf("Oxygen at %v!\n", pos)
	}
}

func (r *robot) outputHandler(wg *sync.WaitGroup, output chan int64) {
	for x := range output {
		r.outputCallback(x)
	}
	wg.Done()
}
