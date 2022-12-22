package partitioner_test

import (
	"testing"

	"github.com/RomanPreuss/adventOfCode2020/pkg/partitioner"
	"github.com/stretchr/testify/assert"
)

func Test_Partition(t *testing.T) {
	t.Run("Partion for 'F'", func(t *testing.T) {
		t.Run("should return lower half", func(t *testing.T) {
			lower, upper, err := partitioner.Partition("F", 0, 127)

			assert.Nil(t, err)
			assert.Equal(t, 0, lower)
			assert.Equal(t, 63, upper)
		})
		t.Run("should return lower half when in the middle", func(t *testing.T) {
			lower, upper, err := partitioner.Partition("F", 32, 63)

			assert.Nil(t, err)
			assert.Equal(t, 32, lower)
			assert.Equal(t, 47, upper)
		})
	})

	t.Run("Partion for 'B'", func(t *testing.T) {
		t.Run("should return upper half", func(t *testing.T) {
			lower, upper, err := partitioner.Partition("B", 0, 127)

			assert.Nil(t, err)
			assert.Equal(t, 64, lower)
			assert.Equal(t, 127, upper)
		})
		t.Run("should return upper half when in the middle", func(t *testing.T) {
			lower, upper, err := partitioner.Partition("B", 32, 47)

			assert.Nil(t, err)
			assert.Equal(t, 40, lower)
			assert.Equal(t, 47, upper)
		})
	})

	t.Run("Partion for 'R'", func(t *testing.T) {
		t.Run("should return upper half", func(t *testing.T) {
			lower, upper, err := partitioner.Partition("R", 0, 7)

			assert.Nil(t, err)
			assert.Equal(t, 4, lower)
			assert.Equal(t, 7, upper)
		})
		t.Run("should return lower equal to equal when range is only 1", func(t *testing.T) {
			lower, upper, err := partitioner.Partition("R", 4, 5)

			assert.Nil(t, err)
			assert.Equal(t, 5, lower)
			assert.Equal(t, 5, upper)
		})
	})

	t.Run("Partion for 'L'", func(t *testing.T) {
		t.Run("should return upper equal to lower when range is only 1", func(t *testing.T) {
			lower, upper, err := partitioner.Partition("L", 4, 5)

			assert.Nil(t, err)
			assert.Equal(t, 4, lower)
			assert.Equal(t, 4, upper)
		})
		t.Run("should return lower half", func(t *testing.T) {
			lower, upper, err := partitioner.Partition("L", 4, 7)

			assert.Nil(t, err)
			assert.Equal(t, 4, lower)
			assert.Equal(t, 5, upper)
		})
	})
}

func Test_BinarySpacePartition(t *testing.T) {
	t.Run("should return correct column and row", func(t *testing.T) {
		testCases := []struct {
			seatCode         string
			expectedRow      int
			expectedColumn   int
			expectedIdSeatID int
		}{
			{
				seatCode:         "FBFBBFFRLR", // 0101100101
				expectedRow:      44,
				expectedColumn:   5,
				expectedIdSeatID: 357,
			},
			{
				seatCode:         "BFFFBBFRRR",
				expectedRow:      70,
				expectedColumn:   7,
				expectedIdSeatID: 567,
			},
			{
				seatCode:         "FFFBBBFRRR",
				expectedRow:      14,
				expectedColumn:   7,
				expectedIdSeatID: 119,
			},
			{
				seatCode:         "BBFFBBFRLL",
				expectedRow:      102,
				expectedColumn:   4,
				expectedIdSeatID: 820,
			},
		}

		for _, testCase := range testCases {
			row, column, err := partitioner.FindSeat(testCase.seatCode)

			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedRow, row)
			assert.Equal(t, testCase.expectedColumn, column)
			seatID := (row * 8) + column
			assert.Equal(t, testCase.expectedIdSeatID, seatID)
		}
	})
}
