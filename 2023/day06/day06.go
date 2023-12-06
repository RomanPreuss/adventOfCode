package day06

import (
	"log"
	"strconv"
	"strings"

	"github.com/RomanPreuss/adventOfCode2023/helper"
)

// starting speed 0 mm/ms
// speed up 1 mm/ms

const (
	ACCELERATION = 1 // mm/ms
)

func GetWinOptions(time, prevDistance int) int {
	winOptions := 0

	// first and last are no wins
	for accelerationTime := 1; accelerationTime < time-1; accelerationTime++ {
		distance := accelerationTime * ACCELERATION * (time - accelerationTime)
		if distance > prevDistance {
			winOptions++
		}
	}
	return winOptions
}

func Level1(input string) int {
	totalWinOptions := 1
	lines := strings.Split(input, "\n")
	times := helper.DigitStringToArray(strings.TrimPrefix(lines[0], "Time:"))
	distances := helper.DigitStringToArray(strings.TrimPrefix(lines[1], "Distance:"))

	for i := 0; i < len(times); i++ {
		totalWinOptions *= GetWinOptions(times[i], distances[i])
	}

	return totalWinOptions
}

func Level2(input string) int {
	lines := strings.Split(input, "\n")

	time := fixKerning(strings.TrimPrefix(lines[0], "Time:"))
	distance := fixKerning(strings.TrimPrefix(lines[1], "Distance:"))

	totalWinOptions := GetWinOptions(time, distance)

	return totalWinOptions
}

func fixKerning(input string) int {
	input = strings.ReplaceAll(input, " ", "")
	val, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
