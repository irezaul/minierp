package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
)

func menuItem() {

	menuItem1 := fyne.NewMenuItem("Light", func() {
		Light := theme.LightTheme()
		myApp.Settings().SetTheme(Light)
	})
	menuItem2 := fyne.NewMenuItem("Dark", func() {
		Dark := theme.DarkTheme()
		myApp.Settings().SetTheme(Dark)
	})
	menuItem3 := fyne.NewMenuItem("About", func() {
		dialog.NewInformation("About", "MT MART \n 2no Gate Nasirabad, Chittagong,4203. \n 0189523690 ", myWindow).Show()
	})
	menuItem4 := fyne.NewMenuItem("logout", func() {
		loginForm(myApp)
	})

	newMenu1 := fyne.NewMenu("File", menuItem3, menuItem4)

	newMenu := fyne.NewMenu("Theme", menuItem1, menuItem2)

	menu := fyne.NewMainMenu(newMenu1, newMenu)
	myWindow.SetMainMenu(menu)

}
