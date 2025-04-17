package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"session28-kafka/internal/services"
)

func connectKafka() {

}

const TOPIC = "logs"

func main() {
	// We will create a kafka client

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", TOPIC, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	kafkaClient := services.KafkaClient{
		KafkaConn: conn,
	}

	err = kafkaClient.PublishMessages("hey there ")
	if err != nil {
		log.Println(err)
		return
	}
}
