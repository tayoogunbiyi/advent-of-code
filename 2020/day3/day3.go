package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	treeMarker       string = "#"
	openSquareMarker string = "."
)

func inBounds(grid [][]string, row, col int) bool {
	return row < len(grid) && col < len(grid[row])
}

func countTreesAlongPathUtil(grid [][]string, row, col int, sl slope) int {
	if !inBounds(grid, row, col) {
		return 0
	}
	treesAtCurrentCell := 0
	if grid[row][col] == treeMarker {
		treesAtCurrentCell = 1
	}

	return treesAtCurrentCell + countTreesAlongPathUtil(grid, row+sl.Down, (col+sl.Right)%len(grid[row]), sl)
}

type slope struct {
	Right int
	Down  int
}

// ProductOfTreesAlongPath calculates the product of the total number of treeMarkers encountered while traversing the input path using each of the slopes given in []slope
func ProductOfTreesAlongPath(input string, slopes []slope) int {
	var grid [][]string

	for _, inputLine := range strings.Split(input, "\n") {
		inputLine = strings.Trim(inputLine, " ")
		if len(inputLine) > 0 {
			var row = []string{}
			for _, ch := range inputLine {
				row = append(row, string(ch))

			}
			grid = append(grid, row)
		}
	}
	totalProduct := 1
	for _, sl := range slopes {
		fmt.Println(countTreesAlongPathUtil(grid, 0, 0, sl))
		totalProduct *= countTreesAlongPathUtil(grid, 0, 0, sl)
	}
	return totalProduct
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	slopes := []slope{
		{Right: 1, Down: 1},
		{Right: 3, Down: 1},
		{Right: 5, Down: 1},
		{Right: 7, Down: 1},
		{Right: 1, Down: 2},
	}

	fmt.Println("The number of encountered trees along the path is ", ProductOfTreesAlongPath(string(data), slopes))
}
