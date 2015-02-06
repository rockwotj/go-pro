package main
import (
        "problem/sudoku"
        "fmt"
        "testing"
)

// BOARD FOR FIRST 5 TESTS FOUND AT: https://www.rose-hulman.edu/class/csse/csse403/Fall2010/Homework/hw14.html

func TestFirst(t *testing.T) {
	initialString := ".98257134542361987713948652954782316286413795371596248829634571137825469465179823"
	solutionString := "698257134542361987713948652954782316286413795371596248829634571137825469465179823"
	runTest(t, initialString, solutionString)
}

func TestSecond(t *testing.T) {
	initialString := "..8257134542361987713948652954782316286413795371596248829634571137825469465179823"
	solutionString := "698257134542361987713948652954782316286413795371596248829634571137825469465179823"
	runTest(t, initialString, solutionString)
}

func TestThird(t *testing.T) {
	// Whole first line or first square needs to be solved
	initialString := "...257134542361987713948652954782316286413795371596248829634571137825469465179823"
	solutionString := "698257134542361987713948652954782316286413795371596248829634571137825469465179823"
	runTest(t, initialString, solutionString)
}

func TestFourth(t *testing.T) {
	// Whole first square needs to be solved
	initialString := "...25713454.361987.1.948652954782316286413795371596248829634571137825469465179823"
	solutionString := "698257134542361987713948652954782316286413795371596248829634571137825469465179823"
	runTest(t, initialString, solutionString)
}

func TestFromClifton(t *testing.T) {
	// Whole puzzle
	initialString := "...2.713.54..6..8..1.9...5295..82..6..6...7..3..59..4882...4.7..3..2..69.651.9..."
	solutionString := "698257134542361987713948652954782316286413795371596248829634571137825469465179823"
	runTest(t, initialString, solutionString)
}

// Next three tests are from http://norvig.com/top95.txt
func FullTest1(t *testing.T) {
    // Whole puzzle
    initialString := ".237....68...6.59.9.....7......4.97.3.7.96..2.........5..47.........2....8......."
    solutionString := "123759486874261593965384721216543978357896142498127365532478619641932857789615234"
    runTest(t, initialString, solutionString)
}

func FullTest2(t *testing.T) {
    // Whole puzzle
    initialString := "53..2.9...24.3..5...9..........1.827...7.........981.............64....91.2.5.43."
    solutionString := "538127946624839751719645382965314827381762594247598163493281675856473219172956438"
    runTest(t, initialString, solutionString)
}

func FullTest3(t *testing.T) {
    // Whole puzzle
    initialString := "..84...3....3.....9....157479...8........7..514.....2...9.6...2.5....4......9..56"
    solutionString := "518476239427359618963821574795248361832617945146935827379564182651782493284193756"
    runTest(t, initialString, solutionString)
}

func runTest(t *testing.T, initialString string, solutionString string) {
    board := sudoku.BoardFromString(initialString)
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
    if (board.ToString() != solutionString) {
        t.Errorf("Board Incorrect!", board.ToString(), solutionString)
    }
}