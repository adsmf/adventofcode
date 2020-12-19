package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	rules, messages := load("input.txt")
	rules.genRegex()
	p1 := part1(rules, messages)
	p2 := part2(rules, messages)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(rules ruleSpecs, messages []string) int {
	requirement := regexp.MustCompile("^" + rules[0].rawRegex + "$")
	countValid := 0
	for _, message := range messages {
		if requirement.Match([]byte(message)) {
			countValid++
		}
	}
	return countValid
}

func part2(rules ruleSpecs, messages []string) int {
	// Original:
	//    8: 42
	//   11: 42 31
	//    0: 8 11   => 42 42 31
	//
	// New:
	//    8: 42 | 42 8        => 42+
	//   11: 42 31 | 42 11 31 => 42{n} 31{n}
	//    0: 8 11             => 42+ 42{n} 31{n}   => 42{n+1,} 31{n}

	rule0 := rules[8].rawRegex + rules[11].rawRegex
	for n := 1; n <= 5; n++ {
		rule0 += fmt.Sprintf("|(%s{%d,}%s{%d})", rules[42].rawRegex, n+1, rules[31].rawRegex, n)
	}
	requirement := regexp.MustCompile("^(?:" + rule0 + ")$")
	countValid := 0
	for _, message := range messages {
		if requirement.Match([]byte(message)) {
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

type ruleSpecs map[int]*ruleSpec

func (r ruleSpecs) genRegex() {
	ruleRequires := make(map[int][]int, len(r))
	for ruleID, spec := range r {
		if spec.raw[0] == '"' {
			r[ruleID].rawRegex = strings.Trim(spec.raw, "\"")
			continue
		}
		ruleRequires[ruleID] = utils.GetInts(spec.raw)
	}

	for len(ruleRequires) > 0 {
		for ruleID, requires := range ruleRequires {
			resolvable := true
			for _, requirement := range requires {
				if _, found := ruleRequires[requirement]; found {
					resolvable = false
					break
				}
			}
			if !resolvable {
				continue
			}
			regexStringMap := map[string]bool{}
			for _, subRegexes := range strings.Split(r[ruleID].raw, " | ") {
				parts := utils.GetInts(subRegexes)
				regexString := ""
				for _, requiredRule := range parts {
					regexString += r[requiredRule].rawRegex
				}
				regexStringMap[regexString] = true
			}
			regexStrings := make([]string, 0, len(regexStringMap))
			for regexString := range regexStringMap {
				regexStrings = append(regexStrings, regexString)
			}
			r[ruleID].rawRegex = "(?:" + strings.Join(regexStrings, "|") + ")"
			delete(ruleRequires, ruleID)
		}
	}
}

type ruleSpec struct {
	raw      string
	rawRegex string
}

var benchmark = false
