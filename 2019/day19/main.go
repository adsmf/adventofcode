package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/adsmf/adventofcode2019/utils/intcode"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	return count5050("input.txt")
}

func part2() int {
	return find1010("input.txt")
}

func count5050(filename string) int {
	count := 0
	input := loadInputString()
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			count += getPoint(input, x, y)
		}
	}
	return count
}

func find1010(filename string) int {
	input := loadInputString()
	// var prevXLeft int
	var prevXRight int
	type row [2]int
	rows := []row{}
	searchSize := 100
	for y := 0; y < 1000; y++ {
		xStart := sort.Search(prevXRight+1, func(x int) bool {
			return getPoint(input, x, y) == 1
		})
		xEnd := sort.Search(10000, func(x int) bool {
			return getPoint(input, x+xStart, y) == 0
		})
		xEnd += xStart - 1
		prevXRight = xEnd
		newRow := row{xStart, xEnd}
		rows = append(rows, newRow)
		if len(rows) < searchSize+1 {
			continue
		}
		up100 := rows[len(rows)-searchSize]
		if up100[1] >= newRow[0]+searchSize-1 {
			return xStart*10000 + (y - searchSize + 1)
		}
	}
	return 0
}

type point struct {
	x, y int
}

func getPoint(program string, x, y int) int {
	t := tractor{}
	t.inputs = []int{x, y}
	m := intcode.NewMachine(intcode.M19(t.inputHandler, nil))
	m.LoadProgram(program)
	m.Run(true)
	return m.Register(intcode.M19RegisterOutput)
}

type tractor struct {
	inputs  []int
	outputs []int
}

func (t *tractor) inputHandler() (int, bool) {
	next := t.inputs[0]
	t.inputs = t.inputs[1:]
	return next, false
}

func loadInputString() string {
	inputRaw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputRaw)

}
