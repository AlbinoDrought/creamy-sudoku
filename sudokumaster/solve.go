package main

import "C"
import "unsafe"
import "github.com/AlbinoDrought/creamy-sudoku/solvers"
import "github.com/AlbinoDrought/creamy-sudoku/sudoku"

const boardSize = 81

func pointerToBoard(start *C.int) *sudoku.Board {
	pointer := unsafe.Pointer(start)
	size := unsafe.Sizeof(C.int(0))

	board := sudoku.MakeBoard(9)
	
	row := 0
	col := 0

	for i := 0; i < boardSize; i++ {
		item := byte(*(*C.int)(unsafe.Pointer(uintptr(pointer) + size*uintptr(i))))

		row = i / board.GridSize
		col = i % board.GridSize

		if item > 0 {
			board.Set(row, col, item)
			board.Sticky[row][col] = true
		}
	}

	return board
}

func boardToPointer(start *C.int, board *sudoku.Board) {
	/*
	nums := make([]C.int, boardSize)

	row := 0
	col := 0
	for i := 0; i < boardSize; i++ {
		row = i / board.GridSize
		col = i % board.GridSize

		nums[i] = C.int(board.Numbers[row][col])
	}

	return C.int(uintptr(unsafe.Pointer(&nums)))
	*/


	pointer := unsafe.Pointer(start)
	size := unsafe.Sizeof(C.int(0))

	row := 0
	col := 0
	for i := 0; i < boardSize; i++ {
		row = i / board.GridSize
		col = i % board.GridSize

		*(*C.int)(unsafe.Pointer(uintptr(pointer) + size*uintptr(i))) = C.int(board.Numbers[row][col])
	}

}

//export solve
func solve(start *C.int) C.int {
	board := pointerToBoard(start)
	solver := &solvers.LoopierBruteforceSolver{}
	solver.Solve(board)
	boardToPointer(start, board)
	return 0
}
