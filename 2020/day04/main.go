package main

import (
	"fmt"
	"strconv"
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

func (p passportData) validatePart2() bool {
	// "byr"
	byr, err := strconv.Atoi(p["byr"])
	if err != nil || byr < 1920 || byr > 2002 {
		return false
	}
	// "iyr"
	iyr, err := strconv.Atoi(p["iyr"])
	if err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}
	// "eyr"
	eyr, err := strconv.Atoi(p["eyr"])
	if err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}
	// "hgt"
	var height int
	var suffix string
	fmt.Sscanf(p["hgt"], "%d%s", &height, &suffix)
	switch {
	case suffix == "in" && height >= 59 && height <= 76:
	case suffix == "cm" && height >= 150 && height <= 193:
	default:
		return false
	}
	// "hcl"
	if len(p["hcl"]) != 7 {
		return false
	}
	for _, ch := range p["hcl"][1:6] {
		switch true {
		case '0' <= ch && ch <= '9', 'a' <= ch && ch <= 'f':
		default:
			return false
		}
	}
	// "ecl"
	allowedECLs := map[string]struct{}{"amb": {}, "blu": {}, "brn": {}, "gry": {}, "grn": {}, "hzl": {}, "oth": {}}
	if _, found := allowedECLs[p["ecl"]]; !found {
		return false
	}
	// "pid"
	if len(p["pid"]) != 9 {
		return false
	}
	_, err = strconv.Atoi(p["pid"])
	return err == nil
}
