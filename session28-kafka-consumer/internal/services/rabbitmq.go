package services

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type RabbitMq struct {
	conn *amqp.Connection
}

func (r *RabbitMq) PublishMessages(data string) error {

	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}

	// if you donot have queue ,then you need to create a queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})

	return nil
}

func (r *RabbitMq) ConsumeMessages() error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}

	// if you donot have queue ,then you need to create a queue
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		return err
	}

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}

	return nil
}

func NewRabbitMqClient(conn *amqp.Connection) Broker {
	return &RabbitMq{
		conn: conn,
	}
}
