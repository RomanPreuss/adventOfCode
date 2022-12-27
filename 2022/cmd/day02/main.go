package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type ruleSet map[rune]map[rune]int

const SCORE_LOOSE = 0
const SCORE_DRAW = 3
const SCORE_WIN = 6
const SCORE_ROCK = 1
const SCORE_PAPER = 2
const SCORE_SCISSOR = 3

const ME_ROCK = 'X'
const ME_PAPER = 'Y'
const ME_SCISSOR = 'Z'

const SHOULD_LOOSE = 'X'
const SHOULD_DRAW = 'Y'
const SHOULD_WIN = 'Z'

const OPP_ROCK = 'A'
const OPP_PAPER = 'B'
const OPP_SCISSOR = 'C'

func main() {
	fmt.Println("Day 02")
	fileLvl1, err := os.Open("lvl1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileLvl1.Close()

	gameScoresLvl1 := play(fileLvl1, evalGameV1)
	totalScoreLvl1 := totalScore(gameScoresLvl1)

	fmt.Println("1. Total score:", totalScoreLvl1)

	fileLvl2, err := os.Open("lvl2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileLvl2.Close()

	gameScoresLvl2 := play(fileLvl2, evalGameV2)
	totalScoreLvl2 := totalScore(gameScoresLvl2)

	fmt.Println("2. Total score:", totalScoreLvl2)
}

func totalScore(instructions []int) int {
	totalScore := 0
	for _, score := range instructions {
		totalScore += score
	}
	return totalScore
}

func play(reader io.Reader, evalGame func(string) int) (gameScores []int) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		game := scanner.Text()
		gameScores = append(gameScores, evalGame(game))
	}
	fmt.Println("games", len(gameScores))
	return
}

func evalGameV2(game string) int {
	game = strings.ReplaceAll(game, " ", "")

	moves := make([]rune, 2)
	for i, r := range game {
		moves[i] = r
	}

	opp := moves[0]
	me := moves[1]
	score := 0
	switch me {
	case SHOULD_DRAW:
		score = SCORE_DRAW
	case SHOULD_LOOSE:
		score = SCORE_LOOSE
	case SHOULD_WIN:
		score = SCORE_WIN
	}

	ruleScore := ruleSet{
		SHOULD_LOOSE: {
			OPP_ROCK:    SCORE_SCISSOR,
			OPP_SCISSOR: SCORE_PAPER,
			OPP_PAPER:   SCORE_ROCK,
		},
		SHOULD_DRAW: {
			OPP_ROCK:    SCORE_ROCK,
			OPP_SCISSOR: SCORE_SCISSOR,
			OPP_PAPER:   SCORE_PAPER,
		},
		SHOULD_WIN: {
			OPP_ROCK:    SCORE_PAPER,
			OPP_SCISSOR: SCORE_ROCK,
			OPP_PAPER:   SCORE_SCISSOR,
		},
	}

	return score + ruleScore[me][opp]
}
func evalGameV1(game string) int {
	game = strings.ReplaceAll(game, " ", "")

	moves := make([]rune, 2)
	for i, r := range game {
		moves[i] = r
	}

	opp := moves[0]
	me := moves[1]

	// rock bets scissor
	ruleScore := ruleSet{
		ME_ROCK: {
			OPP_ROCK:    SCORE_DRAW,
			OPP_SCISSOR: SCORE_WIN,
			OPP_PAPER:   SCORE_LOOSE,
		},
		ME_SCISSOR: {
			OPP_ROCK:    SCORE_LOOSE,
			OPP_SCISSOR: SCORE_DRAW,
			OPP_PAPER:   SCORE_WIN,
		},
		ME_PAPER: {
			OPP_ROCK:    SCORE_WIN,
			OPP_SCISSOR: SCORE_LOOSE,
			OPP_PAPER:   SCORE_DRAW,
		},
	}

	return ruleScore[me][opp] + shapeScore(me)
}

func shapeScore(shape rune) int {
	switch shape {
	case ME_ROCK:
		return SCORE_ROCK
	case ME_PAPER:
		return SCORE_PAPER
	case ME_SCISSOR:
		return SCORE_SCISSOR
	}
	return -1
}
