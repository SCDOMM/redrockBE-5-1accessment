package mq

import (
	"GeneralConfig"
	"Storage/sv"
	"fmt"
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

func init() {
	config := GeneralConfig.GetRabbitMQConfig()
	Url = "amqp://" + config.UserName + ":" + config.Password + "@" + config.Host + ":" + strconv.Itoa(config.Port) + "/" + config.Vhost
}
func NewRabbitStruct(queueName string, exchange string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, key: "", MqUrl: Url}
}
func (r *RabbitMQ) Destroy() error {
	err := r.channel.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	err = r.conn.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func NewRabbitMQSample(queueName string) (*RabbitMQ, error) {
	rabbitmq := NewRabbitStruct(queueName, "")
	var err error
	rabbitmq.conn, err = amqp091.Dial(rabbitmq.MqUrl)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return rabbitmq, nil
}
func (r *RabbitMQ) ConsumeSample() {
	q, err := r.channel.QueueDeclare(r.QueueName, true, false, false, false, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	messages, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	forever := make(chan bool)
	go func() {
		for d := range messages {
			err := sv.StorageMsgHandler(d.Body)
			if err != nil {
				log.Println(err.Error())
				return
			}
			fmt.Printf("Received a message: %s", d.Body)
		}
	}()
	<-forever
}
