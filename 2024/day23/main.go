package main

import (
	_ "embed"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func solve() (int, string) {
	pairs := connectionSet{}
	utils.EachLine(input, func(index int, line string) (done bool) {
		c1, c2 := line[0:2], line[3:5]
		if pairs[c1] == nil {
			pairs[c1] = map[string]bool{}
		}
		if pairs[c2] == nil {
			pairs[c2] = map[string]bool{}
		}
		pairs[c1][c2] = true
		pairs[c2][c1] = true
		return
	})

	p1groups := map[string]bool{}
	for c1, connections := range pairs {
		for c2 := range connections {
			for c3 := range pairs[c2] {
				if pairs[c1][c3] {
					if c1[0] != 't' && c2[0] != 't' && c3[0] != 't' {
						continue
					}
					comps := [3]string{c1, c2, c3}
					slices.Sort(comps[:])
					group := fmt.Sprintf("%s-%s-%s", comps[0], comps[1], comps[2])
					p1groups[group] = true
				}
			}
		}
	}
	p2 := password(pairs)
	return len(p1groups), p2
}

func password(pairs connectionSet) string {
	allComps := computerSet{}
	for c := range pairs {
		allComps[c] = true
	}
	set := pairs.BronKerbosch(computerSet{}, allComps, computerSet{})
	comps := make([]string, 0, len(set))
	for comp := range set {
		comps = append(comps, comp)
	}
	slices.Sort(comps)
	return strings.Join(comps, ",")
}

type connectionSet map[string]computerSet

/*
Bron-Kerbosch to find maximal clique

From https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm

	algorithm BronKerbosch1(R, P, X) is
		if P and X are both empty then
			report R as a maximal clique
		for each vertex v in P do
			BronKerbosch1(R ⋃ {v}, P ⋂ N(v), X ⋂ N(v))
			P := P \ {v}
			X := X ⋃ {v}
*/
func (cs connectionSet) BronKerbosch(R, P, X computerSet) computerSet {
	if len(P) == 0 && len(X) == 0 {
		return R
	}
	best := R
	for v := range P {
		nv := cs[v]
		res := cs.BronKerbosch(R.union(computerSet{v: true}), P.intersection(nv), X.intersection(nv))
		if len(res) > len(best) {
			best = res
		}
		delete(P, v)
		X[v] = true
	}
	return best
}

type computerSet map[string]bool

func (c computerSet) union(o computerSet) computerSet {
	set := maps.Clone(c)
	for oc := range o {
		set[oc] = true
	}
	return set
}

func (c computerSet) intersection(o computerSet) computerSet {
	set := make(computerSet, len(c))
	for comp := range c {
		if o[comp] {
			set[comp] = true
		}
	}
	return set
}

var benchmark = false
