package service

import (
	"log"
	"github.com/streadway/amqp"
	"github.com/DouwaIO/hairtail/src/pipeline"
	task_pipeline "github.com/DouwaIO/hairtail/src/task/pipeline"
)

func MQ(protocol, host, user, pwd, topic, ackPolicy string, data []*pipeline.Container) error {
	if protocol == "amqp" {
		mq_connct := protocol+"://"+user+":"+pwd+"@"+host+"/"
		conn, err := amqp.Dial(mq_connct)
		if err != nil {
			return err
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			return err
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
		    topic, // name

		    false,   // durable
		    false,   // delete when unused
		    false,   // exclusive
		    false,   // no-wait
		    nil,     // arguments
		)
		if err != nil {
			return err
		}

		msgs, err := ch.Consume(
		    q.Name, // queue
		    "",     // consumer
		    true,   // auto-ack
		    false,  // exclusive
		    false,  // no-local
		    false,  // no-wait
		    nil,    // args
		)
		if err != nil {
			return err
		}

		forever := make(chan bool)
		go func() {
		    for d := range msgs {
		        log.Printf("Received a message: %s", d.Body)
			q := task_pipeline.New(data, d.Body)
			q.Pipeline()
		    }
		}()
		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	}
	return nil
}

