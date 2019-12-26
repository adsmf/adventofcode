package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode2018/utils"
)

func main() {
	lines := utils.ReadInputLines("input.txt")
	points := make([]point, 0)
	maxX := 0
	maxY := 0
	for _, line := range lines {
		pointStr := strings.Split(line, ", ")
		x, _ := strconv.Atoi(pointStr[0])
		y, _ := strconv.Atoi(pointStr[1])
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		point := point{x, y}
		points = append(points, point)
	}
	maxX++
	maxY++

	fmt.Printf("Part 1\n======\n")
	part1(points, maxX, maxY)
	fmt.Printf("\nPart 2\n======\n")
	part2(points, maxX, maxY, 10000)
}

type point struct {
	x, y int
}

func part1(points []point, maxX, maxY int) {
	numClosest := closestPointGraph(points, maxX, maxY)
	bestCount := 0
	for _, count := range numClosest {
		if count > bestCount {
			bestCount = count
		}
	}
	fmt.Printf("Best count: %d\n", bestCount)
}

func part2(points []point, maxX, maxY, maxDistance int) {
	numValid := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			if totalDistance(point{x, y}, points) < maxDistance {
				numValid++
			}
		}
	}
	fmt.Printf("Region size: %d\n", numValid)
}

func closestPointGraph(points []point, maxX, maxY int) map[point]int {
	numClose := make(map[point]int)
	invalidPoints := make(map[point]bool)
	invalidPoints[point{0, 0}] = true
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			closest := closestPoint(point{x, y}, points)
			numClose[closest]++
			if x == 0 || y == 0 || x == maxX || y == maxY {
				invalidPoints[closest] = true
			}
		}
	}
	validClose := make(map[point]int)
	for point, count := range numClose {
		if !!!invalidPoints[point] {
			validClose[point] = count
		}
	}
	return validClose
}

func closestPoint(target point, points []point) point {
	minDistance := -1
	closest := points[0]
	for _, curPoint := range points {
		dist := distance(target, curPoint)
		if minDistance == -1 || dist < minDistance {
			minDistance = dist
			closest = curPoint
		} else if dist == minDistance {
			closest = point{0, 0}
		}
	}
	return closest
}

func totalDistance(target point, points []point) int {
	totalDist := 0
	for _, curPoint := range points {
		totalDist += distance(target, curPoint)
	}
	return totalDist
}

func distance(a, b point) int {
	return int(math.Abs(float64(a.x-b.x))) + int(math.Abs(float64(a.y-b.y)))
}
