package sudoku

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type BoardPair struct {
	PuzzleBoard *Board
	SolvedBoard *Board
}

func ImportString(serialized string) (*Board, error) {
	gridSize := int(math.Sqrt(float64(len(serialized))))

	board := MakeBoard(gridSize)

	var (
		row int
		col int
		err error
		num int
	)

	for i, char := range serialized {
		if char == '.' {
			continue
		}

		row = i / board.GridSize
		col = i % board.GridSize

		num, err = strconv.Atoi(string(char))
		if err != nil {
			return nil, err
		}

		board.Set(row, col, byte(num))
		board.Sticky[row][col] = true
	}

	return board, nil
}

func ImportCSV(path string) ([]BoardPair, error) {
	inputBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	inputString := string(inputBytes)
	inputRows := strings.Split(inputString, "\n")
	rows := make([]BoardPair, len(inputRows))

	for rowIndex := range rows {
		columns := strings.Split(inputRows[rowIndex], ",")

		rows[rowIndex].PuzzleBoard, err = ImportString(columns[0])
		if err != nil {
			return nil, err
		}

		if len(columns) > 1 {
			rows[rowIndex].SolvedBoard, err = ImportString(columns[1])
			if err != nil {
				return nil, err
			}
		}
	}

	return rows, nil
}
