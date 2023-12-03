package day01

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
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

func ExtractNumbersV2(input io.Reader) []int {
	scanner := bufio.NewScanner(input)
	result := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		lineNumbers := make([]rune, 0, 2)

		numbers := FindNum(line)

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

	return result
}

func FindNum(input string) []rune {
	result := []rune{}
	// ignore last char to catch edge-cases like 'eighthree' where the 't' is shared by
	// eighT and Three
	pattern := regexp.MustCompile("([0-9])*(on)*(tw)*(thre)*(fou)*(fiv)*(si)*(seve)*(eigh)*(nin)*(zer)*")
	matches := pattern.FindAllSubmatchIndex([]byte(input), len(input))

	for _, m := range matches {

		for i := 2; i < len(m); i++ {
			if m[i] == -1 || m[i+1] == -1 {
				continue
			}
			nextCharIndex := m[i+1]
			if m[i+1]-m[i] == 1 {
				// it's a digit
				digit := input[m[i]:m[i+1]]
				fmt.Println("digit", digit)
				result = append(result, []rune(digit)[0])
				continue
			}
			if nextCharIndex > len(input)-1 {
				// end of line so no match
				continue
			}

			numberCandidate := input[m[i]:m[i+1]]

			switch numberCandidate {
			case "on":
				if input[nextCharIndex] == 'e' {
					result = append(result, '1')
				}
			case "tw":
				if input[nextCharIndex] == 'o' {
					result = append(result, '2')
				}
			case "thre":
				if input[nextCharIndex] == 'e' {
					result = append(result, '3')
				}
			case "fou":
				if input[nextCharIndex] == 'r' {
					result = append(result, '4')
				}
			case "fiv":
				if input[nextCharIndex] == 'e' {
					result = append(result, '5')
				}
			case "si":
				if input[nextCharIndex] == 'x' {
					result = append(result, '6')
				}
			case "seve":
				if input[nextCharIndex] == 'n' {
					result = append(result, '7')
				}
			case "eigh":
				if input[nextCharIndex] == 't' {
					result = append(result, '8')
				}
			case "nin":
				if input[nextCharIndex] == 'e' {
					result = append(result, '9')
				}
			case "zer":
				if input[nextCharIndex] == '0' {
					result = append(result, '0')
				}
			}
		}

	}
	return result
}

func toNumber(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func ConvertToNumber(input []string) []rune {
	result := make([]rune, 0, len(input))
	for _, num := range input {

		switch num {
		case "one":
			result = append(result, '1')
		case "two":
			result = append(result, '2')
		case "three":
			result = append(result, '3')
		case "four":
			result = append(result, '4')
		case "five":
			result = append(result, '5')
		case "six":
			result = append(result, '6')
		case "seven":
			result = append(result, '7')
		case "eight":
			result = append(result, '8')
		case "nine":
			result = append(result, '9')
		case "zero":
			result = append(result, '0')
		default:
			result = append(result, []rune(num)[0])
		}
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
		fmt.Println("window", string(window))

		number, ok := getNumber(window)
		if ok {
			result = append(result, number)
			// reset window and keep last rune to handle edge cases
			window = []rune{window[len(window)-1]}
			fmt.Println("reset window", string(window))
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
