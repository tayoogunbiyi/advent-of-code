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

func ParseInput(input string) ([]Field, []int, [][]int) {
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

	currentLineNumber += 3
	yourTicket := []int{}

	for i := currentLineNumber; i < len(lines) && len(lines[i]) != 0; i++ {
		values := strings.Split(lines[i], ",")
		for _, value := range values {
			yourTicket = append(yourTicket, atoi(value))
		}
		currentLineNumber = i
	}

	currentLineNumber += 3
	nearbyTickets := [][]int{}
	for i := currentLineNumber; i < len(lines) && len(lines[i]) != 0; i++ {
		values := strings.Split(lines[i], ",")
		ticket := []int{}
		for _, value := range values {
			ticket = append(ticket, atoi(value))
		}
		nearbyTickets = append(nearbyTickets, ticket)
	}
	return fields, yourTicket, nearbyTickets

}

func RemoveIndex(s *[]Field, index int) []Field {
	ss := *s
	return append(ss[:index], ss[index+1:]...)
}

func ComputeTicketScanningErrorRate(input string) int {
	fields, _, nearbyTickets := ParseInput(input)
	ht := make(map[int]bool)
	totalErrorRate := 0
	for _, ticket := range nearbyTickets {
		for i, value := range ticket {
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
				ht[i] = false
				totalErrorRate += value
			} else {
				ht[i] = true
			}
		}
	}
	return totalErrorRate
}

func MapContains(ht map[string]int, target int) bool {
	for _, v := range ht {
		if v == target {
			return true
		}
	}
	return false
}

func isSubset(a []int, b []int) bool {
	ht := make(map[int]bool)
	for _, v := range b {
		ht[v] = true
	}
	for _, v := range a {
		if _, ok := ht[v]; !ok {
			return false
		}
	}
	return true
}

func transpose(mat [][]int) [][]int {
	result := [][]int{}
	for i := 0; i < len(mat[0]); i++ {
		row := []int{}
		added := make(map[int]bool)
		for j := 0; j < len(mat); j++ {
			if _, ok := added[mat[j][i]]; !ok {
				row = append(row, mat[j][i])
				added[mat[j][i]] = true
			}
		}
		result = append(result, row)
	}
	return result
}

func includes(arr []int, v int) bool {
	for _, i := range arr {
		if i == v {
			return true
		}
	}
	return false
}

func FindFieldOrdering(input string) int {
	fields, yourTicket, nearbyTicket := ParseInput(input)
	validnearbyTickets := [][]int{}
	numbers := []int{}
	for _, f := range fields {
		for _, fieldRange := range f.Ranges {
			for i := fieldRange.Start; i <= fieldRange.End; i++ {
				if !includes(numbers, i) {
					numbers = append(numbers, i)
				}
			}
		}
	}

	for _, ticket := range nearbyTicket {
		valid := true
		for _, v := range ticket {
			if !includes(numbers, v) {
				valid = false
				break
			}
		}
		if valid {
			validnearbyTickets = append(validnearbyTickets, ticket)
		}
	}
	validnearbyTicketsTranspose := transpose(validnearbyTickets)
	definiteFields := make(map[string]int)
	for len(definiteFields) < len(validnearbyTicketsTranspose) {
		possibleFields := make(map[string][]int)
		for _, field := range fields {
			values := []int{}
			for _, fieldRange := range field.Ranges {
				for j := fieldRange.Start; j <= fieldRange.End; j++ {
					values = append(values, j)
				}
			}
			for i := 0; i < len(validnearbyTicketsTranspose); i++ {
				if !MapContains(definiteFields, i) && isSubset(validnearbyTicketsTranspose[i], values) {
					if !includes(possibleFields[field.Name], i) {
						possibleFields[field.Name] = append(possibleFields[field.Name], i)
					}
				}
			}
		}

		for name, values := range possibleFields {
			if len(values) == 1 {
				definiteFields[name] = values[0]
				possibleFields[name] = append(values[:0], values[1:]...)

				for i, field := range fields {
					if field.Name == name {
						RemoveIndex(&fields, i)
						break
					}
				}
			}
		}
	}

	result := 1
	for k, v := range definiteFields {
		if k[0:3] == "dep" {
			result *= yourTicket[v]
		}
	}
	return result

}
func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ComputeTicketScanningErrorRate(string(data)))
	fmt.Println(FindFieldOrdering(string(data)))
}
