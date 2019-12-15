package main

import (
	"fmt"
	"github.com/adsmf/adventofcode2019/utils/intcode"
	"io/ioutil"
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
	inputString := loadInputString()
	region := exploreAll(inputString, point{0, 0})
	return fillOxygen(region)
}

func fillOxygen(region area) int {
	oxyMap := area{
		grid: grid{},
		minX: region.minX,
		maxX: region.maxX,
		minY: region.minX,
		maxY: region.maxY,
	}
	for pos, tile := range region.grid {
		oxyMap.grid[pos] = tile
	}
	steps := 0
	for ; oxyMap.countEmpty() > 0; steps++ {
		nextFill := []point{}
		for pos, tile := range oxyMap.grid {
			if tile == tileEmpty {
				if oxyMap.grid[point{pos.x + 1, pos.y}] == tileOxygen ||
					oxyMap.grid[point{pos.x - 1, pos.y}] == tileOxygen ||
					oxyMap.grid[point{pos.x, pos.y + 1}] == tileOxygen ||
					oxyMap.grid[point{pos.x, pos.y - 1}] == tileOxygen {
					nextFill = append(nextFill, pos)
				}
			}
		}
		for _, fill := range nextFill {
			oxyMap.set(fill, tileOxygen)
		}
	}
	return steps
}

func exploreAll(program string, start point) area {
	region := area{
		grid: grid{
			point{0, 0}: tileEmpty,
		},
	}
	type routeState struct {
		nextStep      int
		previousState []byte
		lastPos       point
	}

	routes := []routeState{
		routeState{1, []byte{}, start},
		routeState{2, []byte{}, start},
		routeState{3, []byte{}, start},
		routeState{4, []byte{}, start},
	}

	lastRegionLen := -1
	for {
		nextRoutes := []routeState{}
		for _, route := range routes {
			wall, _, state, lastPos := tryRoute(program, &region, []int{route.nextStep}, route.previousState, route.lastPos)
			if !wall {
				if route.nextStep != 1 {
					nextRoutes = append(nextRoutes, routeState{2, state, lastPos})
				}
				if route.nextStep != 2 {
					nextRoutes = append(nextRoutes, routeState{1, state, lastPos})
				}
				if route.nextStep != 3 {
					nextRoutes = append(nextRoutes, routeState{4, state, lastPos})
				}
				if route.nextStep != 4 {
					nextRoutes = append(nextRoutes, routeState{3, state, lastPos})
				}
			}
		}
		if len(nextRoutes) == 0 {
			break
		}
		routes = nextRoutes

		newRegionLen := len(region.grid)
		if lastRegionLen == newRegionLen {
			break
		}
		lastRegionLen = newRegionLen
	}

	return region
}

func findOxygen(program string) int {
	region := area{
		grid: grid{
			point{0, 0}: tileEmpty,
		},
	}

	type routeState struct {
		nextStep      int
		previousState []byte
		lastPos       point
	}

	routes := []routeState{
		routeState{1, []byte{}, point{0, 0}},
		routeState{2, []byte{}, point{0, 0}},
		routeState{3, []byte{}, point{0, 0}},
		routeState{4, []byte{}, point{0, 0}},
	}

	for i := 1; i < 400; i++ {
		nextRoutes := []routeState{}
		for _, route := range routes {
			wall, oxygen, state, lastPos := tryRoute(program, &region, []int{route.nextStep}, route.previousState, route.lastPos)
			if oxygen {
				return i
			}
			if !wall {
				if route.nextStep != 1 {
					nextRoutes = append(nextRoutes, routeState{2, state, lastPos})
				}
				if route.nextStep != 2 {
					nextRoutes = append(nextRoutes, routeState{1, state, lastPos})
				}
				if route.nextStep != 3 {
					nextRoutes = append(nextRoutes, routeState{4, state, lastPos})
				}
				if route.nextStep != 4 {
					nextRoutes = append(nextRoutes, routeState{3, state, lastPos})
				}
			}
		}
		if len(nextRoutes) == 0 {
			panic("No routes left!")
		}
		routes = nextRoutes
	}

	return 0
}

func tryRoute(program string, region *area, inputs []int, previousState []byte, start point) (bool, bool, []byte, point) {
	mapper := robot{
		tiles:     region,
		inputList: inputs,
		position:  start,
	}
	cpu := intcode.NewMachine(intcode.M19(mapper.guided, mapper.outputCallback))
	if len(previousState) > 0 {
		cpu.Restore(previousState)
	} else {
		cpu.LoadProgram(program)
	}

	mapper.cpu = &cpu

	cpu.Run(true)

	state := cpu.Save()
	return mapper.hitWall, mapper.foundOxygen, state, mapper.position
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
					newOutput += fmt.Sprint("⭐")
				} else {
					newOutput += fmt.Sprint("⬜")
				}
			case tileWall:
				newOutput += fmt.Sprint("🟥")
			case tileOxygen:
				newOutput += fmt.Sprint("🟦")
			default:
				newOutput += fmt.Sprint("⬛")
			}
		}
		newOutput += fmt.Sprintln()
	}
	return newOutput
}

func (a *area) countEmpty() int {
	count := 0
	for _, tile := range a.grid {
		if tile == tileEmpty {
			count++
		}
	}
	return count
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
	cpu         *intcode.Machine
	inputList   []int
	heading     direction
	position    point
	hitWall     bool
	foundOxygen bool
	oxygenPos   point
}

func (r *robot) guided() (int, bool) {
	if len(r.inputList) > 0 {
		nextInput := r.inputList[0]
		r.inputList = r.inputList[1:]
		r.heading = direction(nextInput)
		return nextInput, false
	}
	return 99, true
}

func (r *robot) outputCallback(out int) {
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

func loadInputString() string {
	inputRaw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputRaw)

}
