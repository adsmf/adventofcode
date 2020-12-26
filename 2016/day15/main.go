package main

import (
	"fmt"
	"math/big"

	"github.com/adsmf/adventofcode/utils"
	"github.com/adsmf/adventofcode/utils/solvers"
)

func main() {
	discs := load("input.txt")
	p1 := solve(discs)
	discs[7] = discInfo{positions: 11}
	p2 := solve(discs)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve(discs map[int]discInfo) int {
	remainders := []*big.Int{}
	mods := []*big.Int{}
	for discID, disc := range discs {
		remainders = append(remainders, big.NewInt(int64(disc.positions-(discID+disc.rot0)%disc.positions)))
		mods = append(mods, big.NewInt(int64(disc.positions)))
	}
	delay, err := solvers.ChineseRemainderTheorem(remainders, mods)
	if err != nil {
		panic(err)
	}
	return int(delay.Int64())

}

type discInfo struct {
	positions int
	rot0      int
}

func load(filename string) map[int]discInfo {
	lines := utils.ReadInputLines(filename)
	discs := map[int]discInfo{}
	for _, line := range lines {
		stats := utils.GetInts(line)
		discs[stats[0]] = discInfo{
			positions: stats[1],
			rot0:      stats[3],
		}
	}
	return discs
}

var benchmark = false
