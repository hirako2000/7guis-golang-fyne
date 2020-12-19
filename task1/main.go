package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/layout"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("7guis - Task 1")

	var counter int64 = 0
	counterField := widget.NewEntry()
	counterField.Text = strconv.FormatInt(counter, 10)
	counterField.SetReadOnly(true)
  box := widget.NewHBox(
		counterField,
		widget.NewButton("Count", func() {
			counter++
			counterField.SetText(strconv.FormatInt(counter, 10))
		}),
	)
	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(2), box))
	w.ShowAndRun()
}
