package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	adapters := make([]int, len(lines))
	for i := 0; i < len(adapters); i++ {
		adapters[i], _ = strconv.Atoi(lines[i])
	}
	sort.Ints(adapters)

	task1(adapters)
	task2(adapters)
}

func task1(adapters []int) {
	currentJolt := 0
	// storing counts for of 1 and 3 jolts difference. (starting with count 1 for 3 jolts difference)
	diffs := []int{0, 1}
	for _, adapter := range adapters {
		diff := adapter - currentJolt
		if diff == 1 {
			diffs[0]++
		} else if diff == 3 {
			diffs[1]++
		}
		currentJolt = adapter
	}

	task1Result := diffs[0] * diffs[1]
	fmt.Println("Task 1: difference", task1Result)
}

func task2(adapters []int) {
	lastItem := adapters[len(adapters)-1] + 3
	items := append([]int{0}, adapters...)
	items = append(items, lastItem)

	combos := map[int]int{0: 1}
	for _, item := range items[1:] {
		// With some help from https://github.com/mnml/aoc/blob/master/2020/10/2.go
		// The number of combinations is the sum of all possible previous combinations
		// For example
		// (0: 1) -> start
		// (1: 1) -> only one combination possible
		// (2: 2) -> can be reached from 0 and 1 (1+1 = 2)
		// (3: 4) -> can be reached from 0, 1 and 2 (1+1+2 = 4)
		// ....
		combos[item] = combos[item-1] + combos[item-2] + combos[item-3]
	}

	fmt.Println("Task 2: All combinations", combos[lastItem])
}
