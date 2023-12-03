package day03

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"unicode"

	"github.com/RomanPreuss/adventOfCode2023/helper"
)

// 1. iterating through all characters and put them into potentialMaterialNumber
// 2. while doing that always check if there is an adjacent symbol, if yes, set isMaterialNumber to true
// 3. if next character is no number skip it

func Level1(r io.Reader) int {
	field, width, height := Parse(r)
	fmt.Println("Level 1 parameters: width", width, " height: ", height)
	numbers := FindMaterialNumbers(field, width, height)

	return helper.Sum(numbers)
}

func Level2(r io.Reader) int {
	field, width, height := Parse(r)
	fmt.Println("Level 2 parameters: width", width, " height: ", height)
	numbers := FindGearRatio(field, width, height)

	return helper.Sum(numbers)
}

func Parse(r io.Reader) ([]rune, int, int) {
	scanner := bufio.NewScanner(r)
	field := []rune{}
	width := 0
	height := 0
	for scanner.Scan() {
		line := scanner.Text()
		height++
		width = len(line)
		field = append(field, []rune(line)...)
	}

	return field, width, height
}

type position struct {
	x int
	y int
}

func FindGearRatio(field []rune, width, height int) []int {
	gearRatios := map[position][]int{}

	materialNumbers := []int{}

	potentialNumber := []rune{}
	isMaterialNumber := false
	isGearPart := false
	gearPosition := position{}

	for i := 0; i < len(field); i++ {
		x := i % width
		y := i / width

		if x == 0 {
			// reset
			potentialNumber = []rune{}
			isMaterialNumber = false
		}

		val := field[i]
		if unicode.IsDigit(val) {
			potentialNumber = append(potentialNumber, val)

			// search for adjacent symbols
			for x2 := x - 1; x2 <= x+1; x2++ {
				for y2 := y - 1; y2 <= y+1; y2++ {
					adjacentIndex := convert2Dto1D(x2, y2, width, height)
					if x2 < 0 ||
						x2 >= width ||
						adjacentIndex < 0 ||
						adjacentIndex > len(field)-1 ||
						adjacentIndex == i {
						continue
					}
					// fmt.Printf("check neighbor[%v, %v]: (%v, %v) => %v (%v)\n", x, y, x2, y2, string(field[convert2Dto1D(x2, y2, width, height)]), string(field[i]))
					adjacentVal := field[adjacentIndex]
					if adjacentVal != '.' && !unicode.IsDigit(adjacentVal) {
						isMaterialNumber = true
						if adjacentVal == '*' {
							gearPosition = position{
								x: x2,
								y: y2,
							}
							isGearPart = true
						}
					}
				}
			}
		}

		// if not digit and end of line as well as end of list
		if !unicode.IsDigit(val) || i == len(field)-1 || x == width-1 {
			if len(potentialNumber) != 0 && isMaterialNumber {

				materialNumber, err := strconv.Atoi(string(potentialNumber))
				if err != nil {
					log.Fatal(err)
				}

				if isGearPart {
					gearRatios[gearPosition] = append(gearRatios[gearPosition], materialNumber)
				}
				materialNumbers = append(materialNumbers, materialNumber)
				isMaterialNumber = false
				isGearPart = false
			}
			potentialNumber = []rune{}
			continue
		}

	}

	result := []int{}
	for _, vals := range gearRatios {
		if len(vals) != 2 {
			continue
		}
		result = append(result, vals[0]*vals[1])
	}

	return result
}

func FindMaterialNumbers(field []rune, width, height int) []int {
	materialNumbers := []int{}

	potentialNumber := []rune{}
	isMaterialNumber := false

	for i := 0; i < len(field); i++ {
		x := i % width
		y := i / width

		if x == 0 {
			// reset
			potentialNumber = []rune{}
			isMaterialNumber = false
		}

		val := field[i]
		if unicode.IsDigit(val) {
			potentialNumber = append(potentialNumber, val)

			// search for adjacent symbols
			for x2 := x - 1; x2 <= x+1; x2++ {
				for y2 := y - 1; y2 <= y+1; y2++ {
					adjacentIndex := convert2Dto1D(x2, y2, width, height)
					if x2 < 0 ||
						x2 >= width ||
						adjacentIndex < 0 ||
						adjacentIndex > len(field)-1 ||
						adjacentIndex == i {
						continue
					}
					// fmt.Printf("check neighbor[%v, %v]: (%v, %v) => %v (%v)\n", x, y, x2, y2, string(field[convert2Dto1D(x2, y2, width, height)]), string(field[i]))
					adjacentVal := field[adjacentIndex]
					if adjacentVal != '.' && !unicode.IsDigit(adjacentVal) {
						isMaterialNumber = true
					}
				}
			}
		}

		// if not digit and end of line as well as end of list
		if !unicode.IsDigit(val) || i == len(field)-1 || x == width-1 {
			if len(potentialNumber) != 0 && isMaterialNumber {
				materialNumber, err := strconv.Atoi(string(potentialNumber))
				if err != nil {
					log.Fatal(err)
				}
				materialNumbers = append(materialNumbers, materialNumber)
				isMaterialNumber = false
			}
			potentialNumber = []rune{}
			continue
		}

	}

	return materialNumbers
}

func convert2Dto1D(x, y, width, height int) int {
	return x + (y * width)
}
