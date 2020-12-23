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
	cups := newGameFromString(d23input)
	for i := 0; i < 100; i++ {
		cups.iterate()
	}
	return cups.after1()
}

func part2() int {
	cups := newBigGameFromString(d23input)
	for i := 0; i < 10000000; i++ {
		cups.iterate()
	}
	n1 := cups.from1.Next()
	n2 := n1.Next()

	return n1.Value.(int) * n2.Value.(int)
}

type game struct {
	cups      *ring.Ring
	from1     *ring.Ring
	positions map[int]*ring.Ring
}

func newGameFromString(start string) game {
	g := game{
		cups:      ring.New(len(start)),
		positions: map[int]*ring.Ring{},
	}

	for _, char := range start {
		num := int(char - '0')
		g.cups.Value = num
		g.positions[num] = g.cups
		if num == 1 {
			g.from1 = g.cups
		}
		g.cups = g.cups.Next()
	}

	return g
}

func newBigGameFromString(start string) game {
	g := game{
		cups:      ring.New(1000000),
		positions: map[int]*ring.Ring{},
	}

	for _, char := range start {
		num := int(char - '0')
		g.cups.Value = num
		g.positions[num] = g.cups
		if num == 1 {
			g.from1 = g.cups
		}
		g.cups = g.cups.Next()
	}
	for next := 10; next <= 1000000; next++ {
		g.cups.Value = next
		g.positions[next] = g.cups
		g.cups = g.cups.Next()
	}

	return g
}

func (g *game) after1() string {
	aft := ""
	g.from1.Do(func(i interface{}) {
		aft += fmt.Sprintf("%c", i.(int)+'0')
	})
	return aft[1:]
}

func (g *game) iterate() {
	cur := g.cups.Value
	pick := g.cups.Unlink(3)
	target := cur.(int)
	for {
		contains := false
		target--
		if target < 1 {
			target += g.cups.Len() + 3 // Because we picked 3 out
		}
		pick.Do(func(i interface{}) {
			if i == target {
				contains = true
			}
		})
		if !contains {
			break
		}
	}
	insertPos := g.positions[target]
	insertPos.Link(pick)
	g.cups = g.cups.Next()
}

var benchmark = false
