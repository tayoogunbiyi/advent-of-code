// https://adventofcode.com/2020/day/7
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Bag struct {
	colour  string
	content map[string]string
}

type SubBag struct {
	Count  int
	Colour string
}

// func (b Bag) isSame(b2 Bag) bool {
// 	return b.colour == b2.colour
// }

func ParseInput(input string) map[string]map[string]int {
	var bags = map[string]map[string]int{}

	lines := strings.Split(input, "\n")
	for _, line := range lines[:len(lines)-1] {
		words := strings.Split(line, " ")
		key := words[0] + " " + words[1]
		bags[key] = map[string]int{}

		for i := range words {
			if count, err := strconv.Atoi(words[i]); err == nil {
				subBagColor := words[i+1] + " " + words[i+2]
				bags[key][subBagColor] = count
			}
		}
	}
	return bags
}

func dfs(bags map[string]map[string]int, currentSubBags map[string]int, targetBag string) bool {
	if _, ok := currentSubBags[targetBag]; ok {
		return true
	}
	for k := range currentSubBags {
		if dfs(bags, bags[k], targetBag) {
			return true
		}
	}
	return false

}

func CountBagsThatEventuallyContain(input string, targetBagColor string) int {
	bags := ParseInput(input)
	count := 0

	for _, v := range bags {
		if dfs(bags, v, targetBagColor) {
			count += 1
		}
	}
	return count
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the number of bag colors that can eventually contain >= 1 shiny gold bag is", CountBagsThatEventuallyContain(string(data), "shiny gold"))

}
