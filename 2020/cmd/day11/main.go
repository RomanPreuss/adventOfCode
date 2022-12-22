package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const debug bool = false

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	floorPlan, w, h := prepareFloorPlan(string(input))
	printFloor(w, floorPlan)

	floorPlan = applyRules(w, h, floorPlan)
	printFloor(w, floorPlan)

	occupiedSeats := countFreeSeats(floorPlan)
	previousSeats := 0
	for previousSeats != occupiedSeats {
		previousSeats = occupiedSeats
		floorPlan = applyRules(w, h, floorPlan)
		occupiedSeats = countFreeSeats(floorPlan)
		printFloor(w, floorPlan)
	}
	fmt.Println("Number of occupied seats:", occupiedSeats)
}

func applyRules(w, h int, floor string) string {
	newFloorPlan := []rune(floor)
	for i := 0; i < len(floor); i++ {
		if floor[i] == '.' {
			continue
		}
		x, y := to2Dw(i, w)
		if shouldBecomeOccupiedw(x, y, w, h, true, floor) {
			newFloorPlan[i] = '#'
		} else if shouldBecomeFreew(x, y, w, h, true, floor) {
			newFloorPlan[i] = 'L'
		}
	}
	return string(newFloorPlan)
}

func shouldBecomeOccupiedw(x, y, w, h int, includeFar bool, floor string) bool {
	occupiedAdjacents := findOccupiedAdjacents(x, y, w, h, includeFar, floor)
	if floor[to1Dw(x, y, w, h)] == 'L' && occupiedAdjacents == 0 {
		return true
	}
	return false
}

func shouldBecomeFreew(x, y, w, h int, includeFar bool, floor string) bool {
	occupiedAdjacents := findOccupiedAdjacents(x, y, w, h, includeFar, floor)
	if floor[to1Dw(x, y, w, h)] == '#' && occupiedAdjacents >= 5 {
		return true
	}
	return false
}

func to1Dw(x, y, w, h int) int {
	if x < 0 || x >= w || y < 0 || y >= h {
		return -1
	}
	return (y * w) + x
}

func to2Dw(i, w int) (int, int) {
	return (i % w), (i / w)
}

func countFreeSeats(floor string) int {
	sum := 0
	for _, pos := range floor {
		if pos == '#' {
			sum++
		}
	}
	return sum
}

func printFloor(w int, floor string) {
	if !debug {
		return
	}
	for i := 0; i < len(floor); i++ {
		fmt.Printf("%v", string(floor[i]))
		if (i+1)%w == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func prepareFloorPlan(input string) (string, int, int) {
	input = strings.TrimSpace(string(input))
	width := strings.Index(input, "\n")
	height := len(strings.Split(input, "\n"))
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")
	return input, width, height
}

func findOccupiedAdjacents(x, y, w, h int, includeFar bool, floor string) int {
	result := 0

	result += findA(x, y, w, h, floor, includeFar, func(x, y int) (int, int) {
		return x - 1, y - 1 // find all to the top left
	})
	result += findA(x, y, w, h, floor, includeFar, func(x, y int) (int, int) {
		return x, y - 1 // find all to the top
	})
	result += findA(x, y, w, h, floor, includeFar, func(x, y int) (int, int) {
		return x + 1, y - 1 // find all to the top right
	})
	result += findA(x, y, w, h, floor, includeFar, func(x, y int) (int, int) {
		return x + 1, y // find all to the right
	})
	result += findA(x, y, w, h, floor, includeFar, func(x, y int) (int, int) {
		return x + 1, y + 1 // find all to the bottom right
	})
	result += findA(x, y, w, h, floor, includeFar, func(x, y int) (int, int) {
		return x, y + 1 // find all to the bottom
	})
	result += findA(x, y, w, h, floor, includeFar, func(x, y int) (int, int) {
		return x - 1, y + 1 // find all to the bottom left
	})
	result += findA(x, y, w, h, floor, includeFar, func(x, y int) (int, int) {
		return x - 1, y // find all to the left
	})

	return result
}

func findA(x, y, w, h int, floor string, includeFar bool, nextPos func(int, int) (int, int)) int {
	x, y = nextPos(x, y)
	idx := to1Dw(x, y, w, h)
	if idx == -1 {
		return 0
	}
	seat := floor[idx]

	if seat == '#' {
		return 1
	}

	if seat == '.' {
		if includeFar {
			return findA(x, y, w, h, floor, true, nextPos)
		}
	}

	return 0
}
