package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	return runP1("input.txt")
}

const bufferSize = 200

func part2() int {
	queue1 := make(chan int, bufferSize)
	queue2 := make(chan int, bufferSize)

	var p2 int
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		runP2("input.txt", 0, queue2, queue1)
		wg.Done()
	}()
	go func() {
		p2 = runP2("input.txt", 1, queue1, queue2)
		wg.Done()
	}()
	wg.Wait()
	return p2
}

func runP1(filename string) int {
	registers := map[string]int{}
	sound := -1

	getVal := func(input string) int {
		val, err := strconv.Atoi(input)
		if err == nil {
			return val
		}
		return registers[input]
	}

	lines := utils.ReadInputLines(filename)
	for ip := 0; ip >= 0 && ip < len(lines); ip++ {
		parts := strings.Split(lines[ip], " ")
		switch parts[0] {
		case "set":
			registers[parts[1]] = getVal(parts[2])
		case "add":
			registers[parts[1]] += getVal(parts[2])
		case "mul":
			registers[parts[1]] *= getVal(parts[2])
		case "mod":
			registers[parts[1]] %= getVal(parts[2])
		case "snd":
			sound = getVal(parts[1])
		case "rcv":
			if getVal(parts[1]) != 0 {
				return sound
			}
		case "jgz":
			if getVal(parts[1]) > 0 {
				ip += getVal(parts[2]) - 1
			}
		default:
			panic(fmt.Errorf("Unhandled operation: %s", parts[0]))
		}
	}
	return -1
}

func runP2(filename string, id int, input, output chan int) int {
	registers := map[string]int{"p": id}
	sent := 0

	getVal := func(input string) int {
		val, err := strconv.Atoi(input)
		if err == nil {
			return val
		}
		return registers[input]
	}

	lines := utils.ReadInputLines(filename)
	for ip := 0; ip >= 0 && ip < len(lines); ip++ {
		parts := strings.Split(lines[ip], " ")
		switch parts[0] {
		case "set":
			registers[parts[1]] = getVal(parts[2])
		case "add":
			registers[parts[1]] += getVal(parts[2])
		case "mul":
			registers[parts[1]] *= getVal(parts[2])
		case "mod":
			registers[parts[1]] %= getVal(parts[2])
		case "snd":
			toSend := getVal(parts[1])
			select {
			case output <- toSend:
				sent++
			default:
				panic(fmt.Errorf("Buffer too small for %d to send %d", id, toSend))
			}
		case "rcv":
			select {
			case registers[parts[1]] = <-input:
			case <-time.After(1 * time.Millisecond):
				return sent
			}
		case "jgz":
			if getVal(parts[1]) > 0 {
				ip += getVal(parts[2]) - 1
			}
		default:
			panic(fmt.Errorf("Unhandled operation: %s", parts[0]))
		}
	}
	return -1
}

var benchmark = false
