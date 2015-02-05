package sudoku

import "fmt"

/* START BOARD FUNCTIONS */

const BOARD_SIZE int = 9

type Board struct {
    Cells [][]int
}

func (b *Board) Initialize() {
	b.Cells = make([][]int, BOARD_SIZE)
    for i := 0; i < BOARD_SIZE; i++ {
    	b.Cells[i] = make([]int, BOARD_SIZE)
        for j := 0; j < BOARD_SIZE; j++ {
            b.Cells[i][j] = 0
        }
    }
}

func (b *Board) Print() {
    for i := 0; i < BOARD_SIZE; i++ {
    	for j := 0; j < BOARD_SIZE; j++ {
            fmt.Print(b.Cells[i][j], " ")
        }
        fmt.Println()
    }
    fmt.Println()
}

func (b *Board) Copy() Board {
    board := Board{}
    board.Initialize()
    for i := 0; i < BOARD_SIZE; i++ {
        for j := 0; j < BOARD_SIZE; j++ {
            board.Cells[i][j] = b.Cells[i][j]
        }
    }
    return board
}

func (b *Board) IsFull() bool {
	for i := 0; i < BOARD_SIZE; i++ {
    	for j := 0; j < BOARD_SIZE; j++ {
         	if b.Cells[i][j] == 0 {
         		return false   	
         	}
        }
    }
    return true	
}

func (b *Board) IsPositionValid(value int, row int, col int) bool {
	if b.Cells[row][col] != 0 {
        return true
    }
    b.Cells[row][col] = value
    return b.IsRowValid(row) && b.IsColValid(col) && b.IsSquareValid(row, col)
}

func (b *Board) IsRowValid(r int) bool {
    return IsValidSequence(b.Cells[r])
}

func (b *Board) IsColValid(c int) bool {
    column := make([]int, BOARD_SIZE)
    for i := 0; i < BOARD_SIZE; i++ {
        column[i] = b.Cells[i][c]
    }
    return IsValidSequence(column)
}

func (b *Board) IsSquareValid(r int, c int) bool {
    c /= 3
    r /= 3
    c *= 3
    r *= 3
    square := make([]int, 0)
    square = append(square, b.Cells[r][c:c + 3]...)
    square = append(square, b.Cells[r + 1][c:c + 3]...)
    square = append(square, b.Cells[r + 2][c:c + 3]...)
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
/* END BOARD FUNCTIONS */
