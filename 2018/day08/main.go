package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	rawString, _ := ioutil.ReadFile("input.txt")
	dataEntryStrings := strings.Split(string(rawString), " ")
	dataEntries := []int{}
	for _, entry := range dataEntryStrings {
		entryInt, _ := strconv.Atoi(entry)
		dataEntries = append(dataEntries, entryInt)
	}
	root, _ := processEntries(dataEntries)
	total := sumMeta(root)
	value := calcValue(root)

	if !benchmark {
		fmt.Printf("Part 1: %d\n", total)
		fmt.Printf("Part 2: %d\n", value)
	}
}

func calcValue(node Node) int {
	if len(node.Children) == 0 {
		return sumMeta(node)
	}
	value := 0
	for _, meta := range node.Metadata {
		if meta == 0 {
			continue
		}
		if meta <= len(node.Children) {
			value += calcValue(node.Children[meta-1])
		}
	}
	return value
}

func sumMeta(node Node) int {
	total := 0
	for _, child := range node.Children {
		total += sumMeta(child)
	}
	for _, meta := range node.Metadata {
		total += meta
	}
	return total
}

func processEntries(entries []int) (Node, []int) {
	if len(entries) == 0 {
		return Node{}, []int{}
	}
	numChildren := entries[0]
	numMetadata := entries[1]
	remainder := entries[2:]
	node := Node{}
	node.Header = Header{numChildren, numMetadata}
	for child := 0; child < numChildren; child++ {
		var childNode Node
		childNode, remainder = processEntries(remainder)
		node.Children = append(node.Children, childNode)
	}
	node.Metadata = remainder[:numMetadata]
	if len(remainder) == numMetadata {
		remainder = []int{}
	} else {
		remainder = remainder[numMetadata:]
	}
	return node, remainder
}

// Node represents the defined node structure
type Node struct {
	Header   Header
	Children []Node
	Metadata []int
}

// Header is the fixed length portion of a node
type Header struct {
	NumChildren int
	NumMetadata int
}

var benchmark = false
