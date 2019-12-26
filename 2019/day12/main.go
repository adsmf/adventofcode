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
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	planets := loadInput("input.txt")
	planets.run(1000)
	return planets.energy()
}

func part2() int {
	planets := loadInput("input.txt")
	period := planets.calculatePeriod()
	return period
}

func findRepeat(planets system) int {
	initialPositions := planets.positions()

	for i := 2; ; i++ {
		planets.step()
		positions := planets.positions()
		if positions.matches(initialPositions) {
			return i
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

func (s system) axisString(axis int) string {
	retString := ""
	for _, p := range s {
		retString += fmt.Sprintf("%v", p.axisString(axis))
	}
	return retString
}

type axisVec struct {
	axis     int
	position int
	velocity int
}

func (s system) calculatePeriod() int {
	axisPeriods := []int{0, 0, 0}

	previousX := map[string]int{}
	previousY := map[string]int{}
	previousZ := map[string]int{}

	for i := 0; ; i++ {
		s.step()
		if axisPeriods[0] == 0 {
			systemX := s.axisString(0)
			if prev, found := previousX[systemX]; found {
				axisPeriods[0] = i - prev
			}
			previousX[systemX] = i
		}
		if axisPeriods[1] == 0 {
			systemY := s.axisString(1)
			if prev, found := previousY[systemY]; found {
				axisPeriods[1] = i - prev
			}
			previousY[systemY] = i
		}
		if axisPeriods[2] == 0 {
			systemZ := s.axisString(2)
			if prev, found := previousZ[systemZ]; found {
				axisPeriods[2] = i - prev
			}
			previousZ[systemZ] = i
		}
		if axisPeriods[0] > 0 &&
			axisPeriods[1] > 0 &&
			axisPeriods[2] > 0 {
			break
		}
	}

	period := lcm(axisPeriods...)
	return period
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(integers ...int) int {
	if len(integers) < 2 {
		return 0
	}
	a := integers[0]
	b := integers[1]
	integers = integers[2:]
	g := gcd(a, b)
	if g == 0 {
		return 0
	}
	result := a * b / g

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func (s system) positions() positionList {
	positions := positionList{}
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

func (p planet) axisString(axis int) string {
	switch axis {
	case 0:
		return fmt.Sprintf("<%d,%d>", p.position.x, p.velocity.x)
	case 1:
		return fmt.Sprintf("<%d,%d>", p.position.y, p.velocity.y)
	case 2:
		return fmt.Sprintf("<%d,%d>", p.position.z, p.velocity.z)
	}
	panic(fmt.Sprintf("Unknown axis %d", axis))
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
