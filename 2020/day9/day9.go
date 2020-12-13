// https://adventofcode.com/2020/day/10
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func ParseInput(input string) []int {
	lines := strings.Split(input, "\n")
	var result []int
	for _, line := range lines {
		if len(line) > 0 {
			number, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, number)
		}
	}
	return result
}

func hasTwoSum(numbers []int, startIndex int, endIndex int, target int) bool {
	ht := make(map[int]bool)
	for i := startIndex; i <= endIndex; i++ {
		currentTarget := target - numbers[i]
		if _, ok := ht[currentTarget]; ok && currentTarget != numbers[i] {
			return true
		}
		ht[numbers[i]] = true
	}
	return false
}

func DetectFirstViolatingNumber(numbers []int, preambleSize int) int {
	for i := preambleSize; i < len(numbers); i++ {
		target := numbers[i]
		if !hasTwoSum(numbers, i-preambleSize, i-1, target) {
			return target
		}
	}
	return -1
}

func FindMinInRange(numbers []int, lower int, upper int) int {
	minInRange := math.MaxInt32

	for i := lower; i <= upper; i++ {
		if numbers[i] < minInRange {
			minInRange = numbers[i]
		}
	}
	return minInRange
}

func FindMaxInRange(numbers []int, lower int, upper int) int {
	maxInRange := math.MinInt32

	for i := lower; i <= upper; i++ {
		if numbers[i] > maxInRange {
			maxInRange = numbers[i]
		}
	}
	return maxInRange
}

func ComputeEncryptionWeakness(numbers []int, violatingNumber int) int {
	ht := make(map[int]int)
	ht[0] = -1
	currentRunningSum := 0
	for i, number := range numbers {
		currentRunningSum += number
		j, seen := ht[currentRunningSum-violatingNumber]
		if seen && j+1 != i {
			minimumInRange := FindMinInRange(numbers, j+1, i)
			maximumInRange := FindMaxInRange(numbers, j+1, i)
			return minimumInRange + maximumInRange
		}
		ht[currentRunningSum] = i
	}
	return -1
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	numbers := ParseInput(string(data))
	firstViolatingNumber := DetectFirstViolatingNumber(numbers, 25)
	fmt.Printf("the first violating number is %d \n", firstViolatingNumber)
	fmt.Printf("the encryption weakness is is %d \n", ComputeEncryptionWeakness(numbers, firstViolatingNumber))

}
