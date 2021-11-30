package utils

import (
	"fmt"
	"io/ioutil"
)

// ReadInputLines returnes all lines within a file, after trimming start and end whitespace
func ReadInputLines(filename string) []string {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		return []string{}
	}
	return GetLines(string(fileBytes))
}
