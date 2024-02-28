package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/mhope-2/go-services/order-service/shared"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Publish(message shared.Message, queueName, routingKey, exchange string) {
	config := shared.NewEnvConfig()

	// Connect to RabbitMQ server
	conn, err := amqp.Dial(config.AmqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func() {
		if err = conn.Close(); err != nil {
			failOnError(err, "Failed to close RabbitMQ connection")
		}
	}()

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func() {
		if err = ch.Close(); err != nil {
			failOnError(err, "Failed to close RabbitMQ channel")
		}
	}()

	// Declare an exchange
	err = ch.ExchangeDeclare(
		exchange, // exchange name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// Declare a queue
	_, err = ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Serialize the instance to JSON
	body, err := json.Marshal(message)
	failOnError(err, "Failed to serialize message")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Publish the JSON message to the queue
	err = ch.PublishWithContext(ctx,
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent,  msg=%s, exchange=%s, queue=%s", body, exchange, queueName)
}
