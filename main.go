package main

import (
	"database/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	_ "github.com/mattn/go-sqlite3"
)

var (
	myApp    fyne.App    = app.New()
	myWindow fyne.Window = myApp.NewWindow("miniERP v0.0.1")

	db  *sql.DB
	err error
)

func init() {
	dbcon()

}

func main() {
	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.SetFixedSize(true)
	
	loginForm(myApp)

	myWindow.ShowAndRun()
}
