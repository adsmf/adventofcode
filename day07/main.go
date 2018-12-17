package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/adsmf/adventofcode2018/utils"
)

func main() {
	lines := utils.ReadInputLines("example.txt")
	re := regexp.MustCompile("Step (.) must be finished before step (.) can begin.")
	requirements := []requirement{}
	allSteps := make(map[string]*step)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		stepID := matches[2]
		requiresID := matches[1]
		if _, ok := allSteps[stepID]; !!!ok {
			allSteps[stepID] = &step{
				id: stepID,
			}
		}
		if _, ok := allSteps[requiresID]; !!!ok {
			allSteps[requiresID] = &step{
				id: requiresID,
			}
		}
		req := requirement{
			step:     allSteps[stepID],
			requires: allSteps[requiresID],
		}
		requirements = append(requirements, req)
	}
	fmt.Printf("Requirements: %+v\n", requirements)
	fmt.Printf("Known steps: %+v\n", allSteps)
	fmt.Println()
	curChainLength := 0
	unreadySteps := make(map[string]bool)
	for id := range allSteps {
		unreadySteps[id] = true
	}
	for len(unreadySteps) > 0 {
		newSteps := []string{}
		for stepID := range unreadySteps {
			hasDeps := false
			for _, req := range requirements {
				if req.step.id == stepID && !!!req.requires.ready {
					hasDeps = true
					break
				}
			}
			if !!!hasDeps {
				// fmt.Printf("Step ready: %s\n", stepID)
				newSteps = append(newSteps, stepID)
			}
		}
		fmt.Printf("New steps: %v\n", newSteps)
		for _, stepID := range newSteps {
			allSteps[stepID].chainLength = curChainLength
			allSteps[stepID].ready = true
			delete(unreadySteps, stepID)
		}
		curChainLength++
	}

	fmt.Printf("Chain length: %d\nSteps: %+v\n", curChainLength, allSteps)
	for chainOf := curChainLength - 1; chainOf >= 0; chainOf-- {
		availSteps := []string{}
		for id, step := range allSteps {
			if step.chainLength == chainOf {
				availSteps = append(availSteps, id)
			}
		}
		sort.Strings(availSteps)
		fmt.Printf("availSteps: %+v\n", availSteps)
	}
}

type requirement struct {
	step     *step
	requires *step
}

type step struct {
	id          string
	ready       bool
	chainLength int
}
