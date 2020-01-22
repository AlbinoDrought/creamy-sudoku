package main

import (
	"sync"

	"github.com/AlbinoDrought/creamy-sudoku/solvers"
	"github.com/AlbinoDrought/creamy-sudoku/sudoku"
)

const input = "puzzles/17-blank-tips.tsv"
const runs = 25
const parallel = true

func main() {
	board, err := sudoku.ImportTSV(input)
	if err != nil {
		panic(err)
	}
	board.Print()

	sudokuSolvers := []solvers.SudokuSolver{
		&solvers.BruteforceSolver{},
		&solvers.LoopierBruteforceSolver{},
	}

	waitGroup := sync.WaitGroup{}

	for _, solver := range sudokuSolvers {
		waitGroup.Add(1)

		go func(solver solvers.SudokuSolver) {
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
			waitGroup.Done()
		}(solver)

		if !parallel {
			waitGroup.Wait()
		}
	}

	waitGroup.Wait()
}
