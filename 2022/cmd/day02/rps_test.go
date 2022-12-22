package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EvalGame(t *testing.T) {
	t.Run(`rock beats scissor`, func(t *testing.T) {
		assert.Equal(t, 7, evalGame(fmt.Sprintf("%c %c", OPP_SCISSOR, ME_ROCK)))
	})
	t.Run(`paper beats rock`, func(t *testing.T) {
		assert.Equal(t, 8, evalGame(fmt.Sprintf("%c %c", OPP_ROCK, ME_PAPER)))
	})
	t.Run(`scissor beats paper`, func(t *testing.T) {
		assert.Equal(t, 9, evalGame(fmt.Sprintf("%c %c", OPP_PAPER, ME_SCISSOR)))
	})

	t.Run(`rock draw`, func(t *testing.T) {
		assert.Equal(t, 4, evalGame(fmt.Sprintf("%c %c", OPP_ROCK, ME_ROCK)))
	})
	t.Run(`paper draw`, func(t *testing.T) {
		assert.Equal(t, 5, evalGame(fmt.Sprintf("%c %c", OPP_PAPER, ME_PAPER)))
	})
	t.Run(`scissor draw`, func(t *testing.T) {
		assert.Equal(t, 6, evalGame(fmt.Sprintf("%c %c", OPP_SCISSOR, ME_SCISSOR)))
	})

	t.Run(`rock loose`, func(t *testing.T) {
		assert.Equal(t, 1, evalGame(fmt.Sprintf("%c %c", OPP_PAPER, ME_ROCK)))
	})
	t.Run(`paper loose`, func(t *testing.T) {
		assert.Equal(t, 2, evalGame(fmt.Sprintf("%c %c", OPP_SCISSOR, ME_PAPER)))
	})
	t.Run(`scissor loose`, func(t *testing.T) {
		assert.Equal(t, 3, evalGame(fmt.Sprintf("%c %c", OPP_ROCK, ME_SCISSOR)))
	})
}
