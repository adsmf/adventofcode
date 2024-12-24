package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

const (
	writeDotFile = false
)

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func solve() (int, string) {
	dev := device{
		values: make(map[string]bool, 45*2),
		gates:  make(map[string]gate, 45*5),
	}
	utils.EachSectionMB(input, "\n\n", func(secIdx int, section string) (done bool) {
		if secIdx == 0 {
			utils.EachLine(section, func(index int, line string) (done bool) {
				id := line[0:3]
				val := line[5] == '1'
				dev.values[id] = val
				return
			})
			return
		}
		utils.EachLine(section, func(index int, line string) (done bool) {
			parts := strings.Split(line, " ")
			g := gate{
				l: parts[0],
				r: parts[2],
			}
			switch parts[1] {
			case "AND":
				g.op = opAND
			case "OR":
				g.op = opOR
			case "XOR":
				g.op = opXOR
			}
			dev.gates[parts[4]] = g
			return
		})
		return
	})
	p1 := 0
	for i := 0; ; i++ {
		id := gateID('z', i)
		val, err := dev.eval(id)
		if err != nil {
			break
		}
		if val {
			p1 |= 1 << i
		}
	}

	dev.usedBy = map[string]map[string]bool{}
	for tgt, g := range dev.gates {
		if dev.usedBy[g.l] == nil {
			dev.usedBy[g.l] = map[string]bool{}
		}
		if dev.usedBy[g.r] == nil {
			dev.usedBy[g.r] = map[string]bool{}
		}
		dev.usedBy[g.l][tgt] = true
		dev.usedBy[g.r][tgt] = true
	}

	if writeDotFile {
		dev.writeDot()
	}

	swapped := dev.findSwapped()
	slices.Sort(swapped)
	return p1, strings.Join(swapped, ",")
}

type device struct {
	values  map[string]bool
	gates   map[string]gate
	swapped []string
	usedBy  map[string]map[string]bool
}

func (d *device) findSwapped() []string {
	maxBit := 0
	for i := 0; i < 64; i++ {
		_, found := d.gates[gateID('z', i)]
		if !found {
			break
		}
		maxBit = i
	}
	cOut := ""
	for i := 0; i < maxBit; i++ {
		cOut = d.checkAdder(i, cOut)
		if len(d.swapped) == 8 {
			break
		}
	}
	return d.swapped
}

func (d device) findNode(in1, in2 string, op operation) string {
	for tgt := range d.usedBy[in1] {
		if d.gates[tgt].op != op {
			continue
		}
		if d.usedBy[in2][tgt] {
			return tgt
		}
	}
	return ""
}

func (d *device) checkAdder(idx int, cIn string) (cOut string) {
	/*
		Full adder:
			sum = (A ^ B) ^ cIn
			cOut = A*B + cIn*(A ^ B)
		Intermediate gates:
			xor1 => A ^ B
			and1 => AB
			and2 => xor1 * cIn
	*/

	yID := gateID('y', idx)
	zID := gateID('z', idx)
	xID := gateID('x', idx)
	zB := d.gates[zID]
	if idx == 0 {
		cOut = d.findNode("x00", "y00", opAND)
		return
	}
	var xor1, and1, and2 string
	xor1 = d.findNode(xID, yID, opXOR)
	and1 = d.findNode(xID, yID, opAND)
	and2 = d.findNode(xor1, cIn, opAND)
	cOut = d.findNode(and1, and2, opOR)

	// Swap search
	if zB.op != opXOR {
		d.swapped = append(d.swapped, zID)
		switch zB.op {
		case opOR:
			swappedWith := d.findNode(cIn, xor1, opXOR)
			cOut = swappedWith
			d.swapped = append(d.swapped, swappedWith)
		case opAND:
			swappedWith := d.findNode(xor1, cIn, opXOR)
			d.swapped = append(d.swapped, swappedWith)
			if and1 == "" || and1 == zID {
				and1 = swappedWith
			} else {
				and2 = swappedWith
			}
			cOut = d.findNode(and1, and2, opOR)
		}
	} else if cOut == "" {
		d.swapped = append(d.swapped, xor1, and1)
		xor1, and1 = and1, xor1
		and2 = d.findNode(xor1, cIn, opAND)
		cOut = d.findNode(and1, and2, opOR)
	}
	return
}

