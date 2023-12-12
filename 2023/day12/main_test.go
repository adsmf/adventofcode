package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMain() {
	main()
	//Output:
	//Part 1: 7361
	//Part 2: 83317216247365
}

func TestParser(t *testing.T) {
	line := "????.??#?#?#.??? 1,2,7,1"
	conditions, report := parseLine(line)
	assert.Equal(t, []condition{
		condUnknown, condUnknown, condUnknown, condUnknown,
		condOperational,
		condUnknown, condUnknown,
		condDamaged,
		condUnknown,
		condDamaged,
		condUnknown,
		condDamaged,
		condOperational,
		condUnknown, condUnknown, condUnknown,
	}, conditions)
	assert.Equal(t, []int{1, 2, 7, 1}, report)
}

func TestAnswers(t *testing.T) {
}

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		input string
		ways  int
	}
	tests := []testDef{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 4},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 1},
		{"????.######..#####. 1,6,5", 4},
		{"?###???????? 3,2,1", 10},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			contitions, report := parseLine(test.input)
			ways := countArrangements(contitions, report)
			assert.Equal(t, test.ways, ways)
		})
	}
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
