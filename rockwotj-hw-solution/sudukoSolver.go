package main

import "fmt"

type Suduko interface {
	IsFull() (full bool, r, c int) 
	IsValid() bool 
	ValidGuesses(r, c int) 
	Solve() (result *Board, success bool) 
	ConcSolve() (result *Board, success bool) 
}


func (b *Board) Initialize() {
	b.cells = make([][]int, 9)
    for i := 0; i < 9; i++ {
    	b.cells[i] = make([]int, 9)
        for j := 0; j < 9; j++ {
            b.cells[i][j] = 0
        }
    }
}

func (b *Board) Print() {
    for i := 0; i < 9; i++ {
    	for j := 0; j < 9; j++ {
            fmt.Print(b.cells[i][j], " ")
        }
        fmt.Println()
    }
}

func (b *Board) IsFull() bool {
	for i := 0; i < 9; i++ {
    	for j := 0; j < 9; j++ {
         	if b.cells[i][j] == 0 {
         		return false   	
         	}
        }
    }
    return true	
}

func (b *Board) IsValid() bool {
	for i := 0; i < 9; i++ {
    	for j := 0; j < 9; j++ {
         	if b.cells[i][j] == 0 {
         		return false   	
         	}
        }
    }
    return true	
}

// func (b *Board) IsRowValid(r int) bool {
//     nums = make([]bool, 10)
//     for i := 0; i < 10; i++ {
//     	nums = false
//     }
//     for _, e := range b.cells[r] {
//     	if nums[e] {
//     		return false
//     	} else {
//     		nums[e] = true
//     	}
//     }
//     return true
// }

func (b *Board) IsColValid(c int) bool {
    return true	
}

func (b *Board) IsSquareValid(r int, c int) bool {
    return true	
}

type Board struct {
	cells [][]int
}

func main() {
	suduko := Board{}
	suduko.Initialize()
	suduko.cells[1][2] = 1 
	suduko.Print()
	fmt.Println(suduko.IsFull())
}