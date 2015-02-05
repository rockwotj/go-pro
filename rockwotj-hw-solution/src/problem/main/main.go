package main

import (
        "problem/sudoku"
        "fmt"
)


func main() {
    board := sudoku.Board{}
    board.Initialize()
    
    // BOARD FOUND AT: https://www.rose-hulman.edu/class/csse/csse403/Fall2010/Homework/hw14.html

    board.Cells[1][0] = 5
    board.Cells[1][1] = 4
    board.Cells[2][1] = 1
    
    board.Cells[0][3] = 2
    board.Cells[0][5] = 7
    board.Cells[1][4] = 6
    board.Cells[2][3] = 9

    board.Cells[0][6] = 1
    board.Cells[0][7] = 3
    board.Cells[1][7] = 8
    board.Cells[2][7] = 5
    board.Cells[2][8] = 2

    board.Cells[3][0] = 9
    board.Cells[3][1] = 5
    board.Cells[4][2] = 6
    board.Cells[5][0] = 3

    board.Cells[3][4] = 8
    board.Cells[3][5] = 2
    board.Cells[5][3] = 5
    board.Cells[5][4] = 9

    board.Cells[3][8] = 6
    board.Cells[4][6] = 7
    board.Cells[5][7] = 4
    board.Cells[5][8] = 8

    board.Cells[6][0] = 8
    board.Cells[6][1] = 2
    board.Cells[7][1] = 3
    board.Cells[8][1] = 6
    board.Cells[8][2] = 5

    board.Cells[6][5] = 4
    board.Cells[7][4] = 2
    board.Cells[8][3] = 1
    board.Cells[8][5] = 9

    board.Cells[6][7] = 7
    board.Cells[7][7] = 6
    board.Cells[7][8] = 9
    fmt.Println("==Initial Board==")
    board.Print()
    solution := make(chan sudoku.Board)
    go sudokuSolver(board, solution)
    board = <-solution
    fmt.Println("=== Solution! ===")
    board.Print()
}

func sudokuSolver(b sudoku.Board, solution chan sudoku.Board) {
    for v := 1; v <= sudoku.BOARD_SIZE; v++ {
        board := b.Copy() 
        go sudokuSolverHelper(board, solution, 0, 0, v)
    }
}

func sudokuSolverHelper(b sudoku.Board, solution chan sudoku.Board, r int, c int, value int) {
    if r == 9 {
        solution <- b
        return
    }
    if !b.IsPositionValid(value, r, c) {
        return
    }
    row, col := nextCell(r, c)
    for v := 1; v <= sudoku.BOARD_SIZE; v++ {
         board := b.Copy()
         go sudokuSolverHelper(board, solution, row, col, v)
    }
}

func nextCell(r int, c int) (int, int) {
    col := c + 1
    if col == sudoku.BOARD_SIZE {
        return r + 1, 0
    }
    return r, col
}
