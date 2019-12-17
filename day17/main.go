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
	cmds := []string{
		"A,B,A,B,C,C,B,A,B,C\n",
		"L,8,R,12,R,12,R,10\n",
		"R,10,R,12,R,10\n",
		"L,10,R,10,L,6\n",
		"n\n",
	}
	return testDustCount(cmds)
}

func testDustCount(commands []string) int {
	s := scaffold{
		commands: commands,
	}
	input := loadInputString()
	m := intcode.NewMachine(intcode.M19(s.driver, s.outputHandler))
	m.LoadProgram(input)
	m.WriteRAM(0, 2)
	m.Run(false)
	return s.collected

}

func readCamera() scaffold {
	s := scaffold{}
	input := loadInputString()
	m := intcode.NewMachine(intcode.M19(nil, s.outputHandler))
	m.LoadProgram(input)
	m.Run(false)
	s.processCamera()
	return s
}

type scaffold struct {
	cameraViewRaw  string
	grid           map[point]scaffoldTile
	intersections  map[point]intersection
	commands       []string
	currentCommand string
	collected      int
}

func (s *scaffold) driver() (int, bool) {
	// fmt.Printf("Input called!\n")
	if s.currentCommand != "" {
		ch := s.currentCommand[0]
		s.currentCommand = s.currentCommand[1:]
		return int(ch), false
	}
	if len(s.commands) > 0 {
		s.currentCommand = s.commands[0]
		if len(s.commands) > 1 {
			s.commands = s.commands[1:]
		} else {
			s.commands = s.commands[0:0]
		}
		ch := s.currentCommand[0]
		s.currentCommand = s.currentCommand[1:]
		return int(ch), false
	}
	return 0, true
}

func (s *scaffold) outputHandler(out int) {
	if out < 255 {
		s.cameraViewRaw += fmt.Sprintf("%c", out)
	} else {
		s.collected = out
	}
}

func (s *scaffold) sumAlignment() int {
	sum := 0
	for p := range s.intersections {
		sum += p.x * p.y
	}
	return sum
}

func (s *scaffold) processCamera() {
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
