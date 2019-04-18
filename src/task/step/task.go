// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package step

import (
	"log"
//	"os"
//	"github.com/Sirupsen/logrus"
	"github.com/streadway/amqp"
)

// var db *gorm.DB

// New creates a database connection for the given driver and datasource
// and returns a new Store.
func New(config string) error {
	MQ("amqp","aa","aa","aa","aa")
	return nil
	//return &datastore{
	//	DB:     open(driver, config),
	//	driver: driver,
	//	config: config,
	//}
}

func MQ(protocol, host, user, topic, ackPolicy string) error {
	//db,err := gorm.Open("sqlite3", "test.db")
	if protocol == "amqp" {
		conn, err := amqp.Dial("amqp://root:123456@47.97.182.182:32222/")
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
		    "hello", // name
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
		    }
		}()
		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	}
	return nil
}

