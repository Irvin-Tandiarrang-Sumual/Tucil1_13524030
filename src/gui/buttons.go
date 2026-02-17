package gui

import (
	"fmt"
	"queens/board"
	"queens/solver"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func MakeLoadButton(state *AppState, label *widget.Label) *widget.Button {
	return widget.NewButton("Load txt File", func() {
		loadDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, state.Window)
				return
			}

			if reader == nil {
				label.SetText("No file selected")
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
				return
			}
			state.LoadedBoard = b
			RefreshGrid(state)
			label.SetText(fmt.Sprintf("Board loaded successfully from: %s", reader.URI().Name()))
		}, state.Window)

		loadDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		loadDialog.Show()
	})
}

func MakeSaveToTXTButton(state *AppState) *widget.Button {
	return widget.NewButton("Save Solution to TXT", func() {
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
}

func MakeSolveButton(state *AppState, label *widget.Label) *widget.Button {
	return widget.NewButton("Solve", func() {
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
					RefreshGrid(state)
				})
			}

			solution, count, duration := solver.Solve(state.LoadedBoard, updateUI)

			fyne.DoAndWait(func() {
				// final result
				if solution.RowLength == 0 {
					label.SetText("No solution found.")
				} else {
					state.SolutionBoard = solution
					state.FoundSolution = true
					label.SetText(fmt.Sprintf("Solved! %d perms in %dms", count, duration))
					state.LoadedBoard = solution
					RefreshGrid(state)
				}
			})
		}()
	})

}
