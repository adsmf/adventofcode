package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode2019/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	// fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	planets := loadInput("input.txt")
	planets.run(1000)
	return planets.energy()
}

func part2() int {
	seenPositions := map[int][]positionList{}
	planets := loadInput("input.txt")
	initialEnergy := planets.energy()
	seenPositions[initialEnergy] = []positionList{planets.positions()}
	for i := 0; ; i++ {
		planets.step()
		positions := planets.positions()
		energy := planets.energy()
		if previousPositionsList, found := seenPositions[energy]; found {
			for _, previous := range previousPositionsList {
				if previous.matches(positions) {
					return i
				}
			}
			seenPositions[energy] = append(seenPositions[energy], positions)
		} else {
			seenPositions[energy] = []positionList{positions}
		}
	}
}

type positionList []vector

func (p positionList) matches(a positionList) bool {
	if len(p) != len(a) {
		return false
	}
	for i := 0; i < len(p); i++ {
		if p[i] != a[i] {
			return false
		}
	}
	return true
}

type system []*planet

func (s system) run(steps int) {
	for i := 0; i < steps; i++ {
		s.step()
	}
}
func (s system) step() {
	for _, p := range s {
		p.gravity(s)
	}
	for _, p := range s {
		p.move()
	}
}

func (s system) positions() []vector {
	positions := []vector{}
	for _, p := range s {
		positions = append(positions, p.position)
	}
	return positions
}

func (s system) energy() int {
	ret := 0
	for _, p := range s {
		ret += p.energy()
	}
	return ret
}

func (s system) String() string {
	ret := ""
	for _, p := range s {
		ret += fmt.Sprintln(p)
	}
	return ret
}

type planet struct {
	position vector
	velocity vector
}

func (p *planet) gravity(s system) {
	for _, plan := range s {
		if plan == p {
			continue
		}
		velTowards := vector{}
		if plan.position.x > p.position.x {
			velTowards.x = 1
		} else if plan.position.x < p.position.x {
			velTowards.x = -1
		}
		if plan.position.y > p.position.y {
			velTowards.y = 1
		} else if plan.position.y < p.position.y {
			velTowards.y = -1
		}
		if plan.position.z > p.position.z {
			velTowards.z = 1
		} else if plan.position.z < p.position.z {
			velTowards.z = -1
		}
		p.velocity.add(velTowards)
	}
}

func (p *planet) move() {
	p.position.add(p.velocity)
}

func (p planet) energy() int {
	return p.position.energy() * p.velocity.energy()
}

func (p planet) String() string {
	return fmt.Sprintf("pos=%v, vel=%v", p.position, p.velocity)
}

type vector struct {
	x, y, z int
}

func (v *vector) add(a vector) {
	v.x += a.x
	v.y += a.y
	v.z += a.z
}

func (v vector) energy() int {
	return int(math.Abs(float64(v.x)) +
		math.Abs(float64(v.y)) +
		math.Abs(float64(v.z)),
	)
}

func (v vector) String() string {
	return fmt.Sprintf("<x=%3d, y=%3d, z=%3d>", v.x, v.y, v.z)
}

func loadInput(filename string) system {
	planets := system{}
	lines := utils.ReadInputLines(filename)
	for _, line := range lines {
		line = strings.ReplaceAll(line, ",", "")
		line = strings.Trim(line, "<>")
		axes := strings.Split(line, " ")
		p := planet{}
		for _, axis := range axes {
			parts := strings.Split(axis, "=")
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			switch parts[0] {
			case "x":
				p.position.x = value
			case "y":
				p.position.y = value
			case "z":
				p.position.z = value
			}
		}
		planets = append(planets, &p)
	}
	return planets
}
