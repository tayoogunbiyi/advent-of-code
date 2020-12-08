// https://adventofcode.com/2020/day/4
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var fields = []field{
	{name: "ecl", validator: func(v string) bool {
		for _, choices := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
			if v == choices {
				return true
			}
		}
		return false
	}},
	{
		name: "byr", validator: func(v string) bool {
			vInt, err := strconv.Atoi(v)
			return err == nil && isInRange(vInt, 1920, 2002)
		}},
	{
		name: "iyr", validator: func(v string) bool {
			vInt, err := strconv.Atoi(v)
			return err == nil && isInRange(vInt, 2010, 2020)
		}},
	{
		name: "eyr", validator: func(v string) bool {
			vInt, err := strconv.Atoi(v)
			return err == nil && isInRange(vInt, 2020, 2030)
		}},
	{
		name: "cid", validator: func(v string) bool {
			return true
		}},
	{
		name: "hgt", validator: func(v string) bool {
			height, err := strconv.Atoi(v[:len(v)-2])
			unit := v[len(v)-2:]

			lower := 0
			upper := 0

			if unit == "cm" {
				lower = 150
				upper = 193
			} else if unit == "in" {
				lower = 59
				upper = 76
			} else {
				return false
			}
			return err == nil && isInRange(height, lower, upper)
		}},
	{
		name: "hcl", validator: func(v string) bool {
			allowedCharacters := "0123456789abcdef"
			for _, ch := range v[1:] {
				if !strings.Contains(allowedCharacters, string(ch)) {
					return false
				}
			}
			return string(v[0]) == "#"
		}},
	{
		name: "pid", validator: func(v string) bool {
			allowedCharacters := "0123456789"
			for _, ch := range v {
				if !strings.Contains(allowedCharacters, string(ch)) {
					return false
				}
			}
			return true && len(v) == 9
		}},
}

type field struct {
	name      string
	validator func(string) bool
}

type passport struct {
	ht map[string]string
}

func (p passport) updatePassportData(updateData map[string]string) {
	for k, v := range updateData {
		p.ht[k] = v
	}
}

func isBlankLine(line string) bool {
	return len(line) == 0
}

func processLineOfPassportData(line string) map[string]string {
	ht := make(map[string]string)
	for _, dataPair := range strings.Split(line, " ") {
		key, value := strings.Split(dataPair, ":")[0], strings.Split(dataPair, ":")[1]
		ht[key] = value
	}

	return ht
}

func isInRange(n int, lower int, upper int) bool {
	return n >= lower && n <= upper

}

func (p passport) isValidWithConstraints() bool {
	fieldToValidatorFunc := make(map[string]func(string) bool)
	for _, field := range fields {
		fieldToValidatorFunc[field.name] = field.validator
	}

	for fieldName, fieldValue := range p.ht {
		if !fieldToValidatorFunc[fieldName](fieldValue) {
			return false
		}
	}
	return p.isValid()
}

func (p passport) isValid() bool {
	requiredFieldMap := make(map[string]bool)

	for _, field := range fields {
		requiredFieldMap[field.name] = true
	}

	for fieldName := range p.ht {
		if requiredFieldMap[fieldName] {
			delete(requiredFieldMap, fieldName)
		}
	}
	isValidNorthPolePassport := len(requiredFieldMap) == 1 && requiredFieldMap["cid"]
	return len(requiredFieldMap) == 0 || isValidNorthPolePassport
}

func newPassport() *passport {
	return &passport{ht: make(map[string]string)}
}

func parsePassportInput(input string) []*passport {
	var passports []*passport
	passports = append(passports, newPassport())
	inputLines := strings.Split(input, "\n")

	for i, inputLine := range inputLines {
		if isBlankLine(inputLine) {
			if i+1 != len(inputLines) {
				passports = append(passports, newPassport())
			}
		} else {
			latestPassport := passports[len(passports)-1]
			latestPassport.updatePassportData(processLineOfPassportData(inputLine))
		}
	}
	return passports
}

func CountValidPassports(input string) int {
	passports := parsePassportInput(input)
	count := 0
	for _, passport := range passports {
		if passport.isValid() {
			count++
		}
	}
	return count
}

func CountValidPassportsWithExtraConstraints(input string) int {
	passports := parsePassportInput(input)
	count := 0
	for _, passport := range passports {
		if passport.isValidWithConstraints() {
			count++
		}
	}
	return count
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the number of valid passports without extra constraints is %d \n", CountValidPassports(string(data)))
	fmt.Printf("the number of valid passports with  extra constraints is %d \n", CountValidPassportsWithExtraConstraints(string(data)))

}
