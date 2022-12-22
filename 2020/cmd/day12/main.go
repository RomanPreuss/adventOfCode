package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/RomanPreuss/adventOfCode2020/pkg/utils"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	actions := strings.Split(strings.TrimSpace(string(input)), "\n")
	task1(actions)
	task2(actions)
}

type position struct {
	east  float64
	north float64
}

func task1(input []string) {
	shipPos := position{0, 0}
	angle := float64(90)
	for _, line := range input {
		action := string(line[0])

		value, _ := strconv.ParseFloat(line[1:], 64)

		switch action {
		case "N":
			shipPos.north += value
		case "S":
			shipPos.north -= value
		case "E":
			shipPos.east += value
		case "W":
			shipPos.east -= value
		case "F":
			switch angle {
			case 0:
				shipPos.north += value
			case 90:
				shipPos.east += value
			case 180:
				shipPos.north -= value
			case 270:
				shipPos.east -= value
			}
		case "R":
			angle = utils.AbsFloat(math.Mod(angle+value, 360))
		case "L":
			angle = utils.AbsFloat(math.Mod(angle-value+360, 360))
		}
	}

	fmt.Println("Task1: manhatten distance", manhattenDistance(shipPos.east, shipPos.north))
}

func task2(input []string) {
	waypointPos := position{10, 1}
	shipPos := position{0, 0}
	for _, line := range input {
		action := string(line[0])
		value, _ := strconv.ParseFloat(line[1:], 64)

		switch action {
		case "N":
			waypointPos.north += value
		case "S":
			waypointPos.north -= value
		case "E":
			waypointPos.east += value
		case "W":
			waypointPos.east -= value
		case "F":
			shipPos.east += waypointPos.east * value
			shipPos.north += waypointPos.north * value
		case "R":
			waypointPos.east, waypointPos.north = utils.Rotate(float64(waypointPos.east), float64(waypointPos.north), float64(-value))
		case "L":
			waypointPos.east, waypointPos.north = utils.Rotate(float64(waypointPos.east), float64(waypointPos.north), float64(value))
		}
	}

	fmt.Println("Task2: manhatten distance", manhattenDistance(shipPos.east, shipPos.north))
}

func manhattenDistance(x, y float64) float64 {
	return utils.AbsFloat(x) + utils.AbsFloat(y)
}
