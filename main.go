package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/AlbinoDrought/creamy-sudoku/solvers"
	"github.com/AlbinoDrought/creamy-sudoku/sudoku"
)

const input = "puzzles/17-blank-tips.tsv"
const runs = 25

type runResult struct {
	Start time.Time
	End   time.Time
}

func (result *runResult) Duration() time.Duration {
	return result.End.Sub(result.Start)
}

type benchmarkResult struct {
	Solver     solvers.SudokuSolver
	Benchmark  time.Duration
	Runs       int
	RunResults []*runResult
}

func (result *benchmarkResult) LogRun(run *runResult) {
	result.RunResults = append(result.RunResults, run)
	result.Runs = len(result.RunResults)
	result.Benchmark += run.Duration()
}

func (result *benchmarkResult) Bench(callback func()) *runResult {
	run := &runResult{
		Start: time.Now(),
	}
	callback()
	run.End = time.Now()
	result.LogRun(run)
	return run
}

func (result *benchmarkResult) Print() {
	fmt.Printf(
		"%v: %vms (%vms/puzzle)\n",
		reflect.TypeOf(result.Solver).String(),
		result.Benchmark.Milliseconds(),
		result.Benchmark.Milliseconds()/int64(result.Runs),
	)
}

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
			Solver: solver,
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
