package service

import (
	"github.com/DouwaIO/hairtail/src/model"
	task_pipeline "github.com/DouwaIO/hairtail/src/pipeline"
	"github.com/DouwaIO/hairtail/src/store"
	"github.com/DouwaIO/hairtail/src/utils"
	"github.com/DouwaIO/hairtail/src/yaml/pipeline"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func MQ(protocol, host, user, pwd, topic, ackPolicy string, data []*pipeline.Container, service string, v store.Store) error {
	//datas, err := v.GetDataList(service)
	//if err != nil {
	//	log.Printf("no cache data")
	//}
	//for _, d := range datas {
	//	task_pipeline.Pipeline(data, []byte(d.Data))
	//	v.DeleteData(d)
	//}
	if protocol == "amqp" {
		mq_connct := protocol + "://" + user + ":" + pwd + "@" + host + "/"
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

			false, // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
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
				log.Printf("Received a message")
				// gen_id := utils.GeneratorId()
				// newdata := &model.Data{
				// 	ID: gen_id,
				// 	Service: service,
				// 	Data: string(d.Body),
				// }
				// err = v.CreateData(newdata)
				// if err != nil {
				// 	log.Printf("add data error")
				// }
				// v.DeleteData(newdata)
				go func() {
					currentTime := time.Now().Unix()
					gen_id := utils.GeneratorId()
					newdata := &model.Build{
						ID:      gen_id,
						Service: service,
						Data:    string(d.Body),
						//Status: model.StatusPending,
						Status:     model.StatusRunning,
						Timestamp:  currentTime,
						Timestamp2: int64(0),
					}
					err = v.CreateBuild(newdata)
					if err != nil {
						log.Printf("add data error")
					}
					status := task_pipeline.Pipeline(data, d.Body)
					currentTime = time.Now().Unix()
					newdata.Status = status
					newdata.Timestamp2 = currentTime
					err = v.UpdateBuild(newdata)
					if err != nil {
						log.Printf("add data error")
					}
				}()
			}
		}()
		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	}
	return nil
}
