package day04

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Game struct {
	ID      int
	Score   int
	Input   []int
	Winners []int
}

func Parse(input io.Reader) []Game {
	result := []Game{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		gameRaw := strings.Split(line, ":")
		gameID, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(gameRaw[0], "Card ")))
		if err != nil {
			log.Fatal("error parsing game id", err)
		}

		gameSections := strings.Split(gameRaw[1], "|")
		game := Game{
			ID:      gameID,
			Winners: cardInputToIntArray(gameSections[0]),
			Input:   cardInputToIntArray(gameSections[1]),
		}

		result = append(result, game)
	}

	return result
}

func cardInputToIntArray(gameNumbers string) []int {
	result := []int{}
	numbersRaw := strings.Split(strings.TrimSpace(gameNumbers), " ")
	for _, numberRaw := range numbersRaw {
		numberRaw = strings.TrimSpace(numberRaw)
		if numberRaw == "" {
			continue
		}
		numberRaw, err := strconv.Atoi(numberRaw)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, numberRaw)
	}
	return result
}

func Evaluate(game *Game) {
	for _, num := range game.Input {
		for _, winNum := range game.Winners {
			if num != winNum {
				continue
			}

			if game.Score == 0 {
				game.Score = 1
			} else {
				game.Score = game.Score * 2
			}
		}
	}
}

func Level1(r io.Reader) int {
	games := Parse(r)
	totalScore := 0
	for _, g := range games {
		Evaluate(&g)
		totalScore += g.Score
	}

	return totalScore
}
