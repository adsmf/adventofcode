package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

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

func nextLine(in []byte, pos int) int {
	for ; in[pos] != '\n'; pos++ {
	}
	return pos + 1
}

func buildTree(in []byte, nodes *nodePool) int {
	stack := make(nodeList, 10)
	stackPos := -1
	nodesAllocated := 0

	var curNode *node
	for pos := 0; pos < len(in); {
		switch in[pos] {
		case '$':
			if in[pos+2] != 'c' {
				pos += 5
				continue
			}
			if in[pos+5] == '.' {
				stackPos--
				curNode = stack[stackPos]
				pos += 8
				continue
			}
			stackPos++
			curNode = &(nodes)[nodesAllocated]
			nodesAllocated++
			if stackPos > 0 {
				stack[stackPos-1].addChild(curNode)
			}
			stack[stackPos] = curNode
			pos = nextLine(in, pos)
		default:
			size := 0
			size, pos = getInt(in, pos)
			curNode.addFile(fsSize(size))
			pos = nextLine(in, pos)
		}
	}
	return nodesAllocated
}

type nodePool [180]node
type nodeList []*node

type fsSize uint32

type node struct {
	totalFileSize fsSize
	treeSize      fsSize
	children      nodeList
}

func (n *node) addChild(child *node) {
	n.children = append(n.children, child)
}
func (n *node) addFile(size fsSize) {
	n.totalFileSize += size
}

func (n *node) size() fsSize {
	if n.treeSize > 0 {
		return n.treeSize
	}
	size := n.totalFileSize
	for _, child := range n.children {
		size += child.size()
	}
	n.treeSize = size
	return n.treeSize
}

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	for ; pos < len(in) && in[pos] >= '0' && in[pos] <= '9'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	return accumulator, pos
}

var benchmark = false
