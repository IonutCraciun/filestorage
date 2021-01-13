package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

var (
	uri          = "amqp://guest:guest@localhost:5672/"
	exchangeName = "file-updates"
	exchangeType = "direct"
	routingKey   = "file.update"
)

func init() {
	//little hack to add custom routing keys from the command line
	if len(os.Args) > 1 {
		routingKey = os.Args[1]
	}
}

// MessagerRabbitmq used to handle messages
type MessagerRabbitmq struct {
	channel      *amqp.Channel
	exchangeName string
}

// NewMessageHandler s
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

func main() {
	done := make(chan error)
	rabbitMessager := NewMessageHandler() // default
	defer rabbitMessager.closeAll()
	// third true means exclusive true, When the connection that declared it closes, the queue will be deleted because it is declared as exclusive.
	q, err := rabbitMessager.channel.QueueDeclare("", false, false, true, false, nil) // "" empty queue name means rabbitmq will create a new one
	// with a random name

	if err != nil {
		panic("error while declaring the queue: " + err.Error())
	}
	err = rabbitMessager.channel.QueueBind(q.Name, routingKey, exchangeName, false, nil)
	if err != nil {
		panic("error while binding the queue: " + err.Error())
	}

	deliveries, err := rabbitMessager.channel.Consume(
		q.Name, // name
		"",     // consumerTag,
		true,   // noAck
		false,  // exclusive
		false,  // noLocal
		false,  // noWait
		nil,    // arguments
	)

	if err != nil {
		fmt.Printf("Queue Consume err: %s", err)
	}
	forever := make(chan bool)
	go handle(deliveries, done)

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}

func handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		log.Printf(
			"Got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
	}
}

func (m *MessagerRabbitmq) closeAll() {
	m.channel.Close()
}
