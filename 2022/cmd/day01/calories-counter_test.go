package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
This list represents the Calories of the food carried by five Elves:

The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
The second Elf is carrying one food item with 4000 Calories.
The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
The fifth Elf is carrying one food item with 10000 Calories.

In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd
like to know how many Calories are being carried by the Elf carrying the most Calories. In the
example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
*/
func Test_CaloriesCounter(t *testing.T) {

	input := `
	1000
	2000
	3000

	4000

	5000
	6000

	7000
	8000
	9000

	10000
	`

	t.Run(`should get correct amount of elves`, func(t *testing.T) {
		calories := countCalories(strings.NewReader(input))
		assert.Len(t, calories, 5)
	})

	t.Run(`should sum calories per elve`, func(t *testing.T) {
		calories := countCalories(strings.NewReader(input))

		assert.Equal(t, 6000, calories[0])
		assert.Equal(t, 4000, calories[1])
		assert.Equal(t, 11000, calories[2])
		assert.Equal(t, 24000, calories[3])
		assert.Equal(t, 10000, calories[4])
	})

}

func Test_Sum(t *testing.T) {
	input := []int{1, 2, 3}

	assert.Equal(t, 6, sum(input))
}
