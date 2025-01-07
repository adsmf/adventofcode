package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

const (
	writeDotFile = false
	inputBits    = 45
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
		gates: make(map[nameHash]gate, 45*5),
	}

	for pos := 0; pos < len(input)-3; pos++ {
		if input[pos] == '\n' {
			continue
		}
		if input[pos+3] == ':' {
			id := (input[pos+1]-'0')*10 + (input[pos+2] - '0')
			switch input[pos] {
			case 'x':
				dev.xVal[id] = input[pos+5] == '1'
			case 'y':
				dev.yVal[id] = input[pos+5] == '1'
			}
			pos += 5
			continue
		}
		g1 := hash(input[pos : pos+3])
		pos += 3
		g := gate{}
		switch input[pos+1] {
		case 'A':
			g.op = opAND
			pos += 5
		case 'O':
			g.op = opOR
			pos += 4
		case 'X':
			g.op = opXOR
			pos += 5
		}
		g2 := hash(input[pos : pos+3])
		g3 := hash(input[pos+7 : pos+10])
		g.l = g1
		g.r = g2
		dev.gates[g3] = g
		pos += 10
	}

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

	dev.usedBy = make(map[nameHash]map[nameHash]bool, 300)
	for tgt, g := range dev.gates {
		if dev.usedBy[g.l] == nil {
			dev.usedBy[g.l] = map[nameHash]bool{}
		}
		if dev.usedBy[g.r] == nil {
			dev.usedBy[g.r] = map[nameHash]bool{}
		}
		dev.usedBy[g.l][tgt] = true
		dev.usedBy[g.r][tgt] = true
	}

	if writeDotFile {
		dev.writeDot()
	}

	swapped := dev.findSwapped()
	slices.Sort(swapped)
	swappedStr := make([]string, len(swapped))
	for i := range len(swapped) {
		swappedStr[i] = swapped[i].String()
	}
	return p1, strings.Join(swappedStr, ",")
}

type nameHash uint16

func (n nameHash) String() string {
	s := make([]byte, 3)
	for i := range 3 {
		s[2-i] = unhashCh(byte(n % 36))
		n /= 36
	}
	return string(s)
}

func hash(name string) nameHash {
	return nameHash(hashChar(name[0]))*36*36 +
		nameHash(hashChar(name[1]))*36 +
		nameHash(hashChar(name[2]))
}

func hashChar(ch byte) byte {
	if ch >= 'a' {
		return ch - 'a' + 10
	}
	return ch - '0'
}
func unhashCh(ch byte) byte {
	if ch < 10 {
		return ch + '0'
	}
	return ch - 10 + 'a'
}

type device struct {
	xVal, yVal [inputBits]bool
	gates      map[nameHash]gate
	swapped    []nameHash
	usedBy     map[nameHash]map[nameHash]bool
}

func (d *device) findSwapped() []nameHash {
	maxBit := 0
	for i := 0; i < inputBits; i++ {
		_, found := d.gates[gateID('z', i)]
		if !found {
			break
		}
		maxBit = i
	}
	cOut := nameHash(0)
	for i := 0; i < maxBit; i++ {
		cOut = d.checkAdder(i, cOut)
		if len(d.swapped) == 8 {
			break
		}
	}
	return d.swapped
}

func (d device) findNode(in1, in2 nameHash, op operation) nameHash {
	for tgt := range d.usedBy[in1] {
		if d.gates[tgt].op != op {
			continue
		}
		if d.usedBy[in2][tgt] {
			return tgt
		}
	}
	return 0
}

func (d *device) checkAdder(idx int, cIn nameHash) (cOut nameHash) {
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
		cOut = d.findNode(hash("x00"), hash("y00"), opAND)
		return
	}
	var xor1, and1, and2 nameHash
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
			if and1 == 0 || and1 == zID {
				and1 = swappedWith
			} else {
				and2 = swappedWith
			}
			cOut = d.findNode(and1, and2, opOR)
		}
	} else if cOut == 0 {
		d.swapped = append(d.swapped, xor1, and1)
		xor1, and1 = and1, xor1
		and2 = d.findNode(xor1, cIn, opAND)
		cOut = d.findNode(and1, and2, opOR)
	}
	return
}

func (d *device) eval(id nameHash) (bool, error) {
	firstCh := id / 36 / 36
	if firstCh == 'x'-'a'+10 {
		idx := ((id/36)%36)*10 + (id % 36)
		return d.xVal[idx], nil
	} else if firstCh == 'y'-'a'+10 {
		idx := ((id/36)%36)*10 + (id % 36)
		return d.yVal[idx], nil
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
	for i := 0; i < inputBits; i++ {
		xID := gateID('x', i)
		yID := gateID('y', i)
		zID := gateID('z', i)
		if _, found := d.gates[zID]; !found {
			break
		}
		nodeSet := map[nameHash]bool{zID: true}
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
			if node != 0 {
				nodes = append(nodes, node.String())
			}
		}
		dot.WriteString(fmt.Sprintf("\tsubgraph cluster_%02d { %s }\n", i, strings.Join(nodes, "; ")))
	}
	for node, g := range d.gates {
		nodeStyle(node.String())
		nodeStyle(g.l.String())
		nodeStyle(g.r.String())
		dot.WriteString(fmt.Sprintf("\t\t%s [shape=record,label=\"{%s|%s}\"]\n", node, g.op, node))
		dot.WriteString(fmt.Sprintf("\t\t%s -> %s\n", g.l, node))
		dot.WriteString(fmt.Sprintf("\t\t%s -> %s\n", g.r, node))
	}
	dot.WriteString("}\n")
	os.WriteFile("wiring.dot", dot.Bytes(), 0644)
}

func gateID(prefix byte, idx int) nameHash {
	return hash(string([]byte{prefix, byte(idx/10 + '0'), byte(idx%10 + '0')}))
}

type gate struct {
	l, r nameHash
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
