package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	scan := load("input.txt")
	risk := 0
	for y := 0; y <= scan.target.y; y++ {
		for x := 0; x <= scan.target.x; x++ {
			pos := point{x, y}
			risk += int(scan.getErosionLevel(pos)) % 3
		}
	}
	return risk
}

func part2() int {
	return -1
}

type erosionLevel int

const (
	elRocky erosionLevel = iota
	elWet
	elNarrow
)

var elRepr = map[erosionLevel]byte{elRocky: '.', elWet: '=', elNarrow: '|'}

func (el erosionLevel) String() string {
	return string(elRepr[el%3])
}
func (el erosionLevel) asByte() byte {
	return elRepr[el%3]
}

type caveScan struct {
	grid   map[point]int
	depth  int
	target point
}

func (c caveScan) String() string {
	sb := &strings.Builder{}
	sb.Grow((c.target.x + 2) * (c.target.y + 1))
	for y := 0; y <= c.target.y; y++ {
		for x := 0; x <= c.target.x; x++ {
			if x == 0 && y == 0 {
				sb.WriteByte('M')
				continue
			}
			pos := point{x, y}
			if pos == c.target {
				sb.WriteByte('T')
				continue
			}
			sb.WriteByte(c.getErosionLevel(pos).asByte())
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (c *caveScan) getGeoIndex(pos point) int {
	if gi, found := c.grid[pos]; found {
		return gi
	}
	var gi int
	if pos.y == 0 {
		gi = pos.x * 16807
	} else if pos.x == 0 {
		gi = pos.y * 48271
	} else {
		gi = int(c.getErosionLevel(point{pos.x - 1, pos.y})) * int(c.getErosionLevel(point{pos.x, pos.y - 1}))
	}
	c.grid[pos] = gi
	return gi
}

func (c *caveScan) getErosionLevel(pos point) erosionLevel {
	return erosionLevel((c.getGeoIndex(pos) + c.depth) % 20183)
}

type point struct{ x, y int }

func load(filename string) caveScan {
	inputBytes, _ := ioutil.ReadFile(filename)
	ints := utils.GetInts(string(inputBytes))
	depth, target := ints[0], point{ints[1], ints[2]}
	return caveScan{
		grid:   map[point]int{{0, 0}: 0, target: 0},
		depth:  depth,
		target: target,
	}
}

var benchmark = false
