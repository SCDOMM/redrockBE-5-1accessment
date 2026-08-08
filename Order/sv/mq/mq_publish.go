package mq

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func (r *RabbitMQ) PublishSample(message []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, err := r.channel.QueueDeclare(r.QueueName, true, false, false, false, nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = r.channel.Publish(r.Exchange, r.QueueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        message,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
