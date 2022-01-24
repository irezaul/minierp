package main

import (
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

	email := widget.NewEntry()
	email.PlaceHolder = "Enter your email"

	pass := widget.NewPasswordEntry()
	pass.PlaceHolder = "Enter your password"

	emailID := widget.NewFormItem("Email", email)
	password := widget.NewFormItem("Password", pass)

	clientForm := widget.NewForm(emailID, password)

	clientForm.SubmitText = "Login"
	messageLabel := widget.NewLabel("")

	progress := widget.NewProgressBarInfinite()

	clientForm.OnSubmit = func() {

		myemail := "admin"
		myPass := "123"
		if email.Text == myemail && pass.Text == myPass {
			ShowDash(myApp)
			myWindow.Close()
		} else {
			dialog.NewInformation("Invalid Email or password", "Email or Passwor invalid", myWindow).Show()
		}

	}

	myWindow.SetContent(
		container.NewVBox(clientForm, messageLabel, progress),
	)
	myWindow.ShowAndRun()
}
