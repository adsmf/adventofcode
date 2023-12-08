package main

import (
	_ "embed"
	"fmt"

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
	inst, nodes := load()
	aaa := minSteps(inst, makeName("AAA"), nodes)
	lcm := aaa
	for n := range nodes {
		if n[2] == 'A' && n != makeName("AAA") {
			lcm = utils.LowestCommonMultiplePair(lcm, minSteps(inst, n, nodes))
		}
	}
	return aaa, lcm
}

func minSteps(inst string, start nodeName, nodes nodeSet) int {
	curNode := start
	steps := 0
	for {
		for _, ch := range inst {
			steps++
			switch ch {
			case 'L':
				curNode = nodes[curNode].left
			case 'R':
				curNode = nodes[curNode].right
			}
			if curNode[2] == 'Z' {
				return steps
			}
		}
	}
}

func load() (string, nodeSet) {
	instructions := ""
	nodes := make(nodeSet, 720)
	utils.EachLine(input, func(index int, line string) (done bool) {
		if index == 0 {
			instructions = line
			return false
		}
		if len(line) < 16 {
			return false
		}
		parent := makeName(line[0:3])
		left := makeName(line[7:10])
		right := makeName(line[12:15])
		nodes[parent] = node{
			left:  left,
			right: right,
		}
		return false
	})
	return instructions, nodes
}

type nodeSet map[nodeName]node

type node struct{ left, right nodeName }

type nodeName [3]byte

func makeName(in string) nodeName { return nodeName([]byte(in)) }

var benchmark = false
