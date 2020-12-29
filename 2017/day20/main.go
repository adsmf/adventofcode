package main

import (
	"fmt"
	"math"

	"github.com/adsmf/adventofcode/utils"
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
	particles := load("input.txt")
	minDist := 0
	closest := -1
	longTime := 999999
	for id, p := range particles {
		pos := p.positionAfter(longTime)
		dist := pos.magnitude()
		if closest == -1 || dist < minDist {
			closest = id
			minDist = dist
		}
	}
	return closest
}

func part2() int {
	particles := load("input.txt")
	maxSteps := 100
	for step := 0; step < maxSteps; step++ {
		positions := map[string][]int{}
		for id, p := range particles {
			pos := p.positionAfter(step)
			posStr := pos.String()
			positions[posStr] = append(positions[posStr], id)
		}
		newParticles := make([]particle, 0, len(particles))
		for _, list := range positions {
			if len(list) > 1 {
				continue
			}
			newParticles = append(newParticles, particles[list[0]])
		}
		particles = newParticles
	}
	return len(particles)
}

type particle struct {
	position     vector
	velocity     vector
	acceleration vector
}

func (p particle) positionAfter(steps int) vector {
	pos := p.position
	pos = pos.add(p.velocity.times(steps))
	pos = pos.add(p.acceleration.times(steps * (steps + 1) / 2))

	return pos
}

type vector struct{ x, y, z int }

func (v vector) String() string {
	return fmt.Sprintf("%d,%d,%d", v.x, v.y, v.z)
}

func (v vector) magnitude() int {
	mag := 0
	mag += int(math.Abs(float64(v.x)))
	mag += int(math.Abs(float64(v.y)))
	mag += int(math.Abs(float64(v.z)))
	return mag
}

func (v vector) times(scale int) vector {
	return vector{
		v.x * scale,
		v.y * scale,
		v.z * scale,
	}
}
func (v vector) add(a vector) vector {
	return vector{
		v.x + a.x,
		v.y + a.y,
		v.z + a.z,
	}
}

func load(filename string) []particle {
	particles := []particle{}
	for _, line := range utils.ReadInputLines(filename) {
		ints := utils.GetInts(line)
		p := particle{
			position:     vector{(int(ints[0])), (int(ints[1])), (int(ints[2]))},
			velocity:     vector{(int(ints[3])), (int(ints[4])), (int(ints[5]))},
			acceleration: vector{(int(ints[6])), (int(ints[7])), (int(ints[8]))},
		}
		particles = append(particles, p)
	}
	return particles
}

var benchmark = false
