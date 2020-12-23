package main

import (
	"container/ring"
	"fmt"
)

const d23input = "952316487"

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() string {
	game := newGame(d23input, len(d23input))
	for i := 0; i < 100; i++ {
		game.iterate()
	}
	return game.after1()
}

func part2() int {
	game := newGame(d23input, 1000000)
	for i := 0; i < 10000000; i++ {
		game.iterate()
	}
	n1 := game.from1.Next()
	n2 := n1.Next()

	return n1.Value.(int) * n2.Value.(int)
}

type gameData struct {
	cups      *ring.Ring
	from1     *ring.Ring
	positions []*ring.Ring
}

func newGame(start string, numCups int) gameData {
	g := gameData{
		cups:      ring.New(numCups),
		positions: make([]*ring.Ring, numCups),
	}

	for _, char := range start {
		num := int(char - '0')
		g.cups.Value = num
		g.positions[num-1] = g.cups
		if num == 1 {
			g.from1 = g.cups
		}
		g.cups = g.cups.Next()
	}
	for next := 10; next <= numCups; next++ {
		g.cups.Value = next
		g.positions[next-1] = g.cups
		g.cups = g.cups.Next()
	}

	return g
}

func (g *gameData) after1() string {
	aft := ""
	g.from1.Do(func(i interface{}) {
		aft += fmt.Sprintf("%c", i.(int)+'0')
	})
	return aft[1:]
}

func (g *gameData) iterate() {
	target := g.cups.Value.(int)
	pick := g.cups.Unlink(3)
	target--
	for {
		contains := false
		if target < 1 {
			target += g.cups.Len() + 3 // Because we picked 3 out
		}
		pick.Do(func(i interface{}) {
			if i == target {
				contains = true
				target--
			}
		})
		if !contains {
			break
		}
	}
	insertPos := g.positions[target-1]
	insertPos.Link(pick)
	g.cups = g.cups.Next()
}

var benchmark = false
