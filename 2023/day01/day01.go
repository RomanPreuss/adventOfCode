package day01

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/RomanPreuss/adventOfCode2023/helper"
)

func ExtractNumbers(input io.Reader) []int {
	scanner := bufio.NewScanner(input)
	result := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		lineNumbers := make([]rune, 0, 2)
		for _, r := range line {
			if !unicode.IsDigit(r) {
				continue
			}

			if len(lineNumbers) == 0 {
				lineNumbers = append(lineNumbers, r)
			} else if len(lineNumbers) == 1 {
				lineNumbers = append(lineNumbers, r)
			} else {
				lineNumbers[1] = r
			}
		}
		// if only one number is provided duplicated it
		if len(lineNumbers) != 2 {
			lineNumbers = append(lineNumbers, lineNumbers[0])
		}
		number, err := strconv.Atoi(string(lineNumbers))
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, number)

	}

	return result
}

func Level2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	result := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		lineNumbers := make([]rune, 0, 2)

		numbers := ParseNumbers(line)

		if len(numbers) == 0 {
			log.Fatal("no number found in ", line)
		}

		for _, r := range numbers {
			if len(lineNumbers) == 0 {
				lineNumbers = append(lineNumbers, r)
			} else if len(lineNumbers) == 1 {
				lineNumbers = append(lineNumbers, r)
			} else {
				lineNumbers[1] = r
			}
		}

		if len(lineNumbers) != 2 {
			lineNumbers = append(lineNumbers, lineNumbers[0])
		}
		number, err := strconv.Atoi(string(lineNumbers))
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, number)
	}

	return helper.Sum(result)
}

func ParseNumbers(input string) []rune {
	result := []rune{}
	window := []rune{}
	for _, r := range input {
		if unicode.IsDigit(r) {
			result = append(result, r)
			// reset window
			window = []rune{}
			continue
		}
		window = append(window, r)
		if len(window) < 3 {
			continue
		}

		number, ok := getNumber(window)
		if ok {
			result = append(result, number)
			// reset window and keep last rune to handle edge cases
			window = []rune{window[len(window)-1]}
		}

	}
	return result
}

func getNumber(window []rune) (rune, bool) {
	if strings.Contains(string(window), "one") {
		return '1', true
	} else if strings.Contains(string(window), "two") {
		return '2', true
	} else if strings.Contains(string(window), "three") {
		return '3', true
	} else if strings.Contains(string(window), "four") {
		return '4', true
	} else if strings.Contains(string(window), "five") {
		return '5', true
	} else if strings.Contains(string(window), "six") {
		return '6', true
	} else if strings.Contains(string(window), "seven") {
		return '7', true
	} else if strings.Contains(string(window), "eight") {
		return '8', true
	} else if strings.Contains(string(window), "nine") {
		return '9', true
	}
	return '0', false
}
