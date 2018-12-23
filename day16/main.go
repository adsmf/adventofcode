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
	examples := loadInput("input.txt")
	part1(examples)
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

func findBehavesLike(example exampleInput) []operation {
	behavesLike := []operation{}
	debug(
		"Testing example:\nBefore: %v\nInstruction: %+v\nAfter: %v",
		example.input,
		example.instruction,
		example.output)
	for tryOp := opcodeSTART + 1; tryOp < opcodeEND; tryOp++ {
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

func loadInput(filename string) []exampleInput {
	contentBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	parts := strings.SplitN(string(contentBytes), "\n\n\n", 2)
	exampleInputs := parts[0]
	// testProg := parts[1]

	examples := loadExamples(exampleInputs)
	return examples
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
