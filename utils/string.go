package utils

import "strconv"

func GetInts(input string) []int {
	ints := []int{}
	part := ""

	for _, char := range input + "\n" {
		switch {
		case char == '-' && part == "":
			part = "-"
		case char >= '0' && char <= '9':
			part += string(char)
		default:
			if part != "" {
				newInt, err := strconv.Atoi(part)
				if err == nil {
					ints = append(ints, newInt)
				}
				part = ""
			}
		}
	}
	return ints
}
