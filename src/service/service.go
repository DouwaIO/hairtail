package service

import (
	yaml_pipeline "github.com/DouwaIO/hairtail/src/yaml/pipeline"
)

func CallService(service *yaml_pipeline.Container, pipeline2 []*yaml_pipeline.Container) error {
	if service.Type == "MQ" {
		go MQ(service.Settings["protocol"].(string), service.Settings["host"].(string), service.Settings["user"].(string), service.Settings["pwd"].(string), service.Settings["topic"].(string), service.Settings["ackPolicy"].(string), pipeline2)
	}
	if service.Type == "DB" {
		go DB(service.Settings["db_type"].(string), service.Settings["host"].(string), service.Settings["port"].(string), service.Settings["user"].(string), service.Settings["pwd"].(string), service.Settings["name"].(string), service.Settings["table"].(string), service.Settings["column"].(string), 1, 1)
	}
	return nil
}


func Service(service *yaml_pipeline.Container, yaml_pipeline []*yaml_pipeline.Container) error {
	return CallService(service, yaml_pipeline)
}