func (d *device) eval(id string) (bool, error) {
	if v, found := d.values[id]; found {
		return v, nil
	}
	g, found := d.gates[id]
	if !found {
		return false, fmt.Errorf("gate for %s not found", id)
	}
	l, err := d.eval(g.l)
	if err != nil {
		return false, err
	}
	r, err := d.eval(g.r)
	if err != nil {
		return false, err
	}
	switch g.op {
	case opAND:
		return l && r, nil
	case opOR:
		return l || r, nil
	case opXOR:
		return l != r, nil
	}
	return false, fmt.Errorf("no operation result!")
}

func (d device) writeDot() {
	dot := bytes.Buffer{}
	dot.WriteString("digraph {\n")
	nodeStyle := func(id string) {
		switch id[0] {
		case 'x':
			dot.WriteString(fmt.Sprintf("\t\t%s [shape=diamond,style=filled,fillcolor=lightcoral]\n", id))
		case 'y':
			dot.WriteString(fmt.Sprintf("\t\t%s [shape=diamond,style=filled,fillcolor=darkolivegreen1]\n", id))
		case 'z':
			dot.WriteString(fmt.Sprintf("\t\t%s [style=filled,fillcolor=cyan]\n", id))
		}
	}
	for i := 0; i < 64; i++ {
		xID := gateID('x', i)
		yID := gateID('y', i)
		zID := gateID('z', i)
		if _, found := d.gates[zID]; !found {
			break
		}
		nodeSet := map[string]bool{zID: true}
		for n := range d.usedBy[xID] {
			nodeSet[xID] = true
			nodeSet[n] = true
			if i == 0 {
				continue
			}
			for nn := range d.usedBy[n] {
				nodeSet[nn] = true
			}
		}
		for n := range d.usedBy[yID] {
			nodeSet[yID] = true
			nodeSet[n] = true
			if i == 0 {
				continue
			}
			for nn := range d.usedBy[n] {
				nodeSet[nn] = true
			}
		}
		if _, found := d.usedBy[yID]; found {
			nodeSet[yID] = true
		}
		nodes := []string{}
		for node := range nodeSet {
			if node != "" {
				nodes = append(nodes, node)
			}
		}
		dot.WriteString(fmt.Sprintf("\tsubgraph cluster_%02d { %s }\n", i, strings.Join(nodes, "; ")))
	}
	for node, g := range d.gates {
		nodeStyle(node)
		nodeStyle(g.l)
		nodeStyle(g.r)
		dot.WriteString(fmt.Sprintf("\t\t%s [shape=record,label=\"{%s|%s}\"]\n", node, g.op, node))
		dot.WriteString(fmt.Sprintf("\t\t%s -> %s\n", g.l, node))
		dot.WriteString(fmt.Sprintf("\t\t%s -> %s\n", g.r, node))
	}
	dot.WriteString("}\n")
	os.WriteFile("wiring.dot", dot.Bytes(), 0644)
}

func gateID(prefix byte, idx int) string {
	return string([]byte{prefix, byte(idx/10 + '0'), byte(idx%10 + '0')})
}

type gate struct {
	l, r string
	op   operation
}

type operation int

func (o operation) String() string {
	switch o {
	case opAND:
		return "AND"
	case opOR:
		return "OR"
	case opXOR:
		return "XOR"
	default:
		return "???"
	}
}

const (
	opUnknown operation = iota
	opAND
	opOR
	opXOR
)

var benchmark = false
