package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartApp() {
	myApp := app.New()

	state := &AppState{
		Window:      myApp.NewWindow("LinkedIn Queens Solver"),
		GridDisplay: container.NewStack(),
	}

	contentLabel := widget.NewLabel("Select a file to open")

	loadButton := MakeLoadButton(state, contentLabel)

	solveButton := MakeSolveButton(state, contentLabel)

	saveToTXTButton := MakeSaveToTXTButton(state)

	state.Window.SetContent(container.NewBorder(
		contentLabel,
		container.NewHBox(loadButton, solveButton, saveToTXTButton),
		nil, nil,
		state.GridDisplay,
	))

	state.Window.Resize(fyne.NewSize(500, 600))
	state.Window.ShowAndRun()
}
