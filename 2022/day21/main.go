package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	g := loadMonkeys()
	val, _ := g.getValue("root")
	return val
}

func part2() int {
	g := loadMonkeys()
	g["root"].oper = '='
	g["humn"].oper = '?'
	return g.getInverse("root", 0)
}

type monkeyGraph map[string]*monkeyInfo

func (g monkeyGraph) getInverse(node string, target int) int {
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

func (g monkeyGraph) getValue(node string) (int, bool) {
	m := g[node]
	if m == nil {
		panic(node)
	}
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

type monkeyInfo struct {
	value  int
	m1, m2 string
	oper   byte
}

func loadMonkeys() monkeyGraph {
	g := monkeyGraph{}
	for _, line := range utils.GetLines(input) {
		parts := strings.Split(line, " ")
		m := monkeyInfo{}
		name := strings.TrimSuffix(parts[0], ":")
		switch len(parts) {
		case 2:
			m.value = utils.MustInt[int](parts[1])
		case 4:
			m.m1 = parts[1]
			m.m2 = parts[3]
			m.oper = parts[2][0]
		}
		g[name] = &m
	}
	return g
}

var benchmark = false
