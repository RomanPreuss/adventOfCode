package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cleanup(t *testing.T) {
	t.Run(`should not fully contain - no overlap`, func(t *testing.T) {
		assert.False(t, fullyContains("2-4", "6-8"))
		assert.False(t, fullyContains("6-8", "2-4"))
	})

	t.Run(`should not fully contain - with overlap`, func(t *testing.T) {
		assert.False(t, fullyContains("2-6", "4-8"))
		assert.False(t, fullyContains("4-8", "2-6"))
	})

	t.Run(`should fully contain`, func(t *testing.T) {
		assert.True(t, fullyContains("2-8", "3-7"))
		assert.True(t, fullyContains("3-7", "2-8"))
	})
}
