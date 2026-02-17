package board

import (
	"bufio"
	"fmt"
	"os"
)

func TxtToBoard(filename string) (Board, error) {
	file, err := os.Open(filename)
	if err != nil {
		// fmt.Println("Error to read file")
		return Board{}, fmt.Errorf("Error to read file")
	}

	defer file.Close()

	uniqueColours := make(map[rune]bool)
	var board Board

	scanner := bufio.NewScanner(file)

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
			return Board{}, fmt.Errorf("Line %d has length %d, expected %d\n", row+1, lineLength, expectedLength)
		}

		currentRow := make([]Cell, expectedLength)

		for col, ch := range line {
			if ch < 'A' || ch > 'Z' {
				return Board{}, fmt.Errorf("Warna %c tidak valid, silahkan masukkan hanya dari A-Z\n", ch)
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
		return Board{}, fmt.Errorf("Row length (%d) doesn't equal column length %d\n", row, expectedLength)
	}

	countUniqueColours := len(uniqueColours)

	if row != countUniqueColours {
		return Board{}, fmt.Errorf("Colours (%d) not equal to the dimension %d\n", countUniqueColours, row)
	}

	board.RowLength = expectedLength
	board.ColLength = expectedLength

	return board, nil

}

func (board Board) BoardToTxt(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Can't save to txt file")
		return err
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
	return nil
}
