package pipeline

import (
	yaml_pipeline "github.com/DouwaIO/hairtail/src/yaml/pipeline"
	"github.com/DouwaIO/hairtail/src/task"
)



func Pipeline(pipeline []*yaml_pipeline.Container, data []byte) error {
	//parsed, err := yaml_pipeline.ParseString(q.config)
	//if err != nil {
	//	return errors.New("yaml type error")
	//}
	if len(pipeline) > 0 {
		for _, pipeline2 := range pipeline {
			if _, ok := task.Funcs[pipeline2.Type]; ok {
				data2 := task.CallPipeline(pipeline2, data)
				if data2 != nil {
					data = data2
				}
			} else {
				return nil
			}

		}
	}
	return nil
}


