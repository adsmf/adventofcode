package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testOps(t *testing.T, op operation, tests []exampleInput) {

	for idx, test := range tests {
		t.Run(fmt.Sprint(op.toString(), "/", idx), func(t *testing.T) {
			debugLogger = t.Logf
			t.Log("Input:", test.input)
			t.Log("Instruction:", test.instruction)
			correct := checkBehavesLikeOp(test, op)
			assert.True(t, correct)
		})
	}
}

func TestAddr(t *testing.T) {
	tests := []exampleInput{
		exampleInput{
			instruction: instruction{
				input1: 0,
				input2: 0,
				output: 0,
			},
			input:  registers{1, 2, 3, 4},
			output: registers{2, 2, 3, 4},
		},
		exampleInput{
			instruction: instruction{
				input1: 0,
				input2: 1,
				output: 0,
			},
			input:  registers{4, 3, 2, 1},
			output: registers{7, 3, 2, 1},
		},
		exampleInput{
			instruction: instruction{
				input1: 0,
				input2: 1,
				output: 2,
			},
			input:  registers{4, 3, 2, 1},
			output: registers{4, 3, 7, 1},
		},
	}
	testOps(t, addr, tests)
}

func TestAddi(t *testing.T) {
	tests := []exampleInput{
		exampleInput{
			instruction: instruction{
				input1: 0,
				input2: 0,
				output: 0,
			},
			input:  registers{1, 2, 3, 4},
			output: registers{1, 2, 3, 4},
		},
		exampleInput{
			instruction: instruction{
				input1: 0,
				input2: 1,
				output: 0,
			},
			input:  registers{4, 3, 2, 1},
			output: registers{5, 3, 2, 1},
		},
		exampleInput{
			instruction: instruction{
				input1: 0,
				input2: 1,
				output: 2,
			},
			input:  registers{4, 3, 2, 1},
			output: registers{4, 3, 5, 1},
		},
		exampleInput{
			instruction: instruction{
				input1: 1,
				input2: 1,
				output: 2,
			},
			input:  registers{3, 2, 1, 1},
			output: registers{3, 2, 3, 1},
		},
	}
	testOps(t, addi, tests)
}

func TestMulr(t *testing.T) {
	tests := []exampleInput{
		exampleInput{
			instruction: instruction{
				input1: 1,
				input2: 2,
				output: 0,
			},
			input:  registers{1, 2, 3, 4},
			output: registers{6, 2, 3, 4},
		},
		exampleInput{
			instruction: instruction{
				input1: 0,
				input2: 1,
				output: 0,
			},
			input:  registers{4, 3, 2, 1},
			output: registers{12, 3, 2, 1},
		},
		exampleInput{
			instruction: instruction{
				input1: 2,
				input2: 1,
				output: 2,
			},
			input:  registers{3, 2, 1, 1},
			output: registers{3, 2, 2, 1},
		},
	}
	testOps(t, mulr, tests)
}

func TestBanr(t *testing.T) {
	tests := []exampleInput{
		exampleInput{
			instruction: instruction{
				input1: 1,
				input2: 2,
				output: 0,
			},
			input:  registers{1, 2, 3, 4},
			output: registers{2, 2, 3, 4},
		},
		exampleInput{
			instruction: instruction{
				input1: 0,
				input2: 1,
				output: 0,
			},
			input:  registers{4, 3, 2, 1},
			output: registers{0, 3, 2, 1},
		},
	}
	testOps(t, banr, tests)
}

func TestBorr(t *testing.T) {
	tests := []exampleInput{
		exampleInput{
			instruction: instruction{
				input1: 1,
				input2: 2,
				output: 0,
			},
			input:  registers{1, 2, 3, 4},
			output: registers{3, 2, 3, 4},
		},
		exampleInput{
			instruction: instruction{
				input1: 0,
				input2: 1,
				output: 0,
			},
			input:  registers{4, 3, 2, 1},
			output: registers{7, 3, 2, 1},
		},
	}
	testOps(t, borr, tests)
}

func TestSetr(t *testing.T) {
	tests := []exampleInput{
		exampleInput{
			instruction: instruction{
				input1: 1,
				input2: 99,
				output: 0,
			},
			input:  registers{1, 2, 3, 4},
			output: registers{2, 2, 3, 4},
		},
		exampleInput{
			instruction: instruction{
				input1: 2,
				input2: 99,
				output: 0,
			},
			input:  registers{4, 3, 2, 1},
			output: registers{2, 3, 2, 1},
		},
	}
	testOps(t, setr, tests)
}

func TestSeti(t *testing.T) {
	tests := []exampleInput{
		exampleInput{
			instruction: instruction{
				input1: 1,
				input2: 99,
				output: 0,
			},
			input:  registers{1, 2, 3, 4},
			output: registers{1, 2, 3, 4},
		},
		exampleInput{
			instruction: instruction{
				input1: 2,
				input2: 99,
				output: 0,
			},
			input:  registers{4, 3, 2, 1},
			output: registers{2, 3, 2, 1},
		},
		exampleInput{
			instruction: instruction{
				input1: 2,
				input2: 1,
				output: 2,
			},
			input:  registers{3, 2, 1, 1},
			output: registers{3, 2, 2, 1},
		},
	}
	testOps(t, seti, tests)
}

func TestGtrr(t *testing.T) {
	tests := []exampleInput{
		exampleInput{
			instruction: instruction{
				input1: 1,
				input2: 2,
				output: 0,
			},
			input:  registers{1, 2, 3, 4},
			output: registers{0, 2, 3, 4},
		},
		exampleInput{
			instruction: instruction{
				input1: 1,
				input2: 2,
				output: 0,
			},
			input:  registers{4, 3, 2, 1},
			output: registers{1, 3, 2, 1},
		},
		exampleInput{
			instruction: instruction{
				input1: 1,
				input2: 2,
				output: 0,
			},
			input:  registers{4, 4, 4, 4},
			output: registers{0, 4, 4, 4},
		},
	}
	testOps(t, gtrr, tests)
}
