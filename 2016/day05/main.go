package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

var benchmark = false

var doorID = "cxdnnyjw"

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func part1() string {
	code := ""
	start := 0
	for i := 0; i < 8; i++ {
		var char byte
		char, start = findNexChar(start)
		code = fmt.Sprintf("%s%c", code, char)
		start++
	}
	return code
}

func part2() string {
	code := make([]byte, 8)
	start := 0
	done := map[int]bool{}

	for {
		var char byte
		var pos int
		char, pos, start = findNexChar2(start)
		if pos <= 7 {
			if _, found := done[pos]; !found {
				code[pos] = char
				done[pos] = true
			}
		}
		if len(done) == 8 {
			break
		}
		start++
	}
	return string(code)
}

func findNexChar(start int) (byte, int) {
	h := md5.New()
	for i := start; ; i++ {
		io.WriteString(h, fmt.Sprintf("cxdnnyjw%d", i))
		hash := fmt.Sprintf("%x", h.Sum(nil))
		if hash[0:5] == "00000" {
			return hash[5], i
		}
		h.Reset()
	}
}

func findNexChar2(start int) (byte, int, int) {
	h := md5.New()
	for i := start; ; i++ {
		io.WriteString(h, fmt.Sprintf("cxdnnyjw%d", i))
		hash := fmt.Sprintf("%x", h.Sum(nil))
		if hash[0:5] == "00000" {
			return hash[6], int(hash[5] - '0'), i
		}
		h.Reset()
	}
}
