package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	occupied = "#"
	empty    = "L"
	floor    = "."
)

func hashBoard(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func constructNewBoardFromExisting(existingBoard [][]string) [][]string {
	newBoard := make([][]string, len(existingBoard))
	for i := range existingBoard {
		newBoard[i] = make([]string, len(existingBoard[i]))
		copy(newBoard[i], existingBoard[i])
	}
	return newBoard
}

func countOccupiedSeats(board [][]string) int {
	result := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == occupied {
				result++
			}
		}
	}
	return result
}

func inBounds(board [][]string, row int, col int) bool {
	return row >= 0 && row < len(board) && col >= 0 && col < len(board[row])
}

func countOccupiedSeatsAdjacentToCell(board [][]string, row int, col int) int {
	result := 0
	adjacentCellDeltas := [][2]int{{0, -1}, {0, 1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 0}, {-1, 0}}
	for i := range adjacentCellDeltas {
		deltaX := adjacentCellDeltas[i][0]
		deltaY := adjacentCellDeltas[i][1]
		if inBounds(board, row+deltaX, col+deltaY) && board[row+deltaX][col+deltaY] == occupied {
			result++
		}
	}
	return result
}

func countOccupiedSeatsOnEquiliburumUtil(board [][]string) int {
	nextBoard := constructNewBoardFromExisting(board)

	for i := range board {
		for j := range board[i] {
			occupiedSeats := countOccupiedSeatsAdjacentToCell(board, i, j)
			if board[i][j] == occupied && occupiedSeats >= 4 {
				nextBoard[i][j] = empty
			} else if board[i][j] == empty && occupiedSeats == 0 {
				nextBoard[i][j] = occupied
			}
		}
	}
	if hashBoard(nextBoard) == hashBoard(board) {
		return countOccupiedSeats(board)
	}
	return countOccupiedSeatsOnEquiliburumUtil(nextBoard)

}

func CountOccupiedSeatsOnEquiliburum(input string) int {
	inputLines := strings.Split(input, "\n")
	board := make([][]string, len(inputLines))

	for i, line := range inputLines {
		line = strings.Trim(line, " ")
		if len(line) > 0 {
			board[i] = make([]string, len(line))
			for j, ch := range line {
				board[i][j] = string(ch)
			}
		}
	}
	return countOccupiedSeatsOnEquiliburumUtil(board)
}
func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(CountOccupiedSeatsOnEquiliburum(string(data)))
}
