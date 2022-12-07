package main

import (
	_ "embed"
	"fmt"

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
	nodes := buildTree(input)
	size := 0
	for _, node := range nodes {
		if node.size() <= 100000 {
			size += node.treeSize
		}
	}
	return size
}

func part2() int {
	minDir := utils.MaxInt
	nodes := buildTree(input)
	spaceNeeded := -1 * (70000000 - nodes[0].size() - 30000000)
	for _, node := range nodes {
		if node.size() >= spaceNeeded {
			if node.size() < minDir {
				minDir = node.size()
			}
		}
	}
	return minDir
}

func buildTree(input string) []*node {
	stack := []*node{}
	allNodes := []*node{}

	var curNode *node
	for _, line := range utils.GetLines(input) {
		switch line[0] {
		case '$': // command: cd or ls
			if line[2] == 'c' && line[5] == '.' { // Traverse up
				stack = stack[0 : len(stack)-1]
				curNode = stack[len(stack)-1]
			} else if line[2] == 'c' { // cd
				path := line[5:]
				if len(stack) < 1 {
					curNode = &node{
						name:     "/",
						children: map[string]*node{},
						treeSize: -1,
					}
					allNodes = append(allNodes, curNode)
				} else {
					curNode = stack[len(stack)-1].children[path]
				}
				stack = append(stack, curNode)
			}
		case 'd': // Directory
			name := line[4:]
			child := &node{
				name:     name,
				children: map[string]*node{},
				treeSize: -1,
			}
			curNode.children[name] = child
			allNodes = append(allNodes, child)
		default: // file size
			curNode.totalFileSize += getInt(line)
		}
	}
	return allNodes
}

type node struct {
	name          string
	children      map[string]*node
	totalFileSize int
	treeSize      int
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
