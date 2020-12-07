package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var benchmark = false

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	route := load(string(inputBytes))
	p1 := part1(route)
	p2 := part2(route)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(route []vector) int {
	return route[len(route)-1].manhattan()
}

func part2(route []vector) int {
	visited := map[vector]struct{}{}
	for _, pos := range route {
		if _, found := visited[pos]; found {
			return pos.manhattan()
		}
		visited[pos] = struct{}{}
	}
	return -1
}

func load(input string) []vector {
	route := []vector{}
	tur := turtle{}
	for _, command := range strings.Split(input, ", ") {
		command = strings.TrimSpace(command)
		switch command[0] {
		case 'L':
			tur.left()
		case 'R':
			tur.right()
		}
		distance, _ := strconv.Atoi(command[1:])
		for i := 0; i < distance; i++ {
			tur.forwards(1)
			route = append(route, tur.pos)
		}
	}
	return route
}

type turtle struct {
	pos    vector
	facing compassDirection
}

func (t *turtle) right() { t.facing = t.facing.right() }
func (t *turtle) left()  { t.facing = t.facing.left() }
func (t *turtle) forwards(distance int) {
	t.pos = t.pos.add(t.unitFacing().scale(distance))
}
func (t *turtle) unitFacing() vector {
	vec := vector{}
	switch t.facing {
	case compassNorth:
		vec = vector{0, 1}
	case compassEast:
		vec = vector{1, 0}
	case compassSouth:
		vec = vector{0, -1}
	case compassWest:
		vec = vector{-1, 0}
	}
	return vec
}

type vector struct{ x, y int }

func (v vector) scale(mult int) vector   { return vector{v.x * mult, v.y * mult} }
func (v vector) add(other vector) vector { return vector{v.x + other.x, v.y + other.y} }
func (v vector) manhattan() int {
	dist := 0
	if v.x < 0 {
		dist -= v.x
	} else {
		dist += v.x
	}
	if v.y < 0 {
		dist -= v.y
	} else {
		dist += v.y
	}
	return dist
}

type compassDirection int

const (
	compassNorth compassDirection = iota
	compassEast
	compassSouth
	compassWest
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
