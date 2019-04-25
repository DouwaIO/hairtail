package main

import (
    "log"

    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
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

    body := "{\"name\":\"fabric_stock_in\",\"type\":\"add\",\"key\":\"F000323546\",\"time\":\"2019-04-23T03:10:33.775Z\",\"data\":{\"bill_no\":\"string\",\"details\":[{\"grade\":\"string\",\"location_no\":\"string\",\"line\":\"ASDF111\"},{\"fabric_no\":\"string\"}],\"bill_type_name\":\"采购入库\"}}"
    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        })
    log.Printf(" [x] Sent %s", body)
    failOnError(err, "Failed to publish a message")
}
