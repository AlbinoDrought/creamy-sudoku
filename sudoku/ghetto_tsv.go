package sudoku

import (
	"io/ioutil"
	"strconv"
)

// ImportTSV converts a TSV board representation to a Board struct
func ImportTSV(path string) (*Board, error) {
	inputBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	size := 0
	for _, char := range inputBytes {
		if char == '\t' {
			size++
		} else if char == '\n' {
			if size > 0 {
				size++
			}
			break
		}
	}

	board := MakeBoard(size)

	row := 0
	col := 0
	num := 0
	for _, char := range inputBytes {
		if char == '\t' {
			col++
		} else if char == '\n' {
			col = 0
			row++
		} else if char == '\r' {
			continue
		} else {
			num, err = strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			board.Numbers[row][col] = byte(num)
			board.Sticky[row][col] = true
		}
	}

	return board, nil
}
