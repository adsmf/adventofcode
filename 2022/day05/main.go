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
	dock1, dock2 := loadStacks()
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
	crates := make([]byte, _yardStacks)
	for i := 0; i < _yardStacks; i++ {
		crates[i] = d[i].getTop()
	}
	return crates
}

func (d dockyard) String() string {
	sb := strings.Builder{}
	maxHeight := 0
	for _, stack := range d {
		if stack.height > maxHeight {
			maxHeight = stack.height
		}
	}
	for level := maxHeight; level > 0; level-- {
		for _, stack := range d {
			if stack.height < level {
				sb.WriteString("    ")
				continue
			}
			sb.WriteByte('[')
			sb.WriteByte(stack.crates[level-1])
			sb.WriteString("] ")
		}
		sb.WriteByte('\n')
	}
	return sb.String()
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

func loadStacks() (dockyard, dockyard) {
	var layoutRows int
	const lineWidth = 4 * _yardStacks
	for layoutRows = 1; layoutRows*lineWidth < len(input); layoutRows++ {
		if input[layoutRows*lineWidth-1] == '\n' && input[layoutRows*lineWidth] == '\n' {
			break
		}
	}
	dock1, dock2 := dockyard{}, dockyard{}
	for row := 0; row < layoutRows-1; row++ {
		rowOffset := lineWidth * (layoutRows - row - 2)
		for stack := 0; stack < _yardStacks; stack++ {
			char := input[rowOffset+stack*4+1]
			if char >= 'A' && char <= 'Z' {
				dock1[stack].crates[row] = char
				dock1[stack].height = row + 1

				dock2[stack].crates[row] = char
				dock2[stack].height = row + 1
			}
		}
	}
	return dock1, dock2
}

const (
	_maxCrates  = 50
	_yardStacks = 9
)

var benchmark = false
