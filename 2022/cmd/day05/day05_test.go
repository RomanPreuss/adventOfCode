package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartOfPacket(t *testing.T) {

	testCases := []struct {
		input    []rune
		expected int
	}{
		{
			input:    []rune("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			expected: 5,
		},
		{
			input:    []rune("nppdvjthqldpwncqszvftbrmjlhg"),
			expected: 6,
		},
		{
			input:    []rune("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			expected: 10,
		},
		{
			input:    []rune("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			expected: 11,
		},
	}

	for id, testCase := range testCases {
		name := fmt.Sprintf("Find marker %v", id)
		t.Run(name, func(t *testing.T) {
			output := FindMarker(testCase.input, 4)
			fmt.Println(output)
			assert.Equal(t, testCase.expected, output)
		})
	}
}

func TestStartOfMessage(t *testing.T) {

	testCases := []struct {
		input    []rune
		expected int
	}{
		{
			input:    []rune("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
			expected: 19,
		},
		{
			input:    []rune("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			expected: 23,
		},
		{
			input:    []rune("nppdvjthqldpwncqszvftbrmjlhg"),
			expected: 23,
		},
		{
			input:    []rune("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			expected: 29,
		},
		{
			input:    []rune("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			expected: 26,
		},
	}

	for id, testCase := range testCases {
		name := fmt.Sprintf("Find marker %v", id)
		t.Run(name, func(t *testing.T) {
			output := FindMarker(testCase.input, 14)
			fmt.Println(output)
			assert.Equal(t, testCase.expected, output)
		})
	}
}
