package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "> error reading input: %s\n", err)
		os.Exit(1)
	}

	sum := 0
	matchingSum := 0
	// Alternatievely use `re := regexp.MustCompile(`(?m)^\n`)`
	// and `groups := re.Split(string(input), -1)` to split the groups.
	for _, group := range strings.Split(string(input), "\n\n") {
		answerSet := map[rune]int{}
		groupEntries := strings.Fields(group) // alternatively strings.Split(group, "\n")
		numOfPeopleInGroup := len(groupEntries)
		for _, formEntry := range groupEntries {
			for _, answer := range formEntry {
				answerSet[answer]++
				if answerSet[answer] == numOfPeopleInGroup {
					matchingSum++
				}
			}
		}
		sum += len(answerSet)
	}

	fmt.Printf(" > Task 1: sum of all answers %v\n", sum)
	fmt.Printf(" > Task 2: sum of matching answers %v\n", matchingSum)
}
