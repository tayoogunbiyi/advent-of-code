// https://adventofcode.com/2020/day/12
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	directions        = []string{"EAST", "SOUTH", "WEST", "NORTH"}
	inverseDirections = make(map[string]int)
)

func init() {
	for i := range directions {
		inverseDirections[directions[i]] = i
	}
}

type Instruction struct {
	action string
	value  int
}

func (instruction *Instruction) CanChangeDirection() bool {
	return instruction.action == "R" || instruction.action == "L"
}

func (instruction *Instruction) ApplyOnCoordinates(currentDirection string, eastCoordinate, northCoordinate int) (int, int) {
	ht := make(map[string]string)

	for _, direction := range directions {
		ht[direction] = string(direction[0])
	}

	effectiveAction := instruction.action
	if effectiveAction == "F" {
		effectiveAction = ht[currentDirection]
	}

	if effectiveAction == "N" {
		northCoordinate += instruction.value
	} else if effectiveAction == "S" {
		northCoordinate -= instruction.value
	} else if effectiveAction == "E" {
		eastCoordinate += instruction.value
	} else if effectiveAction == "W" {
		eastCoordinate -= instruction.value
	} else {
		log.Fatalf("unaexpected action %v", effectiveAction)
	}

	return eastCoordinate, northCoordinate
}

func ParseInput(input string) []*Instruction {
	lines := strings.Split(input, "\n")
	instructions := []*Instruction{}

	for _, line := range lines {
		if len(line) > 0 {
			action := string(line[0])
			value, err := strconv.Atoi(line[1:])
			if err != nil {
				log.Fatal(err)
			}

			instructions = append(instructions, &Instruction{
				action, value,
			})
		}
	}
	return instructions
}

func FindDistanceMoved(input string) int {
	instructions := ParseInput(input)

	currentDirectionIndex := inverseDirections["EAST"]

	eastCoordinate := 0
	northCoordinate := 0

	for i, instruction := range instructions {
		if instruction.CanChangeDirection() {
			if instruction.action == "R" {
				currentDirectionIndex += instruction.value / 90
				currentDirectionIndex %= len(directions)
			} else {
				currentDirectionIndex -= instruction.value / 90
				if currentDirectionIndex < 0 {
					currentDirectionIndex += len(directions)
				}
				currentDirectionIndex %= len(directions)
			}
		} else {
			eastCoordinate, northCoordinate = instruction.ApplyOnCoordinates(directions[currentDirectionIndex], eastCoordinate, northCoordinate)
		}
		fmt.Println(i, *instruction)
		fmt.Println(eastCoordinate, northCoordinate, currentDirectionIndex, directions[currentDirectionIndex])
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Continue: ")
		// reader.ReadString('\n')
	}

	if eastCoordinate < 0 {
		eastCoordinate = -eastCoordinate
	}
	if northCoordinate < 0 {
		northCoordinate = -northCoordinate
	}

	fmt.Println(eastCoordinate, northCoordinate)
	return eastCoordinate + northCoordinate
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("distance moved - ", FindDistanceMoved(string(data)))
}
