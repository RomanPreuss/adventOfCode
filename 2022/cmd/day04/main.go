package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 04")

	file, err := os.Open("lvl1.txt")
	defer func() { _ = file.Close() }()
	if err != nil {
		log.Fatal(err)
	}

	numOverlaps, numFullyContained := eval(file)
	fmt.Println("1. num of fully contained cleaning assignement pairs:", numFullyContained)
	fmt.Println("2. sum of overlapping cleaning assignement pairs:", numOverlaps)
}

func eval(reader io.Reader) (numOverlaps int, numFullyContained int) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		cleaningRange := strings.Split(scanner.Text(), ",")
		overlaps, fullyContained := detectOverlaps(cleaningRange[0], cleaningRange[1])
		if overlaps {
			numOverlaps++
		}
		if fullyContained {
			numFullyContained++
		}
	}
	return numOverlaps, numFullyContained
}

func detectOverlaps(rangeA, rangeB string) (overlaps bool, fullyContained bool) {

	a := split(rangeA)
	b := split(rangeB)

	// fully contained
	if (a[0] >= b[0] && a[1] <= b[1]) || (b[0] >= a[0] && b[1] <= a[1]) {
		overlaps = true
		fullyContained = true
		return
	}

	// no overlap
	if (a[1] < b[1] && a[1] < b[0]) || (b[1] < a[1] && b[1] < a[0]) {
		overlaps = false
		fullyContained = false
		return
	}

	overlaps = true
	fullyContained = false
	return
}

func split(inRange string) []int {
	r := strings.Split(inRange, "-")
	result := make([]int, 2)

	r0, err := strconv.Atoi(r[0])
	if err != nil {
		panic(fmt.Errorf("error parsing range: %w", err))
	}
	result[0] = r0
	r1, err := strconv.Atoi(r[1])
	if err != nil {
		panic(fmt.Errorf("error parsing range: %w", err))
	}
	result[1] = r1
	return result
}
