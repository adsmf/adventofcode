package main

import (
	_ "embed"
	"fmt"
	"sort"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func main() {
	data := loadData(input)
	p1 := part1(data)
	p2 := part2(data)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(data ventData) int {
	total := 0
	for y := 0; y <= data.maxY; y++ {
		for x := 0; x <= data.maxX; x++ {
			cur := data.grid[point{x, y}]
			if y > 0 && data.grid[point{x, y - 1}] <= cur {
				continue
			}
			if y < data.maxY && data.grid[point{x, y + 1}] <= cur {
				continue
			}
			if x > 0 && data.grid[point{x - 1, y}] <= cur {
				continue
			}
			if x < data.maxX && data.grid[point{x + 1, y}] <= cur {
				continue
			}
			total += int(cur) + 1
		}
	}
	return total
}

func part2(data ventData) int {
	basinData := make(grid, len(data.grid))
	basinID := 0
	merges := map[int]int{}
	basinAreas := []int{}
	for y := 0; y <= data.maxY; y++ {
		for x := 0; x <= data.maxX; x++ {
			pos := point{x, y}
			cur := data.grid[pos]
			if cur == 9 {
				continue
			}
			up, upDef := basinData[point{x, y - 1}]
			left, leftDef := basinData[point{x - 1, y}]
			if upDef && leftDef && up != left {
				basinData[pos] = up
				merges[left] = up
				basinAreas[up]++
				continue
			}
			if upDef {
				basinData[pos] = up
				basinAreas[up]++
				continue
			}
			if leftDef {
				basinData[pos] = left
				basinAreas[left]++
				continue
			}
			basinData[pos] = basinID
			basinAreas = append(basinAreas, 1)
			basinID++
		}
	}
	for from, to := range merges {
		for {
			if replace, found := merges[to]; found {
				to = replace
				continue
			}
			break
		}
		merges[from] = to
	}
	mergedBasinAreas := make(map[int]int, len(basinAreas))
	for basin, count := range basinAreas {
		if to, found := merges[basin]; found {
			basin = to
		}
		mergedBasinAreas[basin] += count
	}
	sums := make([]int, len(mergedBasinAreas))
	for _, count := range mergedBasinAreas {
		sums = append(sums, count)
	}
	sort.Ints(sums)
	total := 1
	for _, sum := range sums[len(sums)-3:] {
		total *= sum
	}
	return total
}

func loadData(in string) ventData {
	g := grid{}
	x, y := 0, 0
	maxX, maxY := 0, 0
	for _, ch := range in {
		switch {
		case ch >= '0':
			g[point{x, y}] = int(ch - '0')
			if x > maxX {
				maxX = x
			}
			x++
		case ch == '\n':
			y++
			x = 0
		}
	}
	maxY = y - 1
	return ventData{grid: g, maxX: maxX, maxY: maxY}
}

type ventData struct {
	grid grid
	maxX int
	maxY int
}
type grid map[point]int
type point struct{ x, y int }

var benchmark = false
