// https://adventofcode.com/2019/day/4
package main

import (
	"fmt"
	"strconv"
)

func hasEqualAdjacentDigits(password int) bool {
	stringifiedPassword := strconv.Itoa(password)

	for i := 1; i < len(stringifiedPassword); i++ {
		if stringifiedPassword[i] == stringifiedPassword[i-1] {
			return true
		}
	}
	return false
}

func isStrictlyNonDecreasing(password int) bool {
	stringifiedPassword := strconv.Itoa(password)

	for i := 1; i < len(stringifiedPassword); i++ {
		currentDigit, _ := strconv.Atoi(string(stringifiedPassword[i]))
		previousDigit, _ := strconv.Atoi(string(stringifiedPassword[i-1]))
		if currentDigit < previousDigit {
			return false
		}
	}
	return true
}

func isValidPassword1(password int) bool {
	return hasEqualAdjacentDigits(password) && isStrictlyNonDecreasing(password)
}

func isValidPassword2(password int) bool {
	return hasEqualAdjacentDigits2(password) && isStrictlyNonDecreasing(password)
}

func hasEqualAdjacentDigits2(password int) bool {
	stringifiedPassword := strconv.Itoa(password)
	runningCount := 1

	for i := 1; i < len(stringifiedPassword); i++ {
		currentDigit, _ := strconv.Atoi(string(stringifiedPassword[i]))
		previousDigit, _ := strconv.Atoi(string(stringifiedPassword[i-1]))
		if previousDigit != currentDigit && runningCount == 2 {
			return true
		} else if previousDigit == currentDigit {
			runningCount++
		} else {
			runningCount = 1
		}
	}
	return runningCount == 2
}

func main() {
	passwordLowerBound := 206938
	passwordUpperBound := 679128
	validPasswordCount := 0
	for currentPassword := passwordLowerBound; currentPassword <= passwordUpperBound; currentPassword++ {
		if isValidPassword1(currentPassword) {
			validPasswordCount++
		}
	}
	fmt.Printf("The number of valid passwords [For Part 1] in the range %d-%d is %d \n", passwordLowerBound, passwordUpperBound, validPasswordCount)

	validPasswordCount = 0
	for currentPassword := passwordLowerBound; currentPassword <= passwordUpperBound; currentPassword++ {
		if isValidPassword2(currentPassword) {
			validPasswordCount++
		}
	}

	fmt.Printf("The number of valid passwords [For Part 2] in the range %d-%d is %d \n", passwordLowerBound, passwordUpperBound, validPasswordCount)

}
