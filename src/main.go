package main

import (
	"fmt"
	"image/color"
	"queens/board"
	"queens/gui"
	"queens/solver"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type AppState struct {
	LoadedBoard   board.Board
	SolutionBoard board.Board
	FoundSolution bool
	GridDisplay   *fyne.Container
	Window        fyne.Window
}

func main() {
	myApp := app.New()

	state := &AppState{
		Window:      myApp.NewWindow("LinkedIn Queens Solver"),
		GridDisplay: container.NewStack(),
	}

	contentLabel := widget.NewLabel("Select a file to open")

	loadButton := widget.NewButton("Load txt File", func() {
		loadDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, state.Window)
				return
			}

			if reader == nil {
				contentLabel.SetText("No file selected")
				return
			}

			defer reader.Close()

			filePath := reader.URI().Path()

			if !strings.HasSuffix(filePath, ".txt") {
				dialog.ShowInformation("Wrong Extension", "File must be with txt extension", state.Window)
				return
			}
			b, parseErr := board.TxtToBoard(filePath)
			if parseErr != nil {
				dialog.ShowError(parseErr, state.Window)
			}
			state.LoadedBoard = b
			refreshGrid(state)
			contentLabel.SetText(fmt.Sprintf("Board loaded successfully from: %s", reader.URI().Name()))
		}, state.Window)

		loadDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		loadDialog.Show()
	})

	solveButton := widget.NewButton("Solve", func() {
		if state.LoadedBoard.RowLength == 0 {
			dialog.ShowInformation("No Board", "Please load a file first!", state.Window)
			return
		}

		// make sure no solution at first
		state.SolutionBoard = board.Board{}
		state.FoundSolution = false

		go func() {

			updateUI := func(b board.Board) {
				fyne.DoAndWait(func() {
					state.LoadedBoard = b
					refreshGrid(state)
				})
			}

			solution, count, duration := solver.Solve(state.LoadedBoard, updateUI)

			fyne.DoAndWait(func() {
				// final result
				if solution.RowLength == 0 {
					contentLabel.SetText("No solution found.")
				} else {
					state.SolutionBoard = solution
					state.FoundSolution = true
					contentLabel.SetText(fmt.Sprintf("Solved! %d perms in %dms", count, duration))
					state.LoadedBoard = solution
					refreshGrid(state)
				}
			})
		}()
	})

	saveToTXTButton := widget.NewButton("Save Solution to TXT", func() {
		if !state.FoundSolution {
			dialog.ShowInformation("No Solution", "There's no solution to be saved", state.Window)
			return
		}
		saveToTXTDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, state.Window)
				return
			}
			if writer == nil {
				return
			}

			saveSolutionErr := state.SolutionBoard.BoardToTxt(writer.URI().Path())
			if saveSolutionErr != nil {
				dialog.ShowError(saveSolutionErr, state.Window)
			} else {
				dialog.ShowInformation("Succes Save", "File saved successfully", state.Window)
			}

			writer.Close()

		}, state.Window)

		saveToTXTDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))

		saveToTXTDialog.Show()
	})

	state.Window.SetContent(container.NewBorder(
		contentLabel,
		container.NewHBox(loadButton, solveButton, saveToTXTButton),
		nil, nil,
		state.GridDisplay,
	))

	state.Window.Resize(fyne.NewSize(500, 600))
	state.Window.ShowAndRun()
}

func refreshGrid(state *AppState) {
	b := state.LoadedBoard
	grid := container.NewGridWithColumns(b.ColLength)

	for r := 0; r < b.RowLength; r++ {
		for c := 0; c < b.ColLength; c++ {

			rect := canvas.NewRectangle(gui.GetColor(b.Grid[r][c].Region))
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
