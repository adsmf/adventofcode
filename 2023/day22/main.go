package main

import (
	_ "embed"
	"fmt"
	"sort"

	"github.com/adsmf/adventofcode/utils"
	"golang.org/x/exp/constraints"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	bricks := load()

	settledBricks := make([]bool, len(bricks))
	grounded := make([]bool, len(bricks))
	supported := make([][]int, len(bricks))
	supporting := make([][]int, len(bricks))

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].min.z < bricks[j].min.z
	})

	allSettled := false
	for !allSettled {
		allSettled = true
		for i := 0; i < len(bricks); i++ {
			if settledBricks[i] {
				continue
			}
			canDrop := true
			isSettled := false
			for j := 0; j < len(bricks); j++ {
				if i == j {
					continue
				}
				if bricks[i].intersects(bricks[j], 1) {
					canDrop = false
					if settledBricks[j] {
						isSettled = true
						supporting[j] = append(supporting[j], i)
						supported[i] = append(supported[i], j)
					}
				}
			}
			if canDrop {
				bricks[i].drop(1)
				allSettled = false
				if bricks[i].min.z == 0 {
					grounded[i] = true
					isSettled = true
				}
			}
			if isSettled {
				settledBricks[i] = true
			}
		}
	}

	totalP1 := 0
	for i := range bricks {
		canRemove := true
		for _, sup := range supporting[i] {
			if len(supported[sup]) == 1 {
				canRemove = false
			}
		}
		if canRemove {
			totalP1++
		}
	}
	totalP2 := 0
	removed := make([]bool, len(bricks))
	for i := range bricks {
		thisRemoves := 0
		clear(removed)
		removed[i] = true

		done := false
		for !done {
			done = true
			for j := 0; j < len(bricks); j++ {
				if removed[j] || grounded[j] {
					continue
				}
				anySupport := false
				for _, sup := range supported[j] {
					if !removed[sup] {
						anySupport = true
						break
					}
				}
				if !anySupport {
					thisRemoves++
					removed[j] = true
					done = false
				}
			}
		}
		totalP2 += thisRemoves
	}
	return totalP1, totalP2
}

func part2() int {
	return -1
}

func load() []brickInfo {
	bricks := []brickInfo{}
	utils.EachLine(input, func(index int, line string) (done bool) {
		coords := utils.GetInts(line)
		bricks = append(bricks, brickInfo{
			point{
				min(coords[0], coords[3]),
				min(coords[1], coords[4]),
				min(coords[2], coords[5]),
			},
			point{
				max(coords[0], coords[3]),
				max(coords[1], coords[4]),
				max(coords[2], coords[5]),
			},
		})
		return false
	})

	return bricks
}

type brickInfo struct {
	min, max point
}

func (b *brickInfo) drop(by int) {
	b.min.z -= by
	b.max.z -= by
}

func (b brickInfo) intersects(o brickInfo, zOffset int) bool {
	if b.min.z-zOffset > o.max.z || b.max.z-zOffset < o.min.z ||
		b.min.x > o.max.x || b.max.x < o.min.x ||
		b.min.y > o.max.y || b.max.y < o.min.y {
		return false
	}
	return true
}

type point struct {
	x, y, z int
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

var benchmark = false
