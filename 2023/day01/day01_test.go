package day01_test

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/RomanPreuss/adventOfCode2023/day01"
	"github.com/RomanPreuss/adventOfCode2023/helper"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	t.Run("level 1", func(t *testing.T) {
		file, err := os.Open("lvl1.txt")
		defer func() { _ = file.Close() }()
		if err != nil {
			log.Fatal(err)
		}

		numbers := day01.ExtractNumbers(file)
		sum := helper.Sum(numbers)

		log.Println("Day01")
		log.Println("Level1: ", sum)

		t.Fail()
	})

	t.Run("level 2", func(t *testing.T) {
		file, err := os.Open("lvl2.txt")
		defer func() { _ = file.Close() }()
		if err != nil {
			log.Fatal(err)
		}

		sum := day01.Level2(file)

		log.Println("Day01")
		log.Println("Level2: ", sum)

		t.Fail()
	})
}

func TestParseNumbers(t *testing.T) {
	t.Run("331s2twonep", func(t *testing.T) {
		res := day01.ParseNumbers("331s2twonep")
		assert.Equal(t, []rune{'3', '3', '1', '2', '2', '1'}, res)
	})
	t.Run("seven", func(t *testing.T) {
		res := day01.ParseNumbers("seven")
		assert.Equal(t, []rune{'7'}, res)
	})
}

func TestLevel2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	result := day01.Level2(strings.NewReader(input))

	assert.Equal(t, 281, result)
}
