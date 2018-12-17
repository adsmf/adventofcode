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
	fmt.Printf("Total: %d\n", total)
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

type Node struct {
	Header   Header
	Children []Node
	Metadata []int
}
type Header struct {
	NumChildren int
	NumMetadata int
}
