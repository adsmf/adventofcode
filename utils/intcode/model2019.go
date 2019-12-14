package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// M19RegisterOutput is the register used to store the last value output
	M19RegisterOutput registerID = iota + registerCommonEnd
)

// M19 sets the behaviour of the intcode machine to AoC 2019 rules
func M19(input <-chan int, output chan<- int) MachineOption {
	return func(m *Machine) { m.model = &m19{machine: m} }
}

type m19 struct {
	machine *Machine

	input  <-chan int
	output chan<- int
}

func (m *m19) name() string {
	return "M19"
}

func (m *m19) parse(program string) error {
	programIntStrings := strings.Split(strings.TrimSpace(program), ",")

	for pos, valString := range programIntStrings {
		value, err := strconv.Atoi(valString)
		if err != nil {
			panic(err)
		}
		decode := &baseInteger{
			machine: m.machine,
			address: address(pos),
			value:   value,
		}
		m.machine.ram[address(pos)] = decode
	}
	m.guessOps()
	return nil
}

func (m *m19) decodeAddress(addr address) operation {
	op := &m19operation{
		baseInteger: &baseInteger{
			machine: m.machine,
			address: addr,
			value:   m.machine.ram[addr].Value(),
		},
	}
	opCode := m19operationCode(op.Value() % 100)
	// opMode := op.Value() / 100
	// fmt.Printf("Decoding operation %d: %d / %d\n", addr, opCode, opMode)
	switch opCode {
	case m19OpAdd:
		op.repr = "ADD"
		numParams := 3
		op.mode = make([]m19opMode, numParams)
		op.numParams = numParams
	case m19OpMultiply:
		op.repr = "MUL"
		numParams := 3
		op.mode = make([]m19opMode, numParams)
		op.numParams = numParams
	case m19OpInput:
		op.repr = "INP"
		numParams := 1
		op.mode = make([]m19opMode, numParams)
		op.numParams = numParams
	case m19OpOutput:
		op.repr = "OUT"
		numParams := 1
		op.mode = make([]m19opMode, numParams)
		op.numParams = numParams
	case 5:
	case 6:
	case 7:
	case 8:
	case m19OpHCF:
		op.repr = "HCF"
		op.numParams = 0
	default:
		op.repr = fmt.Sprintf("UNK-%d", opCode)
		return nil
	}
	return op
}

func (m *m19) guessOps() {
	// Scrolling up to len(ram) is fine in the initial case, will need changing if re-running later
	for addr := address(0); int(addr) < len(m.machine.ram); addr++ {
		op := m.decodeAddress(addr)
		if op == nil {
			return
		}
		m19op := op.(*m19operation)
		m19op.guessed = true

		m.machine.ram[addr] = m19op
		m.machine.operations[addr] = m19op
		addr += address(m19op.numParams)
	}
}

type m19opMode int

const (
	m19opModePositional m19opMode = iota
	m19opModeImmediate
	m19opModeRelative
)

type m19operationCode int

const (
	m19OpNone m19operationCode = iota
	m19OpAdd
	m19OpMultiply
	m19OpInput
	m19OpOutput

	m19OpHCF m19operationCode = 99
)

type m19operation struct {
	baseInteger *baseInteger

	repr      string
	numParams int
	mode      []m19opMode
	guessed   bool
}

func (mo m19operation) Address() address { return mo.baseInteger.Address() }
func (mo m19operation) Value() int       { return mo.baseInteger.Value() }
func (mo *m19operation) Set(value int)   { mo.baseInteger.Set(value) }

func (mo m19operation) Name() string   { return mo.repr }
func (mo m19operation) NumParams() int { return mo.numParams }

func (mo *m19operation) Exec() ExecReturnCode {
	read := mo.baseInteger.machine.readAddress
	write := mo.baseInteger.machine.writeAddress
	paramAddresses := mo.getParamAddresses()
	switch m19operationCode(mo.baseInteger.value) {
	case m19OpAdd:
		a := read(paramAddresses[0]).Value()
		b := read(paramAddresses[1]).Value()

		newVal := a + b
		fmt.Printf("ADD %d + %d => %v\n", a, b, paramAddresses[2])

		write(address(paramAddresses[2]), newVal)
	case m19OpMultiply:
		a := read(paramAddresses[0]).Value()
		b := read(paramAddresses[1]).Value()

		newVal := a * b

		fmt.Printf("MUL %d * %d => %v\n", a, b, paramAddresses[2])

		write(address(paramAddresses[2]), newVal)
	}
	return ExecRCInvalidInstruction
}

func (mo *m19operation) getParamAddresses() []address {
	addrs := make([]address, mo.numParams)

	for i := 0; i < mo.numParams; i++ {
		paramAddress := mo.baseInteger.address + address(i+1)
		indirectAddress := address(mo.baseInteger.machine.readAddress(paramAddress).Value())

		switch mo.mode[i] {
		case m19opModeImmediate:
			addrs[i] = paramAddress
		case m19opModePositional:
			addrs[i] = indirectAddress
		default:
			panic("Unsupported mode")
		}
	}
	return addrs
}

func (mo *m19operation) writeToRAM(addr address, value int) {
	mo.baseInteger.Set(value)
}

func (mo m19operation) String() string {
	retString := ""
	if mo.guessed {
		retString += "?>\t"
	}
	retString += fmt.Sprintf("%s", mo.repr)
	for i := 0; i < mo.numParams; i++ {
		paramAddress := mo.baseInteger.address + address(i+1)
		paramInteger := mo.baseInteger.machine.readAddress(paramAddress)

		switch mo.mode[i] {
		case m19opModeImmediate:
			retString = fmt.Sprintf("%s\t%v'", retString, paramInteger)
		case m19opModePositional:
			dereferenced := mo.baseInteger.machine.readAddress(address(paramInteger.Value())).Value()
			retString = fmt.Sprintf("%s\t#%v (%d)", retString, paramInteger, dereferenced)
		default:
			retString = fmt.Sprintf("%s\t??'%v'", retString, paramInteger)
		}
	}
	return retString
}
