package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func solve() (string, string) {
	blocks := strings.Split(input, "\n\n")
	dock1, dock2 := loadStacks(blocks[1])
	for _, line := range utils.GetLines(blocks[1]) {
		vals := utils.GetInts(line)
		for i := 0; i < vals[0]; i++ {
			dock1[vals[1]-1].sendTo(&(dock1[vals[2]-1]))
		}
		dock2[vals[1]-1].moveNTo(vals[0], &(dock2[vals[2]-1]))
	}
	return string(dock1.readTop()), string(dock2.readTop())
}

type dockyard []crateStack

func (d dockyard) readTop() []byte {
	crates := make([]byte, 9)
	for i := 0; i < 9; i++ {
		crates[i] = d[i].getTop()
	}
	return crates
}

type crateStack struct {
	crates []byte
}

func (c *crateStack) sendTo(toStack *crateStack) {
	c.crates, toStack.crates = c.crates[0:len(c.crates)-1], append(toStack.crates, c.crates[len(c.crates)-1])
}

func (c *crateStack) moveNTo(numCrates int, toStack *crateStack) {
	c.crates, toStack.crates = c.crates[0:len(c.crates)-numCrates], append(toStack.crates, c.crates[len(c.crates)-numCrates:]...)
}

func (c crateStack) String() string {
	sb := strings.Builder{}
	for _, crate := range c.crates {
		sb.WriteByte(crate)
	}
	return sb.String()
}

func (c crateStack) getTop() byte {
	return c.crates[len(c.crates)-1]
}

func loadStacks(input string) (dockyard, dockyard) {
	return parseStacks(input), parseStacks(input)
}

func parseStacks(input string) []crateStack {
	stacks := []crateStack{
		{[]byte{'S', 'T', 'H', 'F', 'W', 'R'}},
		{[]byte{'S', 'G', 'D', 'Q', 'W'}},
		{[]byte{'B', 'T', 'W'}},
		{[]byte{'D', 'R', 'W', 'T', 'N', 'Q', 'Z', 'J'}},
		{[]byte{'F', 'B', 'H', 'G', 'L', 'V', 'T', 'Z'}},
		{[]byte{'L', 'P', 'T', 'C', 'V', 'B', 'S', 'G'}},
		{[]byte{'Z', 'B', 'R', 'T', 'W', 'G', 'P'}},
		{[]byte{'N', 'G', 'M', 'T', 'C', 'J', 'R'}},
		{[]byte{'L', 'G', 'B', 'W'}},
	}
	return stacks
}

var benchmark = false
