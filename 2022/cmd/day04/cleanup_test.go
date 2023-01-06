package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cleanup(t *testing.T) {
	t.Run(`should not overlap`, func(t *testing.T) {
		overlaps, fullContained := detectOverlaps("2-4", "6-8")
		assert.False(t, overlaps)
		assert.False(t, fullContained)

		overlaps, fullContained = detectOverlaps("6-8", "2-4")
		assert.False(t, overlaps)
		assert.False(t, fullContained)
	})

	t.Run(`should overlap but not fully contain`, func(t *testing.T) {
		overlaps, fullContained := detectOverlaps("2-6", "4-8")
		assert.True(t, overlaps)
		assert.False(t, fullContained)

		overlaps, fullContained = detectOverlaps("4-8", "2-6")
		assert.True(t, overlaps)
		assert.False(t, fullContained)
	})

	t.Run(`should fully contain`, func(t *testing.T) {
		overlaps, fullContained := detectOverlaps("2-4", "2-4")
		assert.True(t, overlaps)
		assert.True(t, fullContained)

		overlaps, fullContained = detectOverlaps("2-8", "3-7")
		assert.True(t, overlaps)
		assert.True(t, fullContained)

		overlaps, fullContained = detectOverlaps("3-7", "2-8")
		assert.True(t, overlaps)
		assert.True(t, fullContained)
	})
}
