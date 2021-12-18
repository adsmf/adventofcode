package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	_, p1 := part1(input)
	p2 := part2(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(in string) (string, int) {
	all := parse("")

	for _, line := range utils.GetLines(in) {
		elements := parse(line)
		all = combine(all, elements)
		all = reduce(all)
	}
	final := all.String()
	return final, magnitude(all)
}

func part2(in string) int {
	opts := []elementList{}
	for _, line := range utils.GetLines(in) {
		elements := parse(line)
		opts = append(opts, elements)
	}
	best := 0
	for i := 0; i < len(opts)-1; i++ {
		for j := i + 1; j < len(opts); j++ {
			val := magnitude(reduce(combine(opts[i], opts[j])))
			if val > best {
				best = val
			}
			val2 := magnitude(reduce(combine(opts[j], opts[i])))
			if val2 > best {
				best = val2
			}
		}
	}
	return best
}

func combine(a, b elementList) elementList {
	if len(a) == 0 {
		return b
	}
	newList := make(elementList, len(a)+len(b)+2)
	newList[0] = entityOpen
	copy(newList[1:], a)
	copy(newList[len(a)+1:], b)
	newList[len(newList)-1] = entityClose
	return newList
}

func reduce(elements elementList) elementList {
	changed := true
	for changed {
		elements, changed = process(elements)
	}
	return elements
}

func parse(in string) elementList {
	elements := make(elementList, 0, len(in))
	for i := 0; i < len(in); {
		ch := byte(in[i])
		switch ch {
		case '[':
			elements = append(elements, entityOpen)
			i++
		case ']':
			elements = append(elements, entityClose)
			i++
		case ',':
			i++
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			acc := 0
			for ; in[i] >= '0' && in[i] <= '9'; i++ {
				acc *= 10
				acc += int(in[i] - '0')
			}
			elements = append(elements, entity(acc))
		case '\n':
			i++
		default:
			panic("Unhandled char: '" + string(ch) + "'")
		}
	}
	return elements
}

func process(elements elementList) (elementList, bool) {
	depth := 0
	for i, element := range elements {
		switch element {
		case entityOpen:
			depth++
			if depth == 5 {
				elements = explode(elements, i)
				return elements, true
			}
		case entityClose:
			depth--
		}
	}
	for i, element := range elements {
		if element > 9 {
			elements = split(elements, i)
			return elements, true
		}
	}
	return elements, false
}

func explode(elements elementList, offset int) elementList {
	l, r := elements[offset+1], elements[offset+2]

	for search := offset - 1; search > 0; search-- {
		if elements[search] >= 0 {
			elements[search] = elements[search] + l
			break
		}
	}
	for search := offset + 4; search < len(elements); search++ {
		if elements[search] >= 0 {
			elements[search] = elements[search] + r
			break
		}
	}
	elements = append(elements[:offset], elements[offset+3:]...)
	elements[offset] = 0
	return elements
}

func split(elements elementList, offset int) elementList {
	v := elements[offset]
	elements = append(elements[:offset+4], elements[offset+1:]...)
	elements[offset+0] = entityOpen
	elements[offset+1] = v / 2
	elements[offset+2] = (v + 1) / 2
	elements[offset+3] = entityClose
	return elements
}

func magnitude(origElements elementList) int {
	elements := make(elementList, len(origElements))
	copy(elements, origElements)
	for len(elements) > 1 {
		for i := 0; i < len(elements)-1; i++ {
			left, right := elements[i], elements[i+1]
			if left >= 0 && right >= 0 {
				elements = append(elements[:i], elements[i+3:]...)
				elements[i-1] = left*3 + right*2
			}
		}
	}
	return int(elements[0])
}

type elementList []entity
type entity int16

func (el elementList) String() string {
	sb := strings.Builder{}
	for _, element := range el {
		switch element {
		case entityClose:
			sb.WriteByte(']')
		case entityOpen:
			sb.WriteByte('[')
		default:
			sb.WriteString(strconv.Itoa(int(element)))
		}
	}
	return sb.String()
}

const (
	entityOpen  = -2
	entityClose = -1
)

var benchmark = false
