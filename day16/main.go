package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	mainLogger = fmt.Printf
	examples, instructions := loadInput("input.txt")
	part1(examples)
	part2(examples, instructions)
}

var mainLogger func(string, ...interface{}) (int, error)
var debugLogger func(string, ...interface{})

func logger(format string, args ...interface{}) {
	if mainLogger != nil {
		mainLogger(format, args...)
	} else if debugLogger != nil {
		debugLogger(format, args...)
	}
}

func debug(format string, args ...interface{}) {
	if debugLogger != nil {
		debugLogger(format, args...)
	}
}

func part1(examples []exampleInput) int {
	threeOrMore := 0
	for exampleID, example := range examples {
		behavesLike := findBehavesLike(example)
		debug("Example %d behaves like %d instruction(s)\t%v\n", exampleID, len(behavesLike), example)
		if len(behavesLike) >= 3 {
			threeOrMore++
		}
	}
	logger("%d examples behave like three or more instructions\n", threeOrMore)
	return threeOrMore
}

func part2(examples []exampleInput, instructions []instruction) codemap {

	remainingOps := []operation{}
	for tryOp := operation(0); tryOp < operationEND; tryOp++ {
		remainingOps = append(remainingOps, tryOp)
	}

	unknownExamples := make(map[opcode][]exampleInput)
	for i := opcode(0); i < opcodeEND; i++ {
		unknownExamples[i] = []exampleInput{}
	}
	for _, example := range examples {
		opcode := example.instruction.op
		unknownExamples[opcode] = append(unknownExamples[opcode], example)
	}

	knownOpcodes := make(codemap)
	for {

		foundNew := false
		if len(knownOpcodes) == 16 {
			break
		}

		for curOpcode, opExamples := range unknownExamples {
			opBehavesAs := []operation{}
			checkedAny := false
			for _, example := range opExamples {
				behaviours := findBehavesLikeFromList(example, remainingOps)
				if !!!checkedAny {
					checkedAny = true
					opBehavesAs = behaviours
				} else {
					opBehavesAs = intersectOps(opBehavesAs, behaviours)
				}
			}
			if len(opBehavesAs) == 0 {
				panic("Couldn't find a behaviour that matches")
			} else if len(opBehavesAs) == 1 {
				foundNew = true
				knownOpcodes[curOpcode] = opBehavesAs[0]
				delete(unknownExamples, curOpcode)
				newRemaining := []operation{}
				for _, op := range remainingOps {
					if op != opBehavesAs[0] {
						newRemaining = append(newRemaining, op)
					}
				}
				remainingOps = newRemaining
			}
		}
		if !!!foundNew {
			logger("Found no new mappings; abort!\n")
			break
		}
	}
	fmt.Printf("Codemap (%d): %+v\n", len(knownOpcodes), knownOpcodes)
	reg := registers{0, 0, 0, 0}
	for _, inst := range instructions {
		opcode := inst.op
		op := knownOpcodes[opcode]
		reg = processOp(op, inst, reg)
	}
	fmt.Printf("Registers at end of prog: %+v\n", reg)
	return knownOpcodes
}

func intersectOps(prev, next []operation) []operation {
	set := make(map[operation]bool)
	for _, op := range prev {
		set[op] = true
	}
	for _, op := range next {
		set[op] = true
	}
	intersect := []operation{}
	for op := range set {
		intersect = append(intersect, op)
	}
	return intersect
}

type behaviours map[opcode][]operation

func findBehavesLike(example exampleInput) []operation {
	behavesLike := []operation{}
	debug(
		"Testing example:\nBefore: %v\nInstruction: %+v\nAfter: %v",
		example.input,
		example.instruction,
		example.output)
	for tryOp := operation(0); tryOp < operationEND; tryOp++ {
		correct := checkBehavesLikeOp(example, tryOp)
		if correct {
			debug("Behaves like %s", tryOp.toString())
			behavesLike = append(behavesLike, tryOp)
		}
	}
	return behavesLike
}
func findBehavesLikeFromList(example exampleInput, operations []operation) []operation {
	behavesLike := []operation{}
	debug(
		"Testing example:\nBefore: %v\nInstruction: %+v\nAfter: %v",
		example.input,
		example.instruction,
		example.output)
	for _, tryOp := range operations {
		correct := checkBehavesLikeOp(example, tryOp)
		if correct {
			debug("Behaves like %s", tryOp.toString())
			behavesLike = append(behavesLike, tryOp)
		}
	}
	return behavesLike
}

func checkBehavesLikeOp(example exampleInput, op operation) bool {
	output := processOp(op, example.instruction, example.input)
	if reflect.DeepEqual(output, example.output) {
		return true
	}
	return false
}

type exampleInput struct {
	input       registers
	output      registers
	instruction instruction
}

func loadInput(filename string) ([]exampleInput, []instruction) {
	contentBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	parts := strings.SplitN(string(contentBytes), "\n\n\n", 2)
	exampleInputs := parts[0]
	testProg := parts[1]

	examples := loadExamples(exampleInputs)
	instructions := loadProgram(testProg)
	return examples, instructions
}

func loadProgram(program string) []instruction {
	lines := strings.Split(program, "\n")

	instructions := []instruction{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		instructionParts := strings.Split(line, " ")
		op, _ := strconv.Atoi(instructionParts[0])
		input1, _ := strconv.Atoi(instructionParts[1])
		input2, _ := strconv.Atoi(instructionParts[2])
		output, _ := strconv.Atoi(instructionParts[3])
		newIns := instruction{
			op:     opcode(op),
			input1: argument(input1),
			input2: argument(input2),
			output: argument(output),
		}
		instructions = append(instructions, newIns)
	}
	return instructions
}

func loadExamples(exampleInputs string) []exampleInput {
	examples := []exampleInput{}

	// re := regexp.MustCompile("Step (.) must be finished before step (.) can begin.")
	exampleRE := regexp.MustCompile("(?s)Before: [[]([0-9, ]+)[]]\n([0-9 ]+)\nAfter:  [[]([0-9, ]+)[]]")
	for _, example := range strings.Split(exampleInputs, "\n\n") {
		matches := exampleRE.FindStringSubmatch(example)
		// fmt.Printf("Found matches:\n\t%s\n\t%v\n", example, matches)
		newExample := exampleInput{}

		inputs := strings.Split(matches[1], ", ")
		for _, input := range inputs {
			inputInt, err := strconv.Atoi(input)
			if err != nil {
				panic(err)
			}
			newExample.input = append(newExample.input, inputInt)
		}

		outputs := strings.Split(matches[3], ", ")
		for _, output := range outputs {
			outputInt, err := strconv.Atoi(output)
			if err != nil {
				panic(err)
			}
			newExample.output = append(newExample.output, outputInt)
		}

		instructionParts := strings.Split(matches[2], " ")
		op, _ := strconv.Atoi(instructionParts[0])
		input1, _ := strconv.Atoi(instructionParts[1])
		input2, _ := strconv.Atoi(instructionParts[2])
		output, _ := strconv.Atoi(instructionParts[3])
		newExample.instruction = instruction{
			op:     opcode(op),
			input1: argument(input1),
			input2: argument(input2),
			output: argument(output),
		}

		examples = append(examples, newExample)
	}
	return examples
}
