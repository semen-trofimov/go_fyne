package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	myWindow := myApp.NewWindow("Grid Layout")

	button1 := widget.NewButton("button1", func() {
		log.Println("tapped button1")
	})
	button2 := widget.NewButton("button2", func() {
		log.Println("tapped button2")
	})
	button3 := widget.NewButton("button3", func() {
		log.Println("tapped button3")
	})
	button4 := widget.NewButton("button4", func() {
		log.Println("tapped button4")
	})

	grid := container.New(layout.NewGridLayout(1), input, button1, button2, button3, button4)
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(350, 350))
	myWindow.ShowAndRun()
}
