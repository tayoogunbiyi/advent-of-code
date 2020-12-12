package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	arg       int
}

type BootCode struct {
	instructions []Instruction
	accumulator  int
}

func (bc *BootCode) Execute() {
	processedInstructionIndexes := make(map[int]bool)
	currentInstructionIndex := 0

	for currentInstructionIndex < len(bc.instructions) {
		currentInstruction := bc.instructions[currentInstructionIndex]
		_, processed := processedInstructionIndexes[currentInstructionIndex]
		if processed {
			log.Println("cycle detected")
			return
		}
		processedInstructionIndexes[currentInstructionIndex] = true

		switch currentInstruction.operation {
		case "acc":
			bc.accumulator += currentInstruction.arg
			currentInstructionIndex++
		case "jmp":
			currentInstructionIndex += currentInstruction.arg
		case "nop":
			currentInstructionIndex++
		}
	}
}

func (bc *BootCode) LoadProgram(input string) {
	inputLines := strings.Split(input, "\n")
	for _, line := range inputLines {
		if len(line) > 0 {
			instruction := strings.Split(line, " ")
			arg, _ := strconv.Atoi(instruction[1])

			bc.instructions = append(bc.instructions, Instruction{
				operation: instruction[0],
				arg:       arg,
			})
		}
	}

}

func FindAccumulatorValueAtCycleStart(input string) int {
	bc := BootCode{
		instructions: make([]Instruction, 0),
	}
	bc.LoadProgram(input)
	bc.Execute()
	return bc.accumulator

}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("the accumulator's value right at the start of the cycle is ", FindAccumulatorValueAtCycleStart(string(data)))

}
