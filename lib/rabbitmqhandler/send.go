package rabbitmqhandler

import (
	"github.com/streadway/amqp"
	"log"
)

// Send usage
func (m *MessagerRabbitmq) Send(routingKey string, body ...string) error {
	var bbody string
	for _, val := range body {
		bbody = bbody + " " + val
	}
	if err := m.channel.Publish(
		m.exchangeName, // publish to an exchange
		routingKey,     // routing to 0 or more queues
		false,          // mandatory to
		false,          // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(bbody),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	); err != nil {
		return err
	}
	log.Printf("Message %s was sent", bbody)
	return nil
}
