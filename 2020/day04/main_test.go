package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2Examples(t *testing.T) {
	type testDef struct {
		passport passportData
		good     bool
	}
	tests := []testDef{
		{
			passport: map[string]string{
				"eyr": "1972",
				"cid": "100",
				"hcl": "#18171d",
				"ecl": "amb",
				"hgt": "170",
				"pid": "186cm",
				"iyr": "2018",
				"byr": "1926",
			},
		},
		{
			passport: map[string]string{
				"iyr": "2019",
				"hcl": "#602927",
				"eyr": "1967",
				"hgt": "170cm",
				"ecl": "grn",
				"pid": "012533040",
				"byr": "1946",
			},
		},
		{
			passport: map[string]string{
				"hcl": "dab227",
				"iyr": "2012",
				"ecl": "brn",
				"hgt": "182cm",
				"pid": "021572410",
				"eyr": "2020",
				"byr": "1992",
				"cid": "277",
			},
		},
		{
			passport: map[string]string{
				"hgt": "59cm",
				"ecl": "zzz",
				"eyr": "2038",
				"hcl": "74454a",
				"iyr": "2023",
				"pid": "3556412378",
				"byr": "2007",
			},
		},
	}
	for id, test := range tests {
		t.Run(fmt.Sprintf("Part1-Test%d", id+1), func(t *testing.T) {
			t.Logf("Test def:\n %v", test)
			// Assertions here
			result := test.passport.validatePart2()
			assert.Equal(t, test.good, result)
		})
	}
}

func TestAnswers(t *testing.T) {
}

func ExampleMain() {
	main()
	//Output:
	//Part 1: 192
	//Part 2: 101
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
		part2()
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2()
	}
}
