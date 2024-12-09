package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	p2 := part2()
	p1 := part1()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	p1 := 0
	rFence := len(input) - 1
	for ; input[rFence] == '\n'; rFence-- {
	}
	rFence = rFence &^ 1
	memPos := 0
	for lFence := 0; lFence <= rFence; lFence++ {
		free := int(input[lFence] - '0')
		if lFence&1 == 0 {
			file := (lFence / 2)
			for i := 0; i < free; i++ {
				p1 += (memPos) * file
				memPos++
			}
			continue
		}
		for free > 0 && rFence > 0 {
			file := rFence / 2
			fileSize := int(input[rFence] - '0')
			for i := 0; i < min(fileSize, free); i++ {

				p1 += (memPos) * file
				memPos++
			}
			if fileSize <= free {
				free -= fileSize
				rFence -= 2
				continue
			}
			input[rFence] = byte(fileSize-free) + '0'
			free = 0
		}
	}

	return p1
}

type freeSection uint32

const sizeOffset = 4

func (f freeSection) start() int { return int(f >> sizeOffset) }
func (f freeSection) size() byte { return byte(f & ((1 << sizeOffset) - 1)) }
func makeFreeSection(start int, size byte) freeSection {
	return freeSection(size) | freeSection(start<<sizeOffset)
}

func part2() int {
	freeSections := make([]freeSection, 0, 10000)
	totalDisk := 0
	for input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}
	for pos, free, freeStart := 0, false, 0; pos < len(input); pos++ {
		val := int(input[pos] - '0')
		totalDisk += val
		if !free {
			freeStart += val
			free = true
			continue
		}
		freeSections = append(freeSections, makeFreeSection(freeStart, byte(val)))
		freeStart += val
		free = false
	}
	p2 := 0
	rFence := len(input) - 1
	rFence = rFence &^ 1
	for ; input[rFence] == '\n'; rFence-- {
	}
	calcAt := func(file int, size int, pos int) {
		for i := 0; i < size; i++ {
			p2 += (pos + i) * file
		}
	}
	for pos := rFence; pos >= 0; pos-- {
		size := int(input[pos] - '0')
		totalDisk -= size
		if pos&1 == 1 {
			continue
		}
		added := false
		for freeIdx, free := range freeSections {
			if free.start() > totalDisk {
				break
			}
			if free.size() >= byte(size) {
				calcAt(pos/2, size, free.start())
				added = true
				if free.size() == byte(size) {
					copy(freeSections[freeIdx:], freeSections[freeIdx+1:])
					freeSections = freeSections[0 : len(freeSections)-1]
					break
				}
				freeSections[freeIdx] = makeFreeSection(free.start()+size, free.size()-byte(size))
				break
			}
		}
		if !added {
			calcAt(pos/2, size, totalDisk)
		}
	}
	return p2
}

var benchmark = false
