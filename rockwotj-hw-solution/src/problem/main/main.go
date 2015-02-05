package main

import (
        "fmt"
        "problem/suduko"
)

func main() {
    board := suduko.Board{}
    board.Initialize()
    for i := 0; i < suduko.BOARD_SIZE; i++ {
        board.Cells[0][i] = i + 1;
    }
    for i := 1; i < suduko.BOARD_SIZE; i++ {
        board.Cells[i][0] = i + 1;
    }
    board.Cells[8][8] = 1
    board.Print()
    fmt.Println(board.IsFull())
    fmt.Println(board.IsRowValid(0))
    fmt.Println(board.IsColValid(0))
    board.IsSquareValid(8, 5)
} 
