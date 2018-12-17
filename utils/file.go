package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadInputLines(filename string) []string {

	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
		return []string{}
	}
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")
	return lines
}
