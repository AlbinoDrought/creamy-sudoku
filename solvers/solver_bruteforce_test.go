package solvers

import "testing"

func TestBruteforceSolver(t *testing.T) {
	CheckSolver(t, &BruteforceSolver{})
}
