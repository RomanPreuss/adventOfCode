package partitioner

import (
	"math"

	"github.com/pkg/errors"
)

func Partition(operation string, lower, upper int) (int, int, error) {
	diff := float64(upper-lower) / 2.0
	switch operation {
	case "F", "L":
		newUpper := lower + int(diff)
		return lower, newUpper, nil
	case "B", "R":
		newLower := lower + int(math.Round(diff))
		return newLower, upper, nil
	}
	return 0, 0, errors.Errorf("Unknown operation %v\n", operation)
}

func FindSeat(seatCode string) (int, int, error) {
	runes := []rune(seatCode)
	if len(runes) != 10 {
		return 0, 0, errors.New("invalid seatCode length")
	}
	rows := runes[0:7]
	columns := runes[7:10]

	lowRow := 0
	highRow := 127
	selectedRow := 0
	for _, row := range rows {
		lowRow, highRow, _ = Partition(string(row), lowRow, highRow)
		if row == 'F' {
			selectedRow = lowRow
		} else {
			selectedRow = highRow
		}
	}

	lowColumn := 0
	highColumn := 7
	selectedColumn := 0
	for _, column := range columns {
		lowColumn, highColumn, _ = Partition(string(column), lowColumn, highColumn)
		if column == 'L' {
			selectedColumn = lowColumn
		} else {
			selectedColumn = highColumn
		}
	}

	return selectedRow, selectedColumn, nil
}
