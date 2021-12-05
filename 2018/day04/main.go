package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	lines := utils.ReadInputLines("input.txt")
	logbook := genLogbook(lines)
	// fmt.Print("PART 1\n======\n")
	p1 := part1(logbook)
	// fmt.Print("\nPART 2\n======\n")
	p2 := part2(logbook)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(logbook logbook) int {
	guardsSleepTime := make(map[int]int)
	for id, entry := range logbook.entries {
		if entry.awake {
			startentry := logbook.entries[id-1]
			timediff := entry.minute - startentry.minute
			guardsSleepTime[entry.guard] += timediff
		}
	}
	sleepyGuard := 0
	sleepyGuardTime := 0
	for id, time := range guardsSleepTime {
		if time > sleepyGuardTime {
			sleepyGuardTime = time
			sleepyGuard = id
		}
	}
	sleepyMinutes := make(map[int]int)
	for id, entry := range logbook.entries {
		if entry.guard == sleepyGuard {
			if entry.awake {
				startentry := logbook.entries[id-1]
				for min := startentry.minute; min < entry.minute; min++ {
					sleepyMinutes[min]++
				}
			}
		}
	}
	sleepyMinute := 0
	sleepyMinuteDays := 0
	for minute, days := range sleepyMinutes {
		if days > sleepyMinuteDays {
			sleepyMinute = minute
			sleepyMinuteDays = days
		}
	}
	return sleepyGuard * sleepyMinute
}

func part2(logbook logbook) int {
	guardsSleepMinutes := make(map[int]map[int]int)
	for id, entry := range logbook.entries {
		if entry.awake {
			startentry := logbook.entries[id-1]
			for min := startentry.minute; min < entry.minute; min++ {
				if guardsSleepMinutes[min] == nil {
					guardsSleepMinutes[min] = make(map[int]int)
				}
				guardsSleepMinutes[min][entry.guard]++
			}
		}
	}
	sleepyDays := -1
	sleepyGuard := -1
	sleepyMinute := -1
	for min, guardDays := range guardsSleepMinutes {
		for guard, days := range guardDays {
			if days > sleepyDays {
				sleepyDays = days
				sleepyGuard = guard
				sleepyMinute = min
			}
		}
	}
	return sleepyGuard * sleepyMinute
}

func genLogbook(lines []string) logbook {
	book := logbook{}

	var guard int
	for _, line := range lines {
		date := line[1:11]
		time := line[12:17]
		timeParts := strings.Split(time, ":")
		hour, _ := strconv.Atoi(timeParts[0])
		minute, _ := strconv.Atoi(timeParts[1])
		action := line[19:]
		if strings.HasPrefix(action, "Guard") {
			actionParts := strings.Split(action, " ")
			guardString := strings.TrimPrefix(actionParts[1], "#")
			guard, _ = strconv.Atoi(guardString)
		} else if strings.HasPrefix(action, "wakes") {
			newEntry := status{
				date:   date,
				time:   time,
				hour:   hour,
				minute: minute,
				guard:  guard,
				awake:  true,
			}
			book.entries = append(book.entries, newEntry)
		} else if strings.HasPrefix(action, "falls") {
			newEntry := status{
				date:   date,
				time:   time,
				hour:   hour,
				minute: minute,
				guard:  guard,
				awake:  false,
			}
			book.entries = append(book.entries, newEntry)
		}
	}
	return book
}

type logbook struct {
	entries []status
}
type status struct {
	date   string
	time   string
	hour   int
	minute int
	guard  int
	awake  bool
}

var benchmark = false
