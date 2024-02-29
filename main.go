package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"types",
		true, // durable
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to declare a queue")

	test := [17]string{
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
		`{"data":{"field1": "value1", "field2": "value17"}, "handler":"/Users/fanis/go/src/github.com/Fisher21/rabbit-dispatcher/handler/handler.php"}`,
	}

	for _, v := range test {
		body := v
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         []byte(body),
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
