package sudoku

import "testing"

func TestSolved(t *testing.T) {
	puzzles := []struct {
		path   string
		solved bool
	}{
		{"./../puzzles/wiki.tsv", false},
		{"./../puzzles/wiki-solved.tsv", true},
	}

	var (
		err   error
		board *Board
	)

	for _, puzzle := range puzzles {
		board, err = ImportTSV(puzzle.path)
		if err != nil {
			t.Error(err)
		}

		if solved := board.Solved(); solved != puzzle.solved {
			t.Fatalf("board.Solved() returned %v but we expected %v", solved, puzzle.solved)
		}
	}
}
