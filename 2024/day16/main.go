package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"math"
)

//go:embed input.txt
var input string

const (
	gridAlloc = 20_000
)

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	g := grid{}
	for ; input[g.w] != '\n'; g.w++ {
	}
	g.h = len(input) / (g.w + 1)
	var start, end point
	for i, ch := range input {
		switch ch {
		case 'S':
			start = g.fromIndex(i)
		case 'E':
			end = g.fromIndex(i)
		}
	}
	dists, prevs := dijkstra(g, search{start, 0})
	p1 := math.MaxInt
	var best search
	for dir := range 4 {
		cur := search{end, byte(dir)}
		score := dists[g.searchIndex(cur)]
		if score > 0 && score < p1 {
			best = cur
			p1 = score
		}
	}
	bestSeats := make([]bool, gridAlloc)
	bestSeats[g.index(start)] = true
	bestSeats[g.index(end)] = true

	open := make([]search, 0, 4)
	next := make([]search, 0, 4)
	visited := make([]bool, gridAlloc<<2)

	open = append(open, best)
	for len(open) > 0 {
		for _, cur := range open {
			score := dists[g.searchIndex(cur)]
			for _, p := range prevs[g.searchIndex(cur)] {
				if visited[g.index(p.pos)<<2+int(p.dir)] {
					continue
				}
				visited[g.index(p.pos)<<2+int(p.dir)] = true
				pScore := dists[g.searchIndex(p)]
				if score == pScore+1 || score == pScore+1001 {
					bestSeats[g.index(p.pos)] = true
					next = append(next, p)
				}
			}
		}
		open, next = next, open[0:0]
	}
	p2 := 0
	for _, good := range bestSeats {
		if good {
			p2++
		}
	}
	return p1, p2
}

type search struct {
	pos point
	dir byte
}

func dijkstra(g grid, start search) ([]int, [][]search) {
	queue := searchQueue{}
	dist := make([]int, gridAlloc<<2)
	heap.Push(&queue, queueItem{start, 0})
	prev := make([][]search, gridAlloc<<2)
	for queue.Len() > 0 {
		cur := queue.Pop().(queueItem)
		g.eachNeighbour(cur.node, func(neigh search, cost int) {
			ni := g.searchIndex(neigh)
			alt := dist[g.searchIndex(cur.node)] + cost
			if dist[ni] == 0 || alt < dist[ni] {
				dist[ni] = alt
				prev[ni] = prev[ni][0:0]
				queue.Push(queueItem{neigh, 0})
			}
			prev[ni] = append(prev[ni], cur.node)
		})
	}
	return dist, prev
}

type queueItem struct {
	node     search
	priority int
}

type searchQueue []queueItem

func (pq searchQueue) Len() int           { return len(pq) }
func (pq searchQueue) Less(i, j int) bool { return pq[i].priority > pq[j].priority }
func (pq searchQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *searchQueue) Push(x any)        { *pq = append(*pq, x.(queueItem)) }

func (pq *searchQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type grid struct {
	h, w int
}

func (g grid) inBound(p point) bool     { return p.x >= 0 && p.x < g.w && p.y >= 0 && p.y < g.h }
func (g grid) index(p point) int        { return int(p.x) + int(p.y)*int(g.w+1) }
func (g grid) fromIndex(idx int) point  { return point{idx % (g.w + 1), idx / (g.w + 1)} }
func (g grid) searchIndex(s search) int { return g.index(s.pos)<<2 + int(s.dir) }
func (g grid) valAt(p point) byte {
	if !g.inBound(p) {
		return 0
	}
	return input[g.index(p)]
}

func (g grid) eachNeighbour(s search, callback func(next search, cost int)) {
	try := func(dir byte, cost int) {
		next := s.pos.add(dirs[dir])
		if g.valAt(next) != '#' {
			callback(search{next, dir}, cost)
		}
	}
	try(s.dir, 1)
	try((s.dir+3)%4, 1001)
	try((s.dir+1)%4, 1001)
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
