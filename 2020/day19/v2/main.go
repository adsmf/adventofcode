package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	rules, messages := load("input.txt")
	p1 := part1(rules, messages)
	p2 := part2(rules, messages)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(rules ruleSpecs, messages []string) int {
	rule0 := rules.get(0)
	countValid := 0
	for _, message := range messages {
		if _, found := rule0[message]; found {
			countValid++
		}
	}
	return countValid
}

func part2(rules ruleSpecs, messages []string) int {
	rule42 := rules.get(42)
	rule31 := rules.get(31)

	countValid := 0
	for _, message := range messages {
		translated := ""
		for i := 0; i < len(message); i += 8 {
			part := message[i : i+8]
			if _, found := rule42[part]; found {
				translated += "1"
			} else if _, found := rule31[part]; found {
				translated += "0"
			}
		}
		for strings.Contains(translated, "1100") {
			translated = strings.ReplaceAll(translated, "1100", "10")
		}
		if !strings.Contains(translated, "11") {
			continue
		}
		for strings.Contains(translated, "11") {
			translated = strings.ReplaceAll(translated, "11", "1")
		}
		if translated == "10" {
			countValid++
		}
	}
	return countValid
}

func load(filename string) (ruleSpecs, []string) {
	inputBytes, _ := ioutil.ReadFile(filename)
	blocks := strings.Split(string(inputBytes), "\n\n")
	messages := strings.Split(strings.TrimSpace(blocks[1]), "\n")
	rulesRaw := strings.Split(blocks[0], "\n")
	rules := make(ruleSpecs, len(rulesRaw))
	for _, rule := range rulesRaw {
		specSides := strings.Split(rule, ": ")
		ruleNum, _ := strconv.Atoi(specSides[0])
		rule := ruleSpec{
			raw: specSides[1],
		}
		rules[ruleNum] = &rule
	}

	return rules, messages
}

type ruleMatches map[string]bool

type ruleSpecs map[int]*ruleSpec

func (r ruleSpecs) get(id int) ruleMatches {

	if len(r[id].options) > 0 {
		return r[id].options
	}
	spec := r[id].raw
	if spec[0] == '"' {
		r[id].options = ruleMatches{string(spec[1]): true}
		return r[id].options
	}

	r[id].options = ruleMatches{}
	for _, option := range strings.Split(spec, " | ") {
		options := map[string]bool{"": true}
		parts := utils.GetInts(option)
		for _, subRule := range parts {
			subRuleOptions := r.get(subRule)
			newOptions := make(map[string]bool, len(options)*len(subRuleOptions))
			for baseOpt := range options {
				for subOpt := range subRuleOptions {
					newOptions[baseOpt+subOpt] = true
				}
			}
			options = newOptions
		}
		for option := range options {
			r[id].options[option] = true
		}
	}

	return r[id].options
}

type ruleSpec struct {
	raw     string
	options map[string]bool
}

var benchmark = false
