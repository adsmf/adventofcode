package main

import (
	"fmt"
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
	doorLoop := secret(doorPubKey)
	return genKeyBig(cardPubKey, doorLoop)
}

func secret(public int) int {
	private := 1
	for loopSize := 1; ; loopSize++ {
		private *= 7
		private %= 20201227
		if public == private {
			return loopSize
		}
	}
}

func genKeyBig(public, private int) int {
	return int(new(big.Int).Exp(
		big.NewInt(int64(public)),
		big.NewInt(int64(private)),
		big.NewInt(20201227),
	).Int64())
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
