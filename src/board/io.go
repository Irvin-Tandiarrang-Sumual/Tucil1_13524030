package board

import (
	"bufio"
	"fmt"
	"os"
)

func TxtToBoard(filename string) (Board, bool) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error to read file")
		return Board{}, false
	}

	defer file.Close()

	uniqueColours := make(map[rune]bool)
	var board Board

	scanner := bufio.NewScanner(file)
	// fmt.Print(scanner)

	row := 0
	expectedLength := -1
	for scanner.Scan() {
		line := scanner.Text()
		lineLength := len(line)
		if expectedLength == -1 {
			expectedLength = lineLength
			board.Grid = make([][]Cell, 0, expectedLength)
		}

		if lineLength != expectedLength {
			fmt.Printf("Line %d has length %d, expected %d\n", row+1, lineLength, expectedLength)
			return Board{}, false
		}

		currentRow := make([]Cell, expectedLength)

		for col, ch := range line {
			if ch < 'A' || ch > 'Z' {
				fmt.Printf("Warna %c tidak valid, silahkan masukkan hanya dari A-Z\n", ch)
				return Board{}, false
			}

			currentRow[col].Region = ch
			currentRow[col].IsOccupied = false

			if uniqueColours[ch] {
				continue
			}
			uniqueColours[ch] = true
		}
		board.Grid = append(board.Grid, currentRow)

		row += 1

	}

	if row != expectedLength {
		fmt.Printf("Row length (%d) doesn't equal column length %d\n", row, expectedLength)
		return Board{}, true
	}

	countUniqueColours := len(uniqueColours)

	if row != countUniqueColours {
		fmt.Printf("Colours (%d) not equal to the dimension %d\n", countUniqueColours, row)
		return Board{}, false
	}

	// fmt.Println(data[1])
	// fmt.Printf("%T\n", data)
	board.RowLength = expectedLength
	board.ColLength = expectedLength

	return board, true

}

func (board Board) BoardToTxt(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Can't save to txt file")
		return
	}

	defer file.Close()

	for _, row := range board.Grid {
		line := ""
		for _, cell := range row {
			if cell.IsOccupied {
				line += "#"
			} else {
				line += string(cell.Region)
			}
		}
		file.WriteString(line + "\n")
	}
}
