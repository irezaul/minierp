package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

func ShowDash(a fyne.App) {

	// a:= app.New()

	win := a.NewWindow("dash")
	win.Resize(fyne.NewSize(700, 400))

	headtext := canvas.NewText("Welcome to Dashboard", color.White)
	headtext.TextSize = 25

	btn1 := widget.NewButton("Add Client", func() {

	})
	btn2 := widget.NewButton("All Client", func() {

	})

	win.SetContent(
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(1,
				container.NewVBox(
					container.NewCenter(headtext),
				),
			),

			container.NewGridWithColumns(4,
				container.NewVBox(btn1, btn2),
				container.NewVBox(),
			),
		),
	)

	win.Show()
}
