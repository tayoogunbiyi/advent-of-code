// https://adventofcode.com/2020/day/5
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func isEmptyLine(line string) bool {
	return len(line) == 0
}

func computeSeatRowFromSeatCode(seatCode string, seatCodeIndex int, lowerBound int, upperBound int) int {
	if lowerBound == upperBound {
		return lowerBound
	}
	mid := (lowerBound + upperBound) / 2
	if string(seatCode[seatCodeIndex]) == "F" {
		upperBound = mid
	} else {
		lowerBound = mid + 1
	}
	return computeSeatRowFromSeatCode(seatCode, seatCodeIndex+1, lowerBound, upperBound)
}

func computeSeatColumnFromSeatCode(seatCode string, seatCodeIndex int, lowerBound int, upperBound int) int {
	if lowerBound == upperBound {
		return lowerBound
	}
	mid := (lowerBound + upperBound) / 2
	if string(seatCode[seatCodeIndex]) == "L" {
		upperBound = mid
	} else {
		lowerBound = mid + 1
	}
	return computeSeatColumnFromSeatCode(seatCode, seatCodeIndex+1, lowerBound, upperBound)
}

func convertSeatCodeToID(seatCode string) int {
	totalRowCount := 128
	totalColumnCount := 8
	row := computeSeatRowFromSeatCode(seatCode, 0, 0, totalRowCount-1)
	column := computeSeatColumnFromSeatCode(seatCode, 7, 0, totalColumnCount-1)

	return row*8 + column

}

func GetMySeatID(input string) int {
	maxSeatID := GetHighestSeatID(input)
	seatIDS := make([]bool, maxSeatID+1)

	for _, inputLine := range strings.Split(input, "\n") {
		if !isEmptyLine(inputLine) {
			currentSeatID := convertSeatCodeToID(inputLine)
			seatIDS[currentSeatID] = true
		}
	}

	for i := 1; i < len(seatIDS)-1; i++ {
		if seatIDS[i-1] && seatIDS[i+1] && !seatIDS[i] {
			return i
		}
	}
	return -1
}

func GetHighestSeatID(input string) int {
	maxSeatID := math.MinInt32
	for _, inputLine := range strings.Split(input, "\n") {
		if !isEmptyLine(inputLine) {
			currentSeatID := convertSeatCodeToID(inputLine)
			if currentSeatID > maxSeatID {
				maxSeatID = currentSeatID
			}
		}
	}
	return maxSeatID

}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The highest seat ID is %d\n", GetHighestSeatID(string(data)))
	fmt.Printf("Your seat ID is %d\n", GetMySeatID(string(data)))

}
