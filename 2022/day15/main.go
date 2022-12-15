package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	g := loadInput()
	return len(g.checkRow(2000000))
}

func part2() int {
	g := loadInput()
	pos := g.locate()

	return pos.x*4000000 + pos.y
}

func loadInput() grid {
	g := grid{}
	for _, line := range utils.GetLines(input) {
		vals := utils.GetInts(line)
		sensorPos := point{vals[0], vals[1]}
		closestBeacon := point{vals[2], vals[3]}
		g[sensorPos] = closestBeacon
	}
	return g
}

type grid map[point]point

func (g grid) locate() point {

	maxVal := 4000000
	coverages := map[point]int{}
	for sensor, beacon := range g {
		coverages[sensor] = sensor.manhattan(beacon)
	}
	for sensor := range g {
		coverage := coverages[sensor]
		for x := 0; x <= coverage+1; x++ {
			y := coverage + 1 - x
			tryPoints := []point{
				{sensor.x + x, sensor.y + y},
				{sensor.x - x, sensor.y + y},
				{sensor.x + x, sensor.y - y},
				{sensor.x - x, sensor.y - y},
			}
			for _, pos := range tryPoints {
				if pos.x < 0 || pos.y < 0 || pos.x > maxVal || pos.y > maxVal {
					continue
				}
				valid := true
				for sO := range g {
					if sO.manhattan(pos) <= coverages[sO] {
						valid = false
						break
					}
				}
				if valid {
					return pos
				}
			}
		}
	}
	return point{}
}

func (g grid) checkRow(row int) map[point]bool {
	excluded := map[point]bool{}
	for sensor, beacon := range g {
		coverage := sensor.manhattan(beacon)
		if sensor.y+coverage >= row && sensor.y-coverage <= row {
			rowCenter := point{sensor.x, row}
			for i := 0; i < coverage; i++ {
				rowPos := rowCenter.add(point{i, 0})
				if sensor.manhattan(rowPos) <= coverage {
					excluded[rowPos] = true
					excluded[rowCenter.add(point{-i, 0})] = true
				}
			}
		}
	}
	for sensor, beacon := range g {
		delete(excluded, sensor)
		delete(excluded, beacon)
	}
	return excluded
}

func intAbs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

type point struct{ x, y int }

func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }
func (p point) sub(o point) point { return point{p.x - o.x, p.y - o.y} }
func (p point) manhattan(o point) int {
	diff := p.sub(o)
	return intAbs(diff.x) + intAbs(diff.y)
}

var benchmark = false
