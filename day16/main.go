package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %s\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() string {
	input := loadInput("input.txt")
	result := fftString(input, 100)
	// fftResult := fftTimes(input, signal{0, 1, 0, -1}, 100)
	// result := 0
	// for i := 0; i < 8; i++ {
	// 	result += fftResult[i] * int(math.Pow(10, float64(7-i)))
	// }
	// // result := fftResult[0]*1000 + fftResult[1]*100 + fftResult[2]*10 + fftResult[3]
	// return fmt.Sprintf("%08d", result)
	return result
}

func part2() int {
	return 0
}

func fftString(input string, times int) string {
	inputSig := splitString(input)
	resultSig := fftTimes(inputSig, signal{0, 1, 0, -1}, times)

	result := ""
	for i := 0; i < 8; i++ {
		result += fmt.Sprintf("%d", resultSig[i])
	}
	return result
}

func fftTimes(input, pattern signal, times int) signal {
	next := input
	for i := 0; i < times; i++ {
		next = fft(next, pattern)
	}
	return next
}

func fft(input, pattern signal) []int {
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
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
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
