package board

import "fmt"

type Cell struct {
	Region     rune
	IsOccupied bool
}

type Board struct {
	RowLength int
	ColLength int
	Grid      [][]Cell
}

func (board Board) PrintBoard() {
	for i := 0; i < board.RowLength; i++ {
		for j := 0; j < board.ColLength; j++ {
			if board.Grid[i][j].IsOccupied {
				fmt.Print("#")
			} else {
				fmt.Printf("%c", board.Grid[i][j].Region)
			}
		}
		fmt.Println()
	}

}

func (board *Board) DeepCopy() *Board {
	if board == nil {
		return nil
	}

	newBoard := &Board{
		RowLength: board.RowLength,
		ColLength: board.ColLength,
		Grid:      make([][]Cell, board.RowLength),
	}

	for i := range board.Grid {
		newBoard.Grid[i] = make([]Cell, board.ColLength)
		copy(newBoard.Grid[i], board.Grid[i])
	}
	return newBoard
}

func (board Board) IsSolution() bool {
	usedRegions := make(map[rune]bool)
	countQueen := 0
	for r := 0; r < board.RowLength; r++ {
		for c := 0; c < board.ColLength; c++ {
			// check if a queen is at a particular spot
			if board.Grid[r][c].IsOccupied {
				countQueen += 1
				// check regions
				if exists := usedRegions[board.Grid[r][c].Region]; exists {
					return false
				} else {
					usedRegions[board.Grid[r][c].Region] = true
				}

				// check diagonal
				if !board.checkDiagonalConstraint(r, c) {
					return false
				}
			}
		}
	}
	return countQueen == board.RowLength
}

func (board Board) checkDiagonalConstraint(row int, col int) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 && j != 0 && (row+i >= 0) && (row+i < board.RowLength) && (col+j >= 0) && (col+j < board.ColLength) && board.Grid[row+i][col+j].IsOccupied {
				return false
			}
		}
	}
	return true
}
