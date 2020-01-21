package solvers

import (
	"testing"

	"github.com/AlbinoDrought/creamy-sudoku/sudoku"
)

func CheckSolver(t *testing.T, solver SudokuSolver) {
	puzzleFiles := []string{
		"./../puzzles/qwing1.csv",
		"./../puzzles/qwing2.csv",
		"./../puzzles/qwing3.csv",
		"./../puzzles/qwing-expert.csv",
	}

	for _, puzzleFile := range puzzleFiles {
		t.Run(puzzleFile, func(t *testing.T) {
			puzzles, err := sudoku.ImportCSV(puzzleFile)
			if err != nil {
				t.Error(err)
			}

			for _, puzzle := range puzzles {
				if puzzle.SolvedBoard == nil {
					continue
				}

				err = solver.Solve(puzzle.PuzzleBoard)
				if err != nil {
					t.Error(err)
				}

				for rowIndex, row := range puzzle.SolvedBoard.Numbers {
					for colIndex, value := range row {
						if puzzle.PuzzleBoard.Numbers[rowIndex][colIndex] != value {
							t.Fatalf("Expected %v but got %v at row %v col %v", value, puzzle.PuzzleBoard.Numbers[rowIndex][colIndex], rowIndex, colIndex)
						}
					}
				}
			}
		})
	}
}
