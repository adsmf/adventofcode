package main

import (
	"container/heap"
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := search(1, 3)
	p2 := search(4, 10)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func search(minSteps, maxSteps int) int {
	g, finish := load()
	start := point{0, 0}
	visited := map[searchEntry]bool{}
	open := &searchQueue{}
	heap.Init(open)
	heap.Push(open, searchEntry{pos: start, dir: dirNone})
	for step := 0; len(*open) > 0; step++ {
		cur := heap.Pop(open).(searchEntry)
		if cur.pos == finish {
			return cur.loss
		}
		for cd, offset := range directionOffsets {
			dir := direction(cd)
			if dir.opposite() == cur.dir || dir == cur.dir {
				continue
			}

			checkVisit := searchEntry{
				pos: cur.pos,
				dir: dir,
			}
			if visited[checkVisit] {
				continue
			}

			visited[checkVisit] = true
			neighbour := cur.pos
			losses := 0
			for i := 1; i <= maxSteps; i++ {
				neighbour = neighbour.add(offset)
				nLoss, found := g[neighbour]
				if !found {
					break
				}
				losses += nLoss
				if i < minSteps {
					continue
				}
				next := searchEntry{
					pos:  neighbour,
					loss: cur.loss + losses,
					dir:  dir,
				}
				heap.Push(open, next)
				if neighbour == finish {
					break
				}
			}
		}
	}
	return -1
}

type searchEntry struct {
	pos  point
	loss int
	dir  direction
}

type searchQueue []searchEntry

func (s searchQueue) Len() int           { return len(s) }
func (s searchQueue) Less(i, j int) bool { return s[i].loss < s[j].loss }
func (s searchQueue) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *searchQueue) Push(entry any) {
	*s = append(*s, entry.(searchEntry))
}

func (s *searchQueue) Pop() any {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func load() (grid, point) {
	g := grid{}
	max := point{}
	utils.EachLine(input, func(index int, line string) (done bool) {
		max.y = index
		for x, ch := range line {
			pos := point{x, index}
			g[pos] = int(ch - '0')
			if x > max.x {
				max.x = x
			}
		}
		return false
	})
	return g, max
}

type grid map[point]int

type point struct{ x, y int }

func (p point) add(a point) point { return point{x: p.x + a.x, y: p.y + a.y} }

type direction int

func (d direction) opposite() direction {
	if d < 0 || d >= dirMAX {
		return dirNone
	}
	return (d + 2) % dirMAX
}

const (
	dirUp direction = iota
	dirRight
	dirDown
	dirLeft

	dirMAX
	dirNone = -1
)

var directionOffsets = [dirMAX]point{
	dirUp:    {0, -1},
	dirRight: {1, 0},
	dirDown:  {0, 1},
	dirLeft:  {-1, 0},
}

var benchmark = false
