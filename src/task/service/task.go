package service

import (
	"github.com/DouwaIO/hairtail/src/pipeline"
	"github.com/DouwaIO/hairtail/src/task"
)


type fifo struct {
	service *pipeline.Container
	pipeline []*pipeline.Container
}

func New(service *pipeline.Container, pipeline2 []*pipeline.Container) task.Service  {
	return &fifo{
		service: service,
		pipeline: pipeline2,
	}
}

//func Call_func(service pipeline.Container)  {
//	if service.Type == "MQ" {
//		go MQ(service.Settings["protocol"].(string), service.Settings["host"].(string), service.Settings["user"].(string), service.Settings["pwd"].(string), service.Settings["topic"].(string), service.Settings["ackPolicy"].(string), service.Data)
//	}
//}

func (q *fifo) Service() error {
	if q.service.Type == "MQ" {
		go MQ(q.service.Settings["protocol"].(string), q.service.Settings["host"].(string), q.service.Settings["user"].(string), q.service.Settings["pwd"].(string), q.service.Settings["topic"].(string), q.service.Settings["ackPolicy"].(string), q.pipeline)
	}
	if q.service.Type == "DB" {
		go DB(q.service.Settings["db_type"].(string), q.service.Settings["host"].(string), q.service.Settings["port"].(string), q.service.Settings["user"].(string), q.service.Settings["pwd"].(string), q.service.Settings["name"].(string), q.service.Settings["table"].(string), q.service.Settings["column"].(string), 1, 1)
	}
	return nil
}
