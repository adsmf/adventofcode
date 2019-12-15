package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// M19RegisterOutput is the register used to store the last value output
	M19RegisterOutput registerID = iota + registerCommonEnd

	// M19RelativeBase is the offset used for relative mode operations
	M19RelativeBase
)

// InputCallback is the function to be used by the processor to retrieve new input
type InputCallback func() (int, bool)

// OutputCallback is the function called by the processor when a value is output
type OutputCallback func(int)

// M19 sets the behaviour of the intcode machine to AoC 2019 rules
func M19(inputCallback InputCallback, outputCallback OutputCallback) MachineOption {
	return func(m *Machine) {
		m.model = &m19{
			machine:        m,
			inputCallback:  inputCallback,
			outputCallback: outputCallback,
		}
	}
}

type m19 struct {
	machine *Machine

	relativeBase   int
	inputCallback  InputCallback
	outputCallback OutputCallback
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
			value:   m.machine.readAddress(addr).Value(),
		},
		guessed: false,
	}
	opCode := m19operationCode(op.Value() % 100)
	opMode := op.Value() / 100
	switch opCode {
	case m19OpAdd:
		op.repr = "ADD"
		op.numParams = 3
	case m19OpMultiply:
		op.repr = "MUL"
		op.numParams = 3
	case m19OpInput:
		op.repr = "INP"
		op.numParams = 1
	case m19OpOutput:
		op.repr = "OUT"
		op.numParams = 1
	case m19OpJumpTrue:
		op.repr = "JNZ"
		op.numParams = 2
	case m19OpJumpFalse:
		op.repr = "JEZ"
		op.numParams = 2
	case m19OpLess:
		op.repr = "CLT"
		op.numParams = 3
	case m19OpEqual:
		op.repr = "CEQ"
		op.numParams = 3
	case m19OpAdjustRelativeBase:
		op.repr = "ARB"
		op.numParams = 1
	case m19OpHCF:
		op.repr = "HCF"
		op.numParams = 0
	default:
		op.repr = fmt.Sprintf("UNK-%d", opCode)
		return nil
	}
	op.mode = make([]m19opMode, op.NumParams())
	for i := 0; i < op.numParams; i++ {
		op.mode[i] = m19opMode(opMode % 10)
		opMode /= 10
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
	m19OpJumpTrue
	m19OpJumpFalse
	m19OpLess
	m19OpEqual
	m19OpAdjustRelativeBase

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
	mo.baseInteger.machine.registers[RegisterInstructionPointer] += 1 + mo.numParams
	op := m19operationCode(mo.baseInteger.value % 100)
	switch op {
	case m19OpAdd:
		a := read(paramAddresses[0]).Value()
		b := read(paramAddresses[1]).Value()
		newVal := a + b

		write(address(paramAddresses[2]), newVal)
	case m19OpMultiply:
		a := read(paramAddresses[0]).Value()
		b := read(paramAddresses[1]).Value()
		newVal := a * b

		write(address(paramAddresses[2]), newVal)
	case m19OpOutput:
		newVal := read(paramAddresses[0]).Value()
		mo.baseInteger.machine.setRegister(
			M19RegisterOutput,
			newVal,
		)
		if mo.baseInteger.machine.model.(*m19).outputCallback != nil {
			mo.baseInteger.machine.model.(*m19).outputCallback(newVal)
		}
		return ExecRCInterrupt
	case m19OpInput:
		in, halt := mo.baseInteger.machine.model.(*m19).inputCallback()
		if halt {
			return ExecRCHCF
		}
		write(address(paramAddresses[0]), in)
	case m19OpJumpTrue:
		test := read(paramAddresses[0]).Value()
		jmp := read(paramAddresses[1]).Value()
		if test != 0 {
			mo.baseInteger.machine.setRegister(RegisterInstructionPointer, jmp)
		}
	case m19OpJumpFalse:
		test := read(paramAddresses[0]).Value()
		jmp := read(paramAddresses[1]).Value()
		if test == 0 {
			mo.baseInteger.machine.setRegister(RegisterInstructionPointer, jmp)
		}
	case m19OpLess:
		a := read(paramAddresses[0]).Value()
		b := read(paramAddresses[1]).Value()
		if a < b {
			write(address(paramAddresses[2]), 1)
		} else {
			write(address(paramAddresses[2]), 0)
		}
	case m19OpEqual:
		a := read(paramAddresses[0]).Value()
		b := read(paramAddresses[1]).Value()
		if a == b {
			write(address(paramAddresses[2]), 1)
		} else {
			write(address(paramAddresses[2]), 0)
		}
	case m19OpAdjustRelativeBase:
		value := read(paramAddresses[0]).Value()
		mo.baseInteger.machine.registers[M19RelativeBase] += value
	default:
		return ExecRCInvalidInstruction
	}
	return ExecRCNone
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
		case m19opModeRelative:
			offset := mo.baseInteger.machine.Register(M19RelativeBase)
			addrs[i] = indirectAddress + address(offset)
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
			retString = fmt.Sprintf("%s\t'%v'", retString, paramInteger)
		case m19opModePositional:
			dereferenced := mo.baseInteger.machine.readAddress(address(paramInteger.Value())).Value()
			retString = fmt.Sprintf("%s\t#%v (%d)", retString, paramInteger, dereferenced)
		case m19opModeRelative:
			offset := mo.baseInteger.machine.Register(M19RelativeBase)
			dereferenced := mo.baseInteger.machine.readAddress(address(paramInteger.Value() + offset)).Value()
			retString = fmt.Sprintf("%s\t#%v+%v (%d)", retString, paramInteger, offset, dereferenced)
		default:
			retString = fmt.Sprintf("%s\t??'%v'", retString, paramInteger)
		}
	}
	return retString
}
