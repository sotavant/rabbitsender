package main

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

	test := [16]string{
		`{"taskId":1,"type":"pdf","lastDoc":false,"docCount":3,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e1.pdf"}`,
		`{"taskId":1,"type":"pdf","lastDoc":true,"docCount":3,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e2.pdf"}`,
		`{"taskId":1,"type":"pdf","lastDoc":false,"docCount":3,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e3.pdf"}`,

		`{"taskId":2,"type":"pdf","lastDoc":false,"docCount":8,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e1.pdf"}`,
		`{"taskId":2,"type":"pdf","lastDoc":false,"docCount":8,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e2.pdf"}`,
		`{"taskId":2,"type":"pdf","lastDoc":false,"docCount":8,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e3.pdf"}`,
		`{"taskId":2,"type":"pdf","lastDoc":false,"docCount":8,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e4.pdf"}`,
		`{"taskId":2,"type":"pdf","lastDoc":false,"docCount":8,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e5.pdf"}`,
		`{"taskId":2,"type":"pdf","lastDoc":false,"docCount":8,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e6.pdf"}`,
		`{"taskId":2,"type":"pdf","lastDoc":false,"docCount":8,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e8.pdf"}`,
		`{"taskId":2,"type":"pdf","lastDoc":true,"docCount":8,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e9.pdf"}`,

		`{"taskId":5,"type":"pdf","lastDoc":false,"docCount":4,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e5.pdf"}`,
		`{"taskId":5,"type":"pdf","lastDoc":false,"docCount":4,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e6.pdf"}`,
		`{"taskId":5,"type":"pdf","lastDoc":false,"docCount":4,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e8.pdf"}`,
		`{"taskId":5,"type":"pdf","lastDoc":true,"docCount":4,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e9.pdf"}`,

		`{"taskId":6,"type":"pdf","lastDoc":true,"docCount":1,"fileName":"\u0421\u0435\u043c\u0435\u043d\u043e\u0432 \u041e. \u041d.\u041f\u0435\u0440\u0432\u0438\u0447\u043d\u044b\u0439\u041c\u043e5.pdf"}`,
	}

	for _, v := range test {
		body := v
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
	}

	failOnError(err, "Failed to publish a message")

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
