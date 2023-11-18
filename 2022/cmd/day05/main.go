package main

import (
	"fmt"
	"log"
	"os"
)

func FindMarker(input []rune, windowSize int) int {
	window := []rune{}

	for i, r := range input {

		// check if rune is already in window
		for y := 0; y < len(window); y++ {
			if window[y] == r {
				// if rune is already present remove it from left
				window = window[y+1:]
			}
		}

		if i > windowSize && len(window) > windowSize {
			// remove beginning
			window = window[1:]
		}
		window = append(window, r)

		if len(window) == windowSize {
			return i + 1
		}
	}

	return -1
}

func main() {
	file, err := os.Open("lvl1.txt")
	defer func() { _ = file.Close() }()
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := os.ReadFile("lvl1.txt")
	if err != nil {
		log.Fatalln(err)
	}

	result := FindMarker([]rune(string(bytes)), 4)
	fmt.Println("Result Level1 day05: ", result)

	result = FindMarker([]rune(string(bytes)), 14)
	fmt.Println("Result Level2 day05: ", result)

}
