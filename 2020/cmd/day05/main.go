package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/RomanPreuss/adventOfCode2020/pkg/partitioner"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("======= 🎄 AdventOfCode 2020 - Day 5 🎄 =======")

	input, err := ioutil.ReadFile("input")
	if err != nil {
		return errors.Wrap(err, " > error reading input")
	}

	var seatIDs []int
	for _, seatCode := range strings.Split(string(input), "\n") {
		row, column, _ := partitioner.FindSeat(seatCode)
		seatID := (row * 8) + column
		seatIDs = append(seatIDs, seatID)
	}
	sort.Ints(seatIDs)

	fmt.Println(" > 🎅 Task 1")
	fmt.Println("  Highest seat id:", seatIDs[len(seatIDs)-1])

	mySeatID := -1
	for i := 0; i < len(seatIDs); i++ {
		if i == len(seatIDs)-1 {
			continue
		}
		potentialSeatID := seatIDs[i] + 1
		nextID := seatIDs[i+1]
		if potentialSeatID != nextID {
			mySeatID = potentialSeatID
		}
	}
	fmt.Println(" > 🎅 Task 2")
	fmt.Println("  My seat ID is", mySeatID)

	return nil
}
