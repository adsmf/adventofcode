package main

import (
	"fmt"
	"github.com/adsmf/adventofcode2019/utils"
	"math"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1V2())
	fmt.Printf("Part 2: %d\n", part2V2())
}

func part1V1() int {
	g := loadInputFile("input.txt")
	_, count := g.getBestLineOfSightV1()
	return count
}
func part1V2() int {
	g := loadInputFile("input.txt")
	_, count := g.getBestLineOfSightV2()
	return count
}

func part2V1() int {
	g := loadInputFile("input.txt")
	base, _ := g.getBestLineOfSightV1()
	destroyed := destroyNV1(g, base, 200)
	lastDestroyed := destroyed[len(destroyed)-1]
	return int(lastDestroyed.position.x)*100 + int(lastDestroyed.position.y)
}

func part2V2() int {
	g := loadInputFile("input.txt")
	base, _ := g.getBestLineOfSightV2()
	destroyed := destroyNV2(g, base, 200)
	lastDestroyed := destroyed[len(destroyed)-1]
	return int(lastDestroyed.position.x)*100 + int(lastDestroyed.position.y)
}

func destroyNV1(g grid, base asteroid, num int) []asteroid {
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

func destroyNV2(g grid, base asteroid, num int) []asteroid {
	destroyed := make([]asteroid, num)
	angleIndex := 0
	lines := sortSightlines(base.getLines(g))
	for numDestroyed := 0; numDestroyed < num; numDestroyed++ {
		for len(lines[angleIndex]) == 0 {
			angleIndex++
			angleIndex %= len(lines)
		}
		destroyed[numDestroyed] = lines[angleIndex][0]
		lines[angleIndex] = lines[angleIndex][1:]
		angleIndex++
		angleIndex %= len(lines)
	}
	return destroyed
}

type grid struct {
	asteroids []asteroid
}

func (g grid) getBestLineOfSightV1() (asteroid, int) {
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

func (g grid) getBestLineOfSightV2() (asteroid, int) {
	bestAsteroid := asteroid{}
	bestCount := 0
	for _, ast := range g.asteroids {
		lines := ast.getLines(g)
		count := len(lines)
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

type sightLinesMap map[float64]map[float64]asteroid
type sightLines [][]asteroid

type asteroid struct {
	position vector
}

func (a asteroid) getLines(g grid) sightLinesMap {
	linesMap := make(sightLinesMap)
	for _, other := range g.asteroids {
		if other == a {
			continue
		}
		vecOther := vectorBetween(a, other)
		pAngle := math.Atan2(vecOther.y, vecOther.x) - math.Pi/2
		if pAngle < 0 {
			pAngle += 2 * math.Pi
		}
		if pAngle > 2*math.Pi {
			pAngle -= 2 * math.Pi
		}

		if _, found := linesMap[pAngle]; found {
			linesMap[pAngle][vecOther.length()] = other
		} else {
			linesMap[pAngle] = map[float64]asteroid{
				vecOther.length(): other,
			}
		}
	}
	return linesMap
}

func sortSightlines(linesMap sightLinesMap) sightLines {
	lines := make(sightLines, len(linesMap))
	angles := []float64{}
	for angle := range linesMap {
		angles = append(angles, angle)
	}
	sort.Float64s(angles)
	for angleIndex, angle := range angles {
		astMap := linesMap[angle]

		distances := []float64{}
		for distance := range astMap {
			distances = append(distances, distance)
		}
		sort.Float64s(distances)
		asteroids := make([]asteroid, len(distances))
		for distIndex, distance := range distances {
			asteroids[distIndex] = astMap[distance]
		}
		lines[angleIndex] = asteroids
	}
	return lines
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
