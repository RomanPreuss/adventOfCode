package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const SCORE_LOOSE = 0
const SCORE_DRAW = 3
const SCORE_WIN = 6
const SCORE_ROCK = 1
const SCORE_PAPER = 2
const SCORE_SCISSOR = 3

const ME_ROCK = 'X'
const ME_PAPER = 'Y'
const ME_SCISSOR = 'Z'

const OPP_ROCK = 'A'
const OPP_PAPER = 'B'
const OPP_SCISSOR = 'C'

func main() {
	fmt.Println("Day 02")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gameScores := play(file)
	totalScore := totalScore(gameScores)

	fmt.Println("1. Total score:", totalScore)
}

func totalScore(instructions []int) int {
	totalScore := 0
	for _, score := range instructions {
		totalScore += score
	}
	return totalScore
}

func play(reader io.Reader) (gameScores []int) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		game := scanner.Text()
		gameScores = append(gameScores, evalGame(game))
	}
	fmt.Println("games", len(gameScores))
	return
}

func evalGame(game string) int {
	game = strings.ReplaceAll(game, " ", "")

	moves := make([]rune, 2)
	for i, r := range game {
		moves[i] = r
	}

	opp := moves[0]
	me := moves[1]

	// rock bets scissor
	ruleScore := map[rune]map[rune]int{
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
