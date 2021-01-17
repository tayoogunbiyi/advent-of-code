package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Field struct {
	Name   string
	Ranges []Range
}

type Range struct {
	Start int
	End   int
}

func (r *Range) contains(value int) bool {
	return r.Start <= value && r.End >= value
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func ComputeTicketScanningErrorRate(input string) int {
	currentLineNumber := 0
	lines := strings.Split(input, "\n")

	fields := []Field{}
	for i := currentLineNumber; i < len(lines) && len(lines[i]) != 0; i++ {
		line := lines[i]
		fieldName := strings.Split(line, ":")[0]
		fieldRanges := strings.Split(strings.Split(line, ":")[1], "or")
		field := Field{Name: fieldName}
		for _, fieldRange := range fieldRanges {
			start, end := strings.Split(fieldRange, "-")[0], strings.Split(fieldRange, "-")[1]
			start = strings.Trim(start, " ")
			end = strings.Trim(end, " ")
			field.Ranges = append(field.Ranges, Range{Start: atoi(start), End: atoi(end)})
		}
		currentLineNumber = i
		fields = append(fields, field)
	}

	currentLineNumber += 2
	for i := currentLineNumber; i < len(lines) && len(lines[i]) != 0; i++ {
		currentLineNumber = i
	}

	currentLineNumber += 3
	totalErrorRate := 0
	for i := currentLineNumber; i < len(lines) && len(lines[i]) != 0; i++ {
		values := strings.Split(lines[i], ",")
		ticket := []int{}
		for _, value := range values {
			ticket = append(ticket, atoi(value))
		}

		for _, value := range ticket {
			valueHasMatchingField := false
			for _, field := range fields {
				if valueHasMatchingField {
					break
				}
				for _, fieldRange := range field.Ranges {
					if fieldRange.contains(value) {
						valueHasMatchingField = true
						break
					}
				}
			}
			if !valueHasMatchingField {
				totalErrorRate += value
			}

		}

	}
	return totalErrorRate

}

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ComputeTicketScanningErrorRate(string(data)))
}
