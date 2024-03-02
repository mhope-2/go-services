// Package rabbitmq maintains code for rabbitmq
package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/mhope-2/go-services/order-service/shared"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn *amqp.Connection
}

func NewPublisher(s *shared.EnvConfig) (*Publisher, error) {
	conn, err := amqp.Dial(s.AMQPURI)
	if err != nil {
		log.Println("Error creating amqp connection")
		return nil, err
	}
	return &Publisher{conn: conn}, nil
}

func (p *Publisher) Publish(message shared.Message, queueName, routingKey, exchange string) error {

	// Create a channel
	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

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
	if err != nil {
		return err
	}

	// Declare a queue
	_, err = ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	// Serialize the instance to JSON
	body, err := json.Marshal(message)
	if err != nil {
		log.Println("Failed to serialize message")
		return err
	}

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
	if err != nil {
		log.Println("Failed to publish a message")
		return err
	}

	log.Printf(" [x] Sent,  msg=%s, exchange=%s, queue=%s", body, exchange, queueName)
	return nil
}
