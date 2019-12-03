package main

import (
	"fmt"
	"github.com/adsmf/adventofcode2018/utils"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Day 1: %d\n", day1())
	fmt.Printf("Day 2: %d\n", day2())
}

func day1() int {
	g := loadInput()
	return g.findNearestIntersectionDistance()
}

func day2() int {
	return 0
}

func loadInput() *grid {
	wireRaw := utils.ReadInputLines("input.txt")
	return loadGrid(wireRaw[0], wireRaw[1])
}

type grid struct {
	wire1 *wire
	wire2 *wire
}

func loadGrid(input1, input2 string) *grid {
	newGrid := &grid{}
	newGrid.wire1 = parsePath(input1)
	newGrid.wire2 = parsePath(input2)

	return newGrid
}

func (g *grid) findNearestIntersectionDistance() int {
	x, y := g.findNearestIntersection()
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func (g *grid) findNearestIntersection() (int, int) {
	bestDistance := utils.MaxInt
	bestPoint := point{utils.MaxInt, utils.MaxInt}
	for _, w1point := range g.wire1.points {
		x := w1point.x
		y := w1point.y
		dist := int(math.Abs(float64(x)) + math.Abs(float64(y)))
		if dist >= bestDistance {
			continue
		}
		if g.wire2.occupied(x, y) {
			bestDistance = dist
			bestPoint = w1point
		}
	}
	return bestPoint.x, bestPoint.y

	return 0, 0
}

func (g *grid) bothOccupied(x, y int) bool {
	if g.wire1.occupied(x, y) && g.wire2.occupied(x, y) {
		return true
	}
	return false
}

type wire struct {
	area   map[int]map[int]bool
	points []point
}
type point struct {
	x, y int
}

func NewWire() *wire {
	nw := &wire{}
	nw.area = map[int]map[int]bool{}
	nw.points = []point{}
	return nw
}

func (w *wire) set(x, y int) {
	if w.area[x] == nil {
		w.area[x] = map[int]bool{}
	}
	w.area[x][y] = true
	w.points = append(w.points, point{x, y})
}

func (w *wire) occupied(x, y int) bool {
	if w.area[x] != nil {
		return w.area[x][y]
	}
	return false
}

func parsePath(input string) *wire {
	newWire := NewWire()
	// newWire.set(0, 0)
	x := 0
	y := 0

	instructions := strings.Split(input, ",")
	for _, instruction := range instructions {
		dir := instruction[0]
		xDir := 0
		yDir := 0
		switch dir {
		case 'R':
			xDir = 1
		case 'L':
			xDir = -1
		case 'U':
			yDir = 1
		case 'D':
			yDir = -1
		}
		distString := instruction[1:]
		dist, err := strconv.Atoi(distString)
		if err != nil {
			log.Fatalf("Could not decode %s as int\n", distString)
		}
		for i := 0; i < dist; i++ {
			x = x + xDir
			y = y + yDir
			newWire.set(x, y)
		}
	}
	return newWire
}
