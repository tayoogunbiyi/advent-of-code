// https://adventofcode.com/2018/day/2
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func buildCountMap(id string) map[rune]int {
	countMap := make(map[rune]int)
	for _, ch := range id {
		countMap[ch] += 1
	}
	return countMap
}

func hasKeyWithTargetOccurences(ht map[rune]int, targetOccurences int) bool {
	for _, v := range ht {
		if v == targetOccurences {
			return true
		}
	}
	return false
}

func readInputIntoSlice(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var strings []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings
}

func main() {
	boxIds := readInputIntoSlice("input.txt")
	threeOccurences := 0
	twoOccurences := 0

	for _, id := range boxIds {
		countMap := buildCountMap(id)
		if hasKeyWithTargetOccurences(countMap, 2) {
			twoOccurences += 1
		}

		if hasKeyWithTargetOccurences(countMap, 3) {
			threeOccurences += 1
		}
	}

	fmt.Println(twoOccurences * threeOccurences)
}
