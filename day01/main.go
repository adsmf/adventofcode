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
	fmt.Println(fileString)
	parts := strings.Split(fileString, "\n")
	fmt.Printf("%+v\n", parts)
	freq := 0
	for _, mod := range parts {
		modInt, _ := strconv.Atoi(mod)
		freq = freq + modInt
		fmt.Printf("%s, %d => %d\n", mod, modInt, freq)
	}
}
