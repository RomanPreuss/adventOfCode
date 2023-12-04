package day04

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Card struct {
	ID          int
	Score       int
	CardMatches []int
	Input       []int
	Winners     []int
}

func Parse(input io.Reader) []Card {
	result := []Card{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		gameRaw := strings.Split(line, ":")
		gameID, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(gameRaw[0], "Card ")))
		if err != nil {
			log.Fatal("error parsing game id", err)
		}

		gameSections := strings.Split(gameRaw[1], "|")
		game := Card{
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

func Evaluate(game *Card) {
	winCounter := 1
	for _, num := range game.Input {
		for _, winNum := range game.Winners {
			if num != winNum {
				continue
			}

			game.CardMatches = append(game.CardMatches, game.ID+winCounter)

			if game.Score == 0 {
				game.Score = 1
			} else {
				game.Score = game.Score * 2
			}
			winCounter++
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

func Level2(r io.Reader) int {
	cards := Parse(r)
	// maps card id to number of matches
	matchCount := map[int]int{}
	totalMatches := 0

	for _, card := range cards {
		Evaluate(&card)

		// also count original card
		matchCount[card.ID] = matchCount[card.ID] + 1
		for _, matchID := range card.CardMatches {
			// card match is all counts of the current card + all previous matches
			matchCount[matchID] = matchCount[card.ID] + matchCount[matchID]
		}
	}

	for _, v := range matchCount {
		totalMatches += v
	}

	return totalMatches
}
