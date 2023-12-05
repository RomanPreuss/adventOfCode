package day05_test

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/RomanPreuss/adventOfCode2023/day05"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	file, err := os.Open("lvl1.txt")
	defer func() { _ = file.Close() }()
	if err != nil {
		log.Fatal(err)
	}

	lowestLocation := day05.Level1(file)
	log.Println("Level 1 - lowest location", lowestLocation)
	t.Fail()
}

func TestSeedPlan(t *testing.T) {

	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	t.Run("parse", func(t *testing.T) {
		seeds, mappings := day05.Parse(strings.NewReader(input))

		assert.Equal(t, []int{79, 14, 55, 13}, seeds)

		assert.Len(t, mappings.SeedToSoil, 2)
		assert.Equal(t, []int{50, 98, 2}, mappings.SeedToSoil[0])

		assert.Len(t, mappings.SoilToFertilizer, 3)
		assert.Len(t, mappings.FertilizerToWater, 4)
		assert.Len(t, mappings.WaterToLight, 2)
		assert.Len(t, mappings.LightToTemperature, 3)
		assert.Len(t, mappings.TemperatureToHumidity, 2)
		assert.Len(t, mappings.HumidityToLocation, 2)
	})

	t.Run("traverse mapping", func(t *testing.T) {
		t.Run("not in range too low", func(t *testing.T) {
			seed := 79
			_, ok := day05.TraverseMapping(seed, []int{50, 98, 2})
			assert.False(t, ok)
		})

		t.Run("not in range too high", func(t *testing.T) {
			seed := 100
			_, ok := day05.TraverseMapping(seed, []int{50, 98, 2})
			assert.False(t, ok)
		})

		t.Run("in range", func(t *testing.T) {
			seed := 79
			target, ok := day05.TraverseMapping(seed, []int{52, 50, 48})
			assert.True(t, ok)
			assert.Equal(t, 81, target)
		})

	})
	t.Run("traverse to next", func(t *testing.T) {
		seed := 74
		next := day05.TraverseToNext(seed, [][]int{
			{45, 77, 23},
			{81, 45, 19},
			{68, 64, 13},
		})
		assert.Equal(t, 78, next)
	})

	t.Run("traverse", func(t *testing.T) {
		_, almanac := day05.Parse(strings.NewReader(input))
		seed := 79
		location := day05.Traverse(seed, almanac)
		assert.Equal(t, 82, location)
	})

	t.Run("Level 2 - demo", func(t *testing.T) {
		lowestLocation := day05.Level1(strings.NewReader(input))
		assert.Equal(t, 35, lowestLocation)
	})
}
