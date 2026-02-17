package solver

import (
	"queens/board"
	"time"
)

func Solve(b board.Board, onUpdate func(board.Board)) (board.Board, int, int) {
	firstPermutation := firstPermutation(b.ColLength)
	var solution board.Board
	var countPermutations int
	start := time.Now()
	if !generatePermutation(firstPermutation, b, &solution, &countPermutations, onUpdate) {
		solution = board.Board{}
	}
	finish := time.Now()
	totalTime := finish.Sub(start)
	return solution, countPermutations, int(totalTime.Milliseconds())
}
func firstPermutation(dimension int) []int {
	queenPositions := make([]int, dimension)
	for i := 0; i < dimension; i++ {
		queenPositions[i] = i
	}
	return queenPositions
}

func permutationToBoard(queenPositions []int, template board.Board) board.Board {
	newBoard := template.DeepCopy()
	n := len(queenPositions)
	for row := 0; row < n; row++ {
		newBoard.Grid[row][queenPositions[row]].IsOccupied = true
	}
	return *newBoard
}

func checkPermutation(queenPositions []int, template board.Board, solution *board.Board, countPermutations int, onUpdate func(board.Board)) bool {
	b := permutationToBoard(queenPositions, template)

	if countPermutations%200 == 0 {
		onUpdate(b)

		time.Sleep(time.Millisecond * 2)
	}

	if b.IsSolution() {
		*solution = b
		return true
	}
	return false
}

func generatePermutation(queenPositions []int, template board.Board, solution *board.Board, countPermutations *int, onUpdate func(board.Board)) bool {
	if len(queenPositions) == 0 {
		*countPermutations += 1
		return checkPermutation(queenPositions, template, solution, *countPermutations, onUpdate)
	}

	var rc func(int) bool
	rc = func(np int) bool {
		if np == 1 {
			*countPermutations += 1
			return checkPermutation(queenPositions, template, solution, *countPermutations, onUpdate)
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
