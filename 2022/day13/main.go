package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
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
	p1 := 0
	toSort := make(listSort, 0, 300)
	for i, block := range strings.Split(input, "\n\n") {
		lines := strings.Split(block, "\n")
		parse1, parse2 := parse(lines[0]), parse(lines[1])
		if compareLists(parse1, parse2) != cmpGreater {
			p1 += i + 1
		}
		toSort = append(toSort, &parse1, &parse2)
	}
	divider1, divider2 := parse("[[2]]"), parse("[[6]]")
	toSort = append(toSort, &divider1, &divider2)
	sort.Sort(toSort)
	i1, i2 := 0, 0
	for index, entry := range toSort {
		if entry == &divider1 {
			i1 = index + 1
		}
		if entry == &divider2 {
			i2 = index + 1
		}
		if i1 != 0 && i2 != 0 {
			break
		}
	}
	return p1, i1 * i2
}

type listSort []*interfaceList

func (l listSort) Len() int           { return len(l) }
func (l listSort) Less(i, j int) bool { return compareLists(*l[i], *l[j]) != cmpGreater }
func (l listSort) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

type interfaceList []interface{}

func parse(item string) interfaceList {
	val := interfaceList{}
	_ = json.Unmarshal([]byte(item), &val)
	return val
}

func compareLists(l1, l2 interfaceList) comparison {
	min := len(l1)
	if len(l2) < min {
		min = len(l2)
	}
	for i := 0; i < min; i++ {
		e1, e2 := l1[i], l2[i]
		result := compareElements(e1, e2)
		if result != cmpEqual {
			return result
		}
	}
	if len(l1) == len(l2) {
		return cmpEqual
	}
	if len(l1) < len(l2) {
		return cmpLess
	}
	return cmpGreater
}

func compareElements(e1, e2 interface{}) comparison {
	e1float, e1isFloat := e1.(float64)
	e2float, e2isFloat := e2.(float64)
	if e1isFloat && e2isFloat {
		if e1float < e2float {
			return cmpLess
		}
		if e1float > e2float {
			return cmpGreater
		}
		return cmpEqual
	}
	var e1l, e2l []interface{}
	if e1isFloat {
		e1l = []interface{}{e1float}
	} else {
		e1l = e1.([]interface{})
	}
	if e2isFloat {
		e2l = []interface{}{e2float}
	} else {
		e2l = e2.([]interface{})
	}
	return compareLists(e1l, e2l)
}

type comparison int8

const (
	cmpLess comparison = iota - 1
	cmpEqual
	cmpGreater
)

var benchmark = false
