package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	myApp    fyne.App    = app.New()
	myWindow fyne.Window = myApp.NewWindow("Mini ERP")
)

func main() {

	myWindow.Resize(fyne.NewSize(400, 250))
	myWindow.SetFixedSize(true)

	head := widget.NewLabel("Welcome to miniERP")

	email := widget.NewEntry()
	email.PlaceHolder = "Enter your ID"

	pass := widget.NewPasswordEntry()
	pass.PlaceHolder = "Enter your Password"

	emailID := widget.NewFormItem("Email", email)
	password := widget.NewFormItem("Password", pass)

	clientForm := widget.NewForm(emailID, password)

	clientForm.SubmitText = "Login"
	clientForm.CancelText = "Cancel"

	messageLabel := widget.NewLabel("")

	clientForm.OnSubmit = func() {

		myemail := ""
		myPass := ""
		if email.Text == myemail && pass.Text == myPass {
			ShowDash(myApp)
			fmt.Println("great job ..")
			myWindow.Close()
		} else {
			dialog.NewInformation("Warning !", "Email or Passwor invalid", myWindow).Show()
		}

	}

	myWindow.SetContent(
		container.NewVBox(head, clientForm, messageLabel),
	)
	myWindow.ShowAndRun()
}
