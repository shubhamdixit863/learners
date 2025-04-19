package services

import (
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaClient struct {
	KafkaConn *kafka.Conn
}

func NewKafkaClient(KafkaConn *kafka.Conn) Broker {
	return &KafkaClient{KafkaConn: KafkaConn}
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

func (k *KafkaClient) ConsumeMessages() error {

	data := make([]byte, 20)
	for {
		read, err := k.KafkaConn.Read(data)
		if err != nil {
			return err

		}
		log.Println("Number of bytes read", read)
		log.Println("Messages--", string(data))
	}
}
