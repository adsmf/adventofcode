package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/adsmf/adventofcode/utils"
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
	"byr": regexp.MustCompile("^(19[^01][0-9]|200[012])$"),
	"iyr": regexp.MustCompile("^(201[0-9]|2020)$"),
	"eyr": regexp.MustCompile("^(202[0-9]|2030)$"),
	"hcl": regexp.MustCompile("^#[0-9a-f]{6}$"),
	"pid": regexp.MustCompile("^[0-9]{9}$"),
	"ecl": regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$"),
	"hgt": regexp.MustCompile("^(((59|6[0-9]|7[0123456])in)|((1[5678][0-9]|19[0123])cm))$"),
}

func (p passportData) validatePart2() bool {
	for key, validation := range validationRegexes {
		if !validation.Match([]byte(p[key])) {
			return false
		}
	}
	return true
}
