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

type member struct {
	answers string
}
type group struct {
	members []member
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

	for _, group := range groups {
		groupAnswers := ""
		for _, member := range group.members {
			groupAnswers += member.answers
		}
		result += countUniqueCharacters(groupAnswers)
	}
	return result
}

func CountNumberOfQuestionsAnsweredByEveryone(input string) int {
	groups := buildGroups(input)
	result := 0
	for _, group := range groups {
		groupAnswers := make([]string, 0)
		for _, member := range group.members {
			groupAnswers = append(groupAnswers, member.answers)
		}
		result += len(findIntersection(groupAnswers...))
	}

	return result

}

func buildGroups(input string) []group {
	groups := []group{{members: []member{}}}

	inputLines := strings.Split(input, "\n")
	for i, line := range inputLines {
		if isValidLine(line) {
			newMember := member{answers: line}
			groups[len(groups)-1].members = append(groups[len(groups)-1].members, newMember)
		} else {
			if i != len(inputLines)-1 {
				newGroup := group{members: []member{}}
				groups = append(groups, newGroup)
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

	for i := 1; i < len(questions) && len(result) > 0; i++ {
		currentQuestion := questions[i]
		currentIntersection := ""
		for _, ch := range currentQuestion {
			if strings.Contains(result, string(ch)) {
				currentIntersection += string(ch)
			}
		}
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
	fmt.Printf("the total number of questions answered by everyone within each group is %d\n", CountNumberOfQuestionsAnsweredByEveryone(string(data)))

}
