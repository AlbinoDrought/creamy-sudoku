package sudoku

import (
	"fmt"
	"math"
)

// A Board is a sudoku playing field
type Board struct {
	GridSize int
	BoxSize  int
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
	rowCubeStart := (row / b.BoxSize) * b.BoxSize
	colCubeStart := (col / b.BoxSize) * b.BoxSize
	cubeRow := 0
	cubeCol := 0

	// check row/col
	for i := 0; i < b.GridSize; i++ {
		cubeRow = rowCubeStart + (i / b.BoxSize)
		cubeCol = colCubeStart + (i % b.BoxSize)

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
		// generally 9 for normal puzzles
		GridSize: size,
		// generally 3 for normal puzzles
		BoxSize: int(math.Sqrt(float64(size))),
		Numbers: make([][]byte, size),
		Sticky:  make([][]bool, size),
	}

	for i := 0; i < size; i++ {
		board.Numbers[i] = make([]byte, size)
		board.Sticky[i] = make([]bool, size)
	}

	return board
}
