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

	fmt.Println("1. sum of fully contained cleaning assignement pairs:", lvl1(file))
}

func lvl1(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	sum := 0
	for scanner.Scan() {
		cleaningRange := strings.Split(scanner.Text(), ",")
		fullyContained := fullyContains(cleaningRange[0], cleaningRange[1])
		if fullyContained {
			sum++
		}
	}
	return sum
}

func fullyContains(rangeA, rangeB string) bool {

	a := split(rangeA)
	b := split(rangeB)

	if (a[1] < b[1] && a[0] < b[0]) || (b[1] < a[1] && b[0] < a[0]) {
		return false
	}

	return true
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
