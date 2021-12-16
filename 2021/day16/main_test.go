package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 852
	//Part 2: 19348959966392
}

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		packet     string
		versionSum int
	}
	tests := []testDef{
		{"D2FE28", 6},
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			// Assertions here
			sum, _, err := parse(test.packet)
			assert.NoError(t, err)
			assert.Equal(t, test.versionSum, sum)
		})
	}
}

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		packet string
		value  int
	}
	tests := []testDef{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			_, value, err := parse(test.packet)
			assert.NoError(t, err)
			assert.Equal(t, test.value, value)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
