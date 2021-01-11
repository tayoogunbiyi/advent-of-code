// https://adventofcode.com/2020/day/15/
package main

import (
	"fmt"
	"time"
)

func FindNthNumberSpoken(startingNumbers []int, n int) int {
	ht := make(map[int][]int)

	var mostRecentlySpokenNumber int

	for i, number := range startingNumbers {
		turn := i + 1
		ht[number] = append(ht[number], turn)
		mostRecentlySpokenNumber = number
	}

	for turn := len(startingNumbers) + 1; turn <= n; turn++ {
		turnsWhereMostRecentNumberWasSpoken := ht[mostRecentlySpokenNumber]
		if len(turnsWhereMostRecentNumberWasSpoken) == 1 {
			mostRecentlySpokenNumber = 0
		} else {
			mostRecentlySpokenNumber = turnsWhereMostRecentNumberWasSpoken[len(turnsWhereMostRecentNumberWasSpoken)-1] - turnsWhereMostRecentNumberWasSpoken[len(turnsWhereMostRecentNumberWasSpoken)-2]
		}
		ht[mostRecentlySpokenNumber] = append(ht[mostRecentlySpokenNumber], turn)
	}

	return mostRecentlySpokenNumber
}

func main() {
	startingNumbers := []int{8, 13, 1, 0, 18, 9}
	start := time.Now()
	fmt.Println(FindNthNumberSpoken(startingNumbers, 2020))
	t := time.Now()
	fmt.Println("Duration : ", t.Sub(start))

	start = time.Now()
	fmt.Println(FindNthNumberSpoken(startingNumbers, 30000000))
	t = time.Now()
	fmt.Println("Duration : ", t.Sub(start))
}
