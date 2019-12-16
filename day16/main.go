package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %s\n", part1())
	fmt.Printf("Part 2: %s\n", part2())
}

func part1() string {
	input := loadInput("input.txt")
	return fftString(input, 100)
}

func part2() string {
	input := loadInput("input.txt")
	return fftStringPart2(input)
}

func fftString(input string, times int) string {
	inputSig := splitString(input)
	resultSig := fftTimes(inputSig, times)

	resultBytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		resultBytes[i] = byte(resultSig[i] + '0')
	}
	return string(resultBytes)
}

func fftStringPart2(input string) string {
	input = strings.Repeat(input, 10000)
	inputSig := splitString(input)

	offset, _ := strconv.Atoi(input[0:7])

	fftPartial := fftFromEnd(inputSig, 100)

	resultBytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		resultBytes[i] = byte(fftPartial[i+offset] + '0')
	}
	return string(resultBytes)
}

func fftFromEnd(input signal, phase int) signal {
	result := append(input[0:0], input...)
	for i := 0; i < phase; i++ {
		result = calcFromEnd(result)
	}
	return result
}

func calcFromEnd(input signal) signal {
	result := make(signal, len(input))
	sum := 0
	for i := 0; i < len(input)/2; i++ {
		fromEnd := len(input) - i - 1
		sum += input[fromEnd]
		sum %= 10
		result[len(input)-i-1] = sum
	}

	return result
}

func fftTimes(input signal, times int) signal {
	next := input
	for i := 0; i < times; i++ {
		next = fft(next)
	}
	return next
}

func fft(input signal) []int {
	pattern := signal{0, 1, 0, -1}
	result := make([]int, len(input))
	for idx := range input {
		newVal := 0

		patternReps := 0
		patternPos := 0
		for _, posVal := range input {
			patternReps++
			if patternReps >= idx+1 {
				patternReps = 0
				patternPos++
				patternPos %= len(pattern)
			}
			patternVal := pattern[patternPos]
			posVal *= patternVal

			newVal += posVal
		}
		if newVal < 0 {
			newVal *= -1
		}
		newVal %= 10
		result[idx] = newVal
	}

	return result
}

type signal []int

func loadInput(filename string) string {
	fileBytes, _ := ioutil.ReadFile(filename)
	sigString := string(fileBytes)
	sigString = strings.TrimSpace(sigString)
	return sigString
}

func splitString(input string) signal {
	sig := make(signal, len(input))
	for idx, value := range input {
		sig[idx] = int(value) - int('0')
	}
	return sig
}
