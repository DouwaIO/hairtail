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
	"reflect"
	"errors"
//	"os"
//	"github.com/Sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/DouwaIO/hairtail/src/pipeline"
	"github.com/DouwaIO/hairtail/src/model"
)

var (
	funcs = map[string]interface{}{"MQ_Send": MQSend, "HTTP_Send": HTTPSend,  "test":Test}
)

func Call(m map[string]interface{}, name string, params ... interface{}) (result []reflect.Value, err error) {
    f := reflect.ValueOf(m[name])
    if len(params) != f.Type().NumIn() {
        err = errors.New("The number of params is not adapted.")
        return
    }
    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }
    result = f.Call(in)
    return
}
//Call(funcs, "foo")
//Call(funcs, "bar", 1, 2, 3)

func StartService(service *model.Service) error {
	parsed, err := pipeline.ParseString(service.Data)
	if err != nil {
		return errors.New("yaml type error")
	}
	//log.Printf("type %s", service.Type)
	if service.Type == "MQ" {
		for _, service2 := range parsed.Services {
			if service.Type == service2.Type {
				go MQ(service2.Settings["protocol"].(string), service2.Settings["host"].(string), service2.Settings["user"].(string), service2.Settings["pwd"].(string), service2.Settings["topic"].(string), service2.Settings["ackPolicy"].(string), service.Data)
			}
		}
	}
	return nil
}

func New(config string, data []byte) error {
	parsed, err := pipeline.ParseString(config)
	if err != nil {
		return errors.New("yaml type error")
	}

	if len(parsed.Pipeline) > 0 {
		for _, pipeline := range parsed.Pipeline {
			if _, ok := funcs[pipeline.Type]; ok {
				if pipeline.Type == "MQ_Send" {
					Call(funcs, pipeline.Type, pipeline.Settings["protocol"], pipeline.Settings["host"], pipeline.Settings["user"], pipeline.Settings["pwd"], pipeline.Settings["topic"], data)
				}
				if pipeline.Type == "HTTP_Send" {
					Call(funcs, pipeline.Type, pipeline.Settings["url"], data)
				}
				if pipeline.Type == "test" {
					Call(funcs, pipeline.Type, data)
				}
			} else {
				return nil
			}

		}
	}
	return nil
}


func Test(data []byte) error {
	log.Printf("Data :", string(data))
	return nil
}

func MQ(protocol, host, user, pwd, topic, ackPolicy, data string) error {
	if protocol == "amqp" {
		mq_connct := protocol+"://"+user+":"+pwd+"@"+host+"/"
		//conn, err := amqp.Dial("amqp://root:123456@47.97.182.182:32222/")
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
		    //"hello", // name
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
			err := New(data, d.Body)
			if err != nil {
			    return
			}
		    }
		}()
		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	}
	return nil
}

