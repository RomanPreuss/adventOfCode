package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindGroups(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	assert.Equal(t, 70, lvl2(strings.NewReader(input)))
}

func Test_FindDuplicatesV2(t *testing.T) {
	assert.Equal(t, string('A'), string(findDuplicateV2("Abc", "ABC", "swA")))

	assert.Equal(t, string('r'), string(findDuplicateV2([]string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}...)))
}

func Test_FindDuplicates(t *testing.T) {
	testCases := []struct {
		input    string
		expected rune
	}{
		{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			'p',
		},
		{
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			'L',
		},
		{
			"PmmdzqPrVvPwwTWBwg",
			'P',
		},
		{
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			'v',
		},
		{
			"ttgJtRGJQctTZtZT",
			't',
		},
		{
			"CrZsJsPPZsGzwwsLwLmpwMDw",
			's',
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("should find '%c'", tc.expected)
		t.Run(name, func(t *testing.T) {
			found, ok := findDuplicate(tc.input)
			assert.True(t, ok)
			assert.Equal(t, tc.expected, found)
		})
	}
}

func Test_detectPrio(t *testing.T) {
	assert.Equal(t, 1, getPriority('a'))
	assert.Equal(t, 26, getPriority('z'))
	assert.Equal(t, 27, getPriority('A'))
	assert.Equal(t, 52, getPriority('Z'))
	assert.Equal(t, 16, getPriority('p'))
	assert.Equal(t, 38, getPriority('L'))
	assert.Equal(t, 22, getPriority('v'))
	assert.Equal(t, 20, getPriority('t'))
	assert.Equal(t, 19, getPriority('s'))
}
