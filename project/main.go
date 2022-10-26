package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"fyne_app/internal/fail"

	"github.com/streadway/amqp"
)

func main() {

	myApp := app.New()

	input1 := widget.NewEntry()
	input1.SetPlaceHolder("Enter RabbitMq host...")

	input2 := widget.NewEntry()
	input2.SetPlaceHolder("Enter RabbitMq queue...")

	input3 := widget.NewEntry()
	input3.SetPlaceHolder("Enter Message ..")

	myWindow := myApp.NewWindow("RABBIT CONSUMER")

	button1 := widget.NewButton("Send Message for RabbitMq", func() {
		log.Println("content was", input1.Text, input2.Text, input3.Text)

		connect, err := amqp.Dial(input1.Text)
		fail.OnError(err, "Failed to connect rabbit")
		defer connect.Close()

		ch, err := connect.Channel()
		fail.OnError(err, "Failed to connect a chanel")
		defer ch.Close()

		qa, err := ch.QueueDeclare(
			input2.Text,
			false,
			false,
			false,
			false,
			nil,
		)
		fail.OnError(err, "Failed declara a queue")

		msgs, err := ch.Consume(
			qa.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		fail.OnError(err, "Failed to register a consumer")

		forever := make(chan bool)

		go func() {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
			}
		}()

		// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	})
	contant := widget.NewLabel("Messages RabbitMq")

	grid := container.New(layout.NewGridLayout(1), input1, input2, input3, contant, button1)
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(800, 350))
	myWindow.ShowAndRun()
}
