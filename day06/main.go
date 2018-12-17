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

	part1(points, maxX, maxY)
}

type point struct {
	x, y int
}

func part1(points []point, maxX, maxY int) {
	_, numClosest := closestPointGraph(points, maxX, maxY)
	bestCount := 0
	for _, count := range numClosest {
		if count > bestCount {
			bestCount = count
		}
	}
	fmt.Printf("Best count: %d\n", bestCount)
}

func closestPointGraph(points []point, maxX, maxY int) ([][]point, map[point]int) {
	graph := make([][]point, maxX+1)
	numClose := make(map[point]int)
	invalidPoints := make(map[point]bool)
	invalidPoints[point{0, 0}] = true
	for x := 0; x <= maxX; x++ {
		graph[x] = make([]point, maxY+1)
		for y := 0; y <= maxY; y++ {
			closest := closestPoint(point{x, y}, points)
			numClose[closest]++
			graph[x][y] = closest
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
	return graph, validClose
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

func distance(a, b point) int {
	return int(math.Abs(float64(a.x-b.x))) + int(math.Abs(float64(a.y-b.y)))
}
