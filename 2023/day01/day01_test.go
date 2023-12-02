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

		numbers := day01.ExtractNumbers(file)
		sum := helper.Sum(numbers)

		log.Println("Day01")
		// 55447 too high
		log.Println("Level2: ", sum)

		t.Fail()
	})
}

func TestRealNumbers(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	t.Run("extract numbers", func(t *testing.T) {
		result := day01.ExtractNumbersV2(strings.NewReader(input))

		assert.Equal(t, []int{12, 38, 15, 77}, result)
	})
}

func TestWrittenNumbers(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
seven
kjrqmzv9mmtxhgvsevenhvq7
eighthree`

	result := day01.ExtractNumbersV2(strings.NewReader(input))

	assert.Equal(t, []int{29, 83, 13, 24, 42, 14, 76, 77, 97, 83}, result)
}

func TestFindNumberEdgeCases(t *testing.T) {
	t.Run("multiple matches", func(t *testing.T) {
		res := day01.FindNum("onetwoone")
		assert.Equal(t, []rune{'1', '2', '1'}, res)
	})
	t.Run("merged numbers", func(t *testing.T) {
		res := day01.FindNum("eighthree")
		assert.Equal(t, []rune{'8', '3'}, res)
	})
	t.Run("written numbers and digits", func(t *testing.T) {
		res := day01.FindNum("one2eighthree")
		assert.Equal(t, []rune{'1', '2', '8', '3'}, res)
	})
	t.Run("written numbers and digits 2", func(t *testing.T) {
		res := day01.FindNum("kjrqmzv9mmtxhgvsevenhvq7")
		assert.Equal(t, []rune{'9', '7', '7'}, res)
	})
}
