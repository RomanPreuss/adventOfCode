package utils_test

import (
	"testing"

	"github.com/RomanPreuss/adventOfCode2020/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Abs(t *testing.T) {
	assert.Equal(t, 1, utils.Abs(-1))
	assert.Equal(t, 1, utils.Abs(1))
}

func Test_AbsFloat(t *testing.T) {
	assert.Equal(t, float64(1), utils.AbsFloat(float64(-1)))
	assert.Equal(t, float64(1), utils.AbsFloat(float64(1)))
}

func Test_rotations(t *testing.T) {
	newX, newY := utils.Rotate(10, 0, 90)

	assert.Equal(t, float64(0), newX)
	assert.Equal(t, float64(10), newY)

	newX, newY = utils.Rotate(10, 0, -90)

	assert.Equal(t, float64(0), newX)
	assert.Equal(t, float64(-10), newY)

	newX, newY = utils.Rotate(10, 4, 90)

	assert.Equal(t, float64(-4), newX)
	assert.Equal(t, float64(10), newY)

	newX, newY = utils.Rotate(10, 4, -90)

	assert.Equal(t, float64(4), newX)
	assert.Equal(t, float64(-10), newY)
}
