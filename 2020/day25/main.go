package main

import (
	"fmt"
	"math"
	"math/big"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
	}
}

func part1() int {
	keys := utils.ReadInputLines("input.txt")
	cardPubKey := utils.MustInt(keys[0])
	doorPubKey := utils.MustInt(keys[1])
	doorLoop := findLoopBSGS(doorPubKey)
	return genKeyBig(cardPubKey, doorLoop)
}

// https://en.wikipedia.org/wiki/Baby-step_giant-step
func findLoopBSGS(public int) int {
	mod := int64(20201227)
	m := int64(math.Ceil(math.Sqrt(float64(mod))))
	lookup := make(map[int64]int64, m)
	e := int64(1)
	for j := int64(0); j < m; j++ {
		lookup[e] = j
		e *= 7
		e %= mod
	}
	factor := new(big.Int).Exp(
		big.NewInt(7),
		big.NewInt(int64(mod-m-1)),
		big.NewInt(int64(mod)),
	).Int64()
	e = int64(public)
	for i := int64(0); i < m; i++ {
		if j, found := lookup[e]; found {
			return int(i*m + j)
		}
		e *= factor
		e %= mod
	}
	return -1
}

func genKeyBig(public, private int) int {
	return int(new(big.Int).Exp(
		big.NewInt(int64(public)),
		big.NewInt(int64(private)),
		big.NewInt(20201227),
	).Int64())
}

// Simple implementations (for benchmarks)
func findLoop(public int) int {
	private := 1
	for loopSize := 1; ; loopSize++ {
		private *= 7
		private %= 20201227
		if public == private {
			return loopSize
		}
	}
}

func genKey(public, private int) int {
	key := 1
	for i := 0; i < private; i++ {
		key *= public
		key %= 20201227
	}
	return key
}

var benchmark = false
