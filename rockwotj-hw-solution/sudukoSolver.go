package main

import "fmt"

const BOARD_SIZE int = 9

type Suduko interface {

}


func (b *Board) Initialize() {
	b.cells = make([][]int, BOARD_SIZE)
    for i := 0; i < BOARD_SIZE; i++ {
    	b.cells[i] = make([]int, BOARD_SIZE)
        for j := 0; j < BOARD_SIZE; j++ {
            b.cells[i][j] = 0
        }
    }
}

func (b *Board) Print() {
    for i := 0; i < BOARD_SIZE; i++ {
    	for j := 0; j < BOARD_SIZE; j++ {
            fmt.Print(b.cells[i][j], " ")
        }
        fmt.Println()
    }
}

func (b *Board) IsFull() bool {
	for i := 0; i < BOARD_SIZE; i++ {
    	for j := 0; j < BOARD_SIZE; j++ {
         	if b.cells[i][j] == 0 {
         		return false   	
         	}
        }
    }
    return true	
}

func (b Board) IsPositionValid(value int, row int, col int) bool {
	b.cells[row][col] = value
    return b.IsRowValid(row) && b.IsColValid(col) && b.IsSquareValid(row, col)
}

func (b Board) IsRowValid(r int) bool {
    return IsValidSequence(b.cells[r])
}

func (b Board) IsColValid(c int) bool {
    column := make([]int, BOARD_SIZE)
    for i := 0; i < BOARD_SIZE; i++ {
        column[i] = b.cells[i][c]
    }
    return IsValidSequence(column)
}

// r and c are the location of the changed cell
func (b Board) IsSquareValid(r int, c int) bool {
    c /= 3
    r /= 3
    c *= 3
    r *= 3
    square := append(b.cells[r][c:c + 3], b.cells[r + 1][c:c + 3]...)
    square = append(square, b.cells[r + 2][c:c + 3]...)
    return IsValidSequence(square)	
}

func IsValidSequence(slice []int) bool {
    nums := make([]bool, BOARD_SIZE + 1)
    for i := 0; i < BOARD_SIZE + 1; i++ {
        nums[i] = false
    }
    for _, e := range slice {
        if e == 0 {
            continue
        } else if nums[e] {
            return false
        } else {
            nums[e] = true
        }
    }
    return true
}

type Board struct {
	cells [][]int
}

func main() {
	suduko := Board{}
	suduko.Initialize()
    for i := 0; i < BOARD_SIZE; i++ {
        suduko.cells[0][i] = i + 1;
    }
    for i := 1; i < BOARD_SIZE; i++ {
        suduko.cells[i][0] = i + 1;
    }
    suduko.cells[8][8] = 1
    suduko.Print()
	fmt.Println(suduko.IsFull())
    fmt.Println(suduko.IsRowValid(0))
    fmt.Println(suduko.IsColValid(0))
    suduko.IsSquareValid(8, 5)
} 