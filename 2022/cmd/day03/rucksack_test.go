package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
