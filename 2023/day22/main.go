package main

import (
	_ "embed"
	"fmt"
	"sort"

	"github.com/adsmf/adventofcode/utils"
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

	supported := make([][]int, len(bricks))

	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].min.z < bricks[j].min.z
	})

	for i := 0; i < len(bricks); i++ {
		settleAt := 0
		supported[i] = []int{}
		for j := 0; j < i; j++ {
			if bricks[i].aligns(bricks[j]) {
				newSettle := bricks[j].max.z + 1
				if settleAt < newSettle {
					settleAt = newSettle
					supported[i] = supported[i][0:0]
					supported[i] = append(supported[i], j)
				} else if settleAt == newSettle {
					supported[i] = append(supported[i], j)
				}
			}
		}
		dropBy := bricks[i].min.z - settleAt
		if settleAt == 0 {
			supported[i] = append(supported[i], supportGround)
		}
		bricks[i].drop(dropBy)
	}

	totalP1, totalP2 := 0, 0
	removed := make([]bool, len(bricks))
	for i := range bricks {
		thisRemoves := 0
		clear(removed)
		removed[i] = true

		for done := false; !done; {
			done = true
			for j := 0; j < len(bricks); j++ {
				if removed[j] {
					continue
				}
				anySupport := false
				for _, sup := range supported[j] {
					if sup == supportGround || !removed[sup] {
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
		if thisRemoves == 0 {
			totalP1++
		}
		totalP2 += thisRemoves
	}
	return totalP1, totalP2
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

func (b brickInfo) aligns(o brickInfo) bool {
	if b.min.x > o.max.x || b.max.x < o.min.x ||
		b.min.y > o.max.y || b.max.y < o.min.y {
		return false
	}
	return true
}

type point struct {
	x, y, z int
}

const supportGround = -1

var benchmark = false
