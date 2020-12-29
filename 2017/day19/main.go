package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	tubes := load("input.txt")
	p1, p2 := tubes.follow()
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type tubeMap struct {
	grid     map[point]tubeType
	targets  map[point]byte
	height   int
	width    int
	entrance point
}

func (t tubeMap) follow() (string, int) {
	collected := ""
	pos := t.entrance
	dir := point{0, 1}

	count := 0
	for {
		if target, found := t.targets[pos]; found {
			collected += string(target)
		}
		switch t.grid[pos] {
		case tubeStraight:
			pos = pos.add(dir)
		case tubeCorner:
			nextDir := point{dir.y, dir.x}
			nextPos := pos.add(nextDir)
			if t.grid[nextPos] != tubeNone {
				pos = nextPos
				dir = nextDir
			} else {
				dir = point{-nextDir.x, -nextDir.y}
				pos = pos.add(dir)
			}
		case tubeNone:
			return collected, count
		default:
			panic(fmt.Errorf("Unhandled case: %v", t.grid[pos]))
		}
		count++
	}
}

const (
	tubeNone tubeType = iota
	tubeStraight
	tubeCorner
)

type tubeType int

type point struct{ x, y int }

func (p point) add(a point) point {
	return point{p.x + a.x, p.y + a.y}
}

func load(filename string) tubeMap {
	inputRaw, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(inputRaw), "\n")

	tubes := tubeMap{
		grid:    map[point]tubeType{},
		targets: map[point]byte{},
		height:  len(lines),
	}

	for y, line := range lines {
		for x, char := range line {
			pos := point{x, y}
			if pos.x+1 > tubes.width {
				tubes.width = pos.x + 1
			}
			switch char {
			case '|', '-':
				tubes.grid[pos] = tubeStraight
				if y == 0 {
					tubes.entrance = pos
				}
			case '+':
				tubes.grid[pos] = tubeCorner
			case ' ':
			default:
				if char >= 'A' && char <= 'Z' {
					tubes.grid[pos] = tubeStraight
					tubes.targets[pos] = byte(char)
				} else {
					panic(fmt.Errorf("Unexpected character: %c", char))
				}
			}
		}
	}

	return tubes
}

var benchmark = false
