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

type dockyard [_yardStacks]crateStack

func (d dockyard) readTop() []byte {
	crates := make([]byte, 9)
	for i := 0; i < 9; i++ {
		crates[i] = d[i].getTop()
	}
	return crates
}

type crateStack struct {
	height int
	crates [_maxCrates]byte
}

func (c *crateStack) sendTo(toStack *crateStack) {
	toStack.crates[toStack.height] = c.crates[c.height-1]
	toStack.height++
	c.height--
}

func (c *crateStack) moveNTo(numCrates int, toStack *crateStack) {
	copy(toStack.crates[toStack.height:toStack.height+numCrates], c.crates[c.height-numCrates:c.height])
	toStack.height += numCrates
	c.height -= numCrates
}

func (c crateStack) String() string {
	sb := strings.Builder{}
	for _, crate := range c.crates {
		sb.WriteByte(crate)
	}
	return sb.String()
}

func (c crateStack) getTop() byte {
	return c.crates[c.height-1]
}

func loadStacks(input string) (dockyard, dockyard) {
	return hardcodedStack(input), hardcodedStack(input)
}

func hardcodedStack(input string) dockyard {
	stacks := dockyard{
		{6, [_maxCrates]byte{'S', 'T', 'H', 'F', 'W', 'R'}},
		{5, [_maxCrates]byte{'S', 'G', 'D', 'Q', 'W'}},
		{3, [_maxCrates]byte{'B', 'T', 'W'}},
		{8, [_maxCrates]byte{'D', 'R', 'W', 'T', 'N', 'Q', 'Z', 'J'}},
		{8, [_maxCrates]byte{'F', 'B', 'H', 'G', 'L', 'V', 'T', 'Z'}},
		{8, [_maxCrates]byte{'L', 'P', 'T', 'C', 'V', 'B', 'S', 'G'}},
		{7, [_maxCrates]byte{'Z', 'B', 'R', 'T', 'W', 'G', 'P'}},
		{7, [_maxCrates]byte{'N', 'G', 'M', 'T', 'C', 'J', 'R'}},
		{4, [_maxCrates]byte{'L', 'G', 'B', 'W'}},
	}
	return stacks
}

const (
	_maxCrates  = 50
	_yardStacks = 9
)

var benchmark = false
