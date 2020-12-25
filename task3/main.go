package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

const DEFAULT_DATE_DISPLAY_TEXT = "27.03.2014"
const ONE_WAY_FLIGHT_DISPLAY_TEXT = "one-way flight"
const RETURN_FLIGHT_DISPLAY_TEXT = "return flight"

var inboundEntry *DateEntry
var outboundEntry *DateEntry
var button *widget.Button
var selector *widget.Select

type DateEntry struct {
	widget.Entry
}

func newDateEntry(date string) *DateEntry {
	entry := &DateEntry{}
	entry.ExtendBaseWidget(entry)
	entry.Entry.SetText(date)
	return entry
}

func (e *DateEntry) KeyUp(key *fyne.KeyEvent) {
	updateButton()
}

func main() {
	a := app.New()
	w := a.NewWindow("7guis - Task 3")

	outboundEntry = newDateEntry(DEFAULT_DATE_DISPLAY_TEXT)
	inboundEntry = newDateEntry(DEFAULT_DATE_DISPLAY_TEXT)
	inboundEntry.Entry.SetReadOnly(true)

	options := []string{ONE_WAY_FLIGHT_DISPLAY_TEXT, RETURN_FLIGHT_DISPLAY_TEXT}
	selector = widget.NewSelect(options, func(s string) {
		inboundEntry.Entry.SetReadOnly(s != RETURN_FLIGHT_DISPLAY_TEXT)
		updateButton()
	})
	selector.SetSelected(ONE_WAY_FLIGHT_DISPLAY_TEXT)

	button = widget.NewButton("Book", func() {
		if isOneWay() {
			dialog.NewInformation("Confirm", "You have booked a one-way flight on "+outboundEntry.Entry.Text, w)
		} else {
			dialog.NewInformation("Confirm", "You have booked a return flight on \n Outbound: "+outboundEntry.Entry.Text+"\nReturn: "+inboundEntry.Entry.Text, w)
		}
	})

	box := widget.NewVBox(
		selector,
		outboundEntry,
		inboundEntry,
		button,
	)

	w.SetContent(box)
	w.Resize(fyne.NewSize(400, 150))
	w.ShowAndRun()
}

func updateButton() {
	if button != nil {
		if isAllValid() {
			button.Enable()
		} else {
			button.Disable()
		}
	}
}

func isAllValid() bool {
	if isOneWay() {
		return isWellFormatted(outboundEntry.Entry.Text)
	} else {
		return isWellFormatted(outboundEntry.Entry.Text) && isWellFormatted(inboundEntry.Entry.Text) &&
			!isBefore(inboundEntry.Entry.Text, outboundEntry.Entry.Text)
	}
}

func isOneWay() bool {
	return selector.Selected == ONE_WAY_FLIGHT_DISPLAY_TEXT
}

func isWellFormatted(date string) bool {
	if _, e := parseDate(date); e != nil {
		return false
	} else {
		return true
	}
}

func isBefore(outbound string, inbound string) bool {
	out, _ := parseDate(outbound)
	in, _ := parseDate(inbound)

	return out.Before(in)
}

func parseDate(date string) (time.Time, error) {
	// smoking grimlins how Go handles date formats
	// https://golang.org/src/time/format.go
	layout := "02.01.2006"
	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println(err)
		return time.Now(), err
	}
	fmt.Println(t)
	return t, nil
}
