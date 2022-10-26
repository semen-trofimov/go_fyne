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

	input1 := widget.NewEntry()
	input1.SetPlaceHolder("Enter RabbitMq host...")

	input2 := widget.NewEntry()
	input2.SetPlaceHolder("Enter RabbitMq queue...")

	input3 := widget.NewEntry()
	input3.SetPlaceHolder("Enter Message ..")

	myWindow := myApp.NewWindow("Grid Layout")

	button1 := widget.NewButton("save", func() {
		log.Println("content was", input1.Text, input2.Text, input3.Text)
	})

	// button1 := widget.NewButton("button1", func() {
	// 	log.Println("tapped button1")
	// })

	grid := container.New(layout.NewGridLayout(1), input1, input2, input3, button1)
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(350, 350))
	myWindow.ShowAndRun()
}
