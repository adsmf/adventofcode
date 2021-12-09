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
			cur := data.val(x, y)
			if y > 0 && data.val(x, y-1) <= cur {
				continue
			}
			if y < data.maxY && data.val(x, y+1) <= cur {
				continue
			}
			if x > 0 && data.val(x-1, y) <= cur {
				continue
			}
			if x < data.maxX && data.val(x+1, y) <= cur {
				continue
			}
			total += int(cur) + 1
		}
	}
	return total
}

func part2(data ventData) int {
	basinData := ventData{
		grid: make(grid, len(data.grid)),
		maxX: data.maxX,
		maxY: data.maxY,
	}
	basinID := 0
	merges := map[int]int{}
	basinAreas := []int{}
	for y := 0; y <= data.maxY; y++ {
		for x := 0; x <= data.maxX; x++ {
			cur := data.val(x, y)
			if cur == 9 {
				basinData.set(x, y, -1)
				continue
			}
			up, left := -1, -1
			if y > 0 {
				up = basinData.val(x, y-1)
			}
			if x > 0 {
				left = basinData.val(x-1, y)
			}
			if up >= 0 && left >= 0 && up != left {
				basinData.set(x, y, up)
				merges[left] = up
				basinAreas[up]++
				continue
			}
			if up >= 0 {
				basinData.set(x, y, up)
				basinAreas[up]++
				continue
			}
			if left >= 0 {
				basinData.set(x, y, left)
				basinAreas[left]++
				continue
			}
			basinData.set(x, y, basinID)
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
	g := make(grid, len(in))
	x, y := 0, 0
	data := ventData{grid: g}
	for _, ch := range in {
		switch {
		case ch >= '0':
			data.set(x, y, int(ch-'0'))
			if x > data.maxX {
				data.maxX = x
			}
			x++
		case ch == '\n':
			y++
			x = 0
		}
	}
	data.maxY = y - 1
	data.grid = data.grid[:data.index(data.maxX, data.maxY)+1]
	return data
}

type ventData struct {
	grid grid
	maxX int
	maxY int
}
type grid []int

func (v *ventData) index(x, y int) int {
	return (y)*(v.maxX+1) + x
}
func (v *ventData) val(x, y int) int {
	idx := v.index(x, y)
	return v.grid[idx]
}
func (v *ventData) set(x, y int, val int) {
	idx := v.index(x, y)
	v.grid[idx] = val
}

var benchmark = false
