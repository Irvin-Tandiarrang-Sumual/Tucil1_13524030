package solver

import (
	"queens/board"
)

func FirstPermutation(dimension int) []int {
	queenPositions := make([]int, dimension)
	for i := 0; i < dimension; i++ {
		queenPositions[i] = i
	}
	return queenPositions
}

func PermutationToBoard(queenPositions []int, template board.Board) board.Board {
	newBoard := template.DeepCopy()
	n := len(queenPositions)
	for row := 0; row < n; row++ {
		newBoard.Grid[row][queenPositions[row]].IsOccupied = true
	}
	return *newBoard
}

func CheckPermutation(queenPositions []int, template board.Board, solution *board.Board, countPermutations int) bool {
	b := PermutationToBoard(queenPositions, template)

	if b.IsSolution() {
		*solution = b
		return true
	}
	return false
}

func GeneratePermutation(queenPositions []int, template board.Board, solution *board.Board, countPermutations *int) bool {
	if len(queenPositions) == 0 {
		*countPermutations += 1
		return CheckPermutation(queenPositions, template, solution, *countPermutations)
	}

	var rc func(int) bool
	rc = func(np int) bool {
		if np == 1 {
			*countPermutations += 1
			return CheckPermutation(queenPositions, template, solution, *countPermutations)
		}
		np1 := np - 1
		pp := len(queenPositions) - np1
		// weave
		if rc(np1) {
			return true
		}
		for i := pp; i > 0; i-- {
			queenPositions[i], queenPositions[i-1] = queenPositions[i-1], queenPositions[i]
			if rc(np1) {
				return true
			}
		}
		// restore
		w := queenPositions[0]
		copy(queenPositions, queenPositions[1:pp+1])
		queenPositions[pp] = w

		return false
	}
	return rc(len(queenPositions))
}
