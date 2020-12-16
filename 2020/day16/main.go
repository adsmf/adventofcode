package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1, p2 := analyseTickets()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func analyseTickets() (int, int) {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	blocks := strings.Split(string(inputBytes), "\n\n")

	db := ticketDatabase{
		fields: map[string]validFieldValues{},
	}

	validValues := validFieldValues{}
	for _, line := range strings.Split(blocks[0], "\n") {
		parts := strings.Split(line, ":")
		fieldName := parts[0]
		ints := utils.GetInts(parts[1])
		db.fields[fieldName] = validFieldValues{}
		for i := ints[0]; i <= ints[1]; i++ {
			db.fields[fieldName][i] = true
			validValues[i] = true
		}
		for i := ints[2]; i <= ints[3]; i++ {
			db.fields[fieldName][i] = true
			validValues[i] = true
		}
	}

	errorRate := 0
	ownTicket := utils.GetInts(strings.Split(blocks[1], "\n")[1])
	db.ownTicket = ownTicket
	db.tickets = []ticketValues{ownTicket}
	for _, line := range strings.Split(blocks[2], "\n")[1:] {
		ticket := utils.GetInts(line)
		valid := true
		for _, value := range ticket {
			if !validValues[value] {
				errorRate += value
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		db.tickets = append(db.tickets, ticket)
	}

	possibleFieldPositions := map[string]validFieldValues{}
	for field := range db.fields {
		possibleFieldPositions[field] = validFieldValues{}
		for i := 0; i < len(db.ownTicket); i++ {
			possibleFieldPositions[field][i] = true
		}
	}
	for _, ticket := range db.tickets {
		for pos, value := range ticket {
			for field, fieldPosValid := range possibleFieldPositions {
				if fieldPosValid[pos] {
					if !db.fields[field][value] {
						delete(possibleFieldPositions[field], pos)
					}
				}
			}
		}
	}

	fieldLocations := map[string]int{}
	for len(possibleFieldPositions) > 0 {
		for field, positions := range possibleFieldPositions {
			if len(positions) == 1 {
				position := 0
				for pos := range positions {
					position = pos
					break
				}
				fieldLocations[field] = position
				for otherField := range possibleFieldPositions {
					delete(possibleFieldPositions[otherField], position)
				}
				delete(possibleFieldPositions, field)
			}
		}
	}

	departureFieldMult := 1
	for field := range db.fields {
		if strings.HasPrefix(field, "departure") {
			departureFieldMult *= db.ownTicket[fieldLocations[field]]
		}
	}

	return errorRate, departureFieldMult
}

type ticketValues []int
type validFieldValues map[int]bool
type ticketDatabase struct {
	fields    map[string]validFieldValues
	ownTicket ticketValues
	tickets   []ticketValues
}

var benchmark = false
