package main

import (
	"fmt"
	"github.com/adsmf/adventofcode/utils/vector"
	"io/ioutil"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	v := loadInput("input.txt")
	return len(v)
}

func part2() int {
	return 0
}

type houses map[vector.GridPoint]int

func loadInput(filename string) houses {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	location := vector.GridPoint{X: 0, Y: 0}
	visited := houses{}
	visited[location]++
	for _, char := range input {
		switch char {
		case '>':
			location = location.Right()
		case '<':
			location = location.Left()
		case '^':
			location = location.Up()
		case 'v':
			location = location.Down()
		}
		visited[location]++
	}
	return visited
}
