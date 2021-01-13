package rabbitmqhandler

import (
	"github.com/streadway/amqp"
)

var (
	uri          = "amqp://guest:guest@localhost:5672/"
	exchangeName = "file-updates"
	exchangeType = "direct"
	routingKey   = "file-info"
	//queueName 	 = "test-queue"
)

// MessagerRabbitmq used to handle messages
type MessagerRabbitmq struct {
	channel      *amqp.Channel
	exchangeName string
}

// NewMessageHandler usage
func NewMessageHandler() *MessagerRabbitmq {
	connection, err := amqp.Dial(uri)
	if err != nil {
		panic(err)
	}
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	// needs  somewhere defer connection.Close()

	if err := channel.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		panic(err)
	}
	m := &MessagerRabbitmq{channel, exchangeName}
	return m
}
