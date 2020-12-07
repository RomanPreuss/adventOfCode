package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	amount int
	color  string
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	var colorRe = regexp.MustCompile(`^\w+ \w+`)
	var ruleRe = regexp.MustCompile(`(\d+) (\w+ \w+)`)
	mapping := make(map[string][]rule)
	for _, line := range strings.Split(string(input), "\n") {
		parentColor := colorRe.FindAllString(line, -1)[0]
		mapping[parentColor] = []rule{}

		for _, match := range ruleRe.FindAllStringSubmatch(line, -1) {
			amount, _ := strconv.Atoi(match[1])
			mapping[parentColor] = append(mapping[parentColor], rule{
				amount: amount,
				color:  match[2],
			})
		}
	}

	matchingBags := containsBag("shiny gold", mapping)
	fmt.Printf("Task 1: %v amount of bags contain at least one 'shiny gold' bag\n", len(matchingBags))
	fmt.Printf("Task 2: %v are required to be inside a 'shiny gold' bag\n", amountOfBags(0, 0, "shiny gold", mapping))
}

func amountOfBags(indentation, countSelf int, color string, mapping map[string][]rule) int {
	rules := mapping[color]
	if len(rules) == 0 {
		return 1
	}

	sum := countSelf
	for _, rule := range rules {
		child := amountOfBags(indentation+1, 1, rule.color, mapping)
		sum += rule.amount * child
	}
	return sum
}

func containsBag(color string, mapping map[string][]rule) map[string]string {
	matchingBags := map[string]string{}
	for key, rules := range mapping {
		for _, rule := range rules {
			if rule.color == color {
				matchingBags[key] = key
				for k, v := range containsBag(key, mapping) {
					matchingBags[k] = v + " -> " + key
				}
			}
		}
	}
	return matchingBags
}
