package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	input, _ := ioutil.ReadFile("input.txt")
	return sumNumbers(input, false)
}

func part2() int {
	input, _ := ioutil.ReadFile("input.txt")
	return sumNumbers(input, true)
}

func sumNumbers(input []byte, ignoreRed bool) int {
	var nodes interface{}
	err := json.Unmarshal(input, &nodes)
	if err != nil {
		panic(err)
	}

	sum, _ := sumWalkNode(nodes, ignoreRed)
	return int(sum)
}

func sumWalkNode(node interface{}, ignoreRed bool) (float64, bool) {
	sum := 0.0

	switch val := node.(type) {
	case int:
		sum += float64(val)
	case float64:
		sum += val
	case string:
		if val == "red" {
			return 0, true
		}
		return 0, false
	case map[string]interface{}:
		partSum := 0.0
		for _, child := range val {
			childSum, isRed := sumWalkNode(child, ignoreRed)
			if isRed && ignoreRed {
				return 0, false
			}
			partSum += childSum
		}
		sum += partSum
	case []interface{}:
		for _, child := range val {
			childSum, _ := sumWalkNode(child, ignoreRed)
			sum += childSum
		}
	default:
		fmt.Printf("Unknown type: %T: %v\n", val, val)
	}

	return sum, false
}

var benchmark = false
