package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Day 01")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elveCals := countCalories(file)

	sort.Ints(elveCals)

	fmt.Println("1. Most calories: \t\t", elveCals[len(elveCals)-1:][0])
	fmt.Println("2. Sum calories of top 3:\t", sum(elveCals[len(elveCals)-3:]))
}

func countCalories(data io.Reader) (result []int) {
	scanner := bufio.NewScanner(data)
	calories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			result = append(result, calories)
			calories = 0
			continue
		}
		itemCals, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		calories += itemCals
	}
	return result
}

func sum(input []int) (result int) {
	for _, v := range input {
		result += v
	}
	return
}
