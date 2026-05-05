package mq

import (
	"Order/utils"
	"log"
	"strconv"

	"github.com/rabbitmq/amqp091-go"
)

var Url = ""

type RabbitMQ struct {
	conn      *amqp091.Connection
	channel   *amqp091.Channel
	QueueName string
	Exchange  string //
	key       string //
	MqUrl     string
}

func InitMqUrl() error {
	config := utils.GetRabbitMQConfig()
	Url = "amqp://" + config.UserName + ":" + config.Password + "@" + config.Host + ":" + strconv.Itoa(config.Port) + "/" + config.Vhost
	return nil
}
func NewRabbitStruct(queueName string, exchange string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, key: "", MqUrl: Url}
}
func (r *RabbitMQ) Destroy() error {
	err := r.channel.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = r.conn.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func NewRabbitMQSample(queueName string) (*RabbitMQ, error) {
	rabbitmq := NewRabbitStruct(queueName, "")
	var err error
	rabbitmq.conn, err = amqp091.Dial(rabbitmq.MqUrl)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	return rabbitmq, nil
}
func (r *RabbitMQ) PublishSample(message []byte) error {
	_, err := r.channel.QueueDeclare(r.QueueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	err = r.channel.Publish(r.Exchange, r.QueueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        message,
	})
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	return nil
}
