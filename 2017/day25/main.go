package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/adsmf/adventofcode/utils"
	"gopkg.in/yaml.v3"
)

func main() {
	p1 := part1()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
	}
}

func part1() int {
	state, maxSteps, stateActions := parseSettings("input.txt")
	tape := map[int]tapeValue{}
	cursor := 0
	for i := 0; i < maxSteps; i++ {
		act := stateActions[state][tape[cursor]]
		tape[cursor] = act.write
		cursor += act.offset
		state = act.nextState
	}

	checksum := 0
	for _, loc := range tape {
		if loc == 1 {
			checksum++
		}
	}
	return checksum
}

func parseSettings(filename string) (stateID, int, stateActionMap) {
	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	blocks := strings.Split(string(inputBytes), "\n\n")
	initialBlock, stateBlocks := blocks[0], blocks[1:]
	initialBlockLines := strings.Split(initialBlock, "\n")
	startState := stateID(initialBlockLines[0][len("Begin in state ")] - 'A')
	steps := utils.GetInts(initialBlockLines[1])

	stateActions := stateActionMap{}
	for _, block := range stateBlocks {
		blockInfo := map[string]map[string][]string{}
		yaml.Unmarshal([]byte(block), &blockInfo)
		for stateLine, stateInfo := range blockInfo {
			inState := stateID(stateLine[len("In state ")] - 'A')
			settings := stateSettings{}
			for ruleLine, ruleActions := range stateInfo {
				ruleValue := tapeValue(ruleLine[len(ruleLine)-1] - '0')
				action := action{}
				for _, actionLine := range ruleActions {
					words := strings.Split(actionLine, " ")
					keyWord := words[len(words)-1]
					switch actionLine[0] {
					case 'W': // [W]rite the value ...
						action.write = tapeValue(keyWord[0] - '0')
					case 'M': // [M]ove ...
						switch keyWord {
						case "right.":
							action.offset = 1
						case "left.":
							action.offset = -1
						default:
							panic(fmt.Errorf("Don't know how to move '%s'", keyWord))
						}
					case 'C': // [C]ontinue with...
						action.nextState = stateID(keyWord[0] - 'A')
					}
				}
				settings[ruleValue] = action
			}
			stateActions[inState] = settings
		}
	}
	return startState, steps[0], stateActions
}

type stateActionMap map[stateID]stateSettings
type stateSettings map[tapeValue]action

type action struct {
	write     tapeValue
	offset    int
	nextState stateID
}

type stateID int

const (
	stateA stateID = iota
	stateB
	stateC
	stateD
	stateE
	stateF
)

type tapeValue = int

var benchmark = false
