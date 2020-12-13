// https://adventofcode.com/2020/day/8
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var (
	JMP = "jmp"
	NOP = "nop"
	ACC = "acc"
)

type Instruction struct {
	operation string
	arg       int
}

type BootCode struct {
	instructions []*Instruction
	accumulator  int
}

func NewBootCode() *BootCode {
	return &BootCode{
		instructions: make([]*Instruction, 0),
	}
}

// Execute executes 1the currently loaded program. returns the final accumulator value at the end of execution
func (bc *BootCode) Execute() (int, error) {
	processedInstructions := make(map[int]bool)
	currentInstructionIndex := 0
	accumulator := 0

	for currentInstructionIndex < len(bc.instructions) {
		currentInstruction := bc.instructions[currentInstructionIndex]

		if _, processed := processedInstructions[currentInstructionIndex]; processed {
			return accumulator, errors.New("cycle detected in program")
		}

		processedInstructions[currentInstructionIndex] = true

		switch currentInstruction.operation {
		case ACC:
			accumulator += currentInstruction.arg
			currentInstructionIndex++
		case JMP:
			currentInstructionIndex += currentInstruction.arg
		case NOP:
			currentInstructionIndex++
		}
	}
	return accumulator, nil
}

func (bc *BootCode) LoadProgram(input string) {
	inputLines := strings.Split(input, "\n")
	for _, line := range inputLines {
		if len(line) > 0 {
			instruction := strings.Split(line, " ")
			arg, _ := strconv.Atoi(instruction[1])

			bc.instructions = append(bc.instructions, &Instruction{
				operation: instruction[0],
				arg:       arg,
			})
		}
	}

}

func FindFinalAccumulatorValue(input string) int {
	bc := NewBootCode()
	bc.LoadProgram(input)
	acc, err := bc.Execute()
	if err != nil {
		log.Println(err)
	}
	return acc
}

func operationCanCauseCycle(op string) bool {
	return op == JMP || op == NOP
}

func getReplacementOperation(op string) (string, error) {
	if op == JMP {
		return NOP, nil
	} else if op == NOP {
		return JMP, nil
	} else {
		return "", fmt.Errorf("operation %s has no replacement operation", op)
	}
}

func FindFinalAccumulatorValueWithCycle(input string) int {
	bc := NewBootCode()
	bc.LoadProgram(input)

	for _, instruction := range bc.instructions {
		if operationCanCauseCycle(instruction.operation) {
			replacementOperation, err := getReplacementOperation(instruction.operation)
			if err != nil {
				log.Fatal(err)
			}
			originalOperation := instruction.operation
			instruction.operation = replacementOperation
			acc, err := bc.Execute()
			if err == nil {
				return acc
			}
			instruction.operation = originalOperation
		}
	}
	return -1
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("the accumulator's value right at the start of the cycle is ", FindFinalAccumulatorValue(string(data)))
	fmt.Println("the accumulator's value after fixing the cycle is ", FindFinalAccumulatorValueWithCycle(string(data)))

}
