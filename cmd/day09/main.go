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
	for i := 0; i < len(numbers)-preambleSize-1; i++ {
		window := numbers[i : i+preambleSize]
		target, _ := strconv.Atoi(numbers[i+preambleSize])

		fmt.Println("window", i)
		fmt.Println(window, target)

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
			fmt.Println("not in window:", target)

			contains = false
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
						contains = true
						result := smallest + largest
						fmt.Println("Found range", smallest, largest, result)
						return
					}
				}
			}

			break
		}
	}
}
