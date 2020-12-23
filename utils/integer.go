package utils

import (
	"fmt"
	"strconv"
)

func GreatestCommonDivisorInt(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LowestCommonMultipleInt(integers ...int) int {
	if len(integers) < 2 {
		return 0
	}
	a := integers[0]
	b := integers[1]
	integers = integers[2:]
	g := GreatestCommonDivisorInt(a, b)
	if g == 0 {
		return 0
	}
	result := a * b / g

	for i := 0; i < len(integers); i++ {
		result = LowestCommonMultipleInt(result, integers[i])
	}

	return result
}

func MustInt(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Errorf("Error converting %s to int: %w", input, err))
	}
	return val
}
