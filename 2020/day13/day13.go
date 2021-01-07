// https://www.reddit.com/r/adventofcode/comments/kc4njx/2020_day_13_solutions/?sort=confidence
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func FindFirstDepartingTimeAfterDeadline(busID int, deadline int) int {
	return int(math.Ceil(float64(deadline)/float64(busID))) * busID
}

func FindOptimalTimeIDProduct(input string) int {
	lines := strings.Split(input, "\n")
	earliestDepartingTime, _ := strconv.Atoi(lines[0])

	availableBusIDs := []int{}
	for _, id := range strings.Split(lines[1], ",") {
		if id != "x" {
			integerID, _ := strconv.Atoi(id)
			availableBusIDs = append(availableBusIDs, integerID)
		}
	}

	optimalDepartingTime := math.MaxInt32
	optimalBusID := -1
	for _, busID := range availableBusIDs {
		firstDepartingTime := FindFirstDepartingTimeAfterDeadline(busID, earliestDepartingTime)
		if firstDepartingTime < optimalDepartingTime {
			optimalDepartingTime = firstDepartingTime
			optimalBusID = busID
		}
	}

	waitTime := optimalDepartingTime - earliestDepartingTime
	return waitTime * optimalBusID
}

func FindFirstAligningTimestamp(input string) int {
	lines := strings.Split(input, "\n")

	type Bus struct {
		id    int
		index int
	}
	buses := []Bus{}

	for index, id := range strings.Split(lines[1], ",") {
		if id != "x" {
			integerID, _ := strconv.Atoi(id)
			buses = append(buses, Bus{id: integerID, index: index})
		}
	}
	firstBus := buses[0]
	aligningTimestamp := 0
	period := firstBus.id
	buses = buses[1:]

	/*

		Excellent Explanation - https://www.reddit.com/r/adventofcode/comments/kc4njx/2020_day_13_solutions/gfsc2gg?utm_source=share&utm_medium=web2x&context=3
	*/

	for _, bus := range buses {
		for (aligningTimestamp+bus.index)%bus.id != 0 {
			aligningTimestamp += period
		}
		period *= bus.id
	}
	return aligningTimestamp

}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("optimal time id product - ", FindOptimalTimeIDProduct(string(data)))
	fmt.Println("first aligning timestamp - ", FindFirstAligningTimestamp(string(data)))

}
