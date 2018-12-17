package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/adsmf/adventofcode2018/utils"
)

func main() {
	lines := utils.ReadInputLines("input.txt")
	re := regexp.MustCompile("Step (.) must be finished before step (.) can begin.")
	requirements := []requirement{}
	allSteps := make(map[string]*step)

	stepSequence := ""

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
				newSteps = append(newSteps, stepID)
			}
		}
		if len(newSteps) > 0 {
			sort.Strings(newSteps)
			nextStepID := newSteps[0]
			allSteps[nextStepID].ready = true
			stepSequence = fmt.Sprintf("%s%s", stepSequence, nextStepID)
			delete(unreadySteps, nextStepID)
		}
	}
	fmt.Printf("Sequence: %s\n", stepSequence)
}

type requirement struct {
	step     *step
	requires *step
}

type step struct {
	id    string
	ready bool
}
