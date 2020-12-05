package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jollyra/go-advent-util"
)

func parsePassports(lines []string) (passports []map[string]string) {
	passport := make(map[string]string)
	for _, line := range lines {
		if line == "" {
			passports = append(passports, passport)
			passport = make(map[string]string)
			continue
		}
		fields := strings.Split(line, " ")
		for _, field := range fields {
			keyval := strings.Split(field, ":")
			passport[keyval[0]] = keyval[1]
		}
	}
	passports = append(passports, passport)
	return passports
}

func validate(passport map[string]string) bool {
	val, ok := passport["byr"]
	if !ok {
		return false
	}
	x, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if x < 1920 || x > 2002 {
		return false
	}

	val, ok = passport["iyr"]
	if !ok {
		return false
	}
	x, err = strconv.Atoi(val)
	if err != nil {
		return false
	}
	if x < 2010 || x > 2020 {
		return false
	}

	val, ok = passport["eyr"]
	if !ok {
		return false
	}
	x, err = strconv.Atoi(val)
	if err != nil {
		return false
	}
	if x < 2020 || x > 2030 {
		return false
	}

	val, ok = passport["hgt"]
	if !ok {
		return false
	}
	var h int
	var unit string
	n, err := fmt.Sscanf(val, "%d%s", &h, &unit)
	if err != nil {
		return false
	}
	if n != 2 {
		return false
	}
	if unit == "cm" {
		if h < 150 || h > 193 {
			return false
		}
	} else if unit == "in" {
		if h < 59 || h > 76 {
			return false
		}
	} else {
		return false
	}

	val, ok = passport["hcl"]
	if !ok {
		return false
	}
	re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	if !re.MatchString(val) {
		return false
	}

	val, ok = passport["ecl"]
	if !ok {
		return false
	}
	re = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	if !re.MatchString(val) {
		return false
	}

	val, ok = passport["pid"]
	if !ok {
		return false
	}
	re = regexp.MustCompile(`^[0-9]{9}$`)
	if !re.MatchString(val) {
		return false
	}

	return true
}

func main() {
	lines := advent.InputLines(advent.MustGetArg(1))
	passports := parsePassports(lines)
	count := 0
	for _, passport := range passports {
		if validate(passport) {
			count++
		}
	}
	fmt.Println(count)
}
