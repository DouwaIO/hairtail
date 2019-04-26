package server

import (
//	"github.com/DouwaIO/hairtail/src/task"
	task_pipeline "github.com/DouwaIO/hairtail/src/pipeline"
	task_service "github.com/DouwaIO/hairtail/src/service"
	"net/http"
	"github.com/gin-gonic/gin"
//	"github.com/DouwaIO/hairtail/src/task/step"
	"github.com/DouwaIO/hairtail/src/yaml/schema"
	"github.com/DouwaIO/hairtail/src/yaml/pipeline"
	"github.com/DouwaIO/hairtail/src/model"
	"github.com/DouwaIO/hairtail/src/store"
	"github.com/DouwaIO/hairtail/src/utils"
)

func Schema(c *gin.Context) {
	in := &model.Schema{}

	err := c.Bind(in)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	parsed, err := schema.ParseString(in.Data)
	if err != nil {
		c.String(400, "yaml type error")
		return
	}

	_, err = store.FromContext(c).GetSchema(parsed.Name)
	if err != nil {

		gen_id := utils.GeneratorId()
		newscd := &model.Schema{
			ID: gen_id,
			Name:  parsed.Name,
			Data: in.Data,
		}
		err = store.FromContext(c).CreateSchema(newscd)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		c.JSON(200, newscd)

	} else {
		c.JSON(409, "schema exiting")
	}
}

func Pipeline(c *gin.Context) {
	in := &model.Pipeline{}

	err := c.Bind(in)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	parsed, err := pipeline.ParseString(in.Data)
	if err != nil {
		c.String(400, "yaml type error")
		return
	}

	_, err = store.FromContext(c).GetPipeline(parsed.Name)
	if err != nil {
		gen_id := utils.GeneratorId()
		newscd := &model.Pipeline{
			ID: gen_id,
			Name:  parsed.Name,
			Data: in.Data,
		}
		err = store.FromContext(c).CreatePipeline(newscd)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		if len(parsed.Services) > 0 {
			for _, service := range parsed.Services {
				_, err = store.FromContext(c).GetService(service.Name, newscd.ID)
				if err != nil {
					gen_id := utils.GeneratorId()
					newser := &model.Service{
						ID: gen_id,
						Name: service.Name,
						Type: service.Type,
						Pipeline: newscd.ID,
						Data: in.Data,
					}
					err = store.FromContext(c).CreateService(newser)
					if err != nil {
						c.String(500, err.Error())
						return
					}
					task_service.Service(service, parsed.Pipeline, newser.ID, store.FromContext(c))
					//err = task.StartService(newser)
					//if err != nil {
					//	c.String(500, err.Error())
					//	return
					//}
				}
			}
		} else {
			//step.Send_Message("amqp", "aa","aa","aa")
			task_pipeline.Pipeline(parsed.Pipeline, []byte(in.Context))
			//err = step.New(in.Data, []byte(in.Context))
			//if err != nil {
			//	c.String(500, err.Error())
			//	return
			//}

		}
		c.JSON(200, newscd)

	} else {
		c.JSON(409, "schema exiting")
	}
}

func PostData(c *gin.Context) {
	in := &model.PostData{}
	err := c.Bind(in)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	pipeline2, err := store.FromContext(c).GetPipeline(in.Pipeline)
	if err != nil {
		c.JSON(400, "pipeline find error")
		return
	}

	service2, err := store.FromContext(c).GetService(in.Service, pipeline2.ID)
	if err != nil {
		c.JSON(400, "service find error")
		return
	}
	if service2.Type == "API" {
		parsed, err := pipeline.ParseString(service2.Data)
		if err != nil {
			c.JSON(400, "yaml error")
			return
		}

		task_pipeline.Pipeline(parsed.Pipeline, []byte(in.Context))
		c.JSON(200, "OK")
		return
	}

	c.JSON(400, "service Type error")


}

