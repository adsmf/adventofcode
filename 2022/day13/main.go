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
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	sum := 0
	for i, block := range strings.Split(input, "\n\n") {
		lines := strings.Split(block, "\n")
		if compare([]byte(lines[0]), []byte(lines[1])) {
			sum += i + 1
		}
	}
	return sum
}

func part2() int {
	lines := strings.Split(input, "\n")
	toSort := make(listSort, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		toSort = append(toSort, line)
	}
	divider1, divider2 := "[[2]]", "[[6]]"
	toSort = append(toSort, divider1, divider2)
	sort.Sort(toSort)
	i1, i2 := 0, 0
	for index, entry := range toSort {
		if entry == divider1 {
			i1 = index + 1
		}
		if entry == divider2 {
			i2 = index + 1
		}
		if i1 != 0 && i2 != 0 {
			break
		}
	}
	return i1 * i2
}

type listSort []string

func (l listSort) Len() int           { return len(l) }
func (l listSort) Less(i, j int) bool { return compare([]byte(l[i]), []byte(l[j])) }
func (l listSort) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func compare(l1, l2 []byte) bool {
	v1 := []interface{}{}
	v2 := []interface{}{}

	_ = json.Unmarshal(l1, &v1)
	_ = json.Unmarshal(l2, &v2)

	result := compareLists(v1, v2)
	return result != cmpGreater
}

func compareLists(l1, l2 []interface{}) comparison {
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
