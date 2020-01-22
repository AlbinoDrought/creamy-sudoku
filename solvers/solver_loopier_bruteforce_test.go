package solvers

import "testing"

func TestLoopierBruteforceSolver(t *testing.T) {
	CheckSolver(t, &LoopierBruteforceSolver{}) // 20.460s
}
