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

type boardUpdateFunc func(originalBoard [][]string, updatedBoard [][]string)

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

func countVisibleOccupiedSeats(board [][]string, row int, col int) int {
	result := 0
	adjacentCellDeltas := [][2]int{{0, -1}, {0, 1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 0}, {-1, 0}}
	for i := range adjacentCellDeltas {
		deltaX := adjacentCellDeltas[i][0]
		deltaY := adjacentCellDeltas[i][1]
		targetRow := row + deltaX
		targetCol := col + deltaY
		for inBounds(board, targetRow, targetCol) && board[targetRow][targetCol] == floor {
			targetRow += deltaX
			targetCol += deltaY
		}
		if inBounds(board, targetRow, targetCol) && board[targetRow][targetCol] == occupied {
			result++
		}
	}
	return result
}

func countOccupiedSeatsAtEquiliburumUtil(board [][]string, updateFunc boardUpdateFunc) int {
	nextBoard := constructNewBoardFromExisting(board)

	updateFunc(board, nextBoard)

	if hashBoard(nextBoard) == hashBoard(board) {
		return countOccupiedSeats(board)
	}
	return countOccupiedSeatsAtEquiliburumUtil(nextBoard, updateFunc)
}

func parseInput(input string) [][]string {
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
	return board
}
func CountOccupiedSeatsAtEquiliburum(input string, updateFunc boardUpdateFunc) int {
	board := parseInput(input)
	return countOccupiedSeatsAtEquiliburumUtil(board, updateFunc)
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	updateFunc1 := func(originalBoard [][]string, updatedBoard [][]string) {
		for i := range originalBoard {
			for j := range updatedBoard[i] {
				occupiedSeats := countOccupiedSeatsAdjacentToCell(originalBoard, i, j)
				if originalBoard[i][j] == occupied && occupiedSeats >= 4 {
					updatedBoard[i][j] = empty
				} else if originalBoard[i][j] == empty && occupiedSeats == 0 {
					updatedBoard[i][j] = occupied
				}
			}
		}
	}

	updateFunc2 := func(originalBoard [][]string, updatedBoard [][]string) {
		for i := range originalBoard {
			for j := range updatedBoard[i] {
				occupiedSeats := countVisibleOccupiedSeats(originalBoard, i, j)
				if originalBoard[i][j] == occupied && occupiedSeats >= 5 {
					updatedBoard[i][j] = empty
				} else if originalBoard[i][j] == empty && occupiedSeats == 0 {
					updatedBoard[i][j] = occupied
				}
			}
		}
	}

	fmt.Println(CountOccupiedSeatsAtEquiliburum(string(data), updateFunc1))
	fmt.Println(CountOccupiedSeatsAtEquiliburum(string(data), updateFunc2))

}
