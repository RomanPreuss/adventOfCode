package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shouldFindAdjacents(t *testing.T) {
	t.Run("no adjacent", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`...
										.L.
										...`)

		// when
		adjacents := findOccupiedAdjacents(1, 1, w, h, true, floor)

		// then
		assert.Equal(t, 0, adjacents)
	})

	t.Run("all adjacent for small example", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`###
										#L#
										###`)

		// when
		adjacents := findOccupiedAdjacents(1, 1, w, h, true, floor)

		// then
		assert.Equal(t, 8, adjacents)
	})

	t.Run("only diagonal adjacent", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`#.#
										.L.
										#.#`)

		// when
		adjacents := findOccupiedAdjacents(1, 1, w, h, true, floor)

		// then
		assert.Equal(t, 4, adjacents)
	})

	t.Run("all far adjacent", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`.......#.
										...#.....
										.#.......
										.........
										..#L....#
										....#....
										.........
										#........
										...#.....`)

		// when
		adjacents := findOccupiedAdjacents(3, 4, w, h, true, floor)

		// then
		assert.Equal(t, 8, adjacents)
	})

	t.Run("only close adjacent", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`.......#.
										...#.....
										.#.......
										.........
										..#L....#
										....#....
										.........
										#........
										...#.....`)

		// when
		adjacents := findOccupiedAdjacents(3, 4, w, h, false, floor)

		// then
		assert.Equal(t, 2, adjacents)
	})
}

func Test_to2D(t *testing.T) {
	t.Run("should return correct coordinates", func(t *testing.T) {
		w := 3

		x, y := to2Dw(0, w)
		assert.Equal(t, 0, x)
		assert.Equal(t, 0, y)

		x, y = to2Dw(3, w)
		assert.Equal(t, 0, x)
		assert.Equal(t, 1, y)

		x, y = to2Dw(5, w)
		assert.Equal(t, 2, x)
		assert.Equal(t, 1, y)
	})
}

func Test_to1D(t *testing.T) {
	t.Run("should return correct index", func(t *testing.T) {
		w := 3
		h := 3

		assert.Equal(t, 0, to1Dw(0, 0, w, h))
		assert.Equal(t, 2, to1Dw(2, 0, w, h))
		assert.Equal(t, 5, to1Dw(2, 1, w, h))
		assert.Equal(t, -1, to1Dw(2, 3, w, h))
		assert.Equal(t, -1, to1Dw(8, 7, w, h))
	})

	t.Run("should return correct dimensions for floor plan", func(t *testing.T) {
		_, w, h := prepareFloorPlan(`...
								...
								...
								..A`)

		assert.Equal(t, 3, w)
		assert.Equal(t, 4, h)
	})

	t.Run("should return correct rune for small floor plan", func(t *testing.T) {
		floor, w, h := prepareFloorPlan(`...
									..A`)

		idx := to1Dw(2, 1, w, h)
		assert.Equal(t, 'A', rune(floor[idx]))
	})

	t.Run("should return correct rune for larger floor plan", func(t *testing.T) {
		floor, w, h := prepareFloorPlan(`.......#.
									...#.....
									.#.......
									.........
									..#L....#
									....#....
									.........
									#........
									...#....A`)

		idx := to1Dw(8, 8, w, h)
		assert.Equal(t, 'A', rune(floor[idx]))
	})

	t.Run("should return correct string", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`.......#.
										...#.....
										.#.......
										.........
										..#L....#
										....#....
										.........
										#........
										...#.....`)
		var result []rune

		// when
		for y := 0; y < w; y++ {
			for x := 0; x < h; x++ {
				result = append(result, rune(floor[to1Dw(x, y, w, h)]))
			}
		}

		// then
		assert.Equal(t, floor, string(result))
	})

}

func Test_rule1(t *testing.T) {
	t.Run("should become occupied when no adjacents", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`...
										.L.
										...`)

		// when
		result := shouldBecomeOccupiedw(1, 1, w, h, true, floor)

		// then
		assert.True(t, result)
	})

	t.Run("should become occupied when no far adjacents", func(t *testing.T) {
		// given
		// Note: the # is not a diagonal adjacent
		floor, w, h := prepareFloorPlan(`..#.
										....
										.L..
										....`)

		// when
		result := shouldBecomeOccupiedw(1, 2, w, h, true, floor)

		// then
		assert.True(t, result)
	})

	t.Run("should not become occupied when adjacents present", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`..#
										.L.
										...`)

		// when
		result := shouldBecomeOccupiedw(1, 1, w, h, true, floor)

		// then
		assert.False(t, result)
	})

	t.Run("should not become occupied when far adjacents present", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`....
										..L.
										....
										#...`)

		// when
		result := shouldBecomeOccupiedw(2, 1, w, h, true, floor)

		// then
		assert.False(t, result)
	})

	t.Run("should become occupied when far adjacents present but flag set to no", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`....
										..L.
										....
										#...`)

		// when
		result := shouldBecomeOccupiedw(2, 1, w, h, false, floor)

		// then
		assert.True(t, result)
	})
}

func Test_rule2(t *testing.T) {
	t.Run("should become free when too crowded", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`###
										###
										###`)

		// when
		result := shouldBecomeFreew(1, 1, w, h, true, floor)

		// then
		assert.True(t, result)
	})

	t.Run("should become free when too far crowded", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`#.#.#
										.....
										#.#.#
										.....
										#.#.#`)

		// when
		result := shouldBecomeFreew(2, 2, w, h, true, floor)

		// then
		assert.True(t, result)
	})

	t.Run("should become not free when far seats not considered", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`#.#.#
										.....
										#.#.#
										.....
										#.#.#`)

		// when
		result := shouldBecomeFreew(2, 2, w, h, false, floor)

		// then
		assert.False(t, result)
	})

	t.Run("should stay occupied when not to crowded", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`.LL
										##.
										##L`)

		// when
		result := shouldBecomeFreew(1, 1, w, h, true, floor)

		// then
		assert.False(t, result)
	})

	t.Run("should stay occupied when not to far crowded", func(t *testing.T) {
		// given
		floor, w, h := prepareFloorPlan(`#.L..
										.....
										L.#.L
										.....
										#.#..`)

		// when
		result := shouldBecomeFreew(2, 2, w, h, true, floor)

		// then
		assert.False(t, result)
	})
}
