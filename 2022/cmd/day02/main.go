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

	if me == ME_ROCK && opp == OPP_SCISSOR {
		return SCORE_WIN + shapeScore(me)
	}
	if me == ME_PAPER && opp == OPP_ROCK {
		return SCORE_WIN + shapeScore(me)
	}
	if me == ME_SCISSOR && opp == OPP_PAPER {
		return SCORE_WIN + shapeScore(me)
	}

	if me == ME_ROCK && opp == OPP_ROCK {
		return SCORE_DRAW + shapeScore(me)
	}
	if me == ME_PAPER && opp == OPP_PAPER {
		return SCORE_DRAW + shapeScore(me)
	}
	if me == ME_SCISSOR && opp == OPP_SCISSOR {
		return SCORE_DRAW + shapeScore(me)
	}

	if me == ME_ROCK && opp == OPP_PAPER {
		return SCORE_LOOSE + shapeScore(me)
	}
	if me == ME_PAPER && opp == OPP_SCISSOR {
		return SCORE_LOOSE + shapeScore(me)
	}
	if me == ME_SCISSOR && opp == OPP_ROCK {
		return SCORE_LOOSE + shapeScore(me)
	}

	return -1
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
