package main

import (
	"embed"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed results
var resultsFS embed.FS

func main() {
	benchmarkTime, benchmarkMemory, err := loadBenchmarks()
	if err != nil {
		panic(err)
	}
	table := strings.Builder{}
	table.WriteString("\n## CPU time\n\n")
	table.WriteString(makeTimeTable(benchmarkTime))
	table.WriteString("\n\n")
	table.WriteString("## Heap memory\n\n")
	table.WriteString(makeMemoryTable(benchmarkMemory))
	_ = os.WriteFile("benchmarks.md", []byte(table.String()), 0644)
}

func makeTimeTable(benchmarks benchmarkTimeData) string {
	sb := strings.Builder{}

	years := []int{}
	for year := range benchmarks {
		years = append(years, year)
	}
	sort.Ints(years)
	sb.WriteString(" &nbsp; ")
	yearRuntimes := make([]time.Duration, len(years))
	for i, year := range years {
		sb.WriteString(" | ")
		sb.WriteString(strconv.Itoa(year))
		totalRuntime := time.Duration(0)
		for _, runtime := range benchmarks[year] {
			totalRuntime += runtime
		}
		yearRuntimes[i] = totalRuntime
	}
	sb.WriteString("\n ---: ")
	for range years {
		sb.WriteString(" | ---: ")
	}
	sb.WriteByte('\n')
	for day := 1; day <= 25; day++ {
		sb.WriteString("Day ")
		sb.WriteString(strconv.Itoa(day))
		for i, year := range years {
			sb.WriteString(" | ")
			runtime := benchmarks[year][day]
			propTotal := float64(runtime) / float64(yearRuntimes[i])
			if runtime > 0 {
				strength := ""
				prefix := ""
				if propTotal > 0.2 {
					strength = "**"
					prefix = "🔴 "
				}
				sb.WriteString(strength)
				sb.WriteString(prefix)
				sb.WriteString(formatDuration(runtime))
				sb.WriteString(strength)
			} else {
				sb.WriteByte('-')
			}
		}
		sb.WriteByte('\n')
	}
	sb.WriteString("*Total*")
	for _, year := range years {
		totalRuntime := time.Duration(0)
		for _, runtime := range benchmarks[year] {
			totalRuntime += runtime
		}
		sb.WriteString(" | *")
		sb.WriteString(formatDuration(totalRuntime))
		sb.WriteString("*")

	}

	return sb.String()
}

func makeMemoryTable(benchmarks benchmarkMemoryData) string {
	sb := strings.Builder{}

	years := []int{}
	for year := range benchmarks {
		years = append(years, year)
	}
	sort.Ints(years)
	sb.WriteString(" &nbsp; ")
	yearMemories := make([]int, len(years))
	for i, year := range years {
		sb.WriteString(" | ")
		sb.WriteString(strconv.Itoa(year))
		totalMemory := 0
		for _, memory := range benchmarks[year] {
			totalMemory += memory
		}
		yearMemories[i] = totalMemory
	}
	sb.WriteString("\n ---: ")
	for range years {
		sb.WriteString(" | ---: ")
	}
	sb.WriteByte('\n')
	for day := 1; day <= 25; day++ {
		sb.WriteString("Day ")
		sb.WriteString(strconv.Itoa(day))
		for i, year := range years {
			sb.WriteString(" | ")
			memory := float64(benchmarks[year][day])
			propTotal := memory / float64(yearMemories[i])
			if memory > 0 {
				strength := ""
				prefix := ""
				if propTotal > 0.2 {
					strength = "**"
					prefix = "🔴 "
				}
				sb.WriteString(strength)
				sb.WriteString(prefix)
				sb.WriteString(formatMemory(memory))
				sb.WriteString(strength)
			} else {
				sb.WriteByte('-')
			}
		}
		sb.WriteByte('\n')
	}
	sb.WriteString("*Total*")
	for _, year := range years {
		totalMemory := 0.0
		for _, memory := range benchmarks[year] {
			totalMemory += float64(memory)
		}
		sb.WriteString(" | *")
		sb.WriteString(formatMemory(totalMemory))
		sb.WriteString("*")

	}

	return sb.String()
}

func loadBenchmarks() (benchmarkTimeData, benchmarkMemoryData, error) {
	benchmarkTime := benchmarkTimeData{}
	benchmarkMemory := benchmarkMemoryData{}
	yearDirs, err := resultsFS.ReadDir("results")
	if err != nil {
		panic(err)
	}

	for _, yearDir := range yearDirs {
		year, _ := strconv.Atoi(yearDir.Name())
		benchmarkTime[year] = map[int]time.Duration{}
		benchmarkMemory[year] = map[int]int{}
		dayResults, err := resultsFS.ReadDir("results/" + yearDir.Name())
		if err != nil {
			return nil, nil, err
		}
		for _, dayResult := range dayResults {
			parts := strings.SplitN(dayResult.Name(), "-", 2)
			if len(parts) != 2 {
				continue
			}
			if !strings.HasPrefix(parts[0], "day") {
				continue
			}
			day, err := strconv.Atoi(strings.TrimPrefix(parts[0], "day"))
			if err != nil {
				return nil, nil, err
			}
			path := fmt.Sprintf("results/%s/%s", yearDir.Name(), dayResult.Name())
			result, err := resultsFS.ReadFile(path)
			if err != nil {
				return nil, nil, err
			}
			benchmark, err := strconv.ParseFloat(strings.TrimSpace(string(result)), 64)
			if err != nil {
				continue
			}
			switch parts[1] {
			case "ns":
				benchmarkTime[year][day] = time.Duration(int64(math.Round(benchmark)))
			case "mem-b":
				benchmarkMemory[year][day] = int(benchmark)
			}
		}
	}
	return benchmarkTime, benchmarkMemory, nil
}

func formatDuration(dur time.Duration) string {
	for interval := time.Nanosecond; interval <= time.Second; interval *= 10 {
		if dur >= 100*interval {
			dur = dur.Round(interval)
		}
	}
	return dur.String()
}

func formatMemory(mem float64) string {
	units := []string{"B", "KB", "MB", "GB"}
	for _, unit := range units {
		if mem < 100 {
			return fmt.Sprintf("%.1f %s", math.Round(mem*10)/10, unit)
		}
		if mem < 1000 {
			return fmt.Sprintf("%.0f %s", math.Round(mem), unit)
		}
		mem /= 1000
	}
	unit := units[len(units)-1]
	if mem < 100 {
		return fmt.Sprintf("%f %s", math.Round(mem*10)/10, unit)
	}
	return fmt.Sprintf("%f %s", math.Round(mem), unit)
}

type benchmarkTimeData map[int]map[int]time.Duration
type benchmarkMemoryData map[int]map[int]int
