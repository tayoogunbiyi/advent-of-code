// https://adventofcode.com/2019/day/2

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const additionOpCode = 1
const multiplicationOpCode = 2
const haltingOpCode = 99

func applyOpCode(opCode int, operand1 int, operand2 int) (int, error) {
	if opCode == additionOpCode {
		return operand1 + operand2, nil
	} else if opCode == multiplicationOpCode {
		return operand1 * operand2, nil
	} else {
		return -1, errors.New("unsupported opCode")
	}
}

// ProcessInstructions takes in a slice of instructions and processes each of the opCodes along with it's params
func ProcessInstructions(originalInstructions []int, noun int, verb int) int {
	instructions := make([]int, len(originalInstructions))
	copy(instructions, originalInstructions)

	instructions[1] = noun
	instructions[2] = verb
	for i := 0; i < len(instructions); i += 4 {
		opCode := instructions[i]
		if opCode == haltingOpCode {
			return instructions[0]
		}
		operandIndex1, operandIndex2 := instructions[i+1], instructions[i+2]
		destinationIndex := instructions[i+3]

		result, err := applyOpCode(opCode, instructions[operandIndex1], instructions[operandIndex2])
		if err != nil {
			log.Fatal(err)
		}
		instructions[destinationIndex] = result
	}
	return instructions[0]
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(data)

	var instructions []int

	for _, op := range strings.Split(strings.TrimRight(text, "\n"), ",") {
		integerOp, _ := strconv.Atoi(op)
		instructions = append(instructions, integerOp)
	}

	NOUN := 12
	VERB := 2

	fmt.Println("Value left at position[0] is ", ProcessInstructions(instructions, NOUN, VERB))
	noun, verb := ReverseProcessInstructions(instructions, 19690720)

	fmt.Printf("To obtain output %d , noun = %d and verb = %d\n", 19690720, noun, verb)
	fmt.Println(100*noun + verb)
}

// ReverseProcessInstructions takes in a slice of instructions and a desired output and attempted to find a noun and verb that yields the desired output
func ReverseProcessInstructions(originalInstructions []int, desiredOutput int) (int, int) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if ProcessInstructions(originalInstructions, noun, verb) == desiredOutput {
				return noun, verb
			}
		}
	}
	return -1, -1
}
