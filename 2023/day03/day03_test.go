package day03_test

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/RomanPreuss/adventOfCode2023/day03"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	file, err := os.Open("lvl1.txt")
	defer func() { _ = file.Close() }()
	if err != nil {
		log.Fatal(err)
	}

	sum := day03.Level1(file)
	// 941896 - too high
	// 941896
	// 547774 - also wrong
	// 549908
	log.Println("Level1: ", sum)
	t.Fail()
}

func TestFindMaterialNumbers(t *testing.T) {

	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	t.Run("parse input field", func(t *testing.T) {
		field, width, height := day03.Parse(strings.NewReader(input))
		assert.Len(t, field, 100)
		assert.Equal(t, 10, width)
		assert.Equal(t, 10, height)
	})

	t.Run("find material numbers", func(t *testing.T) {
		field, width, height := day03.Parse(strings.NewReader(input))
		numbers := day03.FindMaterialNumbers(field, width, height)

		assert.Equal(t, []int{
			467, 35, 633, 617, 592, 755, 664, 598,
		}, numbers)

	})

	t.Run("find material numbers edge case 1", func(t *testing.T) {
		edgeCaseInput := `1.1
.$.
1.1`
		field, width, height := day03.Parse(strings.NewReader(edgeCaseInput))
		numbers := day03.FindMaterialNumbers(field, width, height)

		assert.Equal(t, []int{
			1, 1, 1, 1,
		}, numbers)

	})

	t.Run("no adjacent match larger than width", func(t *testing.T) {
		edgeCaseInput := `..1
$..
...`
		field, width, height := day03.Parse(strings.NewReader(edgeCaseInput))
		numbers := day03.FindMaterialNumbers(field, width, height)

		assert.Equal(t, []int{}, numbers)
	})

	t.Run("no adjacent match smaller than 0", func(t *testing.T) {
		edgeCaseInput := `..$
1..
...`
		field, width, height := day03.Parse(strings.NewReader(edgeCaseInput))
		numbers := day03.FindMaterialNumbers(field, width, height)

		assert.Equal(t, []int{}, numbers)
	})

	t.Run("match material number only end of line", func(t *testing.T) {
		edgeCaseInput := `..$1
2...
....`
		field, width, height := day03.Parse(strings.NewReader(edgeCaseInput))
		numbers := day03.FindMaterialNumbers(field, width, height)

		assert.Equal(t, []int{1}, numbers)
	})

	t.Run("Level 1", func(t *testing.T) {
		result := day03.Level1(strings.NewReader(input))
		assert.Equal(t, 4361, result)
	})

}
