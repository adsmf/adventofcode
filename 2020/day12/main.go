package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	commands := utils.ReadInputLines("input.txt")
	p1 := part1(commands)
	p2 := part2(commands)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(commands []string) int {
	ship := turtle{}
	for _, command := range commands {
		char := command[0]
		arg, _ := strconv.Atoi(command[1:])
		switch char {
		case 'N', 'E', 'S', 'W':
			ship.move(charDirs[char], arg)

		case 'F':
			ship.move(ship.facing, arg)

		case 'L':
			arg /= 90
			for i := 0; i < arg; i++ {
				ship.left()
			}
		case 'R':
			arg /= 90
			for i := 0; i < arg; i++ {
				ship.right()
			}
		}
	}
	return ship.pos.manhattan()
}

func part2(commands []string) int {
	ship := turtle{}
	waypoint := turtle{pos: vector{10, 1}}
	for _, command := range commands {
		char := command[0]
		arg, _ := strconv.Atoi(command[1:])
		switch char {
		case 'N', 'E', 'S', 'W':
			waypoint.move(charDirs[char], arg)

		case 'F':
			ship.pos = ship.pos.add(waypoint.pos.scale(arg))

		case 'L':
			arg /= 90
			for i := 0; i < arg; i++ {
				waypoint.pos = waypoint.pos.rotateCCW()
			}
		case 'R':
			arg /= 90
			for i := 0; i < arg; i++ {
				waypoint.pos = waypoint.pos.rotateCW()
			}
		}
	}
	return ship.pos.manhattan()
}

var charDirs = map[byte]compassDirection{
	'N': compassNorth,
	'E': compassEast,
	'S': compassSouth,
	'W': compassWest,
}

type turtle struct {
	pos    vector
	facing compassDirection
}

func (t *turtle) right() { t.facing = t.facing.right() }
func (t *turtle) left()  { t.facing = t.facing.left() }
func (t *turtle) move(dir compassDirection, distance int) {
	t.pos = t.pos.add(dir.unitVector().scale(distance))
}

type vector struct{ x, y int }

func (v vector) scale(mult int) vector   { return vector{v.x * mult, v.y * mult} }
func (v vector) add(other vector) vector { return vector{v.x + other.x, v.y + other.y} }
func (v vector) manhattan() int          { return int(math.Abs(float64(v.x)) + math.Abs(float64(v.y))) }
func (v vector) rotateCW() vector        { return vector{v.y, -v.x} }
func (v vector) rotateCCW() vector       { return vector{-v.y, v.x} }

type compassDirection int

const (
	compassEast compassDirection = iota
	compassSouth
	compassWest
	compassNorth
)

func (c compassDirection) right() compassDirection {
	return (c + 1) % 4
}
func (c compassDirection) left() compassDirection {
	new := (c - 1)
	if new < 0 {
		new += 4
	}
	return new
}

var unitVectors = map[compassDirection]vector{
	compassNorth: {0, 1},
	compassEast:  {1, 0},
	compassSouth: {0, -1},
	compassWest:  {-1, 0},
}

func (c compassDirection) unitVector() vector { return unitVectors[c] }

var benchmark = false
