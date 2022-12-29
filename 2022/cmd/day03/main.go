package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	fmt.Println("Day 03")

	file, err := os.Open("lvl1.txt")
	defer func() { _ = file.Close() }()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("1. sum of priorities:", lvl1(file))

	file, err = os.Open("lvl2.txt")
	defer func() { _ = file.Close() }()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("2. sum of priorities:", lvl2(file))
}

func lvl1(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	sum := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		item, ok := findDuplicate(rucksack)
		if !ok {
			log.Fatal("Error finding duplicate in", rucksack)
		}
		sum += getPriority(item)
	}
	return sum
}

func lvl2(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	i := 1
	sum := 0
	group := make([]string, 3)
	for scanner.Scan() {
		group[i%3] = scanner.Text()
		if i%3 == 0 {
			r := findDuplicateV2(group...)
			sum += getPriority(r)
		}
		i++
	}
	return sum
}

func findDuplicateV2(inputs ...string) rune {
	uniqueItems := []map[rune]bool{{}, {}, {}}
	var uniqueItem rune
	for i, rucksack := range inputs {
		for _, r := range rucksack {
			uniqueItems[i][r] = true
		}
	}
	set := map[rune]int{}
	for _, rucksack := range uniqueItems {
		for r, _ := range rucksack {
			set[r]++
			if set[r] == 3 {
				return r
			}
		}
	}

	return uniqueItem
}

func findDuplicate(input string) (rune, bool) {
	numChars := utf8.RuneCountInString(input)
	half := numChars / 2
	uniqueItems := map[rune]bool{}
	for _, r := range input[:half] {
		uniqueItems[r] = true
	}
	for _, r := range input[half:] {
		if uniqueItems[r] {
			return r, true
		}
	}
	return 0x0, false
}

func getPriority(in rune) int {
	asciiCode := int(in)
	if asciiCode <= int('Z') {
		return asciiCode - int('A') + 27
	}
	return int(in) - int('a') + 1
}
