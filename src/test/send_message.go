package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Detail struct {
	FabricNo string `json:"fabric_no"`
	Line     string `json:"line"`
	Quantity int    `json:"quantity"`
}

type Bill struct {
	BillNo  string    `json:"bill_no"`
	Details []*Detail `json:"details"`
}

// 只能在安装 rabbitmq 的服务器上操作
func main() {
	conn, err := amqp.Dial("amqp://root:123456@47.97.182.182:32222/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	start := time.Now().Unix()
	log.Println("start is ", start)

	for i := 1; i <= 10; i++ {
		var details []*Detail
		for j := 1; j <= 100; j++ {
			detail := &Detail{
				FabricNo: fmt.Sprintf("f%d", j),
				Line:     fmt.Sprintf("l%d", j),
				Quantity: j,
			}
			details = append(details, detail)
		}
		bill := Bill{
			BillNo:  fmt.Sprintf("b%d", i),
			Details: details,
		}
		body, _ := json.Marshal(bill)
		log.Printf("No. %d\nBody: %s\n", i, body)

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		failOnError(err, "Failed to publish a message")
		log.Println("send success")
	}
}
