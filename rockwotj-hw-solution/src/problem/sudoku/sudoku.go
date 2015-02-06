package sudoku

import (
    "fmt"
    "bytes"
    "strconv"
)

const BOARD_SIZE int = 9

type Board struct {
    Cells [][]int
}

func NextCell(r int, c int) (int, int) {
    col := c + 1
    if col == BOARD_SIZE {
        return r + 1, 0
    }
    return r, col
}

func BoardFromString(s string) Board {
    board := Board{}
    board.Initialize()
    row := 0
    col := 0
    for _, c := range s {
        switch c {
        case '1':
            board.Cells[row][col] = 1
        case '2':
            board.Cells[row][col] = 2
        case '3':
            board.Cells[row][col] = 3
        case '4':
            board.Cells[row][col] = 4
        case '5':
            board.Cells[row][col] = 5
        case '6':
            board.Cells[row][col] = 6
        case '7':
            board.Cells[row][col] = 7
        case '8':
            board.Cells[row][col] = 8
        case '9':
            board.Cells[row][col] = 9
        }
        row, col = NextCell(row, col)
    }
    return board
}

/* START BOARD FUNCTIONS */

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
            value := b.Cells[i][j]
            if value == 0 {
                fmt.Print(". ")
            } else {
                fmt.Print(value, " ")
            }
            if j == 2 || j == 5 {
                fmt.Print("| ")
            }
        }
        fmt.Println()
        if i == 2 || i == 5 {
            fmt.Println("------+-------+------")
        }
    }
    fmt.Println()
}

func (b *Board) ToString() string {
    var buffer bytes.Buffer
    for i := 0; i < BOARD_SIZE; i++ {
        for j := 0; j < BOARD_SIZE; j++ {
            buffer.WriteString(strconv.Itoa(b.Cells[i][j]))
        }
    }
    return buffer.String()
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

// WARNING: Side effects, modifies the board if that spot is not taken.
// Even if it is an invalid position
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
