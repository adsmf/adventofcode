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
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	m := load()
	toRx := m["rx"].received
	checkCycles := map[string]int{}
	for toOnceRemoved := range toRx {
		for rec := range m[toOnceRemoved].received {
			checkCycles[rec] = 0
		}
	}
	highPulses, lowPulses := 0, 0
	p1 := 0
	for i := 1; ; i++ {
		pulses := []signal{{"button", "broadcaster", false}}
		nextPulses := []signal{}
		for len(pulses) > 0 {
			for _, pulse := range pulses {
				if pulse.value {
					highPulses++
				} else {
					lowPulses++
				}
				done := true
				for tgt, cyc := range checkCycles {
					if cyc == 0 && pulse.target == tgt && !pulse.value {
						checkCycles[tgt] = i
						continue
					}
					if cyc == 0 {
						done = false
					}
				}
				if done {
					lcm := 1
					for _, v := range checkCycles {
						lcm = utils.LowestCommonMultiplePair(lcm, v)
					}
					return p1, lcm
				}
				toAdd := m[pulse.target].recieve(pulse.value, pulse.source)
				nextPulses = append(nextPulses, toAdd...)
			}
			pulses, nextPulses = nextPulses, pulses[0:0]
		}
		if i == 1000 {
			p1 = highPulses * lowPulses
		}
	}
}

type signal struct {
	source string
	target string
	value  bool
}

type moduleMap map[string]*module

type module struct {
	id       string
	modType  moduleType
	value    bool
	wires    []string
	received map[string]bool
}

func (m *module) recieve(value bool, from string) []signal {
	if m == nil {
		return nil
	}
	outputs := []signal{}
	switch m.modType {
	case modBroadcaster:
		for _, wire := range m.wires {
			outputs = append(outputs, signal{m.id, wire, value})
		}
	case modFlipFlop:
		if value {
			return nil
		}
		m.value = !m.value
		for _, wire := range m.wires {
			outputs = append(outputs, signal{m.id, wire, m.value})
		}
	case modConjunction:
		m.received[from] = value
		allHigh := true
		for _, v := range m.received {
			if !v {
				allHigh = false
				break
			}
		}
		for _, wire := range m.wires {
			outputs = append(outputs, signal{m.id, wire, !allHigh})
		}
	case modReceiver:
		// Do nothing?
	default:
		panic("Unhandled receiver: " + m.modType.String())
	}
	return outputs
}

func (m module) String() string {
	return fmt.Sprintf("(%s[%v] (%v)->%v)", m.modType, m.value, m.received, m.wires)
}
func (m module) GoString() string {
	return fmt.Sprintf("(%s[%v]->%v)", m.modType, m.value, m.wires)
}

type moduleType int

func (m moduleType) String() string {
	switch m {
	case modButton:
		return "button"
	case modBroadcaster:
		return "broadcaster"
	case modFlipFlop:
		return "flipflop"
	case modConjunction:
		return "conjunction"
	}
	return "unknown"
}

func load() moduleMap {
	modules := moduleMap{
		"button": &module{
			id:      "button",
			modType: modButton,
			wires:   []string{"broadcaster"},
		},
		"rx": &module{
			id:       "rx",
			modType:  modReceiver,
			wires:    []string{},
			received: map[string]bool{},
		},
	}

	utils.EachLine(input, func(lineIndex int, line string) (done bool) {
		mod := module{
			received: map[string]bool{},
		}
		utils.EachSection(line, ' ', func(fieldIndex int, field string) (done bool) {
			switch fieldIndex {
			case 0:
				switch {
				case field[0] == '%':
					mod.modType = modFlipFlop
					mod.id = field[1:]
				case field[0] == '&':
					mod.modType = modConjunction
					mod.id = field[1:]
				case field == "broadcaster":
					mod.id = field
					mod.modType = modBroadcaster
				default:
					panic(fmt.Sprintf("Unhandled module: %s", field))
				}
			case 1:
				// Ignore separator
			default:
				mod.wires = append(mod.wires, strings.TrimSuffix(field, ","))
			}
			return false
		})
		modules[mod.id] = &mod
		return false
	})
	for id, mod := range modules {
		for _, tgt := range mod.wires {
			modules[tgt].received[id] = false
		}
	}
	return modules
}

const (
	modUnknown moduleType = iota
	modButton
	modBroadcaster
	modFlipFlop
	modConjunction
	modReceiver
)

var benchmark = false
