package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	e := echo.New()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"queue-to-say", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	e.GET("/says/:msg", func(c echo.Context) error {
		msg := c.Param("msg")
		ctx := c.Request().Context()

		ctxMQ, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		err = ch.PublishWithContext(ctxMQ,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf("Sent: %s\n", msg)

		return c.String(http.StatusOK, msg)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
