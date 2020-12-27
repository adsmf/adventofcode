package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const (
	input = "qzthpkfp"
)

func main() {
	p1, p2 := solve(input)
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve(seed string) (string, int) {
	start := vaultState{0, 0, seed}
	shortestPath := ""
	longestPath := ""
	openStates := []vaultState{start}
	for i := 0; ; i++ {
		nextOpen := []vaultState{}
		for _, state := range openStates {
			for _, path := range state.paths() {
				if path.x == 3 && path.y == 3 {
					route := strings.TrimPrefix(path.route, seed)
					if shortestPath == "" {
						shortestPath = route
					}
					longestPath = route
				} else {
					nextOpen = append(nextOpen, path)
				}
			}
		}
		openStates = nextOpen
		if len(openStates) == 0 {
			break
		}
	}
	return shortestPath, len(longestPath)
}

type vec struct{ x, y int }

type vaultState struct {
	x, y  int
	route string
}

var dirChars = []string{"U", "D", "L", "R"}
var dirVectors = []vec{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func (v vaultState) paths() []vaultState {
	openPaths := []vaultState{}
	hash := fmt.Sprintf("%0x", md5.Sum([]byte(v.route)))
	for offset, dir := range dirChars {
		if hash[offset] >= 'b' {
			dirVec := dirVectors[offset]
			newX, newY := v.x+dirVec.x, v.y+dirVec.y
			if newX >= 0 && newX < 4 && newY >= 0 && newY < 4 {
				openPaths = append(openPaths, vaultState{
					newX, newY,
					v.route + dir,
				})
			}
		}
	}
	return openPaths
}

var benchmark = false
