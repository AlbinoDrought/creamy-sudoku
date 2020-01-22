package solvers

import (
	"errors"

	"github.com/AlbinoDrought/creamy-sudoku/sudoku"
)

// LoopierBruteforceSolver solves a sudoku board using raw power but more loops
type LoopierBruteforceSolver struct{}

// Solve a sudoku power with sheer force and loops
func (solver *LoopierBruteforceSolver) Solve(board *sudoku.Board) error {
	byteGridSize := byte(board.GridSize % 255)
	i := 0
	row := 0
	col := 0
	found := false
	for i < board.GridSize*board.GridSize {
		row = i / board.GridSize
		col = i % board.GridSize

		// skip the unchangeable numbers
		if board.Sticky[row][col] {
			i++
			continue
		}

		found = false
		for guess := board.Numbers[row][col] + 1; guess <= byteGridSize; guess++ {
			if board.Valid(row, col, guess) {
				board.Set(row, col, guess)
				i++
				found = true
				break
			}
		}

		if !found {
			board.Set(row, col, 0)

			for {
				i--

				if i < 0 {
					return errors.New("Unsolvable puzzle?")
				}

				row = i / board.GridSize
				col = i % board.GridSize
				if !board.Sticky[row][col] {
					break
				}
			}
		}
	}

	return nil
}
