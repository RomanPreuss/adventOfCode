package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	numbers := strings.Split(strings.TrimSpace(string(input)), "\n")
	preambleSize := 25
	for i := preambleSize; i < len(numbers); i++ {
		window := numbers[i-preambleSize : i]
		target, _ := strconv.Atoi(numbers[i])

		contains := false
		for x := 0; x < len(window)-1; x++ {
			for y := 1; y < len(window); y++ {
				part1, _ := strconv.Atoi(window[x])
				part2, _ := strconv.Atoi(window[y])
				sum := part1 + part2
				if sum == target {
					contains = true
					break
				}
			}
		}

		if !contains {
			fmt.Println("Task1: not in window:", target)

			for x := 0; x < len(numbers)-1; x++ {
				for y := x + 1; y < len(numbers); y++ {
					window = numbers[x:y]
					sum := 0
					smallest := target
					largest := 0
					for _, n := range window {
						num, _ := strconv.Atoi(n)
						if num < smallest {
							smallest = num
						}
						if num > largest {
							largest = num
						}
						sum += num
					}
					if sum == target {
						result := smallest + largest
						fmt.Println("Task2: Found range", smallest, largest, "result:", result)
						return
					}
				}
			}

			break
		}
	}
}
