package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"math"
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
	for p, score := range dists {
		if p.pos == end {
			if score < p1 {
				best = p
				p1 = score
			}
		}
	}
	bestSeats := map[point]bool{start: true, end: true}
	open := []search{best}
	next := []search{}
	visited := map[search]bool{}
	for len(open) > 0 {
		for _, cur := range open {
			score := dists[cur]
			for _, p := range prevs[cur] {
				if visited[p] {
					continue
				}
				visited[p] = true
				pScore := dists[p]
				if score == pScore+1 || score == pScore+1001 {
					bestSeats[p.pos] = true
					next = append(next, p)
				}
			}
		}
		open, next = next, open[0:0]
	}

	return p1, len(bestSeats)
}

type search struct {
	pos point
	dir byte
}

func (s search) neighbours(g grid) ([]search, []int) {
	n := make([]search, 0, 3)
	cost := make([]int, 0, 3)
	dir := s.dir
	dirL := (s.dir + 3) % 4
	dirR := (s.dir + 1) % 4

	next := s.pos.add(dirs[dir])
	nextL := s.pos.add(dirs[dirL])
	nextR := s.pos.add(dirs[dirR])

	if g.valAt(next) != '#' {
		n = append(n, search{next, dir})
		cost = append(cost, 1)
	}
	if g.valAt(nextL) != '#' {
		n = append(n, search{nextL, dirL})
		cost = append(cost, 1001)
	}
	if g.valAt(nextR) != '#' {
		n = append(n, search{nextR, dirR})
		cost = append(cost, 1001)
	}
	return n, cost
}

func dijkstra(g grid, start search) (map[search]int, map[search][]search) {
	queue := searchQueue{}
	dist := map[search]int{}
	heap.Push(&queue, queueItem{start, 0})
	prev := map[search][]search{}
	for queue.Len() > 0 {
		cur := queue.Pop().(queueItem)
		neighbours, costs := cur.node.neighbours(g)
		for i, neigh := range neighbours {
			alt := dist[cur.node] + costs[i]
			prev[neigh] = append(prev[neigh], cur.node)
			if dist[neigh] == 0 || alt < dist[neigh] {
				dist[neigh] = alt
				queue.Push(queueItem{neigh, 0})
			}
		}
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

func (g grid) inBound(p point) bool    { return p.x >= 0 && p.x < g.w && p.y >= 0 && p.y < g.h }
func (g grid) index(p point) int       { return int(p.x) + int(p.y)*int(g.w+1) }
func (g grid) fromIndex(idx int) point { return point{idx % (g.w + 1), idx / (g.w + 1)} }
func (g grid) valAt(p point) byte {
	if !g.inBound(p) {
		return 0
	}
	return input[g.index(p)]
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
