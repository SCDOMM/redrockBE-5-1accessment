package mq

import (
	"Storage/sv"
	"fmt"
	"log"
)

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
