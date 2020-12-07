package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	type testDef struct {
		enc      string
		sector   int
		checksum string
		valid    bool
	}
	tests := []testDef{
		{"aaaaa-bbb-z-y-x", 123, "abxyz", true},
		{"a-b-c-d-e-f-g-h", 987, "abcde", true},
		{"not-a-real-room", 404, "oarel", true},
		{"totally-real-room", 200, "decoy", false},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			valid := room{
				enc:      test.enc,
				sector:   test.sector,
				checksum: test.checksum,
			}.validate()
			assert.Equal(t, test.valid, valid)
		})
	}
}

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		room      room
		decrypted string
	}
	tests := []testDef{
		{
			room:      room{enc: "qzmt-zixmtkozy-ivhz", sector: 343},
			decrypted: "very encrypted name",
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			assert.Equal(t, test.decrypted, test.room.decrypt())
		})
	}
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 158835
	//Part 2: 993
}
