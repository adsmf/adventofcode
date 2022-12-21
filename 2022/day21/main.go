package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	g := loadMonkeys()
	p1 := part1(g)
	p2 := part2(g)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(g monkeyGraph) int {
	val, _ := g.getValue(nameHash([]byte("root")))
	return val
}

func part2(g monkeyGraph) int {
	g.clearCache()
	g[nameHash([]byte("root"))].oper = '='
	g[nameHash([]byte("humn"))].oper = '?'
	return g.getInverse(nameHash([]byte("root")), 0)
}

type nodeID uint32
type monkeyGraph map[nodeID]*monkeyInfo

func (g monkeyGraph) getInverse(node nodeID, target int) int {
	m := g[node]
	if m.oper == '?' {
		return target
	}
	m1val, m1f := g.getValue(m.m1)
	m2val, _ := g.getValue(m.m2)
	switch m.oper {
	case '=':
		if m1f {
			return g.getInverse(m.m2, m1val)
		}
		return g.getInverse(m.m1, m2val)
	case '+':
		if m1f {
			return g.getInverse(m.m2, target-m1val)
		}
		return g.getInverse(m.m1, target-m2val)
	case '-':
		if m1f {
			return g.getInverse(m.m2, m1val-target)
		}
		return g.getInverse(m.m1, target+m2val)
	case '/':
		if m1f {
			return g.getInverse(m.m2, target/m1val)
		}
		return g.getInverse(m.m1, target*m2val)
	case '*':
		if m1f {
			return g.getInverse(m.m2, target/m1val)
		}
		return g.getInverse(m.m1, target/m2val)
	default:
		panic(string(m.oper))
	}
}

func (g monkeyGraph) getValue(node nodeID) (int, bool) {
	m := g[node]
	if m.oper == '?' {
		return 0, false
	}
	if m.value != 0 {
		return m.value, true
	}
	m1val, m1f := g.getValue(m.m1)
	m2val, m2f := g.getValue(m.m2)
	if !(m1f && m2f) {
		return 0, false
	}
	val := int(0)
	switch m.oper {
	case '+':
		val = int(m1val + m2val)
	case '-':
		val = int(m1val - m2val)
	case '*':
		val = int(m1val * m2val)
	case '/':
		val = int(m1val / m2val)
	default:
		panic(m.oper)
	}
	g[node].value = val
	return val, true
}

func (g monkeyGraph) clearCache() {
	for m := range g {
		if g[m].oper != 0 {
			g[m].value = 0
		}
	}
}

type monkeyInfo struct {
	value  int
	m1, m2 nodeID
	oper   byte
}

func loadMonkeys() monkeyGraph {
	pool := make([]monkeyInfo, 2200)
	g := make(monkeyGraph, 2200)
	idx := 0
	for pos := 0; pos < len(input); pos++ {
		name := nameHash(input[pos : pos+4])
		m := &pool[idx]
		idx++
		val := 0
		val, pos = getInt(input, pos+6)
		if val != 0 {
			m.value = val
			g[name] = m
			continue
		}
		m.m1 = nameHash(input[pos : pos+4])
		m.m2 = nameHash(input[pos+7 : pos+11])
		m.oper = input[pos+5]
		pos += 11
		g[name] = m
	}
	return g
}

func nameHash(name []byte) nodeID {
	return nodeID(name[0]-'a') |
		(nodeID(name[1]-'a') << 5) |
		(nodeID(name[2]-'a') << 10) |
		(nodeID(name[3]-'a') << 15)
}

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	for ; in[pos] >= '0' && in[pos] <= '9'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	return accumulator, pos
}

var benchmark = false
