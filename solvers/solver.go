package solvers

import "github.com/AlbinoDrought/creamy-sudoku/sudoku"

type SudokuSolver interface {
	Solve(board *sudoku.Board) error
}
