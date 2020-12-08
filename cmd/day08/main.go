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

	instructions := strings.Split(strings.TrimSpace(string(input)), "\n")
	accumulator, _ := exec(instructions, -1)
	fmt.Printf("Task 1: Detected loop at accumulator: %v\n", accumulator)
	fmt.Println("Task 2: Try to fix program")
	accumulator = 0
	for i := range instructions {
		acc, terminated := exec(instructions, i)
		if terminated {
			accumulator = acc
			break
		}
	}
	fmt.Printf("Task 2: Program successfully terminated with accumulator value: %v\n", accumulator)
}

func exec(input []string, fixAt int) (int, bool) {
	instructions := make([]string, len(input))
	copy(instructions, input)
	if fixAt != -1 {
		instruction := instructions[fixAt]
		if strings.Contains(instruction, "nop") {
			instructions[fixAt] = strings.Replace(instruction, "nop", "jmp", -1)
		} else if strings.Contains(instruction, "jmp") {
			instructions[fixAt] = strings.Replace(instruction, "jmp", "nop", 1)
		}
	}

	stacktrace := map[int]int{}
	currInstructionIdx := 0
	accumulator := 0
	for ok := true; ok; ok = stacktrace[currInstructionIdx] < 1 {
		if currInstructionIdx >= len(instructions) {
			fmt.Println("Programm successfully terminated")
			return accumulator, true
		}

		var instruction, argumentStr string
		fmt.Sscanf(instructions[currInstructionIdx], "%s %s", &instruction, &argumentStr)
		argument, _ := strconv.ParseInt(argumentStr, 10, 64)
		// fmt.Printf("execute %v with %v\n", instruction, argument)

		stacktrace[currInstructionIdx]++
		switch instruction {
		case "acc":
			currInstructionIdx++
			accumulator += int(argument)
		case "nop":
			currInstructionIdx++
		case "jmp":
			currInstructionIdx += int(argument)
		}
	}
	return accumulator, false
}
