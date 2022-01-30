package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func loginForm(a fyne.App) {
	win := myWindow


	head := widget.NewLabel("Welcome to miniERP")
	head.Alignment = fyne.TextAlignCenter

	idEntry := widget.NewEntry()
	idEntry.PlaceHolder = "Enter your ID"

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.PlaceHolder = "Enter your password"

	id := widget.NewFormItem("ID", idEntry)
	password := widget.NewFormItem("Password", passwordEntry)

	loginForm := widget.NewForm(id, password)

	loginForm.SubmitText = "Login"
	loginForm.CancelText = "Quit"

	loginForm.OnCancel = func() {
		myWindow.Close()
	}

	loginForm.OnSubmit = func() {

		appEmail := ""
		appPass := ""
		if idEntry.Text == appEmail && passwordEntry.Text == appPass {
			Dashbord(myApp)
		} else {
			dialog.NewInformation("Warning !", "ID or Password invalid..", myWindow).Show()
			notify := fyne.NewNotification("Warning..", "Your banned for 1hour")
			myApp.SendNotification(notify)
		}

	}

	win.SetContent(
		container.NewVBox(head, loginForm),
	)

}
