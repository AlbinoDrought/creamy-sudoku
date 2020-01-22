package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/AlbinoDrought/creamy-sudoku/solvers"
	"github.com/AlbinoDrought/creamy-sudoku/sudoku"
)

const input = "puzzles/17-blank-tips.tsv"
const runs = 10

func main() {
	solvers := []solvers.SudokuSolver{
		&solvers.BruteforceSolver{},
		&solvers.LoopierBruteforceSolver{},
	}

	for _, solver := range solvers {
		benchmark := time.Duration(0)

		for i := 0; i < runs; i++ {
			board, err := sudoku.ImportTSV(input)
			if err != nil {
				panic(err)
			}

			start := time.Now()
			solver.Solve(board)
			end := time.Now()

			difference := end.Sub(start)
			benchmark += difference
		}

		fmt.Printf("%v: %vms (%vms/puzzle)\n", reflect.TypeOf(solver).String(), benchmark.Milliseconds(), benchmark.Milliseconds()/runs)
	}
}
