package main

import (
	"github.com/AlbinoDrought/creamy-sudoku/solvers"
	"github.com/AlbinoDrought/creamy-sudoku/sudoku"
)

const input = "puzzles/17-blank-tips.tsv"
const runs = 25

func main() {
	board, err := sudoku.ImportTSV(input)
	if err != nil {
		panic(err)
	}
	board.Print()

	solvers := []solvers.SudokuSolver{
		&solvers.BruteforceSolver{},
		&solvers.LoopierBruteforceSolver{},
	}

	for _, solver := range solvers {
		benchmark := &benchmarkResult{
			Target: solver,
			Runs:   runs,
		}

		for i := 0; i < runs; i++ {
			board, err := sudoku.ImportTSV(input)
			if err != nil {
				panic(err)
			}

			benchmark.Bench(func() {
				solver.Solve(board)
			})
		}

		benchmark.Print()
	}
}
