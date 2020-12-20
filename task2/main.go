package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var fahrenheitField *FahrenheitEntry
var celsiusField *CelciusEntry

type CelciusEntry struct {
	widget.Entry
}

func newCelciusEntry() *CelciusEntry {
	entry := &CelciusEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *CelciusEntry) KeyUp(key *fyne.KeyEvent) {
	s := e.Entry.Text

	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		e.Entry.KeyUp(key)
		fahrenheitField.Entry.SetText(strconv.FormatInt(InFahrenheit(n), 10))
	} else {
		fmt.Println(s, "is not an integer.")
	}
}

func InCelcius(fahrenheit int64) int64 {
	return int64((float64(fahrenheit) - 32) * (float64(5) / 9))
}

type FahrenheitEntry struct {
	widget.Entry
}

func newFahrenheitEntry() *FahrenheitEntry {
	entry := &FahrenheitEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *FahrenheitEntry) KeyUp(key *fyne.KeyEvent) {
	s := e.Entry.Text

	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		e.Entry.KeyUp(key)
		celsiusField.Entry.SetText(strconv.FormatInt(InCelcius(n), 10))
	} else {
		fmt.Println(s, "is not an integer.")
	}
}

func InFahrenheit(celsius int64) int64 {
	return celsius*(9/5) + 32
}

func main() {
	a := app.New()
	w := a.NewWindow("7guis - Task 2")

	celsiusField = newCelciusEntry()
	fahrenheitField = newFahrenheitEntry()
	celsiusLabel := widget.NewLabel("Celsius = ")
	fahrenheitLabel := widget.NewLabel("Fahrenheit")
	box := widget.NewHBox(
		celsiusField,
		celsiusLabel,
		fahrenheitField,
		fahrenheitLabel,
	)
	w.SetContent(box)
	w.ShowAndRun()
}
