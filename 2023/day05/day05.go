package day05

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

const (
	SEED_TO_SOIL_MODE            = iota
	SOIL_TO_FERTILIZER_MODE      = iota
	FERTILIZER_TO_WATER_MODE     = iota
	WATER_TO_LIGHT_MODE          = iota
	LIGHT_TO_TEMPERATURE_MODE    = iota
	TEMPERATURE_TO_HUMIDITY_MODE = iota
	HUMIDITY_TO_LOCATION_MODE    = iota
)

type Almanac struct {
	SeedToSoil            [][]int
	SoilToFertilizer      [][]int
	FertilizerToWater     [][]int
	WaterToLight          [][]int
	LightToTemperature    [][]int
	TemperatureToHumidity [][]int
	HumidityToLocation    [][]int
}

func Parse(r io.Reader) ([]int, Almanac) {
	var seeds []int
	almanac := Almanac{}
	parsingMode := -1

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds: ") {
			line = strings.TrimPrefix(line, "seeds: ")
			seeds = digitStringToArray(line)
			continue
		}

		mode, ok := detectParsingMode(line)
		if ok {
			parsingMode = mode
			continue
		}

		switch parsingMode {
		case SEED_TO_SOIL_MODE:
			if mapping := digitStringToArray(line); len(mapping) > 0 {
				almanac.SeedToSoil = append(almanac.SeedToSoil, mapping)
			}
		case SOIL_TO_FERTILIZER_MODE:
			if mapping := digitStringToArray(line); len(mapping) > 0 {
				almanac.SoilToFertilizer = append(almanac.SoilToFertilizer, mapping)
			}
		case FERTILIZER_TO_WATER_MODE:
			if mapping := digitStringToArray(line); len(mapping) > 0 {
				almanac.FertilizerToWater = append(almanac.FertilizerToWater, mapping)
			}
		case WATER_TO_LIGHT_MODE:
			if mapping := digitStringToArray(line); len(mapping) > 0 {
				almanac.WaterToLight = append(almanac.WaterToLight, mapping)
			}
		case LIGHT_TO_TEMPERATURE_MODE:
			if mapping := digitStringToArray(line); len(mapping) > 0 {
				almanac.LightToTemperature = append(almanac.LightToTemperature, mapping)
			}
		case TEMPERATURE_TO_HUMIDITY_MODE:
			if mapping := digitStringToArray(line); len(mapping) > 0 {
				almanac.TemperatureToHumidity = append(almanac.TemperatureToHumidity, mapping)
			}
		case HUMIDITY_TO_LOCATION_MODE:
			if mapping := digitStringToArray(line); len(mapping) > 0 {
				almanac.HumidityToLocation = append(almanac.HumidityToLocation, mapping)
			}
		}

	}

	return seeds, almanac
}

func detectParsingMode(input string) (int, bool) {
	switch input {
	case "seed-to-soil map:":
		return SEED_TO_SOIL_MODE, true
	case "soil-to-fertilizer map:":
		return SOIL_TO_FERTILIZER_MODE, true
	case "fertilizer-to-water map:":
		return FERTILIZER_TO_WATER_MODE, true
	case "water-to-light map:":
		return WATER_TO_LIGHT_MODE, true
	case "light-to-temperature map:":
		return LIGHT_TO_TEMPERATURE_MODE, true
	case "temperature-to-humidity map:":
		return TEMPERATURE_TO_HUMIDITY_MODE, true
	case "humidity-to-location map:":
		return HUMIDITY_TO_LOCATION_MODE, true
	default:
		return -1, false
	}
}

func digitStringToArray(input string) []int {
	if input == "" {
		return []int{}
	}
	result := []int{}
	digits := strings.Split(input, " ")
	for _, d := range digits {
		if d == "" {
			continue
		}

		v, err := strconv.Atoi(d)
		if err != nil {
			log.Fatalf("Error converting to int array '%v': %v \n", d, err)
		}
		result = append(result, v)
	}
	return result
}

func TraverseMapping(seed int, mapping []int) (int, bool) {
	destinationRangeStart := mapping[0]
	sourceRangeStart := mapping[1]
	// range is zero based
	rangeLength := mapping[2] - 1

	if seed < sourceRangeStart || seed > sourceRangeStart+rangeLength {
		return -1, false
	}

	diff := seed - sourceRangeStart
	return destinationRangeStart + diff, true
}

func Traverse(seed int, almanac Almanac) int {
	soil := TraverseToNext(seed, almanac.SeedToSoil)
	fertilizer := TraverseToNext(soil, almanac.SoilToFertilizer)
	water := TraverseToNext(fertilizer, almanac.FertilizerToWater)
	light := TraverseToNext(water, almanac.WaterToLight)
	temperature := TraverseToNext(light, almanac.LightToTemperature)
	humidity := TraverseToNext(temperature, almanac.TemperatureToHumidity)
	location := TraverseToNext(humidity, almanac.HumidityToLocation)

	return location
}

func TraverseToNext(seed int, mapping [][]int) int {
	for _, m := range mapping {
		next, ok := TraverseMapping(seed, m)
		if ok {
			return next
		}
	}
	// Any source numbers that aren't mapped correspond to the same destination number.
	return seed
}

func Level1(reader io.Reader) int {
	seeds, almanac := Parse(reader)

	lowestLocation := Traverse(seeds[0], almanac)
	for i := 1; i < len(seeds); i++ {
		newLocation := Traverse(seeds[i], almanac)
		if newLocation < lowestLocation {
			lowestLocation = newLocation
		}
	}
	return lowestLocation
}