package server

import (
//	"github.com/DouwaIO/hairtail/src/task"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/DouwaIO/hairtail/src/task/step"
	"github.com/DouwaIO/hairtail/src/schema"
	"github.com/DouwaIO/hairtail/src/pipeline"
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
					err = step.StartService(newser)
					if err != nil {
						c.String(500, err.Error())
						return
					}
				}
			}
		} else {
			//step.Send_Message("amqp", "aa","aa","aa")
			err = step.New(in.Data, []byte(in.Context))
			if err != nil {
				c.String(500, err.Error())
				return
			}

		}
		c.JSON(200, newscd)

	} else {
		c.JSON(409, "schema exiting")
	}
}

func PostData(c *gin.Context) {
	c.String(200, "hello world")
	//go step.New("test")
	//go task.FromContext(c).MQ("amqp","aa","aa","aa","aa")
	//c.JSON(200, "test teste")
	//user, err := store.FromContext(c).GetUser(1)
	//if err != nil {
	//	new_user := &model.User{
	//		Login:  "test",
	//		Token: "test test",
	//	}
	//	err = store.FromContext(c).CreateUser(new_user)
	//	if err != nil {
	//		c.String(500, err.Error())
	//		return
	//	}
	//	c.JSON(200, new_user)

	//} else {
	//	c.JSON(200, user)
	//}
}

