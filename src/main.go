package main

import (
	"fmt"
	"queens/board"
	"queens/solver"
	"time"
)

func main() {
	b, _ := board.TxtToBoard("test.txt")
	arr := solver.FirstPermutation(b.ColLength)
	var sol board.Board
	var cP int

	start := time.Now()
	fmt.Println("Start")
	if solver.GeneratePermutation(arr, b, &sol, &cP) {
		fmt.Printf("Solution is found\n")
		sol.PrintBoard()
	} else {
		sol.PrintBoard()
		fmt.Println("Nice try")
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Used time : %d ms\n", elapsed.Milliseconds())
	fmt.Printf("Traversed case : %d \n", cP)
}
