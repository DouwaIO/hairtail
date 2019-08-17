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

	log.Debugf("mq protocal: %s", protocol)
	if protocol == "amqp" {
		conn, err := amqp.Dial(connectStr)
		if err != nil {
			log.Errorf("mq connect error: %s", err)
			return err
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Errorf("mq get channel error: %s", err)
			return err
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			topic, // name
			// false, // durable
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			log.Errorf("mq decalre queue error: %s", err)
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
			log.Errorf("mq get consume error: %s", err)
			return err
		}
		log.Infof("mq service connected")

		forever := make(chan bool)
		go func() {
			for d := range msgs {
				// log.Debugf("MQ received a message: %s", d.Body)
				log.Debug("mq received a new message")

				err := s.RunPipeline(d.Body)
				if err != nil {
					d.Ack(false)
					continue
				}

				d.Ack(true)
			}
		}()
		log.Debugf("mq waiting for messages")
		<-forever
	}
	return nil
}
