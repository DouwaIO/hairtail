package service

import (
	"log"
	"fmt"

	"github.com/streadway/amqp"
)

// func MQ(protocol, host, user, pwd, topic, ackPolicy string, data []*pipeline.Task, service string, v store.Store) error {
func MQ(s *Service) error {
    consumerName := ""
    protocol := s.Settings["protocol"].(string)
    topic := s.Settings["topic"].(string)
    // ackPolicy := settings["ackPolicy"].(string)

    connectStr := fmt.Sprintf("%s://%s:%s@%s/",
        protocol,
        s.Settings["user"].(string),
        s.Settings["pwd"].(string),
        s.Settings["host"].(string))

	if protocol == "amqp" {
		conn, err := amqp.Dial(connectStr)
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
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
            // Failed to declare a queue
			return err
		}

		msgs, err := ch.Consume(
			q.Name,       // queue
			consumerName, // consumer
			false,        // auto-ack
			false,        // exclusive
			false,        // no-local
			false,        // no-wait
			nil,          // args
		)
		if err != nil {
            // Failed to register a consumer
			return err
		}

		forever := make(chan bool)
		go func() {
			for d := range msgs {
                log.Printf("Received a message: %s", d.Body)

				go func() {
					err := s.RunStep(d.Body)
                    if err != nil {
                        log.Printf("Pipeline step error: %s", err)
                        d.Ack(false)
                        return
                    }

                    d.Ack(true)
				}()
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	}
	return nil
}
