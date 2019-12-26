package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ReadInputLines returnes all lines within a file, after trimming start and end whitespace
func ReadInputLines(filename string) []string {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		return []string{}
	}
	fileString := string(fileBytes)
	lines := strings.Split(strings.TrimSpace(fileString), "\n")

	return lines
}
