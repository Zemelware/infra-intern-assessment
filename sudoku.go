package main

import (
	"fmt"
)

// TODO:
// 	- document functions
// 	- write comments explaining code

var puzzle = [][]int{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

// SolveSudoku takes a Sudoku puzzle and returns the solved puzzle
func SolveSudoku(puzzle [][]int) [][]int {
	// Contains positions of all numbers 1-9, in the following format:
	// {
	// 	number: [[row, col], [row, col], ...],
	// 	...
	// }
	// Example:
	// {
	// 	1: [[0, 0], [3, 1], ...],
	// 	2: [[4, 5], [6, 6], ...],
	// 	...
	// }
	var pos map[int][][]int

	// Count of the number of remaining occurrences of each number, in the following format:
	// {
	// 	number: remaining count,
	// 	...
	// }
	// Example:
	// {
	// 	1: 7,
	// 	2: 5,
	// 	...
	// }
	var remaining map[int]int

	// Graph representing the valid positions for each number, stored in the following format:
	// Format:
	// {
	// 	number: {
	// 		row_i: [col_i, ...],
	// 		row_j: [col_j, ...],
	// 		...
	// 	},
	// 	...
	// }
	// Example:
	// {
	// 	1: {
	// 		0: [0, 1, 2, 6, 7],
	// 		3: [0, 1, 2, 6, 7],
	// 		...
	// 	},
	// 	...
	// }
	var graph map[int]map[int][]int

	// pos, remaining = populatePosAndRemaining(puzzle)

	// TODO: Sort the remaining map in ascending order by value

	// graph = populateGraph(puzzle)

	printPuzzle(puzzle)

	solvedPuzzle := fillPuzzle(puzzle, pos, remaining, graph)

	printPuzzle(solvedPuzzle)
	
	return solvedPuzzle
}

// fillPuzzle fills in the Sudoku puzzle with the correct numbers using a backtracking algorithm with crosshatching
func fillPuzzle(puzzle [][]int, pos map[int][][]int, remaining map[int]int, graph map[int]map[int][]int) [][]int {
	// Add logic
	return puzzle
}

// validPlacement checks if the number in the position puzzle[row][col] is in a valid spot according to the rules of Sudoku
func validPlacement(puzzle [][]int, row int, col int) bool {
	num := puzzle[row][col]
	// Check if the number is in the same row or column
	for i := 0; i < len(puzzle); i++ {
		// Check if any numbers in the same row equal num
		if puzzle[row][i] == num && i != col {
			return false
		}
		// Check if any numbers in the same column equal num
		if puzzle[i][col] == num && i != row {
			return false
		}
	}

	// Check if there are any duplicate numbers in the 3x3 box
	boxRowStart := row / 3 * 3
	boxColStart := col / 3 * 3

	for i := boxRowStart; i < boxRowStart + 3; i++ {
		for j := boxColStart; j < boxColStart + 3; j++ {
			if puzzle[i][j] == num && i != row && j != col {
				return false
			}
		}
	}

	return true
}

// populatePosAndRemaining populates the pos and remaining maps with the appropriate values
// func populatePosAndRemaining(puzzle [][]int) (map[int][][]int, map[int]int) {
// }

// populateGraph populates the graph map with the appropriate values
// func populateGraph(puzzle [][]int) map[int]map[int][]int {
// }

// printPuzzle prints a Sudoku puzzle
func printPuzzle(puzzle [][]int) {
	for row := 0; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[row]); col++ {
			fmt.Printf("%d ", puzzle[row][col])
		}
		fmt.Println()
	}
}

func main() {
	SolveSudoku(puzzle)
}