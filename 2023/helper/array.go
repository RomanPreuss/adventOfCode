package helper

import (
	"log"
	"strconv"
	"strings"
)

func DigitStringToArray(input string) []int {
	if input == "" {
		return []int{}
	}
	result := []int{}
	digits := strings.Split(input, " ")
	for _, d := range digits {
		if d == "" {
			continue
		}

		v, err := strconv.Atoi(d)
		if err != nil {
			log.Fatalf("Error converting to int array '%v': %v \n", d, err)
		}
		result = append(result, v)
	}
	return result
}
