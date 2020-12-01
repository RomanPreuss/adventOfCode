package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("======= 🎄 AdventOfCode 2020 🎄 =======")
	expenses, err := readExpenses("expenses.txt")
	if err != nil {
		return errors.Wrap(err, " > reading expenses")
	}

	task1(expenses)
	task2(expenses)

	return nil
}

func task1(expenses []int) {
	for i1, n1 := range expenses {
		for i2 := i1; i2 < len(expenses); i2++ {
			n2 := expenses[i2]
			if (n1 + n2) == 2020 {
				result := n1 * n2
				fmt.Printf(" > 🎅 Task1 answer: %v ( = %v * %v)\n", result, n1, n2)
			}
		}
	}
}

func task2(expenses []int) {
	for i1, n1 := range expenses {
		for i2 := i1; i2 < len(expenses); i2++ {
			for i3 := i2; i3 < len(expenses); i3++ {
				n2 := expenses[i2]
				n3 := expenses[i3]
				if (n1 + n2 + n3) == 2020 {
					result := n1 * n2 * n3
					fmt.Printf(" > 🎅 Task2 answer: %v ( = %v * %v * %v)\n", result, n1, n2, n3)
				}
			}
		}
	}
}

func readExpenses(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, " > open file")
	}
	defer file.Close()

	var expenses []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expsenseStr := scanner.Text()
		number, err := strconv.Atoi(expsenseStr)
		if err != nil {
			return nil, errors.Wrap(err, " > converting to number")
		}
		expenses = append(expenses, number)
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, " > scanning file")
	}
	return expenses, nil
}
