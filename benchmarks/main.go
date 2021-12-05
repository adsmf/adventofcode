package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hako/durafmt"
)

//go:embed results
var resultsFS embed.FS

func main() {
	benchmarks, err := loadBenchmarks()
	if err != nil {
		panic(err)
	}
	table := makeTable(benchmarks)
	ioutil.WriteFile("benchmarks.md", []byte(table), 0644)
}

func makeTable(benchmarks benchmarkData) string {
	sb := strings.Builder{}

	years := []int{}
	for year := range benchmarks {
		years = append(years, year)
	}
	sort.Ints(years)
	sb.WriteString(" &nbsp; ")
	for _, year := range years {
		sb.WriteString(" | ")
		sb.WriteString(strconv.Itoa(year))
	}
	sb.WriteString("\n ---: ")
	for range years {
		sb.WriteString(" | ---: ")
	}
	sb.WriteByte('\n')
	units := durafmt.Units{
		Hour:        durafmt.Unit{"h", "h"},
		Minute:      durafmt.Unit{"m", "m"},
		Second:      durafmt.Unit{"s", "s"},
		Millisecond: durafmt.Unit{"ms", "ms"},
		Microsecond: durafmt.Unit{"µs", "µs"},
	}
	for day := 1; day <= 25; day++ {
		sb.WriteString("Day ")
		sb.WriteString(strconv.Itoa(day))
		for _, year := range years {
			sb.WriteString(" | ")
			runtime := benchmarks[year][day]
			if runtime > 0 {
				dur := durafmt.ParseShort(runtime).LimitFirstN(1).LimitToUnit("milliseconds")
				durString := dur.Format(units)
				if durString == "" {
					sb.WriteString("<1 µs")
				} else {
					sb.WriteString(dur.Format(units))
				}
			} else {
				sb.WriteByte('-')
			}
		}
		sb.WriteByte('\n')
	}
	sb.WriteString("Total")
	for _, year := range years {
		count := 0
		totalRuntime := time.Duration(0)
		for _, runtime := range benchmarks[year] {
			count++
			totalRuntime += runtime
		}
		sb.WriteString(" | ")
		dur := durafmt.Parse(totalRuntime).LimitFirstN(1).LimitToUnit("milliseconds")
		sb.WriteString(dur.Format(units))

	}

	return sb.String()
}

func loadBenchmarks() (benchmarkData, error) {
	benchmarks := benchmarkData{}
	yearDirs, err := resultsFS.ReadDir("results")
	if err != nil {
		panic(err)
	}

	for _, yearDir := range yearDirs {
		year, _ := strconv.Atoi(yearDir.Name())
		benchmarks[year] = map[int]time.Duration{}
		dayResults, err := resultsFS.ReadDir("results/" + yearDir.Name())
		if err != nil {
			return nil, err
		}
		for _, dayResult := range dayResults {
			day, err := strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(dayResult.Name(), "-ns"), "day"))
			if err != nil {
				return nil, err
			}
			path := fmt.Sprintf("results/%s/%s", yearDir.Name(), dayResult.Name())
			result, err := resultsFS.ReadFile(path)
			if err != nil {
				return nil, err
			}
			benchmark, err := strconv.ParseFloat(strings.TrimSpace(string(result)), 64)
			if err == nil {
				benchmarks[year][day] = time.Duration(int64(math.Round(benchmark)))
			}
		}
	}
	return benchmarks, nil
}

type benchmarkData map[int]map[int]time.Duration
