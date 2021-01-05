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

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("optimal time id product - ", FindOptimalTimeIDProduct(string(data)))

}
