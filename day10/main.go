package main

import (
	"fmt"
	"github.com/adsmf/adventofcode2019/utils"
	"math"
	"sort"
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
	destroyed := destroyNV(g, base, 200)
	lastDestroyed := destroyed[len(destroyed)-1]
	return int(lastDestroyed.position.x)*100 + int(lastDestroyed.position.y)
}

func destroyNV(g grid, base asteroid, num int) []asteroid {
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

func (g grid) getBestLineOfSight() (asteroid, int) {
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

type vector struct {
	x, y float64
}

func (v vector) length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
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
