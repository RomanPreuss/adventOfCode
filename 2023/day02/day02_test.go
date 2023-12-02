package day02_test

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/RomanPreuss/adventOfCode2023/day02"
	"github.com/stretchr/testify/assert"
)

func TestDay02(t *testing.T) {
	t.Run("Level 1", func(t *testing.T) {
		bag := day02.GameBag{
			Red:   12,
			Green: 13,
			Blue:  14,
		}

		file, err := os.Open("lvl1.txt")
		defer func() { _ = file.Close() }()
		if err != nil {
			log.Fatal(err)
		}
		result := day02.Level1(file, bag)
		log.Println("Level 1 result: ", result)
		t.Fail()
	})

	t.Run("Level 2", func(t *testing.T) {
		bag := day02.GameBag{
			Red:   12,
			Green: 13,
			Blue:  14,
		}

		file, err := os.Open("lvl2.txt")
		defer func() { _ = file.Close() }()
		if err != nil {
			log.Fatal(err)
		}
		result := day02.Level2(file, bag)
		log.Println("Level 2 result: ", result)
		t.Fail()
	})
}

func TestGame(t *testing.T) {

	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	t.Run("get max cubes", func(t *testing.T) {
		games := day02.ParseGames(strings.NewReader(input))

		maxCubes := []day02.GameStats{}
		for _, g := range games {
			maxCubes = append(maxCubes, g.Stats)
		}

		assert.Equal(t, []day02.GameStats{
			{
				MaxRed:   4,
				MaxGreen: 2,
				MaxBlue:  6,
			},
			{
				MaxRed:   1,
				MaxGreen: 3,
				MaxBlue:  4,
			},
			{
				MaxRed:   20,
				MaxGreen: 13,
				MaxBlue:  6,
			},
			{
				MaxRed:   14,
				MaxGreen: 3,
				MaxBlue:  15,
			},
			{
				MaxRed:   6,
				MaxGreen: 3,
				MaxBlue:  2,
			},
		}, maxCubes)
	})

	t.Run("Level 1 - Get sum of possible games", func(t *testing.T) {
		bag := day02.GameBag{
			Red:   12,
			Green: 13,
			Blue:  14,
		}
		result := day02.Level1(strings.NewReader(input), bag)
		assert.Equal(t, 8, result)
	})

	t.Run("Level 2 - Get power of games", func(t *testing.T) {
		bag := day02.GameBag{
			Red:   12,
			Green: 13,
			Blue:  14,
		}
		result := day02.Level2(strings.NewReader(input), bag)
		assert.Equal(t, 2286, result)
	})

	t.Run("parse game", func(t *testing.T) {
		input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"

		games := day02.ParseGame(input)

		assert.Equal(t, day02.Game{
			ID: 1,
			Rounds: []day02.Round{
				{
					Red:  4,
					Blue: 3,
				},
				{
					Red:   1,
					Green: 2,
					Blue:  6,
				},
				{
					Green: 2,
				},
			},
			Stats: day02.GameStats{
				MaxRed:   4,
				MaxGreen: 2,
				MaxBlue:  6,
			},
		}, games)
	})

}
