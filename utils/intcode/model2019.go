package intcode

import (
	"fmt"
	"strconv"
	"strings"
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
		decode := baseInteger{
			machine: m.machine,
			address: address(pos),
			value:   value,
		}
		m.machine.ram[address(pos)] = decode
	}
	m.guessOps()
	return nil
}

func (m *m19) guessOps() {
	// Scrolling up to len(ram) is fine in the initial case, will need changing if re-running later
	for addr := address(0); int(addr) < len(m.machine.ram); addr++ {
		op := m19operation{
			baseInteger: m.machine.ram[addr].(baseInteger),
		}
		opCode := op.Value() % 100
		// opMode := op.Value() / 100
		// fmt.Printf("Decoding operation %d: %d / %d\n", addr, opCode, opMode)
		switch opCode {
		case 1:
			op.repr = "ADD"
			numParams := 3
			op.mode = make([]m19opMode, numParams)
			op.numParams = numParams
		case 2:
			op.repr = "MUL"
			numParams := 3
			op.mode = make([]m19opMode, numParams)
			op.numParams = numParams
		case 3:
			op.repr = "INP"
			numParams := 1
			op.mode = make([]m19opMode, numParams)
			op.numParams = numParams
		case 4:
			op.repr = "OUT"
			numParams := 1
			op.mode = make([]m19opMode, numParams)
			op.numParams = numParams
		case 5:
		case 6:
		case 7:
		case 8:
		case 99:
			op.repr = "HCF"
			op.numParams = 0
		default:
			op.repr = fmt.Sprintf("UNK-%d", opCode)
			return
		}
		m.machine.ram[addr] = op
		m.machine.operations[addr] = op
		addr += address(op.numParams)
	}
}

type m19opMode int

const (
	m19opModePositional m19opMode = iota
	m19opModeImmediate
	m19opModeRelative
)

type m19operation struct {
	baseInteger baseInteger

	repr      string
	numParams int
	mode      []m19opMode
}

func (mo m19operation) Address() address         { return mo.baseInteger.Address() }
func (mo m19operation) IntegerType() integerType { return mo.baseInteger.IntegerType() }
func (mo m19operation) Value() int               { return mo.baseInteger.Value() }

func (mo m19operation) Name() string   { return mo.repr }
func (mo m19operation) NumParams() int { return mo.numParams }

func (mo m19operation) Exec() {}

func (mo m19operation) String() string {
	retString := fmt.Sprintf("%s", mo.repr)
	for i := 0; i < mo.numParams; i++ {
		paramAddress := mo.baseInteger.address + address(i+1)
		paramInteger := mo.baseInteger.machine.readAddress(paramAddress)

		switch mo.mode[i] {
		case m19opModeImmediate:
			retString = fmt.Sprintf("%s '%v'", retString, paramInteger)
		case m19opModePositional:
			dereferenced := mo.baseInteger.machine.readAddress(address(paramInteger.Value())).Value()
			retString = fmt.Sprintf("%s #%v (%d)", retString, paramInteger, dereferenced)
		default:
			retString = fmt.Sprintf("%s ??'%v'", retString, paramInteger)
		}
	}
	return retString
}
