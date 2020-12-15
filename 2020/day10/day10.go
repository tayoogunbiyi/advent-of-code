package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

/*
- Each adapter can take 1,2 or 3 jolts < it's rating and still produce desired op
- Device has a built in adapter rated 3 + max(ip list)


*/

func maxKey(nums map[int]bool) int {
	result := math.MinInt32
	for k := range nums {
		if result < k {
			result = k
		}
	}
	return result
}

func ComputeJoltProduct(input string) int {
	lines := strings.Split(input, "\n")
	ratings := make(map[int]bool)

	for _, line := range lines {
		if len(line) > 0 {
			n, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			ratings[n] = true
		}
	}

	currentTerminatingOutletRating := 0
	oneJoltDifferences := 0
	threeJoltDifferences := 0
	maxAdapterRating := maxKey(ratings)

	for currentTerminatingOutletRating < maxAdapterRating {
		for _, delta := range []int{1, 2, 3} {
			potentialOutletRating := currentTerminatingOutletRating + delta
			if _, ok := ratings[potentialOutletRating]; ok {
				currentTerminatingOutletRating = potentialOutletRating
				if delta == 1 {
					oneJoltDifferences++
				} else if delta == 3 {
					threeJoltDifferences++
				}
				break
			}
		}
	}
	return oneJoltDifferences * (threeJoltDifferences + 1)
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the number of 1-jolt differences multiplied by 3-jolt differences is", ComputeJoltProduct(string(data)))
}
