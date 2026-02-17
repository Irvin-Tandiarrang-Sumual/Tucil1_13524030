package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func RefreshGrid(state *AppState) {
	b := state.LoadedBoard
	grid := container.NewGridWithColumns(b.ColLength)

	for r := 0; r < b.RowLength; r++ {
		for c := 0; c < b.ColLength; c++ {

			rect := canvas.NewRectangle(GetColor(b.Grid[r][c].Region))
			rect.SetMinSize(fyne.NewSize(40, 40))
			rect.StrokeWidth = 1
			rect.StrokeColor = color.Black

			if b.Grid[r][c].IsOccupied {
				queen := canvas.NewText("👑", color.Black)
				queen.TextSize = 40
				grid.Add(container.NewStack(rect, container.NewCenter(queen)))
			} else {
				grid.Add(rect)
			}
		}
	}

	state.GridDisplay.Objects = []fyne.CanvasObject{grid}
	state.GridDisplay.Refresh()
}
