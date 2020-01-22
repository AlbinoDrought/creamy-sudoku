package sudoku

import (
	"fmt"
	"math"
)

// A Board is a sudoku playing field
type Board struct {
	GridSize int
	Numbers  [][]byte
	Sticky   [][]bool
}

// Set a number on the board
func (b *Board) Set(row, col int, number byte) {
	b.Numbers[row][col] = number
}

// Print the board to standard out
func (b *Board) Print() {
	for _, row := range b.Numbers {
		fmt.Printf("%+v\n", row)
	}
}

// Valid returns true if the number could be used at the given row/col
func (b *Board) Valid(row, col int, number byte) bool {
	// check square
	sqrt := int(math.Sqrt(float64(b.GridSize)))
	rowCube := row / sqrt
	rowCubeStart := rowCube * sqrt
	colCube := col / sqrt
	colCubeStart := colCube * sqrt

	// check row/col
	for i := 0; i < b.GridSize; i++ {
		cubeRow := rowCubeStart + (i / sqrt)
		cubeCol := colCubeStart + (i % sqrt)

		if cubeRow != row && cubeCol != col && b.Numbers[cubeRow][cubeCol] == number {
			return false
		}

		if i != row && b.Numbers[i][col] == number {
			return false
		}

		if i != col && b.Numbers[row][i] == number {
			return false
		}
	}

	return true
}

// Solved returns true if the board is solved, but very slowly...
func (b *Board) Solved() bool {
	maxNumberSize := byte(b.GridSize % 255)
	for row, rowValues := range b.Numbers {
		for col, colValue := range rowValues {
			if colValue == 0 || colValue > maxNumberSize {
				return false
			}

			if !b.Valid(row, col, colValue) {
				return false
			}
		}
	}

	return true
}

// MakeBoard initializes a board of the given size and returns it
func MakeBoard(size int) *Board {
	board := &Board{
		GridSize: size,
		Numbers:  make([][]byte, size),
		Sticky:   make([][]bool, size),
	}

	for i := 0; i < size; i++ {
		board.Numbers[i] = make([]byte, size)
		board.Sticky[i] = make([]bool, size)
	}

	return board
}
