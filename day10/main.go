package main

import (
	"fmt"
	"github.com/adsmf/adventofcode2019/utils"
	"math"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	g := loadInputFile("input.txt")
	_, count := g.getBestLineOfSight()
	return count
}

func part2() int {
	g := loadInputFile("input.txt")
	base, _ := g.getBestLineOfSight()
	destroyed := destroyN(g, base, 200)
	lastDestroyed := destroyed[len(destroyed)-1]
	return int(lastDestroyed.position.x)*100 + int(lastDestroyed.position.y)
}

func destroyN(g grid, base asteroid, num int) []asteroid {
	destroyed := []asteroid{}

	offsetAngle := 0.000001

	laserDirection := math.Pi/2 - offsetAngle
	for numDestroyed := 0; numDestroyed < num; numDestroyed++ {
		asteroids := g.withLineOfSight(base)
		bestAngle := float64(9999999)
		var nextTarget asteroid
		for _, potential := range asteroids {
			if potential == base {
				continue
			}
			vecPotential := vectorBetween(base, potential)
			pAngle := math.Atan2(vecPotential.y, vecPotential.x)
			angleDiff := pAngle - laserDirection
			if angleDiff < 0 {
				angleDiff += 2 * math.Pi
			}
			if angleDiff < bestAngle {
				bestAngle = angleDiff
				nextTarget = potential
			}
		}

		destroyed = append(destroyed, nextTarget)
		vecTarget := vectorBetween(base, nextTarget)
		laserDirection = math.Atan2(vecTarget.y, vecTarget.x) + offsetAngle
		g.remove(nextTarget)
	}
	return destroyed
}

type grid struct {
	asteroids []asteroid
}

func (g grid) getBestLineOfSight() (asteroid, int) {
	bestAsteroid := asteroid{}
	bestCount := 0
	for _, ast := range g.asteroids {
		count := ast.countLineOfSight(g)
		if count > bestCount {
			bestCount = count
			bestAsteroid = ast
		}
	}
	return bestAsteroid, bestCount
}

func (g grid) withLineOfSight(base asteroid) []asteroid {
	with := []asteroid{}
	for _, ast := range g.asteroids {
		if hasLineOfSight(g, base, ast) {
			with = append(with, ast)
		}
	}
	return with
}

func (g *grid) remove(target asteroid) {
	n := 0
	for _, ast := range g.asteroids {
		if ast != target {
			g.asteroids[n] = ast
			n++
		}
	}
	g.asteroids = g.asteroids[:n]
}

type asteroid struct {
	position vector
}

func (a asteroid) countLineOfSight(g grid) int {
	count := 0
	for _, other := range g.asteroids {
		if a == other {
			continue
		}
		if hasLineOfSight(g, a, other) {
			count++
		}
	}
	return count
}

func hasLineOfSight(g grid, a, b asteroid) bool {
	vecBetween := vectorBetween(a, b)
	unitBetween := vecBetween.unit()
	for _, other := range g.asteroids {
		if other == a || other == b {
			continue
		}
		vecOther := vectorBetween(a, other)
		if vecOther.unit() == unitBetween && vecOther.length() < vecBetween.length() {
			return false
		}
	}
	return true
}

type vector struct {
	x, y float64
}

func (v vector) length() float64 {
	return math.Sqrt(v.dot(v))
}

func (v vector) unit() vector {
	length := v.length()
	unit := vector{
		x: math.Round(1000*v.x/length) / 1000,
		y: math.Round(1000*v.y/length) / 1000,
	}
	return unit
}

func (v vector) dot(b vector) float64 {
	return v.x*b.x + v.y*b.y
}

func vectorBetween(a, b asteroid) vector {
	v := vector{
		x: float64(a.position.x - b.position.x),
		y: float64(a.position.y - b.position.y),
	}
	return v
}

func loadInputFile(filename string) grid {
	lines := utils.ReadInputLines(filename)
	g := grid{
		asteroids: []asteroid{},
	}
	for row, line := range lines {
		for col, symbol := range line {
			if symbol == '#' {
				ast := asteroid{
					position: vector{
						x: float64(col),
						y: float64(row),
					},
				}
				g.asteroids = append(g.asteroids, ast)
			}
		}
	}
	return g
}
