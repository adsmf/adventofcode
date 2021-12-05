package main

import (
	"fmt"
	"io/ioutil"

	"github.com/adsmf/adventofcode/utils/vector"
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
	v := loadInput("input.txt", 1)
	return len(v)
}

func part2() int {
	v := loadInput("input.txt", 2)
	return len(v)
}

type houses map[vector.GridPoint]int

func loadInput(filename string, numSantas int) houses {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	locations := make([]vector.GridPoint, numSantas)
	visited := houses{}
	for id := range locations {
		visited[locations[id]]++
	}
	curMover := 0
	for _, char := range input {
		switch char {
		case '>':
			locations[curMover] = locations[curMover].Right()
		case '<':
			locations[curMover] = locations[curMover].Left()
		case '^':
			locations[curMover] = locations[curMover].Up()
		case 'v':
			locations[curMover] = locations[curMover].Down()
		}
		visited[locations[curMover]]++
		curMover++
		curMover %= len(locations)
	}
	return visited
}

var benchmark = false
