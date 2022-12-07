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

func solve() (fsSize, fsSize) {
	nodes := nodePool{}
	numNodes := buildTree(input, &nodes)
	p1 := fsSize(0)
	p2 := fsSize(0)
	spaceNeeded := (nodes[0].size() + 30000000 - 70000000)
	for i := 0; i < numNodes; i++ {
		node := nodes[i]
		if node.size() <= 100000 {
			p1 += node.treeSize
		}
		if node.size() >= spaceNeeded {
			if node.size() < p2 || p2 == 0 {
				p2 = node.size()
			}
		}
	}

	return p1, p2
}

func buildTree(in string, nodes *nodePool) int {
	stack := nodeList{}
	stackPos := -1
	nodesAllocated := 0

	var curNode *node
	for _, line := range utils.GetLines(in) {
		switch line[0] {
		case '$': // command: cd or ls
			if line[2] != 'c' {
				continue
			}
			if line[5] == '.' { // Traverse up
				stackPos--
				curNode = stack[stackPos]
				continue
			}
			stackPos++
			curNode = &(nodes)[nodesAllocated]
			nodesAllocated++
			if stackPos > 0 {
				stack[stackPos-1].addChild(curNode)
			}
			stack[stackPos] = curNode
		case 'd': // Directory
		default: // file size
			curNode.addFile(fsSize(getInt(line)))
		}
	}
	return nodesAllocated
}

type nodePool [180]node
type nodeList [10]*node

type fsSize uint32

type node struct {
	totalFileSize fsSize
	treeSize      fsSize
	children      nodeList
	numChildren   byte
}

func (n *node) addChild(child *node) {
	n.children[n.numChildren] = child
	n.numChildren++
}
func (n *node) addFile(size fsSize) {
	n.totalFileSize += size
}

func (n *node) size() fsSize {
	if n.treeSize > 0 {
		return n.treeSize
	}
	size := n.totalFileSize
	for i := byte(0); i < n.numChildren; i++ {
		size += n.children[i].size()
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
