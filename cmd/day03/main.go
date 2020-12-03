package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

type slope struct {
	right int
	down  int
}

func run() error {
	rawMap, err := ioutil.ReadFile("input")
	if err != nil {
		return errors.Wrap(err, " > reading input")
	}
	terrain := strings.Split(string(rawMap), "\n")

	slopes := []slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	treesPerSlope := evaluateSlopes(terrain, slopes)

	fmt.Printf(" > 🎅 Task1 - Encountered trees:\t\t%v\n", treesPerSlope[1])
	fmt.Printf(" > 🎅 Task2 - Product of encountered trees:\t%v\n", product(treesPerSlope))

	return nil
}

func product(in []int) int {
	res := 1
	for _, x := range in {
		res *= x
	}
	return res
}

func evaluateSlopes(terrain []string, slopes []slope) []int {
	var treesPerSlope []int
	for _, slope := range slopes {
		posX := 0
		numberOfTrees := 0

		for currentY := 0; currentY < len(terrain); currentY += slope.down {
			row := terrain[currentY]
			rowIdx := posX % len(row)
			currentPos := string([]rune(row)[rowIdx])
			if currentPos == "#" {
				numberOfTrees++
			}
			posX += slope.right
		}
		treesPerSlope = append(treesPerSlope, numberOfTrees)
	}
	return treesPerSlope
}
