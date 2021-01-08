package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the sum of all values left in memory after completion is ", SumOfValuesLeftInMemory(string(data)))
}

func applyMask(mask string, value int) int {
	tempResult := ""
	for i := len(mask) - 1; i >= 0; i-- {
		currentMaskCharacter := string(mask[i])
		if currentMaskCharacter == "X" {
			tempResult = fmt.Sprint(value&1) + tempResult
		} else {
			tempResult = currentMaskCharacter + tempResult
		}
		value >>= 1
	}

	result, err := strconv.ParseInt(tempResult, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(result)
}

func SumOfValuesLeftInMemory(input string) int {
	addressMap := make(map[int]int)

	currentMask := ""
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, "mask") {
			currentMask = strings.Split(line, "= ")[1]
		} else {
			valueToWriteToMemory := strings.Trim(strings.Split(line, " =")[1], " ")
			address := strings.Trim(strings.Split(strings.Split(line, "]")[0], "[")[1], " ")

			intAddress, err := strconv.Atoi(address)
			if err != nil {
				log.Fatal(err)
			}
			value, err := strconv.Atoi(valueToWriteToMemory)
			if err != nil {
				log.Fatal(err)
			}
			maskedValue := applyMask(currentMask, value)
			addressMap[intAddress] = maskedValue
		}
	}
	result := 0
	for k := range addressMap {
		result += addressMap[k]
	}
	return result
}
