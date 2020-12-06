//https://adventofcode.com/2020/day/2
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	Password      string
	MinOccurences int
	MaxOccurences int
	Letter        string
}

func (pwPolicy PasswordPolicy) isValid() bool {
	occurences := strings.Count(pwPolicy.Password, pwPolicy.Letter)
	return occurences >= pwPolicy.MinOccurences && occurences <= pwPolicy.MaxOccurences
}

func readInputIntoPolicySlice(input string) []PasswordPolicy {
	var passwordPolicies []PasswordPolicy

	for _, s := range strings.Split(input, "\n") {
		if len(s) > 0 {
			occurencesAndLetter, password := strings.Split(s, ":")[0], strings.Split(s, ":")[1]
			occurences, letter := strings.Split(occurencesAndLetter, " ")[0], strings.Split(occurencesAndLetter, " ")[1]

			minAndMaxOccurences := strings.Split(occurences, "-")
			minOccurences, _ := strconv.Atoi(minAndMaxOccurences[0])
			maxOccurences, _ := strconv.Atoi(minAndMaxOccurences[1])

			letter = string(letter)
			pwPolicy := PasswordPolicy{Password: strings.Trim(password, " "), MinOccurences: minOccurences, MaxOccurences: maxOccurences, Letter: letter}
			passwordPolicies = append(passwordPolicies, pwPolicy)
		}
	}
	return passwordPolicies
}

func CountValidPasswords(passwordPolicies []PasswordPolicy) int {
	count := 0
	for _, pwPolicy := range passwordPolicies {
		if pwPolicy.isValid() {
			count++
		}
	}
	return count
}

func CountValidPasswordsGivenNewInterpretation(passwordPolicies []PasswordPolicy) int {
	count := 0
	for _, pwPolicy := range passwordPolicies {
		firstIndexContainsLetter := string(pwPolicy.Password[pwPolicy.MinOccurences-1]) == pwPolicy.Letter
		secondIndexContainsLetter := string(pwPolicy.Password[pwPolicy.MaxOccurences-1]) == pwPolicy.Letter
		// != simulates the XOR operation in this case as we need only 1 to be true
		if firstIndexContainsLetter != secondIndexContainsLetter {
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
	passwordPolicies := readInputIntoPolicySlice(string(data))
	fmt.Printf("%d passwords are valid\n", CountValidPasswords(passwordPolicies))
	fmt.Printf("%d passwords are valid given new interpretation\n", CountValidPasswordsGivenNewInterpretation(passwordPolicies))

}
