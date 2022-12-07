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
	nodes := make([]*node, 0, 1000)
	buildTree(input, &nodes)
	p1 := 0
	p2 := utils.MaxInt
	spaceNeeded := -1 * (70000000 - nodes[0].size() - 30000000)
	for _, node := range nodes {
		if node.size() <= 100000 {
			p1 += node.treeSize
		}
		if node.size() >= spaceNeeded {
			if node.size() < p2 {
				p2 = node.size()
			}
		}
	}

	return p1, p2
}

func buildTree(in string, allNodes *[]*node) {
	stack := []*node{}

	var curNode *node
	for _, line := range utils.GetLines(in) {
		switch line[0] {
		case '$': // command: cd or ls
			if line[2] != 'c' {
				continue
			}
			if line[5] == '.' { // Traverse up
				stack = stack[0 : len(stack)-1]
				curNode = stack[len(stack)-1]
				continue
			}
			curNode = newNode()
			*allNodes = append(*allNodes, curNode)
			if len(stack) > 0 {
				stack[len(stack)-1].addChild(curNode)
			}
			stack = append(stack, curNode)
		case 'd': // Directory
		default: // file size
			curNode.totalFileSize += getInt(line)
		}
	}
}

func newNode() *node {
	return &node{
		children: make([]*node, 0),
		treeSize: -1,
	}
}

type nodeList []*node

type node struct {
	children      nodeList
	numChildren   int
	totalFileSize int
	treeSize      int
}

func (n *node) addChild(child *node) {
	n.numChildren++
	n.children = append(n.children, child)
}

func (n *node) size() int {
	if n.treeSize > -1 {
		return n.treeSize
	}
	size := n.totalFileSize
	for _, child := range n.children {
		size += child.size()
	}
	n.treeSize = size
	return n.treeSize
}

func getInt(in string) int {
	accumulator := 0
	for pos := 0; pos < len(in) && in[pos] >= '0' && in[pos] <= '9'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	return accumulator
}

var benchmark = false
