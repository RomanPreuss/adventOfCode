package helper_test

import (
	"testing"

	"github.com/RomanPreuss/adventOfCode2023/helper"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := helper.Sum([]int{12, 38, 15, 77})
	assert.Equal(t, 142, result)
}
