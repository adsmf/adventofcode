package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	lines := utils.ReadInputLines("input.txt")
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
	fmt.Printf("Part 1\n======\n")
	part1(allSteps, requirements)
	for _, step := range allSteps {
		step.ready = false
	}
	fmt.Printf("\nPart 2\n======\n")
	part2(allSteps, requirements)
}

func part1(allSteps map[string]*step, requirements []requirement) {
	stepSequence := ""
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

func part2(allSteps map[string]*step, requirements []requirement) {
	numWorkers := 5
	baseTime := 60
	stepSequence := ""
	unreadySteps := make(map[string]bool)
	for id := range allSteps {
		unreadySteps[id] = true
	}
	tick := 0
	var workers []worker

	for i := 0; i < numWorkers; i++ {
		workers = append(workers, worker{})
	}
	for len(unreadySteps) > 0 {
		availableWorkers := []int{}
		for workerID, worker := range workers {
			if workers[workerID].busyUntil == tick {
				workerTask := worker.task
				if workerTask != "" {
					allSteps[workerTask].ready = true
					delete(unreadySteps, workerTask)
					stepSequence = fmt.Sprintf("%s%s", stepSequence, workerTask)
				}
			}
			if tick >= worker.busyUntil {
				availableWorkers = append(availableWorkers, workerID)
			}
		}

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
		sort.Strings(newSteps)
		for _, workerID := range availableWorkers {
			for _, step := range newSteps {
				if !!!allSteps[step].started {
					taskTime := baseTime + int(step[0]-"A"[0]+1)
					busyUntil := taskTime + tick
					allSteps[step].started = true
					workers[workerID].busyUntil = busyUntil
					workers[workerID].task = step
					break
				}
			}
		}
		if len(unreadySteps) == 0 {
			break
		}
		tick++
		if tick > 999 {
			fmt.Printf("%d\tuh oh\n", tick)
			break
		}
	}
	fmt.Printf("Time taken: %d (seq: %s)\n", tick, stepSequence)
}

type requirement struct {
	step     *step
	requires *step
}

type step struct {
	id      string
	ready   bool
	started bool
}

type worker struct {
	task      string
	busyUntil int
}
