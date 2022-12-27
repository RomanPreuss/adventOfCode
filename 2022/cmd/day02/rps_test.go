package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EvalGame(t *testing.T) {
	t.Run(`rock beats scissor`, func(t *testing.T) {
		assert.Equal(t, 7, evalGameV1(fmt.Sprintf("%c %c", OPP_SCISSOR, ME_ROCK)))
	})
	t.Run(`paper beats rock`, func(t *testing.T) {
		assert.Equal(t, 8, evalGameV1(fmt.Sprintf("%c %c", OPP_ROCK, ME_PAPER)))
	})
	t.Run(`scissor beats paper`, func(t *testing.T) {
		assert.Equal(t, 9, evalGameV1(fmt.Sprintf("%c %c", OPP_PAPER, ME_SCISSOR)))
	})

	t.Run(`rock draw`, func(t *testing.T) {
		assert.Equal(t, 4, evalGameV1(fmt.Sprintf("%c %c", OPP_ROCK, ME_ROCK)))
	})
	t.Run(`paper draw`, func(t *testing.T) {
		assert.Equal(t, 5, evalGameV1(fmt.Sprintf("%c %c", OPP_PAPER, ME_PAPER)))
	})
	t.Run(`scissor draw`, func(t *testing.T) {
		assert.Equal(t, 6, evalGameV1(fmt.Sprintf("%c %c", OPP_SCISSOR, ME_SCISSOR)))
	})

	t.Run(`rock loose`, func(t *testing.T) {
		assert.Equal(t, 1, evalGameV1(fmt.Sprintf("%c %c", OPP_PAPER, ME_ROCK)))
	})
	t.Run(`paper loose`, func(t *testing.T) {
		assert.Equal(t, 2, evalGameV1(fmt.Sprintf("%c %c", OPP_SCISSOR, ME_PAPER)))
	})
	t.Run(`scissor loose`, func(t *testing.T) {
		assert.Equal(t, 3, evalGameV1(fmt.Sprintf("%c %c", OPP_ROCK, ME_SCISSOR)))
	})
}

func Test_EvalGameV2(t *testing.T) {
	t.Run(`should be draw when rock`, func(t *testing.T) {
		assert.Equal(t, 4, evalGameV2(fmt.Sprintf("%c %c", OPP_ROCK, SHOULD_DRAW)))
	})
	t.Run(`should loose when paper`, func(t *testing.T) {
		assert.Equal(t, 1, evalGameV2(fmt.Sprintf("%c %c", OPP_PAPER, SHOULD_LOOSE)))
	})
	t.Run(`should win when scissor`, func(t *testing.T) {
		assert.Equal(t, 7, evalGameV2(fmt.Sprintf("%c %c", OPP_SCISSOR, SHOULD_WIN)))
	})
}
