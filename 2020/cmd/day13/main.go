package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	busplan := strings.Split(strings.TrimSpace(string(input)), "\n")
	// task1(busplan)
	task2(busplan)
}

func task2(busplan []string) {
	busses := []int{}
	for _, busID := range strings.Split(busplan[1], ",") {
		if busID == "x" {
			busses = append(busses, -1)
		} else {
			id, _ := strconv.Atoi(busID)
			busses = append(busses, id)
		}
	}

	cache := make([]int, len(busses))
	match := false
	for t := 0; !match; t++ {
		for i, busID := range busses {
			fmt.Println(i, busID)
			depature := 0
			if busID == -1 {
				depature = cache[i-1] + 1
			} else {
				depature = busID * (t + i)
			}
			if i > 0 {
				if depature-cache[i-1] != 1 {
					break
				}
				if i == len(busses)-1 {
					match = true
				}
			}

			cache[i] = depature
			fmt.Println("time", t, cache)
			if depature > 3420 {
				return
			}
		}
		fmt.Println("")
	}

}

func task1(busplan []string) {
	depature, _ := strconv.Atoi(busplan[0])
	shortestTime := depature
	for _, f := range strings.Split(busplan[1], ",") {
		if f == "x" {
			continue
		}
		busID, _ := strconv.Atoi(f)

		seaportDepature := busID * (int(depature/busID) + 1)
		timeToArrival := seaportDepature - depature
		if timeToArrival < shortestTime {
			shortestTime = timeToArrival
			result := busID * shortestTime
			fmt.Printf("Bus %v arrives in %v minutes = %v\n", busID, timeToArrival, result)
		}
	}
}
