package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

const MAX_SLIDER_IN_SEC = 30
const MAX_PROGRESS_VALUE = 100
const UPDATE_EVERY_MS = 100

var button *widget.Button
var elapseTimeValue *widget.Label
var elapsed float64 = 0
var progressBar *widget.ProgressBar
var slider *widget.Slider
var ticker *time.Ticker

func main() {
	a := app.New()
	w := a.NewWindow("Timer")

	var elapseTimeLabel = widget.NewLabel("Elapsed Time:")
	elapseTimeValue = widget.NewLabel("0s")
	slider = widget.NewSlider(1, MAX_SLIDER_IN_SEC)
	slider.SetValue(MAX_SLIDER_IN_SEC)

	progressBar = widget.NewProgressBar()

	button = widget.NewButton("Reset", func() {
		kickNewTicker()
	})

	box := widget.NewHBox(
		elapseTimeLabel,
		progressBar,
	)

	boxV := widget.NewVBox(
		box,
		elapseTimeValue,
		slider,
		button,
	)

	kickNewTicker()

	w.SetContent(boxV)
	w.Resize(fyne.NewSize(150, 120))
	w.ShowAndRun()
}

func kickNewTicker() {
	initValues()
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				elapsed = elapsed + (float64(UPDATE_EVERY_MS) / 1000)
				progressBar.SetValue((slider.Value - elapsed) / slider.Value * MAX_PROGRESS_VALUE)
				elapseTimeValue.SetText(fmt.Sprintf("%.1f", elapsed) + "s")
				if elapsed > slider.Value {
					ticker.Stop()
					done <- true
				}
			}
		}
	}()
}

func initValues() {
	elapsed = 0
	if ticker != nil {
		ticker.Stop()
	}
	progressBar.Min = 0
	progressBar.Max = MAX_PROGRESS_VALUE
	progressBar.SetValue(MAX_PROGRESS_VALUE)
	ticker = time.NewTicker(UPDATE_EVERY_MS * time.Millisecond)
}
