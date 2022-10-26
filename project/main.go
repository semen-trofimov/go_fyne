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

	myWindow := myApp.NewWindow("Grid Layout")

	button1 := widget.NewButton("Send Message for RabbitMq", func() {
		log.Println("content was", input1.Text, input2.Text, input3.Text)

		conn, err := amqp.Dial(input1.Text)
		fail.OnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		fail.OnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			input2.Text, // name
			false,       // durable
			false,       // delete when unused
			false,       // exclusive
			false,       // no-wait
			nil,         // arguments
		)
		fail.OnError(err, "Failed to declare a queue")

		// body := "Hello World!"
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(input3.Text),
			})
		log.Printf(" [x] Sent %s", input3.Text)
		fail.OnError(err, "Failed to publish a message")
	})
	grid := container.New(layout.NewGridLayout(1), input1, input2, input3, button1)
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(800, 350))
	myWindow.ShowAndRun()
}
