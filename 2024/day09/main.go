package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	mem := []int{}
	file := 0
	firstFree := 0
	for pos := 0; pos < len(input); pos++ {
		if input[pos] == '\n' {
			break
		}
		val := int(input[pos] - '0')
		if pos&1 == 0 {
			for i := 0; i < val; i++ {
				mem = append(mem, file)
			}
			file++
			continue
		}
		for i := 0; i < val; i++ {
			mem = append(mem, empty)
		}
	}
	for pos := len(mem) - 1; pos >= firstFree; pos-- {
		if mem[pos] == empty {
			mem = mem[:len(mem)-1]
			continue
		}
		for ; mem[firstFree] != empty && firstFree < pos; firstFree++ {
		}
		if mem[firstFree] != empty {
			break
		}
		mem[firstFree] = mem[len(mem)-1]
		mem = mem[:len(mem)-1]
	}
	p1 := 0
	for pos, file := range mem {
		p1 += pos * int(file)
	}
	return p1
}

func part2() int {
	sections := []section{}

	for pos, file, free := 0, 0, false; pos < len(input); pos, free = pos+1, !free {
		if input[pos] == '\n' {
			break
		}
		val := int(input[pos] - '0')
		sec := section{
			size: val,
		}
		if free {
			sec.file = empty
		} else {
			sec.file = file
			file++
		}
		sections = append(sections, sec)
	}
	for pos := len(sections) - 1; pos > 0; pos-- {
		sec := sections[pos]
		if sec.file == empty {
			continue
		}
		for prevPos := 0; prevPos < pos; prevPos++ {
			prevSec := sections[prevPos]
			if prevSec.file != empty {
				continue
			}
			if prevSec.size == sec.size {
				sections[prevPos].file, sections[pos].file = sections[pos].file, sections[prevPos].file
				break
			}
			if prevSec.size > sec.size {
				remaining := section{
					file: empty,
					size: prevSec.size - sec.size,
				}

				newSections := make([]section, 0, len(sections))
				newSections = append(newSections, sections[:prevPos]...)
				newSections = append(newSections, sec)
				newSections = append(newSections, remaining)
				newSections = append(newSections, sections[prevPos+1:pos]...)
				newSections = append(newSections, section{file: empty, size: sec.size})
				newSections = append(newSections, sections[pos+1:]...)
				sections = newSections
				break
			}
		}
	}
	pos := 0
	p2 := 0
	for _, sec := range sections {
		if sec.file == empty {
			pos += sec.size
			continue
		}
		for i := 0; i < sec.size; i++ {
			p2 += (pos + i) * sec.file

		}
		pos += sec.size
	}
	return p2
}

type section struct {
	file int
	size int
}

const empty = -1

var benchmark = false
