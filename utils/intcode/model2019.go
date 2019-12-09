package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

// Model2019 sets the behaviour of the intcode machine to AoC 2019 rules
func Model2019(input <-chan int, output chan<- int) MachineOption {
	return func(m *Machine) { m.model = &model2019{machine: m} }
}

type model2019 struct {
	machine *Machine

	input  <-chan int
	output chan<- int
}

func (m *model2019) name() string {
	return "Model2019"
}

func (m *model2019) parse(program string) error {
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
	for addr := address(0); int(addr) < len(programIntStrings); addr++ {
		op := model2019operation{
			baseInteger: m.machine.ram[addr].(baseInteger),
		}
		opCode := op.Value() % 100
		// opMode := op.Value() / 100
		// fmt.Printf("Decoding operation %d: %d / %d\n", addr, opCode, opMode)
		switch opCode {
		case 1:
			op.repr = "ADD"
			op.numParams = 3
		case 2:
			op.repr = "MUL"
			op.numParams = 3
		case 3:
			op.repr = "INP"
			op.numParams = 1
		case 4:
			op.repr = "OUT"
			op.numParams = 1
		case 5:
		case 6:
		case 7:
		case 8:
		case 99:
			op.repr = "HCF"
			op.numParams = 0
		default:
			op.repr = fmt.Sprintf("UNK-%d", opCode)
			// return fmt.Errorf("Op code %d, not implemented", op.Value())
		}
		m.machine.ram[addr] = op
		m.machine.operations[addr] = op
		addr += address(op.numParams)
	}
	// return fmt.Errorf("Not implemented")
	return nil
}

type model2019operation struct {
	baseInteger baseInteger

	repr      string
	numParams int
	immediate []bool
}

func (mo model2019operation) Address() address         { return mo.baseInteger.Address() }
func (mo model2019operation) IntegerType() integerType { return mo.baseInteger.IntegerType() }
func (mo model2019operation) Value() int               { return mo.baseInteger.Value() }

func (mo model2019operation) Name() string   { return mo.repr }
func (mo model2019operation) NumParams() int { return mo.numParams }

func (mo model2019operation) Exec() {}

func (mo model2019operation) String() string {
	retString := fmt.Sprintf("%s", mo.repr)
	for i := 0; i < mo.numParams; i++ {
		paramAddress := mo.baseInteger.address + address(i+1)
		paramInteger := mo.baseInteger.machine.readAddress(paramAddress)
		if i < len(mo.immediate) && mo.immediate[i] {
			retString = fmt.Sprintf("%s '%v'", retString, paramInteger)
		} else {
			dereferenced := mo.baseInteger.machine.readAddress(address(paramInteger.Value())).Value()
			retString = fmt.Sprintf("%s #%v (%d)", retString, paramInteger, dereferenced)
		}
	}
	return retString
}
