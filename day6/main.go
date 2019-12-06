package main

import (
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode2018/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	g := newGraph(utils.ReadInputLines("input.txt"))
	return g.countOrbits()
}

func part2() int {
	g := newGraph(utils.ReadInputLines("input.txt"))
	return g.calcRouteLength("YOU", "SAN")
}

type object struct {
	name     string
	parent   *object
	children graph
}

func (o *object) getParents() []*object {
	if o.parent == nil {
		return []*object{}
	}
	upstream := o.parent.getParents()
	return append(upstream, o.parent)
}

type graph map[string]*object

func newGraph(definitions []string) graph {
	g := graph{}
	for _, def := range definitions {
		g.addOrbit(def)
	}
	return g
}

func (g graph) calcRouteLength(a, b string) int {
	aParents := g[a].getParents()
	bParents := g[b].getParents()
	for aParents[0] == bParents[0] {
		aParents = aParents[1:]
		bParents = bParents[1:]
	}
	return len(aParents) + len(bParents)
}

func (g graph) addOrbit(definition string) {
	objectNames := strings.Split(definition, ")")
	parentName := objectNames[0]
	childName := objectNames[1]
	parent := g.getsert(parentName)
	child := g.getsert(childName)
	child.parent = parent
	parent.children[childName] = child
}

func (g graph) getsert(name string) *object {
	if obj, ok := g[name]; ok {
		return obj
	}
	obj := &object{
		name:     name,
		parent:   nil,
		children: graph{},
	}
	g[name] = obj
	return obj
}

func (g graph) countOrbits() int {
	return countChildOrbits(g.root(), 0)
}

func countChildOrbits(node *object, parents int) int {
	if node == nil {
		return 0
	}
	orbits := parents
	for _, child := range node.children {
		orbits = orbits + countChildOrbits(child, parents+1)
	}
	return orbits
}

func (g graph) root() *object {
	var root *object
	for _, m := range g {
		root = m
		break
	}
	for root.parent != nil {
		root = root.parent
	}
	return root
}
