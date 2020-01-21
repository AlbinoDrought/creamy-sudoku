package solvers

import "github.com/AlbinoDrought/creamy-sudoku/sudoku"

// BruteforceSolver solves a sudoku board using raw power
type BruteforceSolver struct{}

// Solve a sudoku power with sheer force
func (solver *BruteforceSolver) Solve(board *sudoku.Board) error {
	byteGridSize := byte(board.GridSize % 255)
	i := 0
	row := 0
	col := 0
	for i < board.GridSize*board.GridSize {
		row = i / board.GridSize
		col = i % board.GridSize

		// skip the unchangeable numbers
		if board.Sticky[row][col] {
			i++
			continue
		}

		if board.Numbers[row][col] >= byteGridSize {
			// we must backtrack
			board.Numbers[row][col] = 0

			for {
				i--
				row = i / board.GridSize
				col = i % board.GridSize
				if !board.Sticky[row][col] {
					break
				}
			}
			// restart at this backtracked row/col position
			continue
		}

		if board.Numbers[row][col] < byteGridSize {
			board.Numbers[row][col]++
			if board.Valid(row, col, board.Numbers[row][col]) {
				i++
			}
		}
	}

	return nil
}
