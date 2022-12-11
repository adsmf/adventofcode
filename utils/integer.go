package utils

import (
	"fmt"
	"strconv"
)

type anyInteger interface {
	int | int8 | int16 | int32 | int64
}

func GreatestCommonDivisorInt[T anyInteger](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LowestCommonMultiplePair[T anyInteger](a, b T) T {
	g := GreatestCommonDivisorInt(a, b)
	return a * b / g
}

func LowestCommonMultipleInt[T anyInteger](integers ...T) T {
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

func MustInt[T anyInteger](input string) T {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Errorf("Error converting %s to int: %w", input, err))
	}
	return T(val)
}

func IntAbs[T anyInteger](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
