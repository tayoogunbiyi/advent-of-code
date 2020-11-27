// https://adventofcode.com/2019/day/1

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetTotalFuelRequirements(rocketMasses []int) int {
	totalFuelRequirements := 0

	for _, rocketMass := range rocketMasses {
		currentFuelRequirement := (rocketMass / 3) - 2
		if currentFuelRequirement < 0 {
			currentFuelRequirement = 0
		}
		totalFuelRequirements += currentFuelRequirement
	}

	return totalFuelRequirements
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var masses []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		masses = append(masses, mass)
	}

	result := GetTotalFuelRequirements(masses)

	fmt.Println(result)

}
