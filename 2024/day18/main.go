package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

const (
	gridMax  = 70
	loadOnly = 1024
	maxBuf   = 1 << 6
)

type gridWeight [(gridMax + 1) * (gridMax + 1)]int

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d,%d\n", p2.x, p2.y)
	}
}

func solve() (int, point) {
	g := grid{}
	cur := point{}
	blockedPoints := make([]point, 0, 3500)
	utils.EachInteger(input, func(index, value int) (done bool) {
		if index&1 == 0 {
			cur.x = value
		} else {
			cur.y = value
			g.blocked[g.index(cur)] = (index / 2) + 1
			blockedPoints = append(blockedPoints, cur)
		}
		return false
	})
	start, end := point{0, 0}, point{gridMax, gridMax}

	// Part 1
	prevs := dijkstra(g, start, loadOnly)
	p1 := 0
	for cur := g.index(end); cur != g.index(start); cur = prevs[cur] {
		p1++
	}

	// Part 2
	p2 := loadOnly + 1
	low := loadOnly + 1
	high := len(blockedPoints)
	for low <= high {
		mid := (low + high) / 2
		prevs := dijkstra(g, start, mid)
		if prevs[g.index(end)] != 0 {
			low = mid + 1
		} else {
			p2 = mid
			high = mid - 1
		}
	}
	return p1, blockedPoints[p2]
}

func dijkstra(g grid, start point, t int) gridWeight {
	queue := cirQueue[point]{}
	dist := gridWeight{}
	queue.add(start)
	prev := gridWeight{}
	for queue.Len() > 0 {
		cur := queue.pop()
		ci := g.index(cur)
		g.eachNeighbour(cur, func(neigh point, cost int) {
			ni := g.index(neigh)
			if cost > 0 && cost <= t+1 {
				return
			}
			alt := dist[ci] + 1
			if dist[ni] == 0 || alt < dist[ni] {
				dist[ni] = alt
				prev[ni] = ci
				queue.add(neigh)
			}
		})
	}
	return prev
}

type cirQueue[T any] struct {
	buf        [maxBuf]T
	start, end int
}

func (c cirQueue[T]) Len() int {
	end := c.end
	if end < c.start {
		end += len(c.buf)
	}
	return end - c.start
}

func (c *cirQueue[T]) add(p T) {
	c.buf[c.end] = p
	c.end++
	c.end &= maxBuf - 1
}

func (c *cirQueue[T]) pop() T {
	if c.start == c.end {
		panic("Pop from empty queue")
	}
	v := c.buf[c.start]
	c.start++
	c.start &= maxBuf - 1
	return v
}

type grid struct {
	blocked gridWeight
}

func (g grid) inBound(p point) bool    { return p.x >= 0 && p.x <= gridMax && p.y >= 0 && p.y <= gridMax }
func (g grid) index(p point) int       { return int(p.x) + int(p.y)*int(gridMax+1) }
func (g grid) fromIndex(idx int) point { return point{idx % (gridMax + 1), idx / (gridMax + 1)} }
func (g grid) eachNeighbour(p point, callback func(next point, cost int)) {
	for _, dir := range dirs {
		next := p.add(dir)
		if !g.inBound(next) {
			continue
		}
		callback(next, g.blocked[g.index(next)])
	}
}

type queueItem struct {
	node     point
	priority int
}

type point struct{ x, y int }

func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }

var dirs = [4]point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

var benchmark = false
