package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

type passwordPolicy struct {
	min       int
	max       int
	character string
}

type passwordEntry struct {
	policy   passwordPolicy
	password string
}

func run() error {
	rawDatabase, err := load("input")
	if err != nil {
		return errors.Wrap(err, " > open file")
	}

	passwordDatabase, err := parse(rawDatabase)
	if err != nil {
		return errors.Wrap(err, " > open file")
	}

	taskExecutor(passwordDatabase, task1Validation, func(validPasswords []passwordEntry, invalidPasswords []passwordEntry) {
		fmt.Println(" > 🎅 Task1 🎅")
		fmt.Printf("  > Number of valid password: %v\n", len(validPasswords))
		fmt.Printf("  > Number of invalid password: %v\n", len(invalidPasswords))
	})

	taskExecutor(passwordDatabase, task2Validation, func(validPasswords []passwordEntry, invalidPasswords []passwordEntry) {
		fmt.Println(" > 🎅 Task2 🎅")
		fmt.Printf("  > Number of valid password: %v\n", len(validPasswords))
		fmt.Printf("  > Number of invalid password: %v\n", len(invalidPasswords))
	})

	return nil
}

func taskExecutor(input []passwordEntry, validate func(entry passwordEntry) bool, result func(validPasswords []passwordEntry, invalidPasswords []passwordEntry)) {
	var validPasswords []passwordEntry
	var invalidPasswords []passwordEntry
	for _, entry := range input {
		if validate(entry) {
			validPasswords = append(validPasswords, entry)
		} else {
			invalidPasswords = append(invalidPasswords, entry)

		}
	}
	result(validPasswords, invalidPasswords)
}

func task1Validation(entry passwordEntry) bool {
	actualOccurance := strings.Count(entry.password, entry.policy.character)
	if actualOccurance >= entry.policy.min && actualOccurance <= entry.policy.max {
		return true
	}
	return false
}

func task2Validation(entry passwordEntry) bool {
	pos1 := entry.policy.min - 1
	pos2 := entry.policy.max - 1
	if pos1 > len(entry.password) || pos2 > len(entry.password) {
		return false
	}

	runeAtPos1 := string([]rune(entry.password)[pos1])
	runeAtPos2 := string([]rune(entry.password)[pos2])

	if runeAtPos1 == entry.policy.character && runeAtPos2 != entry.policy.character ||
		runeAtPos1 != entry.policy.character && runeAtPos2 == entry.policy.character {
		return true
	}
	return false
}

func transform(input []string, t func(s string) interface{}) []interface{} {
	var converted []interface{}
	for _, entry := range input {
		converted = append(converted, t(entry))
	}
	return converted
}

func parse(input []string) ([]passwordEntry, error) {
	var passwordDatabase []passwordEntry
	for _, entry := range input {

		var min, max int
		var password string
		var policyCharacter byte
		fmt.Sscanf(entry, "%v-%v %c: %v", &min, &max, &policyCharacter, &password)

		passwordDatabase = append(passwordDatabase, passwordEntry{
			policy: passwordPolicy{
				min:       min,
				max:       max,
				character: string(policyCharacter),
			},
			password: password,
		})
	}
	return passwordDatabase, nil
}

func load(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, " > open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var passwords []string
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, " > scanning file")
	}

	return passwords, nil
}
