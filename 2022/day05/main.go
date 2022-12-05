package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	d1, d2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %s\n", d1.readTop())
		fmt.Printf("Part 2: %s\n", d2.readTop())
	}
}

func solve() (dockyard, dockyard) {
	dock1, dock2, offset := loadStacks()
	var num, from, to int
	for offset < len(input) {
		num, from, to, offset = getInstruction(offset)
		if num == 0 {
			break
		}
		for i := 0; i < num; i++ {
			dock1[from-1].sendTo(&(dock1[to-1]))
		}
		dock2[from-1].moveNTo(num, &(dock2[to-1]))
	}
	return dock1, dock2
}

func getInstruction(offset int) (int, int, int, int) {
	var num, from, to int
	num, offset = getInt(offset)
	from, offset = getInt(offset)
	to, offset = getInt(offset)
	return num, from, to, offset
}

func getInt(offset int) (int, int) {
	for ; offset < len(input) && (input[offset] < '0' || input[offset] > '9'); offset++ {
	}
	accumulator := 0
	for ; offset < len(input) && input[offset] >= '0' && input[offset] <= '9'; offset++ {
		accumulator *= 10
		accumulator += int(input[offset] & 0xf)
	}
	return accumulator, offset
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

func loadStacks() (dockyard, dockyard, int) {
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
	return dock1, dock2, layoutRows*lineWidth + 1
}

const (
	_maxCrates  = 50
	_yardStacks = 9
)

var benchmark = false
