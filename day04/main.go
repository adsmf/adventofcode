package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode2018/utils"
)

func main() {
	lines := utils.ReadInputLines("input.txt")
	logbook := genLogbook(lines)
	part1(logbook)
}

func part1(logbook logbook) {
	fmt.Printf("Logbook: %+v\n", logbook.entries[0])
	guardsSleepTime := make(map[int]int)
	for id, entry := range logbook.entries {
		if entry.awake {
			startentry := logbook.entries[id-1]
			timediff := entry.minute - startentry.minute
			guardsSleepTime[entry.guard] += timediff
		}
	}
	fmt.Printf("Sleep times: %+v\n", guardsSleepTime)
	sleepyGuard := 0
	sleepyGuardTime := 0
	for id, time := range guardsSleepTime {
		if time > sleepyGuardTime {
			sleepyGuardTime = time
			sleepyGuard = id
		}
	}
	fmt.Printf("Sleepy guard is %d with %d mins asleep\n", sleepyGuard, sleepyGuardTime)
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
	fmt.Printf("Guard is most sleepy at %d (%d days)\n", sleepyMinute, sleepyMinuteDays)

	fmt.Printf("Part 1 answer: %d\n", sleepyGuard*sleepyMinute)
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
