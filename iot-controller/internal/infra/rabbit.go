package infra

import "github.com/streadway/amqp"

type Rabbit struct {
	Channel   *amqp.Channel
	MainQueue amqp.Queue
}

func NewRabbit(channel *amqp.Channel, queueName string) (*Rabbit, error) {
	queue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &Rabbit{Channel: channel, MainQueue: queue}, nil
}
