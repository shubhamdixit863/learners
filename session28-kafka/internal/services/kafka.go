package services

import (
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaClient struct {
	KafkaConn *kafka.Conn
}

func (k *KafkaClient) PublishMessages(data string) error {
	// invoke kafka client methods
	write, err := k.KafkaConn.Write([]byte(data))
	if err != nil {
		return err
	}

	log.Println("messages published", write)
	return nil
}
