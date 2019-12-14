package intcode

import (
	"fmt"
	"sort"
	"strings"
)

// NewMachine creates a new intcode machine
func NewMachine(options ...MachineOption) Machine {
	m := Machine{
		ram:        map[address]integer{},
		operations: map[address]operation{},
	}
	for _, option := range options {
		option(&m)
	}
	return m
}

// Machine is a virtual machine capable of running intcode (e.g. https://adventofcode.com/2019/day/2)
type Machine struct {
	model      model
	ram        ram
	operations operationMap
}

// LoadProgram wipes the machine and loads a new program from an input string
func (m *Machine) LoadProgram(program string) error {
	if m.model == nil {
		return fmt.Errorf("Cannot parse program: No intcode machine model defined")
	}
	return m.model.parse(program)
}

func (m Machine) String() string {
	state := []string{
		"Model: " + m.model.name(),
		"Decode:",
	}
	opAddresses := addressList{}
	for opAddress := range m.operations {
		opAddresses = append(opAddresses, opAddress)
	}
	sort.Sort(opAddresses)

	ramAddresses := addressList{}
	for ramAddress := range m.ram {
		ramAddresses = append(ramAddresses, ramAddress)
	}
	sort.Sort(ramAddresses)

	lastRAMAccounted := address(-1)
	for _, ramAddress := range ramAddresses {
		if ramAddress <= lastRAMAccounted {
			continue
		}
		lastRAMAccounted = ramAddress
		if len(opAddresses) == 0 || ramAddress < opAddresses[0] {
			// Not an op
			val := m.ram[ramAddress].Value()
			state = append(state, fmt.Sprintf("\t%v:\tDATA\t%d\t(%x)", address(ramAddress), val, val))
		} else {
			// Is op
			op := m.operations[ramAddress]
			state = append(state, fmt.Sprintf("\t%v:\tOPER\t%v", address(ramAddress), op))
			lastRAMAccounted += address(op.NumParams())
			opAddresses = opAddresses[1:]
		}
	}

	state = append(state, fmt.Sprintf("RAM: %v", m.ram))

	stateString := ""
	for _, line := range state {
		stateString += line + "\n"
	}

	return strings.TrimSpace(stateString)
}

func (m *Machine) readAddress(addr address) integer {
	return m.ram[addr]
}

type ram map[address]integer

func (r ram) String() string {
	retString := ""
	ramAddresses := addressList{}
	for ramAddress := range r {
		ramAddresses = append(ramAddresses, ramAddress)
	}
	sort.Sort(ramAddresses)
	lastAddress := address(-1)
	for _, ramAddress := range ramAddresses {
		val := r[ramAddress].Value()
		if ramAddress == lastAddress+1 {
			if retString != "" {
				retString += ","
			}
			retString += fmt.Sprintf("%d", val)
			lastAddress = ramAddress
		} else {
			if val != 0 {
				retString += fmt.Sprintf(",...,#%d=%d", ramAddress, val)
				lastAddress = ramAddress
			}
		}
	}
	return retString
}

type operationMap map[address]operation

func (om operationMap) String() string {
	state := []string{
		fmt.Sprintf("\tNum operations: %d", len(om)),
	}

	operationAddresses := []int{}
	for addr := range om {
		operationAddresses = append(operationAddresses, int(addr))
	}
	operationAddresses = sort.IntSlice(operationAddresses)

	for _, addr := range operationAddresses {
		state = append(state, fmt.Sprintf("\t%v: %v", address(addr), om[address(addr)]))
	}

	stateString := ""
	for _, line := range state {
		stateString += line + "\n"
	}
	return stateString

}

type model interface {
	name() string
	parse(program string) error
}

// MachineOption defines configuration options that can be applied to an intcode machine
type MachineOption func(*Machine)

type address int

func (a address) String() string {
	return fmt.Sprintf("#%04d", a)
}

type addressList []address

func (a addressList) Len() int           { return len(a) }
func (a addressList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a addressList) Less(i, j int) bool { return a[i] < a[j] }

type integer interface {
	Address() address
	IntegerType() integerType
	Value() int
}

type baseInteger struct {
	machine     *Machine
	address     address
	integerType integerType
	value       int
}

func (i baseInteger) Address() address {
	return i.address
}

func (i baseInteger) IntegerType() integerType {
	return i.integerType
}

func (i baseInteger) Value() int {
	return i.value
}

func (i baseInteger) String() string {
	return fmt.Sprintf("%d", i.value)
}

type integerType int

const (
	integerTypeInstruction integerType = iota
	integerTypeData
)

type operation interface {
	Exec()
	Name() string
	NumParams() int
}
