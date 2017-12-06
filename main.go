package main

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://fans:123456@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close();

	q, err := ch.QueueDeclare(
		"gendoc2",
		true, // durable
 		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to declare a queue")

	body := "[\"body\"]"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body: []byte(body),
		})

	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
