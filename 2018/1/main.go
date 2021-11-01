// https://adventofcode.com/2018/day/1
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInputIntoSlice(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}
	return numbers
}

func main() {
	numbers := readInputIntoSlice("input.txt")

	seenFrequencies := make(map[int]bool)
	currentFrequency := 0
	seenFrequencies[currentFrequency] = true

	for i := 0; i < len(numbers); i = (i + 1) % len(numbers) {
		currentFrequency += numbers[i]
		if _, ok := seenFrequencies[currentFrequency]; ok {
			fmt.Println(currentFrequency)
			return
		}
		seenFrequencies[currentFrequency] = true
	}
}
