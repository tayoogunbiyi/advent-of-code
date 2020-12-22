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

type Coordinates struct {
	north int
	east  int
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

func rotateToNewDirectionIndex(instruction *Instruction, currentDirectionIndex int) int {
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
	return currentDirectionIndex
}

func FindDistanceMoved(input string) int {
	instructions := ParseInput(input)

	currentDirectionIndex := inverseDirections["EAST"]

	eastCoordinate := 0
	northCoordinate := 0

	for _, instruction := range instructions {
		if instruction.CanChangeDirection() {
			currentDirectionIndex = rotateToNewDirectionIndex(instruction, currentDirectionIndex)
		} else {
			eastCoordinate, northCoordinate = instruction.ApplyOnCoordinates(directions[currentDirectionIndex], eastCoordinate, northCoordinate)
		}
	}
	return abs(eastCoordinate) + abs(northCoordinate)
}

func turnWayPoint(instruction *Instruction, waypointCoordinates *Coordinates) *Coordinates {
	newWaypointCoordinates := &Coordinates{north: waypointCoordinates.north, east: waypointCoordinates.east}

	if instruction.action == "R" {
		if instruction.value == 90 {
			newWaypointCoordinates.north = -waypointCoordinates.east
			newWaypointCoordinates.east = waypointCoordinates.north
		} else if instruction.value == 180 {
			newWaypointCoordinates.north = -waypointCoordinates.north
			newWaypointCoordinates.east = -waypointCoordinates.east
		} else if instruction.value == 270 {
			newWaypointCoordinates.north = waypointCoordinates.east
			newWaypointCoordinates.east = -waypointCoordinates.north
		}
	} else if instruction.action == "L" {
		return turnWayPoint(&Instruction{action: "R", value: 360 - instruction.value}, waypointCoordinates)
	} else {
		log.Fatalf("unexpected instruction %v", instruction)
	}
	return newWaypointCoordinates
}

func FindDistanceMovedWithWaypoint(input string) int {
	instructions := ParseInput(input)

	waypointCoordinates := &Coordinates{north: 1, east: 10}
	shipCoordinates := &Coordinates{north: 0, east: 0}

	for _, instruction := range instructions {
		if instruction.action == "F" {
			shipCoordinates.east += waypointCoordinates.east * instruction.value
			shipCoordinates.north += waypointCoordinates.north * instruction.value
		} else if instruction.action == "N" || instruction.action == "S" {
			multiplier := 1
			if instruction.action == "S" {
				multiplier *= -1
			}
			waypointCoordinates.north += multiplier * instruction.value
		} else if instruction.action == "W" || instruction.action == "E" {
			multiplier := 1
			if instruction.action == "W" {
				multiplier *= -1
			}
			waypointCoordinates.east += multiplier * instruction.value
		} else if instruction.CanChangeDirection() {
			waypointCoordinates = turnWayPoint(instruction, waypointCoordinates)

		} else {
			log.Fatalf("unexpected instruction %v", instruction)
		}
	}
	return abs(shipCoordinates.north) + abs(shipCoordinates.east)
}

func abs(v int) int {
	if v < 0 {
		v *= -1
	}
	return v
}
func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("distance moved - ", FindDistanceMoved(string(data)))
	fmt.Println("distance moved with waypoint - ", FindDistanceMovedWithWaypoint(string(data)))
}
