// https://adventofcode.com/2020/day/4
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

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

type passport struct {
	ht map[string]string
}

func (p passport) updatePassportData(updateData map[string]string) {
	for k, v := range updateData {
		p.ht[k] = v
	}
}

func (p passport) isValid() bool {
	requiredFieldMap := make(map[string]bool)
	for _, field := range []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "cid", "hgt"} {
		requiredFieldMap[field] = true
	}

	for k := range p.ht {
		if requiredFieldMap[k] {
			delete(requiredFieldMap, k)
		}
	}
	isValidNorthPolePassport := len(requiredFieldMap) == 1 && requiredFieldMap["cid"]
	return len(requiredFieldMap) == 0 || isValidNorthPolePassport
}

func newPassport() *passport {
	return &passport{ht: make(map[string]string)}
}

func CountValidPassports(input string) int {
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

	count := 0
	for _, passport := range passports {
		if passport.isValid() {
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
	fmt.Printf("the number of valid passports is %d \n", CountValidPassports(string(data)))
}
