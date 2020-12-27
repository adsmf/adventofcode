package main

import (
	"crypto/md5"
	"fmt"
)

const salt = "ngcjuoqr"

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	cache = []keyData{}
	lastKey := -1
	for i := 0; i < 64; i++ {
		lastKey = nextKey(salt, lastKey+1, false)
	}
	return lastKey
}

func part2() int {
	cache = []keyData{}
	lastKey := -1
	for i := 0; i < 64; i++ {
		lastKey = nextKey(salt, lastKey+1, true)
	}
	return lastKey
}

type keyData struct {
	firstTriple byte
	quintuples  map[byte]bool
	checked     bool
}

var cache = []keyData{}

func nextKey(salt string, start int, stretched bool) int {
	for i := start; ; i++ {
		data := findReps(salt, i, stretched)
		if data.firstTriple == 0 {
			continue
		}

		for offset := 1; offset <= 1000; offset++ {
			searchData := findReps(salt, i+offset, stretched)
			if searchData.quintuples[data.firstTriple] {
				return i
			}
		}
	}
}

func findReps(salt string, num int, stretched bool) keyData {
	if len(cache) > num && cache[num].checked {
		return cache[num]
	}
	for len(cache) <= num {
		cache = append(cache, keyData{})
	}
	data := keyData{
		quintuples: make(map[byte]bool),
		checked:    true,
	}
	hash := fmt.Sprintf("%s%d", salt, num)
	iterations := 1
	if stretched {
		iterations = 2017
	}
	for i := 0; i < iterations; i++ {
		md5bytes := md5.Sum([]byte(hash))
		hash = fmt.Sprintf("%0x", md5bytes)
	}
	lastC := rune(0)
	matchCount := 0
	for _, c := range hash {
		if c == lastC {
			matchCount++
			if matchCount == 3 && data.firstTriple == 0 {
				data.firstTriple = byte(c)
			}
			if matchCount == 5 {
				data.quintuples[byte(c)] = true
			}
		} else {
			matchCount = 1
			lastC = c
		}
	}
	cache[num] = data
	return data
}

var benchmark = false
