package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type passport struct {
	BirthYear      int    `aoc:"byr"`
	IssueYear      int    `aoc:"iyr"`
	ExpirationYear int    `aoc:"eyr"`
	Height         string `aoc:"hgt"`
	HairColor      string `aoc:"hcl"`
	EyeColor       string `aoc:"ecl"`
	PassportID     string `aoc:"pid"`
	CountryID      string `aoc:"cid"`
}

func main() {
	fmt.Println("======= 🎄 AdventOfCode 2020 - Day 4 🎄 =======")
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	passportData, err := ioutil.ReadFile("input")
	if err != nil {
		return errors.Wrap(err, " > error reading input")
	}

	passports := parsePassports(string(passportData))

	var validPassports []passport
	var invalidPassports []passport
	for _, passport := range passports {
		if passport.BirthYear != 0 &&
			passport.IssueYear != 0 &&
			passport.ExpirationYear != 0 &&
			passport.PassportID != "" &&
			passport.EyeColor != "" &&
			passport.Height != "" &&
			passport.HairColor != "" {
			validPassports = append(validPassports, passport)
		} else {
			invalidPassports = append(invalidPassports, passport)
		}
	}

	fmt.Println(" > 🎅 Task 1")
	fmt.Println("  # detected passports:", len(passports))
	fmt.Println("  # valid passports:", len(validPassports))
	fmt.Println("  # invalid passports:", len(invalidPassports))

	var verifiedPassports []passport
	for _, passport := range validPassports {
		var height int
		var unit string
		fmt.Sscanf(passport.Height, "%d%s", &height, &unit)
		validHeight := (unit == "cm" && inRange(height, 150, 193)) || (unit == "in" && inRange(height, 59, 76))

		validHairColor, _ := regexp.MatchString(`#[0-9a-f]{6}$`, passport.HairColor)
		validPassportID, _ := regexp.MatchString(`^[0-9]{9}$`, passport.PassportID)

		if validHeight &&
			validHairColor &&
			validPassportID &&
			inRange(passport.BirthYear, 1920, 2002) &&
			inRange(passport.IssueYear, 2010, 2020) &&
			inRange(passport.ExpirationYear, 2020, 2030) &&
			contains(passport.EyeColor, []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}) {
			verifiedPassports = append(verifiedPassports, passport)
		}
	}

	fmt.Println(" > 🎅 Task 2")
	fmt.Println("  # valid and verified passports:", len(verifiedPassports))

	return nil
}

func inRange(value, lower, upper int) bool {
	if value >= lower && value <= upper {
		return true
	}
	return false
}

func contains(str string, elements []string) bool {
	for _, element := range elements {
		if element == str {
			return true
		}
	}
	return false
}

func parsePassports(input string) []passport {
	pass := passport{}
	var passports []passport
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			// next passport
			passports = append(passports, pass)
			pass = passport{}
		}
		for _, part := range strings.Split(row, " ") {
			marshal(part, &pass)
		}
	}
	passports = append(passports, pass)
	return passports
}

func marshal(in string, pass *passport) {
	if in == "" {
		return
	}

	passportVal := reflect.ValueOf(pass)
	passportType := reflect.TypeOf(*pass)

	parts := strings.Split(in, ":")
	// first part of input is the type and second the value e.g. `pid:261384974`
	t := parts[0]
	val := parts[1]
	for i := 0; i < passportType.NumField(); i++ {
		field := passportVal.Elem().Field(i)
		fieldType := passportType.Field(i)

		tag, _ := fieldType.Tag.Lookup("aoc")
		if t == tag {
			switch fieldType.Type.Kind() {
			case reflect.Int:
				intVal, _ := strconv.ParseInt(val, 10, 64)
				field.SetInt(intVal)
			case reflect.String:
				field.SetString(val)
			default:
				fmt.Println("Type not mapped")
			}
		}
	}
}
