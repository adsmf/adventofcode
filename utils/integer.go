package utils

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

func GreatestCommonDivisorInt[T constraints.Integer](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func ExtendedGreatestCommonDivisor[T constraints.Integer](a, b T) (gcd T, x T, y T) {
	if b == 0 {
		return a, 1, 0
	}
	y = 1
	aOld, b := a, b
	for b != 0 {
		q := aOld / b
		aOld, b = b, aOld-q*b
		x, y = y, x-T(q)*y
	}
	return aOld, x, y
}

func LowestCommonMultiplePair[T constraints.Integer](a, b T) T {
	g := GreatestCommonDivisorInt(a, b)
	return a * b / g
}

func LowestCommonMultipleInt[T constraints.Integer](integers ...T) T {
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

func MustInt[T constraints.Integer](input string) T {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Errorf("Error converting %s to int: %w", input, err))
	}
	return T(val)
}

func IntAbs[T constraints.Integer](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
