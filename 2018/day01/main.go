package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	fileString := string(fileBytes)
	parts := strings.Split(fileString, "\n")
	freq := 0
	seenFreqs := make(map[int]bool)
	for {
		for _, mod := range parts {
			modInt, _ := strconv.Atoi(mod)
			freq = freq + modInt
			if seenFreqs[freq] {
				fmt.Printf("Been here before!: %d\n", freq)
				return
			}
			seenFreqs[freq] = true
		}
	}
}
