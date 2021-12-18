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
	newList := make(elementList, len(a)+len(b)+3)
	newList[0] = "["
	copy(newList[1:], a)
	newList[len(a)+1] = ","
	copy(newList[len(a)+2:], b)
	newList[len(newList)-1] = "]"
	return newList
}

func reduce(elements elementList) elementList {
	changed := true
	for changed {
		// fmt.Println("Reduce", elements)
		elements, changed = process(elements)
	}
	return elements
}

func parse(in string) elementList {
	elements := elementList{}
	for i := 0; i < len(in); {
		ch := byte(in[i])
		switch ch {
		case '[', ']', ',':
			elements = append(elements, string(ch))
			i++
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			acc := 0
			for ; in[i] >= '0' && in[i] <= '9'; i++ {
				acc *= 10
				acc += int(in[i] - '0')
			}
			elements = append(elements, acc)
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
		case "[":
			depth++
			if depth == 5 {
				elements = explode(elements, i)
				return elements, true
			}
		case "]":
			depth--
		}
	}
	for i, element := range elements {
		switch e := element.(type) {
		case int:
			if e > 9 {
				elements = split(elements, i)
				return elements, true
			}
		}
	}
	return elements, false
}

func explode(elements elementList, offset int) elementList {
	l, r := elements[offset+1], elements[offset+3]
	done := false
	for search := offset - 1; search > 0 && !done; search-- {
		switch e := elements[search].(type) {
		case int:
			elements[search] = e + l.(int)
			done = true
		}
	}
	done = false
	for search := offset + 5; search < len(elements) && !done; search++ {
		switch e := elements[search].(type) {
		case int:
			elements[search] = e + r.(int)
			done = true
		}
	}
	elements = append(elements[:offset], elements[offset+4:]...)
	elements[offset] = 0
	return elements
}

func split(elements elementList, offset int) elementList {
	v := elements[offset].(int)
	elements = append(elements[:offset+5], elements[offset+1:]...)
	elements[offset+0] = "["
	elements[offset+1] = v / 2
	elements[offset+2] = ","
	elements[offset+3] = (v + 1) / 2
	elements[offset+4] = "]"
	return elements
}

func magnitude(origElements elementList) int {
	elements := make(elementList, len(origElements))
	copy(elements, origElements)
	for len(elements) > 1 {
		changed := false
		for i := 0; i < len(elements)-2; i++ {
			elemLeft, elemRight := elements[i], elements[i+2]
			valLeft, isIntLeft := elemLeft.(int)
			valRight, isIntRight := elemRight.(int)
			if isIntLeft && isIntRight {
				elements = append(elements[:i], elements[i+4:]...)
				elements[i-1] = valLeft*3 + valRight*2
				changed = true
				break
			}
		}
		if !changed {
			fmt.Println("No change", elements)
			return -1
		}
	}
	return elements[0].(int)
}

type elementList []interface{}

func (el elementList) String() string {
	sb := strings.Builder{}
	for _, element := range el {
		switch e := element.(type) {
		case int:
			sb.WriteString(strconv.Itoa(e))
		case string:
			sb.WriteString(e)
			if e == "" {
				sb.WriteString("?")
			}
		default:
			sb.WriteByte('?')
		}
	}
	return sb.String()
}

var benchmark = false
