package main

import (
	"container/heap"
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

const (
	gridMax  = 70
	loadOnly = 1024
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d,%d\n", p2.x, p2.y)
	}
}

func part1() int {
	vals := utils.GetInts(input)
	g := grid{
		blocked: map[point]int{},
	}
	for i := 0; i < loadOnly; i++ {
		g.blocked[point{vals[i*2], vals[i*2+1]}] = i
	}
	start, end := point{0, 0}, point{gridMax, gridMax}
	_, prevs := dijkstra(g, start, loadOnly*2)
	p1 := 0
	for cur := end; cur != start; cur = prevs[cur] {
		p1++
	}
	return p1
}

func part2() point {
	vals := utils.GetInts(input)
	g := grid{
		blocked: map[point]int{},
	}
	for i := range len(vals) / 2 {
		g.blocked[point{vals[i*2], vals[i*2+1]}] = i
	}

	start, end := point{0, 0}, point{gridMax, gridMax}

	p2 := loadOnly + 1
	low := loadOnly + 1
	high := len(vals) / 2
	for low <= high {
		mid := (low + high) / 2
		_, prevs := dijkstra(g, start, mid)
		if prevs[end] != (point{}) {
			low = mid + 1
		} else {
			p2 = mid
			high = mid - 1
		}
	}
	return point{vals[p2*2], vals[p2*2+1]}
}

func dijkstra(g grid, start point, t int) (map[point]int, map[point]point) {
	queue := searchQueue{}
	dist := map[point]int{}
	heap.Push(&queue, queueItem{start, 0})
	prev := map[point]point{}
	for queue.Len() > 0 {
		cur := queue.Pop().(queueItem)
		neighbours := cur.node.neighbours(g, t)
		for _, neigh := range neighbours {
			alt := dist[cur.node] + 1
			if dist[neigh] == 0 || alt < dist[neigh] {
				dist[neigh] = alt
				prev[neigh] = cur.node
				queue.Push(queueItem{neigh, 0})
			}
		}
	}
	return dist, prev
}

type grid struct {
	blocked map[point]int
}

func (g grid) inBound(p point) bool {
	return p.x >= 0 && p.x <= gridMax && p.y >= 0 && p.y <= gridMax
}

type queueItem struct {
	node     point
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

type point struct{ x, y int }

func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }
func (p point) neighbours(g grid, t int) []point {
	n := make([]point, 0, 4)
	for _, dir := range dirs {
		next := p.add(dir)
		if !g.inBound(next) {
			continue
		}
		if blockedAt, found := g.blocked[next]; !found || blockedAt > t {
			n = append(n, next)
		}
	}
	return n
}

var dirs = [4]point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

var benchmark = false
