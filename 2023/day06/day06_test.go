package day06_test

import (
	"log"
	"os"
	"testing"

	"github.com/RomanPreuss/adventOfCode2023/day06"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	t.Run("Level 1", func(t *testing.T) {
		input, err := os.ReadFile("lvl1.txt")
		if err != nil {
			log.Fatal(err)
		}

		totalWinOptions := day06.Level1(string(input))
		log.Println("Level 1", totalWinOptions)
		t.Fail()
	})

	t.Run("Level 2", func(t *testing.T) {
		input, err := os.ReadFile("lvl2.txt")
		if err != nil {
			log.Fatal(err)
		}

		totalWinOptions := day06.Level2(string(input))
		log.Println("Level 2", totalWinOptions)
		t.Fail()
	})
}

func TestCalculateOptions(t *testing.T) {

	t.Run("number of wins - 1", func(t *testing.T) {
		numberOfWinOptions := day06.GetWinOptions(7, 9)
		assert.Equal(t, 4, numberOfWinOptions)
	})
	t.Run("number of wins - 2", func(t *testing.T) {
		numberOfWinOptions := day06.GetWinOptions(15, 40)
		assert.Equal(t, 8, numberOfWinOptions)
	})
	t.Run("number of wins - 3", func(t *testing.T) {
		numberOfWinOptions := day06.GetWinOptions(30, 200)
		assert.Equal(t, 9, numberOfWinOptions)
	})

	t.Run("Level 1 - demo", func(t *testing.T) {
		input := `Time:      7  15   30
Distance:  9  40  200`

		totalWinOptions := day06.Level1(input)
		assert.Equal(t, 288, totalWinOptions)
	})

	t.Run("Level 2 - demo", func(t *testing.T) {
		input := `Time:      7  15   30
Distance:  9  40  200`

		totalWinOptions := day06.Level2(input)
		assert.Equal(t, 71503, totalWinOptions)
	})
}
