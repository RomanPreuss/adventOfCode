package day02

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Round struct {
	Red   int
	Green int
	Blue  int
}
type Game struct {
	ID     int
	Rounds []Round
	Stats  GameStats
}

type GameStats struct {
	MaxRed   int
	MaxGreen int
	MaxBlue  int
}

// GameBag how many different colored cubes are in the bag
type GameBag struct {
	Red   int
	Green int
	Blue  int
}

func Level1(input io.Reader, bag GameBag) int {
	games := ParseGames(input)
	sumOfPossibleGameIDs := 0
	for _, g := range games {
		if isGamePossible(g, bag) {
			sumOfPossibleGameIDs += g.ID
		}
	}
	return sumOfPossibleGameIDs
}

func Level2(input io.Reader, bag GameBag) int {
	games := ParseGames(input)
	gamePower := 0
	for _, g := range games {
		power := maxInt(g.Stats.MaxRed, 1) *
			maxInt(g.Stats.MaxGreen, 1) *
			maxInt(g.Stats.MaxBlue, 1)
		gamePower += power
	}

	return gamePower
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isGamePossible(game Game, bag GameBag) bool {
	return game.Stats.MaxRed <= bag.Red &&
		game.Stats.MaxGreen <= bag.Green &&
		game.Stats.MaxBlue <= bag.Blue
}

func ParseGames(input io.Reader) []Game {
	scanner := bufio.NewScanner(input)
	games := []Game{}
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, ParseGame(line))
	}
	return games
}

func ParseGame(input string) Game {
	parts := strings.Split(input, ":")
	gameID, err := strconv.Atoi(strings.TrimPrefix(parts[0], "Game "))
	if err != nil {
		log.Fatal(err)
	}

	rawRounds := strings.Split(parts[1], ";")
	game := Game{
		ID:     gameID,
		Rounds: []Round{},
	}
	for _, r := range rawRounds {
		rawCubes := strings.Split(r, ",")
		round := Round{}
		for _, c := range rawCubes {
			if strings.HasSuffix(c, "red") {
				round.Red = toNumber(strings.TrimSpace(strings.TrimSuffix(c, " red")))
				if round.Red > game.Stats.MaxRed {
					game.Stats.MaxRed = round.Red
				}
				continue
			}
			if strings.HasSuffix(c, "green") {
				round.Green = toNumber(strings.TrimSpace(strings.TrimSuffix(c, " green")))
				if round.Green > game.Stats.MaxGreen {
					game.Stats.MaxGreen = round.Green
				}
				continue
			}
			if strings.HasSuffix(c, "blue") {
				round.Blue = toNumber(strings.TrimSpace(strings.TrimSuffix(c, " blue")))
				if round.Blue > game.Stats.MaxBlue {
					game.Stats.MaxBlue = round.Blue
				}
				continue
			}
		}
		game.Rounds = append(game.Rounds, round)
	}

	return game
}

func toNumber(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
