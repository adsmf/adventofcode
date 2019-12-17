package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/adsmf/adventofcode2019/utils/intcode"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	s := readCamera()
	return s.sumAlignment()
}

func part2() int {
	return 0
}

func readCamera() scaffold {
	s := scaffold{}
	input := loadInputString()
	m := intcode.NewMachine(intcode.M19(nil, s.cameraOutputHandler))
	m.LoadProgram(input)
	m.Run(false)
	s.processCamera()
	return s
}

type scaffold struct {
	cameraViewRaw string
	grid          map[point]scaffoldTile
	intersections map[point]intersection
}

func (s *scaffold) sumAlignment() int {
	sum := 0
	for p := range s.intersections {
		sum += p.x * p.y
	}
	return sum
}

func (s *scaffold) processCamera() {
	// fmt.Printf("%v\n", s.cameraViewRaw)
	s.intersections = map[point]intersection{}
	s.grid = map[point]scaffoldTile{}

	lines := strings.Split(s.cameraViewRaw, "\n")
	for y, line := range lines {
		for x, inp := range line {
			switch inp {
			case '#':
				s.grid[point{x, y}] = scaffoldTile{}
			}
		}
	}
	for p := range s.grid {
		_, up := s.grid[point{p.x, p.y - 1}]
		_, down := s.grid[point{p.x, p.y + 1}]
		_, left := s.grid[point{p.x - 1, p.y}]
		_, right := s.grid[point{p.x + 1, p.y}]
		if up && down && left && right {
			s.intersections[p] = intersection{}
		}
	}
	// fmt.Printf("Scaffold: %v\n", s.grid)
	// fmt.Printf("Intersections: %v\n", s.intersections)
}

func (s *scaffold) cameraOutputHandler(out int) {
	s.cameraViewRaw += fmt.Sprintf("%c", out)
}

type point struct {
	x, y int
}

type intersection struct{}
type scaffoldTile struct{}

func loadInputString() string {
	inputRaw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputRaw)

}
