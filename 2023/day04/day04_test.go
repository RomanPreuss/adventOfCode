package day04_test

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/RomanPreuss/adventOfCode2023/day04"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	t.Run("Level 1", func(t *testing.T) {
		file, err := os.Open("lvl1.txt")
		defer func() { _ = file.Close() }()
		if err != nil {
			log.Fatal(err)
		}

		totalScore := day04.Level1(file)
		log.Println("Level1 total score:", totalScore)
		t.Fail()
	})

	t.Run("Level 2", func(t *testing.T) {
		file, err := os.Open("lvl2.txt")
		defer func() { _ = file.Close() }()
		if err != nil {
			log.Fatal(err)
		}

		totalCards := day04.Level2(file)
		log.Println("Level2 total cards:", totalCards)
		t.Fail()
	})
}

func TestDay04(t *testing.T) {

	t.Run("parse", func(t *testing.T) {
		input := `Card 1: 42  1 | 34 42
Card 2: 13 61 | 61 30`
		game := day04.Parse(strings.NewReader(input))

		assert.Equal(t, day04.Card{
			ID:      1,
			Input:   []int{34, 42},
			Winners: []int{42, 1},
		}, game[0])
		assert.Equal(t, day04.Card{
			ID:      2,
			Input:   []int{61, 30},
			Winners: []int{13, 61},
		}, game[1])
	})

	t.Run("evaluate Game", func(t *testing.T) {
		game := day04.Card{
			ID:      1,
			Input:   []int{41, 48, 83, 86, 17},
			Winners: []int{83, 86, 6, 31, 17, 9, 48, 53},
		}
		day04.Evaluate(&game)

		assert.Equal(t, 8, game.Score)
	})

	t.Run("Test Level 1", func(t *testing.T) {
		input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

		totalScore := day04.Level1(strings.NewReader(input))

		assert.Equal(t, 13, totalScore)
	})

	t.Run("Test Level 2", func(t *testing.T) {
		input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

		totalScore := day04.Level2(strings.NewReader(input))

		assert.Equal(t, 30, totalScore)
	})

}
