package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		item, ok := findDuplicate(rucksack)
		if !ok {
			log.Fatal("Error finding duplicate in", rucksack)
		}
		sum += getPriority(item)
	}

	fmt.Println("1. sum of priorities:", sum)
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
