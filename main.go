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
		"types",
		true, // durable
 		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to declare a queue")

	test := [16]string{
		`{"type":"pdf","data":{"field1": "value1", "field2": "value17"}}`,
		`{"type":"doc","data":{"field1": "value2", "field2": "value18"'}}`,
		`{"type":"pdf","data":{"field1": "value3", "field2": "value19"}}`,
		`{"type":"pdf","data":{"field1": "value4", "field2": "value20"'}}`,
		`{"type":"pdf","data":{"field1": "value5", "field2": "value21"}}`,
		`{"type":"doc","data":{"field1": "value6", "field2": "value22"}}`,
		`{"type":"pdf","data":{"field1": "value7", "field2": "value23"}}`,
		`{"type":"doc","data":{"field1": "value8", "field2": "value24"}}`,
		`{"type":"pdf","data":{"field1": "value9", "field2": "value25"}}`,
		`{"type":"txt","data":{"field1": "value10", "field2": "value26"}}`,
		`{"type":"pdf","data":{"field1": "value11", "field2": "value27"}}`,
		`{"type":"txt","data":{"field1": "value12", "field2": "value28"'}}`,
		`{"type":"pdf","data":{"field1": "value13", "field2": "value29"}}`,
		`{"type":"txt","data":{"field1": "value14", "field2": "value30"'}}`,
		`{"type":"pdf","data":{"field1": "value15", "field2": "value31"}}`,
		`{"type":"pdf","data":{"field1": "value16", "field2": "value32"}}`,
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
