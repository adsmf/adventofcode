package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2:\n%s\n", part2())
}

func part1() int {
	inputString := loadInputString()
	hull := runPainter(inputString, 0)
	return hull.countVisited()
}

func part2() string {
	inputString := loadInputString()
	hull := runPainter(inputString, 1)
	return hull.print()
}

type shipHull struct {
	paintColour map[int]map[int]int
	visited     map[int]map[int]bool
}

func (h *shipHull) print() string {
	printout := ""
	var minX, maxX int
	var minY, maxY int
	for x, cols := range h.paintColour {
		for y, tile := range cols {
			if tile > 0 {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}
	for y := minY; y < maxY+1; y++ {
		for x := minX + 1; x < maxX+1; x++ {
			if h.painted(x, y) == 1 {
				printout += fmt.Sprint("#")
			} else {
				printout += fmt.Sprint(".")
			}
		}
		printout += fmt.Sprintln()
	}
	return printout
}

func (h *shipHull) visit(x, y int) {
	if h.visited == nil {
		h.visited = map[int]map[int]bool{}
	}
	if h.visited[x] == nil {
		h.visited[x] = map[int]bool{}
	}
	h.visited[x][y] = true
}

func (h *shipHull) paint(x, y, colour int) {
	h.visit(x, y)
	if h.paintColour == nil {
		h.paintColour = map[int]map[int]int{}
	}
	if h.paintColour[x] == nil {
		h.paintColour[x] = map[int]int{}
	}
	h.paintColour[x][y] = colour
}

func (h *shipHull) painted(x, y int) int {
	if h.paintColour == nil {
		return 0
	}
	if h.paintColour[x] == nil {
		return 0
	}
	return h.paintColour[x][y]
}

func (h *shipHull) countVisited() int {
	count := 0
	for _, cols := range h.visited {
		for _, tile := range cols {
			if tile {
				count++
			}
		}
	}
	return count
}

type robot struct {
	x, y   int
	facing facing
}

func (r *robot) turnRight() {
	r.facing = r.facing.right()
}
func (r *robot) turnLeft() {
	r.facing = r.facing.left()
}
func (r *robot) move() {
	switch r.facing {
	case facingUp:
		r.y--
	case facingDown:
		r.y++
	case facingRight:
		r.x++
	case facingLeft:
		r.x--
	}
}

type facing int

const (
	facingUp    facing = 0
	facingRight facing = 1
	facingDown  facing = 2
	facingLeft  facing = 3
)

func (f facing) right() facing {
	return facing((f + 1) % 4)
}

func (f facing) left() facing {
	if f == facingUp {
		return facingLeft
	}
	return facing(f - 1)
}

func runPainter(program string, startingPanel int64) shipHull {
	hull := shipHull{}
	robo := robot{}
	output := make(chan int64)
	input := make(chan int64, 1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		nextInstructionIsPaint := true
		for op := range output {
			if nextInstructionIsPaint {
				// We're painting
				hull.paint(robo.x, robo.y, int(op))
			} else {
				// We're moving
				if op == 1 {
					robo.turnRight()
				} else {
					robo.turnLeft()
				}
				robo.move()
				input <- int64(hull.painted(robo.x, robo.y))
			}
			nextInstructionIsPaint = !nextInstructionIsPaint
		}
		wg.Done()
	}()

	input <- startingPanel
	tape := newMachine(program, input, output)
	tape.run()

	wg.Wait()
	return hull
}
