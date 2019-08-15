package service

import (
	"fmt"

	log "github.com/sirupsen/logrus"

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

	log.Debugf("MQ protocal: %s", protocol)
	if protocol == "amqp" {
		conn, err := amqp.Dial(connectStr)
		if err != nil {
			log.Errorf("MQ connect error: %s", err)
			return err
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Errorf("MQ get channel error: %s", err)
			return err
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			topic, // name
			false, // durable
			// true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			log.Errorf("MQ decalre queue error: %s", err)
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
			log.Errorf("MQ get consume error: %s", err)
			return err
		}
		log.Debugf("MQ service starting...")

		forever := make(chan bool)
		go func() {
			for d := range msgs {
				log.Debugf("Received a message: %s", d.Body)

				go func() {
					log.Debugf("go to pipeline")
					err := s.RunStep(d.Body)
					if err != nil {
						log.Errorf("Pipeline step error: %s", err)
						d.Ack(false)
						return
					}

					d.Ack(true)
				}()
			}
		}()

		log.Info("MQ waiting for messages.")
		<-forever
	}
	return nil
}
