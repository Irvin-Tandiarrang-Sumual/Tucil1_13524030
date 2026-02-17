package gui

import (
	"queens/board"

	"fyne.io/fyne/v2"
)

type AppState struct {
	LoadedBoard   board.Board
	SolutionBoard board.Board
	FoundSolution bool
	GridDisplay   *fyne.Container
	Window        fyne.Window
}
