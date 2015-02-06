package main

import (
        "problem/sudoku"
        "fmt"
)

var routines int

func main() {
    board := sudoku.BoardFromString("")
    routines = 0
    fmt.Println("=== Initial Board ===")
    board.Print()
    solution := make(chan sudoku.Board)
    sudokuSolver(board, solution)
    board = <-solution
    fmt.Println("====  Solution!  ====")
    board.Print()
    fmt.Println()
    fmt.Println("This problem created ", routines, " routines")
    fmt.Println(board.ToString())
}

func sudokuSolver(b sudoku.Board, solution chan sudoku.Board) {
    if b.Cells[0][0] != 0 {
        // We just start one thread
        routines++
        go sudokuSolverHelper(b.Copy(), solution, 0, 0, 0)
    } else {
        // Spawn all possibilities
        for v := 1; v <= sudoku.BOARD_SIZE; v++ {
            board := b.Copy()
            routines++
            go sudokuSolverHelper(board, solution, 0, 0, v)
        }  
    }
}

func sudokuSolverHelper(b sudoku.Board, solution chan sudoku.Board, r int, c int, value int) {
    if !b.IsPositionValid(value, r, c) {
        return
    }
    row, col := sudoku.NextCell(r, c)
    if row == 9 {
        solution <- b
        return
    }
    if b.Cells[row][col] != 0 {
        sudokuSolverHelper(b, solution, row, col, 0)
    } else {
        for v := 1; v <= sudoku.BOARD_SIZE; v++ {
            board := b.Copy()
            routines++
            go sudokuSolverHelper(board, solution, row, col, v)
        }  
    }
}

