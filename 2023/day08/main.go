package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	inst, nodes := load()
	p1 := part1(inst, nodes)
	p2 := part2(inst, nodes)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(inst string, nodes nodeSet) int {
	return minSteps(inst, nodeName([]byte("AAA")), nodes)
}

func part2(inst string, nodes nodeSet) int {
	startNodes := []nodeName{}
	for n := range nodes {
		if n[2] == 'A' {
			startNodes = append(startNodes, n)
		}
	}
	mins := make([]int, 0, len(startNodes))
	for _, start := range startNodes {
		mins = append(mins, minSteps(inst, start, nodes))
	}
	return utils.LowestCommonMultipleInt(mins...)
}

func minSteps(inst string, start nodeName, nodes nodeSet) int {
	curNode := start
	steps := 0
	for {
		for _, ch := range inst {
			steps++
			switch ch {
			case 'L':
				curNode = nodes[curNode].l
			case 'R':
				curNode = nodes[curNode].r
			}
			if curNode[2] == 'Z' {
				return steps
			}
		}
	}
}

func load() (string, nodeSet) {
	instructions := ""
	nodes := nodeSet{}
	utils.EachLine(input, func(index int, line string) (done bool) {
		if index == 0 {
			instructions = line
			return false
		}
		if len(line) < 16 {
			return false
		}
		parName := nodeName([]byte(line[0:3]))
		c1Name := nodeName([]byte(line[7:10]))
		c2Name := nodeName([]byte(line[12:15]))
		nodes[parName] = node{
			l: c1Name,
			r: c2Name,
		}
		return false
	})
	return instructions, nodes
}

type nodeSet map[nodeName]node

type node struct{ l, r [3]byte }

type nodeName [3]byte

var benchmark = false
