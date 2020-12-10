// https://adventofcode.com/2020/day/6#
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func isValidLine(line string) bool {
	return len(line) > 0
}

type group struct {
	MemberData map[int]string
}

func (g group) update(groupID int, data string) {
	g.MemberData[groupID] += data
}

func countUniqueCharacters(s string) int {
	ht := make(map[rune]int)
	for _, ch := range s {
		ht[ch]++
	}
	return len(ht)

}

func CountNumberOfQuestionsAnsweredByAnyone(input string) int {
	groups := buildGroups(input)
	result := 0
	for _, g := range groups {
		for _, v := range g.MemberData {
			result += countUniqueCharacters(v)
		}
	}
	return result
}

func buildGroups(input string) []group {
	groups := []group{{MemberData: make(map[int]string)}}
	inputLines := strings.Split(input, "\n")
	for groupID, line := range inputLines {
		if isValidLine(line) {
			groups[len(groups)-1].update(len(groups)-1, line)
		} else {
			if groupID != len(inputLines)-1 {
				groups = append(groups, group{MemberData: make(map[int]string)})
			}
		}
	}
	return groups
}

/*

ab
ac

*/

func findIntersection(questions ...string) string {
	result := questions[0]
	fmt.Println(questions)

	for i := 1; i < len(questions); i++ {
		currentQuestion := questions[i]
		currentIntersection := ""
		for _, ch := range currentQuestion {
			if strings.Contains(result, string(ch)) {
				currentIntersection += string(ch)
			}
		}
		fmt.Println(currentIntersection)
		result = currentIntersection
	}

	return result
}


func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the total number of questions answered by anyone within each group is %d\n", CountNumberOfQuestionsAnsweredByAnyone(string(data)))
	
}
