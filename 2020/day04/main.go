package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/adsmf/adventofcode/utils"
	rex "github.com/adsmf/adventofcode/utils/regextools"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	passports := load("input.txt")
	count := 0
	for _, pass := range passports {
		if pass.validatePart1() {
			count++
		}
	}
	return count
}

func part2() int {
	passports := load("input.txt")
	count := 0
	for _, pass := range passports {
		if pass.validatePart2() {
			count++
		}
	}
	return count
}

func load(filename string) []passportData {
	lines := utils.ReadInputLines(filename)
	curPass := passportData{}
	passports := []passportData{}
	for _, line := range lines {
		if line == "" {
			passports = append(passports, curPass)
			curPass = passportData{}
			continue
		}
		fields := strings.Split(line, " ")
		for _, field := range fields {
			parts := strings.SplitN(field, ":", 2)
			curPass[parts[0]] = parts[1]
		}
	}
	if len(curPass) > 0 {
		passports = append(passports, curPass)
	}
	return passports
}

type passportData map[string]string

func (p passportData) validatePart1() bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, required := range requiredFields {
		if _, found := p[required]; !found {
			return false
		}
	}
	return true
}

var validationRegexes = map[string]*regexp.Regexp{
	"byr": rex.Anchor(rex.NumberBetween(1920, 2002)),
	"iyr": rex.Anchor(rex.NumberBetween(2010, 2020)),
	"eyr": rex.Anchor(rex.NumberBetween(2020, 2030)),
	"hcl": rex.Anchor(rex.Literal("#"), rex.Times(6, rex.HexChar)),
	"pid": rex.Anchor(rex.Times(9, rex.Digit)),
	"ecl": rex.Anchor(rex.AnyLit("amb", "blu", "brn", "gry", "grn", "hzl", "oth")),
	"hgt": rex.Anchor(rex.Any(
		rex.Group(rex.NumberBetween(150, 193), rex.Literal("cm")),
		rex.Group(rex.NumberBetween(59, 76), rex.Literal("in")),
	)),
}

func (p passportData) validatePart2() bool {
	for key, validation := range validationRegexes {
		if !validation.Match([]byte(p[key])) {
			return false
		}
	}
	return true
}
