package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
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
	p1total := 0
	rules := workflowRules{}
	utils.EachSectionMB(input, "\n\n", func(sectionIdx int, section string) (done bool) {
		if sectionIdx == 0 {
			// Parse rules
			utils.EachLine(section, func(index int, value string) (done bool) {
				ruleName, rule := parseRule(value)
				rules[ruleName] = rule
				return false
			})

			return false
		}
		// Process parts
		utils.EachLine(section, func(partIdx int, def string) (done bool) {
			values := utils.GetInts(def)
			part := partInfo{}
			for i := 0; i < 4; i++ {
				part[i] = values[i]
			}
			result := evaluatePart(part, rules)
			if result {
				for _, val := range part {
					p1total += val
				}
			}
			return false
		})
		return false
	})
	bounds := boundaryList{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}
	p2total := evalBounds("in", bounds, rules)
	return p1total, p2total
}

func evalBounds(ruleName string, bounds boundaryList, rules workflowRules) int {
	switch ruleName {
	case "A":
		return bounds.len()
	case "R":
		return 0
	default:
		total := 0
		rule := rules[ruleName]
		for _, cond := range rule.conditions {
			condBound := bounds[cond.check]
			switch cond.op {
			case '>':
				if condBound.max <= cond.value {
					continue
				}
				sub := bounds
				if sub[cond.check].min <= cond.value {
					sub[cond.check].min = cond.value + 1
				}
				total += evalBounds(cond.target, sub, rules)
				bounds[cond.check].max = cond.value
			case '<':
				if condBound.min >= cond.value {
					continue
				}
				sub := bounds
				if sub[cond.check].max >= cond.value {
					sub[cond.check].max = cond.value - 1
				}
				total += evalBounds(cond.target, sub, rules)
				bounds[cond.check].min = cond.value

			}
		}
		total += evalBounds(rule.final, bounds, rules)
		return total
	}
}

func evaluatePart(part partInfo, rules workflowRules) bool {
	ruleName := "in"
	var result bool
	for {
		rule := rules[ruleName]
		ruleName = rule.evaluate(part)
		switch ruleName {
		case "A":
			return true
		case "R":
			return false
		}
		if ruleName == "A" {
			return result
		}
	}
}

type workflowRules map[string]workflowRule
type workflowRule struct {
	conditions []workflowCondition
	final      string
}

func (wr workflowRule) evaluate(part partInfo) string {
	for _, cond := range wr.conditions {
		value := part[cond.check]
		switch cond.op {
		case '>':
			if value > cond.value {
				return cond.target
			}
		case '<':
			if value < cond.value {
				return cond.target
			}
		}
	}
	return wr.final
}

func parseRule(def string) (string, workflowRule) {
	defStart := strings.Index(def, "{")
	name := def[0:defStart]
	def = def[defStart+1 : len(def)-1]
	rule := workflowRule{}
	checkSymbols := [...]byte{'x': 0, 'm': 1, 'a': 2, 's': 3}
	utils.EachSection(def, ',', func(index int, section string) (done bool) {
		if strings.Index(section, ":") != -1 {
			cond := workflowCondition{}
			parts := strings.Split(section, ":")
			cond.check = checkSymbols[parts[0][0]]
			cond.op = parts[0][1]
			cond.value, _ = strconv.Atoi(parts[0][2:])
			cond.target = parts[1]
			rule.conditions = append(rule.conditions, cond)
			return false
		}
		rule.final = section
		return false
	})
	return name, rule
}

type workflowCondition struct {
	check  byte
	op     byte
	value  int
	target string
}

type partInfo [4]int

type boundaryList [4]boundary

func (b boundaryList) len() int {
	return b[0].len() * b[1].len() * b[2].len() * b[3].len()
}

type boundary struct {
	min, max int
}

func (b boundary) len() int { return b.max - b.min + 1 }

var benchmark = false
