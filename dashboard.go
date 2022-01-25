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

	win := a.NewWindow("miniERP Dashboard")
	win.Resize(fyne.NewSize(700, 400))

	headtext := widget.NewCard("Welcome to miniERP Dashboard", "",
		canvas.NewImageFromFile(""),
	)

	// btn1 := widget.NewButton("Add Client", func() {

	// })
	// btn2 := widget.NewButton("All Client", func() {

	// })

	card5 := widget.NewCard(
		"Add",
		"Click here to add a new client.",
		canvas.NewText("New Client", color.White),
	)
	card2 := widget.NewCard(
		"Search",
		"Click here to search client.",
		canvas.NewText("New Client", color.White),
	)
	card3 := widget.NewCard(
		"Product",
		"Click here to see all product.",
		canvas.NewText("New Client", color.White),
	)
	// win.SetContent(
	// 	headtext,

	// )

	win.SetContent(

		container.NewVBox(headtext),

		container.NewGridWithColumns(1,
		container.NewGridWithColumns(3,
			container.NewVBox(card2,card3,),
		),
		),
	)

	win.Show()
}
