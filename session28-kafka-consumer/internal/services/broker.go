package services

type Broker interface {
	PublishMessages(data string) error
	ConsumeMessages() error
}
