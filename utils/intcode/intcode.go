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
		registers: registerList{
			RegisterInstructionPointer: 0,
		},
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
	registers  registerList
}

// LoadProgram wipes the machine and loads a new program from an input string
func (m *Machine) LoadProgram(program string) error {
	if m.model == nil {
		return fmt.Errorf("Cannot parse program: No intcode machine model defined")
	}
	return m.model.parse(program)
}

// Register reads the value from a machine register
func (m *Machine) Register(reg registerID) int {
	return m.registers[reg]
}

// Step executes a single operation on the processor
func (m *Machine) Step() ExecReturnCode {
	ip := address(m.registers[RegisterInstructionPointer])
	op := m.model.decodeAddress(ip)
	if op == nil {
		panic(fmt.Sprintf("Unable to decode op att address %v", ip))
	}
	m.operations[ip] = op

	return op.Exec()
}

// Run runs the processor until a halt signal is hit
func (m *Machine) Run(stopOnInterrupt bool) {
	for {
		rc := m.Step()
		switch rc {
		case ExecRCNone:
		case ExecRCInterrupt:
			if stopOnInterrupt {
				return
			}
		default:
			return
		}
	}
}

func (m Machine) String() string {
	state := []string{
		"Model: " + m.model.name(),
	}

	state = append(state, "Registers:")
	for reg, value := range m.registers {
		state = append(state, fmt.Sprintf("\t%d: %d", reg, value))
	}

	state = append(state, "Decode:")
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
	// return m.ram[addr]
	if value, found := m.ram[addr]; found {
		return value
	}
	return &baseInteger{
		machine: m,
		address: addr,
	}
}

func (m *Machine) writeAddress(addr address, value int) {
	if m.ram[addr] == nil {
		m.ram[addr] = &baseInteger{
			machine: m,
			address: addr,
			value:   value,
		}
	} else {
		m.ram[addr].Set(value)
	}
	delete(m.operations, addr)
	// if _,found := m.operations[addr]; found {
	// 	// TODO decode?
	// }
}

// Register reads the value from a machine register
func (m *Machine) setRegister(reg registerID, value int) {
	m.registers[reg] = value
}

type registerList map[registerID]int

const (
	_ registerID = iota // Skip 0

	// RegisterInstructionPointer register stores the current processor instruction pointer
	RegisterInstructionPointer

	registerCommonEnd // Used for derived models to continue numbering
)

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
	decodeAddress(addr address) operation
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
	Value() int
	Set(int)
}

type registerID int

type baseInteger struct {
	machine *Machine
	address address
	value   int
}

func (i baseInteger) Address() address {
	return i.address
}

func (i baseInteger) Value() int {
	return i.value
}

func (i *baseInteger) Set(value int) {
	i.value = value
}

func (i baseInteger) String() string {
	return fmt.Sprintf("%d", i.value)
}

type operation interface {
	Exec() ExecReturnCode
	Name() string
	NumParams() int
}

// ExecReturnCode represents the return code from executing an operation
type ExecReturnCode int

const (
	// ExecRCNone indicates a normal (non-halting) operation
	ExecRCNone ExecReturnCode = iota

	// ExecRCInvalidInstruction indicates an invalid instruction execution was attempted
	ExecRCInvalidInstruction

	// ExecRCHCF indicates that the machine should Halt and Catch Fire
	ExecRCHCF

	// ExecRCInterrupt indicates that operation triggered an interrupt
	ExecRCInterrupt
)
